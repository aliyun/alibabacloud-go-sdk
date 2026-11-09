// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegisterCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegisterCode(v string) *GetRegisterCodeResponseBody
	GetRegisterCode() *string
	SetRequestId(v string) *GetRegisterCodeResponseBody
	GetRequestId() *string
}

type GetRegisterCodeResponseBody struct {
	// The registration code of the CA certificate.
	//
	// example:
	//
	// 13274673-8f90-4630-bea1-9cccb25756ad2089******
	RegisterCode *string `json:"RegisterCode,omitempty" xml:"RegisterCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRegisterCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRegisterCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegisterCodeResponseBody) GetRegisterCode() *string {
	return s.RegisterCode
}

func (s *GetRegisterCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRegisterCodeResponseBody) SetRegisterCode(v string) *GetRegisterCodeResponseBody {
	s.RegisterCode = &v
	return s
}

func (s *GetRegisterCodeResponseBody) SetRequestId(v string) *GetRegisterCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegisterCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
