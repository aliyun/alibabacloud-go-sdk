// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpotPriceHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSpotPriceHistoryResponseBody
	GetRequestId() *string
	SetSpotPriceHistory(v []*SpotPriceItem) *GetSpotPriceHistoryResponseBody
	GetSpotPriceHistory() []*SpotPriceItem
	SetTotalCount(v int32) *GetSpotPriceHistoryResponseBody
	GetTotalCount() *int32
}

type GetSpotPriceHistoryResponseBody struct {
	// example:
	//
	// 8BDA4440-DD3C-5F4B-BBDD-94A9CE1E75C7
	RequestId        *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SpotPriceHistory []*SpotPriceItem `json:"SpotPriceHistory,omitempty" xml:"SpotPriceHistory,omitempty" type:"Repeated"`
	// example:
	//
	// 194
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSpotPriceHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSpotPriceHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpotPriceHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSpotPriceHistoryResponseBody) GetSpotPriceHistory() []*SpotPriceItem {
	return s.SpotPriceHistory
}

func (s *GetSpotPriceHistoryResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetSpotPriceHistoryResponseBody) SetRequestId(v string) *GetSpotPriceHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpotPriceHistoryResponseBody) SetSpotPriceHistory(v []*SpotPriceItem) *GetSpotPriceHistoryResponseBody {
	s.SpotPriceHistory = v
	return s
}

func (s *GetSpotPriceHistoryResponseBody) SetTotalCount(v int32) *GetSpotPriceHistoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetSpotPriceHistoryResponseBody) Validate() error {
	if s.SpotPriceHistory != nil {
		for _, item := range s.SpotPriceHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
