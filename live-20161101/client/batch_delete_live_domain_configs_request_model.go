// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteLiveDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchDeleteLiveDomainConfigsRequest
	GetDomainNames() *string
	SetFunctionNames(v string) *BatchDeleteLiveDomainConfigsRequest
	GetFunctionNames() *string
	SetOwnerAccount(v string) *BatchDeleteLiveDomainConfigsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BatchDeleteLiveDomainConfigsRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchDeleteLiveDomainConfigsRequest
	GetSecurityToken() *string
}

type BatchDeleteLiveDomainConfigsRequest struct {
	// The ingest domain or streaming domain. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com,example.aliyundoc.com,example.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The names of the features. Separate multiple features with commas (,). For more information, see **Features specified by the Functions parameter**.
	//
	// This parameter is required.
	//
	// example:
	//
	// referer_white_list_set,ip_black_list_set
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchDeleteLiveDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteLiveDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteLiveDomainConfigsRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchDeleteLiveDomainConfigsRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *BatchDeleteLiveDomainConfigsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BatchDeleteLiveDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchDeleteLiveDomainConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchDeleteLiveDomainConfigsRequest) SetDomainNames(v string) *BatchDeleteLiveDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchDeleteLiveDomainConfigsRequest) SetFunctionNames(v string) *BatchDeleteLiveDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *BatchDeleteLiveDomainConfigsRequest) SetOwnerAccount(v string) *BatchDeleteLiveDomainConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchDeleteLiveDomainConfigsRequest) SetOwnerId(v int64) *BatchDeleteLiveDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchDeleteLiveDomainConfigsRequest) SetSecurityToken(v string) *BatchDeleteLiveDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchDeleteLiveDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
