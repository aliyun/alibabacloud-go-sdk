// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnDeliverTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliverId(v int64) *DeleteDcdnDeliverTaskRequest
	GetDeliverId() *int64
}

type DeleteDcdnDeliverTaskRequest struct {
	// The IDs of the tracking tasks that you want to delete. You can call the [DescribeCdnDeliverList](https://help.aliyun.com/document_detail/270043.html) operation to query task IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92
	DeliverId *int64 `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
}

func (s DeleteDcdnDeliverTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnDeliverTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnDeliverTaskRequest) GetDeliverId() *int64 {
	return s.DeliverId
}

func (s *DeleteDcdnDeliverTaskRequest) SetDeliverId(v int64) *DeleteDcdnDeliverTaskRequest {
	s.DeliverId = &v
	return s
}

func (s *DeleteDcdnDeliverTaskRequest) Validate() error {
	return dara.Validate(s)
}
