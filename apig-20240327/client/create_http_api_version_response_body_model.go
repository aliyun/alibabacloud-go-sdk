// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpApiVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHttpApiVersionResponseBody
	GetCode() *string
	SetMessage(v string) *CreateHttpApiVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHttpApiVersionResponseBody
	GetRequestId() *string
}

type CreateHttpApiVersionResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 3ACFC7A7-45A9-58CF-B2D5-765B60254695
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateHttpApiVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpApiVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHttpApiVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHttpApiVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHttpApiVersionResponseBody) SetCode(v string) *CreateHttpApiVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHttpApiVersionResponseBody) SetMessage(v string) *CreateHttpApiVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHttpApiVersionResponseBody) SetRequestId(v string) *CreateHttpApiVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHttpApiVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
