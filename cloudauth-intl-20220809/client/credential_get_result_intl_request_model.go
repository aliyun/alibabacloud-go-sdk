// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialGetResultIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTransactionId(v string) *CredentialGetResultIntlRequest
	GetTransactionId() *string
}

type CredentialGetResultIntlRequest struct {
	// Unique identifier for the authentication request
	//
	// This parameter is required.
	//
	// example:
	//
	// 4ab0b***abde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CredentialGetResultIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialGetResultIntlRequest) GoString() string {
	return s.String()
}

func (s *CredentialGetResultIntlRequest) GetTransactionId() *string {
	return s.TransactionId
}

func (s *CredentialGetResultIntlRequest) SetTransactionId(v string) *CredentialGetResultIntlRequest {
	s.TransactionId = &v
	return s
}

func (s *CredentialGetResultIntlRequest) Validate() error {
	return dara.Validate(s)
}
