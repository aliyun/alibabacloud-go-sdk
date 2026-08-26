// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCsrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCsrId(v int64) *UpdateCsrRequest
	GetCsrId() *int64
	SetKey(v string) *UpdateCsrRequest
	GetKey() *string
}

type UpdateCsrRequest struct {
	// The unique identifier of the CSR. The CsrId is generated when you upload the CSR. You can obtain this value by querying the CSR list. For more information, see [ListCsr](https://help.aliyun.com/document_detail/2709717.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5209
	CsrId *int64 `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
	// The certificate private key content in PEM format. This private key must match the public key cryptography contained in the CSR referenced by CsrId. Otherwise, the API returns the NotMatch.CsrAndPrivateKey error.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- MII.... -----END RSA PRIVATE KEY-----
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s UpdateCsrRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCsrRequest) GoString() string {
	return s.String()
}

func (s *UpdateCsrRequest) GetCsrId() *int64 {
	return s.CsrId
}

func (s *UpdateCsrRequest) GetKey() *string {
	return s.Key
}

func (s *UpdateCsrRequest) SetCsrId(v int64) *UpdateCsrRequest {
	s.CsrId = &v
	return s
}

func (s *UpdateCsrRequest) SetKey(v string) *UpdateCsrRequest {
	s.Key = &v
	return s
}

func (s *UpdateCsrRequest) Validate() error {
	return dara.Validate(s)
}
