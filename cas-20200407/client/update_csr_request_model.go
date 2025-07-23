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
	// The ID of the CSR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5209
	CsrId *int64 `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
	// The private key content of the certificate in the PEM format.
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
