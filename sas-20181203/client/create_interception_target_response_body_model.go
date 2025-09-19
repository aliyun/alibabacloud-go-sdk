// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInterceptionTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateInterceptionTargetResponseBody
	GetRequestId() *string
	SetResult(v bool) *CreateInterceptionTargetResponseBody
	GetResult() *bool
}

type CreateInterceptionTargetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5B9ECCC0-38F7-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateInterceptionTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInterceptionTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInterceptionTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInterceptionTargetResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CreateInterceptionTargetResponseBody) SetRequestId(v string) *CreateInterceptionTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInterceptionTargetResponseBody) SetResult(v bool) *CreateInterceptionTargetResponseBody {
	s.Result = &v
	return s
}

func (s *CreateInterceptionTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
