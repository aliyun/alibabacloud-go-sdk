// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDetectionTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetImageDetectionTaskResultRequest
	GetClientToken() *string
	SetTaskId(v string) *GetImageDetectionTaskResultRequest
	GetTaskId() *string
}

type GetImageDetectionTaskResultRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The task ID returned by `CreateImageDetectionTask`.
	//
	// This parameter is required.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetImageDetectionTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetImageDetectionTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageDetectionTaskResultRequest) SetClientToken(v string) *GetImageDetectionTaskResultRequest {
	s.ClientToken = &v
	return s
}

func (s *GetImageDetectionTaskResultRequest) SetTaskId(v string) *GetImageDetectionTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetImageDetectionTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
