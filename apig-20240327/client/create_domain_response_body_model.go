// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDomainResponseBody
	GetCode() *string
	SetData(v *CreateDomainResponseBodyData) *CreateDomainResponseBody
	GetData() *CreateDomainResponseBodyData
	SetMessage(v string) *CreateDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDomainResponseBody
	GetRequestId() *string
}

type CreateDomainResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *CreateDomainResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 0C2D1C68-0D93-5561-8EE6-FDB7BF067A30
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDomainResponseBody) GetData() *CreateDomainResponseBodyData {
	return s.Data
}

func (s *CreateDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDomainResponseBody) SetCode(v string) *CreateDomainResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDomainResponseBody) SetData(v *CreateDomainResponseBodyData) *CreateDomainResponseBody {
	s.Data = v
	return s
}

func (s *CreateDomainResponseBody) SetMessage(v string) *CreateDomainResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDomainResponseBody) SetRequestId(v string) *CreateDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDomainResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateDomainResponseBodyData struct {
	// The ID of the domain name.
	//
	// example:
	//
	// d-cpu1aullhtgkidg7sa4g
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
}

func (s CreateDomainResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBodyData) GetDomainId() *string {
	return s.DomainId
}

func (s *CreateDomainResponseBodyData) SetDomainId(v string) *CreateDomainResponseBodyData {
	s.DomainId = &v
	return s
}

func (s *CreateDomainResponseBodyData) Validate() error {
	return dara.Validate(s)
}
