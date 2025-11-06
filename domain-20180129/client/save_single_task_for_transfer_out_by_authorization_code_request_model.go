// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForTransferOutByAuthorizationCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationCode(v string) *SaveSingleTaskForTransferOutByAuthorizationCodeRequest
	GetAuthorizationCode() *string
	SetDomainName(v string) *SaveSingleTaskForTransferOutByAuthorizationCodeRequest
	GetDomainName() *string
}

type SaveSingleTaskForTransferOutByAuthorizationCodeRequest struct {
	// Schema of Response
	//
	// This parameter is required.
	//
	// example:
	//
	// Test2o#Lck
	AuthorizationCode *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	// The transfer key.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s SaveSingleTaskForTransferOutByAuthorizationCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForTransferOutByAuthorizationCodeRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeRequest) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeRequest) SetAuthorizationCode(v string) *SaveSingleTaskForTransferOutByAuthorizationCodeRequest {
	s.AuthorizationCode = &v
	return s
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeRequest) SetDomainName(v string) *SaveSingleTaskForTransferOutByAuthorizationCodeRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeRequest) Validate() error {
	return dara.Validate(s)
}
