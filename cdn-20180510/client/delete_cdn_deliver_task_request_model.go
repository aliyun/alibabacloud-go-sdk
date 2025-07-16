// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCdnDeliverTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliverId(v int64) *DeleteCdnDeliverTaskRequest
	GetDeliverId() *int64
}

type DeleteCdnDeliverTaskRequest struct {
	// The ID of the tracking task that you want to delete. You can call the [DescribeCdnDeliverList](https://help.aliyun.com/document_detail/270877.html) operation to query task IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DeliverId *int64 `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
}

func (s DeleteCdnDeliverTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCdnDeliverTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteCdnDeliverTaskRequest) GetDeliverId() *int64 {
	return s.DeliverId
}

func (s *DeleteCdnDeliverTaskRequest) SetDeliverId(v int64) *DeleteCdnDeliverTaskRequest {
	s.DeliverId = &v
	return s
}

func (s *DeleteCdnDeliverTaskRequest) Validate() error {
	return dara.Validate(s)
}
