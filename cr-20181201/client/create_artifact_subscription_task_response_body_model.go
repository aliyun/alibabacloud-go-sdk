// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactSubscriptionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateArtifactSubscriptionTaskResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateArtifactSubscriptionTaskResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateArtifactSubscriptionTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateArtifactSubscriptionTaskResponseBody
	GetTaskId() *string
}

type CreateArtifactSubscriptionTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 132D314B-BDD4-564C-89FE-3E2BAE115239
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// crast-40le4es9yh0p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateArtifactSubscriptionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactSubscriptionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactSubscriptionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateArtifactSubscriptionTaskResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateArtifactSubscriptionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateArtifactSubscriptionTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateArtifactSubscriptionTaskResponseBody) SetCode(v string) *CreateArtifactSubscriptionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateArtifactSubscriptionTaskResponseBody) SetIsSuccess(v bool) *CreateArtifactSubscriptionTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateArtifactSubscriptionTaskResponseBody) SetRequestId(v string) *CreateArtifactSubscriptionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArtifactSubscriptionTaskResponseBody) SetTaskId(v string) *CreateArtifactSubscriptionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateArtifactSubscriptionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
