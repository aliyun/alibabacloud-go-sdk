// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendHotlineServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SuspendHotlineServiceResponseBody
	GetCode() *string
	SetMessage(v string) *SuspendHotlineServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *SuspendHotlineServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendHotlineServiceResponseBody
	GetSuccess() *bool
}

type SuspendHotlineServiceResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendHotlineServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendHotlineServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendHotlineServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SuspendHotlineServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SuspendHotlineServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendHotlineServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendHotlineServiceResponseBody) SetCode(v string) *SuspendHotlineServiceResponseBody {
	s.Code = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) SetMessage(v string) *SuspendHotlineServiceResponseBody {
	s.Message = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) SetRequestId(v string) *SuspendHotlineServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) SetSuccess(v bool) *SuspendHotlineServiceResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
