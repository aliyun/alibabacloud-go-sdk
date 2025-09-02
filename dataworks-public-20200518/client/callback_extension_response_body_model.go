// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallbackExtensionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CallbackExtensionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CallbackExtensionResponseBody
	GetSuccess() *string
}

type CallbackExtensionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF020E7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true
	//
	// false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CallbackExtensionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CallbackExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *CallbackExtensionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CallbackExtensionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CallbackExtensionResponseBody) SetRequestId(v string) *CallbackExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CallbackExtensionResponseBody) SetSuccess(v string) *CallbackExtensionResponseBody {
	s.Success = &v
	return s
}

func (s *CallbackExtensionResponseBody) Validate() error {
	return dara.Validate(s)
}
