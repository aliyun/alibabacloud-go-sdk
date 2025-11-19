// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CredentialResult
	GetCode() *string
	SetData(v *Credential) *CredentialResult
	GetData() *Credential
	SetRequestId(v string) *CredentialResult
	GetRequestId() *string
}

type CredentialResult struct {
	Code      *string     `json:"code,omitempty" xml:"code,omitempty"`
	Data      *Credential `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CredentialResult) String() string {
	return dara.Prettify(s)
}

func (s CredentialResult) GoString() string {
	return s.String()
}

func (s *CredentialResult) GetCode() *string {
	return s.Code
}

func (s *CredentialResult) GetData() *Credential {
	return s.Data
}

func (s *CredentialResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CredentialResult) SetCode(v string) *CredentialResult {
	s.Code = &v
	return s
}

func (s *CredentialResult) SetData(v *Credential) *CredentialResult {
	s.Data = v
	return s
}

func (s *CredentialResult) SetRequestId(v string) *CredentialResult {
	s.RequestId = &v
	return s
}

func (s *CredentialResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
