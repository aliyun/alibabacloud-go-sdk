// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInventoryJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *CheckInventoryJobRequest
	GetJobId() *int64
}

type CheckInventoryJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1001
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CheckInventoryJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckInventoryJobRequest) GoString() string {
	return s.String()
}

func (s *CheckInventoryJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *CheckInventoryJobRequest) SetJobId(v int64) *CheckInventoryJobRequest {
	s.JobId = &v
	return s
}

func (s *CheckInventoryJobRequest) Validate() error {
	return dara.Validate(s)
}
