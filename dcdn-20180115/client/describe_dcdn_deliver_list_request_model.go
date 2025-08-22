// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDeliverListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliverId(v int64) *DescribeDcdnDeliverListRequest
	GetDeliverId() *int64
}

type DescribeDcdnDeliverListRequest struct {
	// The ID of the tracking task that you want to query. If you do not specify an ID, all tracking tasks are queried.
	//
	// example:
	//
	// 92
	DeliverId *int64 `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
}

func (s DescribeDcdnDeliverListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDeliverListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeliverListRequest) GetDeliverId() *int64 {
	return s.DeliverId
}

func (s *DescribeDcdnDeliverListRequest) SetDeliverId(v int64) *DescribeDcdnDeliverListRequest {
	s.DeliverId = &v
	return s
}

func (s *DescribeDcdnDeliverListRequest) Validate() error {
	return dara.Validate(s)
}
