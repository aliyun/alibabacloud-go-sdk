// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDefectDetectionTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetImageDefectDetectionTaskResultRequest
	GetClientToken() *string
	SetTaskId(v string) *GetImageDefectDetectionTaskResultRequest
	GetTaskId() *string
}

type GetImageDefectDetectionTaskResultRequest struct {
	// The idempotency parameter.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The task ID returned by `CreateImageDefectDetectionTask`.
	//
	// This parameter is required.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetImageDefectDetectionTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageDefectDetectionTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetImageDefectDetectionTaskResultRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetImageDefectDetectionTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageDefectDetectionTaskResultRequest) SetClientToken(v string) *GetImageDefectDetectionTaskResultRequest {
	s.ClientToken = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultRequest) SetTaskId(v string) *GetImageDefectDetectionTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
