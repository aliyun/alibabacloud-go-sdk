// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCustomDomainResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteCustomDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomDomainResponseBody
	GetRequestId() *string
}

type DeleteCustomDomainResponseBody struct {
	// The response code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B5AD8B54-4358-5F5B-ACAA-52F2016459C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteCustomDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCustomDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomDomainResponseBody) SetCode(v string) *DeleteCustomDomainResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomDomainResponseBody) SetMessage(v string) *DeleteCustomDomainResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomDomainResponseBody) SetRequestId(v string) *DeleteCustomDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
