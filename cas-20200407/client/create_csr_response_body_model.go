// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCsrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCsr(v string) *CreateCsrResponseBody
	GetCsr() *string
	SetCsrId(v int64) *CreateCsrResponseBody
	GetCsrId() *int64
	SetRequestId(v string) *CreateCsrResponseBody
	GetRequestId() *string
}

type CreateCsrResponseBody struct {
	// The content of the CSR.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The unique identifier of the CSR. You can use this value to obtain the content of the CSR. For more information about the operation that you can call to obtain the content of a CSR, see [GetCsrDetail](https://help.aliyun.com/document_detail/2709720.html).
	//
	// example:
	//
	// 3365
	CsrId *int64 `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCsrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCsrResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCsrResponseBody) GetCsr() *string {
	return s.Csr
}

func (s *CreateCsrResponseBody) GetCsrId() *int64 {
	return s.CsrId
}

func (s *CreateCsrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCsrResponseBody) SetCsr(v string) *CreateCsrResponseBody {
	s.Csr = &v
	return s
}

func (s *CreateCsrResponseBody) SetCsrId(v int64) *CreateCsrResponseBody {
	s.CsrId = &v
	return s
}

func (s *CreateCsrResponseBody) SetRequestId(v string) *CreateCsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCsrResponseBody) Validate() error {
	return dara.Validate(s)
}
