// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeliveryHistoryJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int32) *DeleteDeliveryHistoryJobRequest
	GetJobId() *int32
}

type DeleteDeliveryHistoryJobRequest struct {
	// The ID of the historical event delivery task to be deleted.
	//
	// You can call the [ListDeliveryHistoryJobs](https://help.aliyun.com/document_detail/188101.html) operation to query task IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16602
	JobId *int32 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteDeliveryHistoryJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeliveryHistoryJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryHistoryJobRequest) GetJobId() *int32 {
	return s.JobId
}

func (s *DeleteDeliveryHistoryJobRequest) SetJobId(v int32) *DeleteDeliveryHistoryJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteDeliveryHistoryJobRequest) Validate() error {
	return dara.Validate(s)
}
