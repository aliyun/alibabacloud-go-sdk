// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSensitiveScanResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetSensitiveScanResultRequest
	GetTaskId() *string
}

type GetSensitiveScanResultRequest struct {
	// The task ID returned by `CreateSensitiveScanTask`.
	//
	// This parameter is required.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetSensitiveScanResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSensitiveScanResultRequest) GoString() string {
	return s.String()
}

func (s *GetSensitiveScanResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetSensitiveScanResultRequest) SetTaskId(v string) *GetSensitiveScanResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetSensitiveScanResultRequest) Validate() error {
	return dara.Validate(s)
}
