// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSilencePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *DeleteSilencePolicyResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteSilencePolicyResponseBody
	GetRequestId() *string
}

type DeleteSilencePolicyResponseBody struct {
	// Indicates whether the silence policy was deleted successfully. Valid values:
	//
	// 	- `true`: The silence policy was deleted successfully.
	//
	// 	- `false`: The silence policy failed to be deleted.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The operation that you want to perform. Set the value to **DeleteSilencePolicy**.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSilencePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSilencePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSilencePolicyResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteSilencePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSilencePolicyResponseBody) SetIsSuccess(v bool) *DeleteSilencePolicyResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteSilencePolicyResponseBody) SetRequestId(v string) *DeleteSilencePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSilencePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
