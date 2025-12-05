// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateServiceVersionResponseBody
	GetCode() *string
	SetMessage(v string) *CreateServiceVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceVersionResponseBody
	GetRequestId() *string
}

type CreateServiceVersionResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0B373A13-9BB8-5068-9C94-AD6D39E6BBA2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateServiceVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceVersionResponseBody) SetCode(v string) *CreateServiceVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceVersionResponseBody) SetMessage(v string) *CreateServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceVersionResponseBody) SetRequestId(v string) *CreateServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
