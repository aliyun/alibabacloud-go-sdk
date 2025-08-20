// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepoTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRepoTriggerResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *UpdateRepoTriggerResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateRepoTriggerResponseBody
	GetRequestId() *string
}

type UpdateRepoTriggerResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 32535049-ED91-4589-98C0-7C88766EDF1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRepoTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepoTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepoTriggerResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRepoTriggerResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateRepoTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRepoTriggerResponseBody) SetCode(v string) *UpdateRepoTriggerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRepoTriggerResponseBody) SetIsSuccess(v bool) *UpdateRepoTriggerResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateRepoTriggerResponseBody) SetRequestId(v string) *UpdateRepoTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRepoTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
