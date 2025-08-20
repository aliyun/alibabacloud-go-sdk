// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRepoTriggerResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateRepoTriggerResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateRepoTriggerResponseBody
	GetRequestId() *string
	SetTriggerId(v string) *CreateRepoTriggerResponseBody
	GetTriggerId() *string
}

type CreateRepoTriggerResponseBody struct {
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
	// B79F5E0E-8770-407D-BCB6-ECF4BA9C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the trigger.
	//
	// example:
	//
	// crw-0z4pf81pgz35****
	TriggerId *string `json:"TriggerId,omitempty" xml:"TriggerId,omitempty"`
}

func (s CreateRepoTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoTriggerResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRepoTriggerResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateRepoTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRepoTriggerResponseBody) GetTriggerId() *string {
	return s.TriggerId
}

func (s *CreateRepoTriggerResponseBody) SetCode(v string) *CreateRepoTriggerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoTriggerResponseBody) SetIsSuccess(v bool) *CreateRepoTriggerResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoTriggerResponseBody) SetRequestId(v string) *CreateRepoTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoTriggerResponseBody) SetTriggerId(v string) *CreateRepoTriggerResponseBody {
	s.TriggerId = &v
	return s
}

func (s *CreateRepoTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
