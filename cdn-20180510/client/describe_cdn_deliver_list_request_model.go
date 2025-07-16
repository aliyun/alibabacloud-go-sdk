// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDeliverListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliverId(v int64) *DescribeCdnDeliverListRequest
	GetDeliverId() *int64
}

type DescribeCdnDeliverListRequest struct {
	// The ID of the tracking task that you want to query. If you do not specify an ID, all tracking tasks are queried.
	//
	// example:
	//
	// 3
	DeliverId *int64 `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
}

func (s DescribeCdnDeliverListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeliverListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeliverListRequest) GetDeliverId() *int64 {
	return s.DeliverId
}

func (s *DescribeCdnDeliverListRequest) SetDeliverId(v int64) *DescribeCdnDeliverListRequest {
	s.DeliverId = &v
	return s
}

func (s *DescribeCdnDeliverListRequest) Validate() error {
	return dara.Validate(s)
}
