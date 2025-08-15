// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeliveryHistoryJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *GetDeliveryHistoryJobRequest
	GetJobId() *int64
}

type GetDeliveryHistoryJobRequest struct {
	// The ID of the historical event delivery task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16602
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetDeliveryHistoryJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryHistoryJobRequest) GoString() string {
	return s.String()
}

func (s *GetDeliveryHistoryJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *GetDeliveryHistoryJobRequest) SetJobId(v int64) *GetDeliveryHistoryJobRequest {
	s.JobId = &v
	return s
}

func (s *GetDeliveryHistoryJobRequest) Validate() error {
	return dara.Validate(s)
}
