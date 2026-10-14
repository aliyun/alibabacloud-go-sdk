// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageIncomeTaxDetectionTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetImageIncomeTaxDetectionTaskResultRequest
	GetClientToken() *string
	SetTaskId(v string) *GetImageIncomeTaxDetectionTaskResultRequest
	GetTaskId() *string
}

type GetImageIncomeTaxDetectionTaskResultRequest struct {
	// The idempotency token.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The task ID returned by `CreateImageIncomeTaxDetectionTask`.
	//
	// This parameter is required.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetImageIncomeTaxDetectionTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageIncomeTaxDetectionTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetImageIncomeTaxDetectionTaskResultRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetImageIncomeTaxDetectionTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageIncomeTaxDetectionTaskResultRequest) SetClientToken(v string) *GetImageIncomeTaxDetectionTaskResultRequest {
	s.ClientToken = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultRequest) SetTaskId(v string) *GetImageIncomeTaxDetectionTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetImageIncomeTaxDetectionTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
