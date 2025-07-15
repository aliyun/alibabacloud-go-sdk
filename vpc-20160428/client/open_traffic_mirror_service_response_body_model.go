// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenTrafficMirrorServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OpenTrafficMirrorServiceResponseBody
	GetCode() *string
	SetMessage(v string) *OpenTrafficMirrorServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *OpenTrafficMirrorServiceResponseBody
	GetRequestId() *string
}

type OpenTrafficMirrorServiceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information returned after traffic mirror is enabled.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4FCCF008-4C13-4231-BE77-D5203801A9E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenTrafficMirrorServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenTrafficMirrorServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenTrafficMirrorServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *OpenTrafficMirrorServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenTrafficMirrorServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenTrafficMirrorServiceResponseBody) SetCode(v string) *OpenTrafficMirrorServiceResponseBody {
	s.Code = &v
	return s
}

func (s *OpenTrafficMirrorServiceResponseBody) SetMessage(v string) *OpenTrafficMirrorServiceResponseBody {
	s.Message = &v
	return s
}

func (s *OpenTrafficMirrorServiceResponseBody) SetRequestId(v string) *OpenTrafficMirrorServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenTrafficMirrorServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
