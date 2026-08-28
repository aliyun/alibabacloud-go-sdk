// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateNetworkAccessResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateNetworkAccessResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateNetworkAccessResponseBody
	GetRequestId() *string
}

type UpdateNetworkAccessResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 019F3F7D-9EC4-5F8B-A3F7-97E1369C31BD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateNetworkAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkAccessResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAccessResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateNetworkAccessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNetworkAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNetworkAccessResponseBody) SetCode(v string) *UpdateNetworkAccessResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNetworkAccessResponseBody) SetMessage(v string) *UpdateNetworkAccessResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNetworkAccessResponseBody) SetRequestId(v string) *UpdateNetworkAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNetworkAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
