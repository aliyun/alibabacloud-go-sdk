// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDomainResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDomainResponseBody
	GetRequestId() *string
}

type DeleteDomainResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used to trace the API call link.
	//
	// example:
	//
	// A60EE5CA-1294-532A-9775-8D2FD1C6EFBF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDomainResponseBody) SetCode(v string) *DeleteDomainResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDomainResponseBody) SetMessage(v string) *DeleteDomainResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
