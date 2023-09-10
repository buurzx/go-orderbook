package orderbook

import (
	"strings"

	"github.com/shopspring/decimal"
)

// ProcessPostOnlyOrder processes a post only order.
func (ob *OrderBook) ProcessPostOnlyOrderAppendAmount(orderID, traderID string, side Side, amount, price decimal.Decimal) ([]*Trade, error) {
	defer func() {
		ob.version++
		ob.Unlock()
	}()

	ob.Lock()

	if strings.TrimSpace(orderID) == "" {
		return nil, ErrInvalidOrderID
	}

	if ob.orders[orderID] != nil {
		return nil, ErrOrderAlreadyExists
	}

	if strings.TrimSpace(traderID) == "" {
		return nil, ErrInvalidTraderID
	}

	if amount.LessThanOrEqual(decimal.Zero) {
		return nil, ErrInvalidAmount
	}

	if price.LessThanOrEqual(decimal.Zero) {
		return nil, ErrInvalidPrice
	}

	if side == Buy {
		ob.bids.UpdateAmount(ob.orders[orderID], amount)
	} else {
		ob.asks.UpdateAmount(ob.orders[orderID], amount)
	}

	return make([]*Trade, 0), nil
}
