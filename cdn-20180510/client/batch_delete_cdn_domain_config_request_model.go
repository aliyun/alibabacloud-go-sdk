// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteCdnDomainConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchDeleteCdnDomainConfigRequest
	GetDomainNames() *string
	SetFunctionNames(v string) *BatchDeleteCdnDomainConfigRequest
	GetFunctionNames() *string
	SetOwnerAccount(v string) *BatchDeleteCdnDomainConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BatchDeleteCdnDomainConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchDeleteCdnDomainConfigRequest
	GetSecurityToken() *string
}

type BatchDeleteCdnDomainConfigRequest struct {
	// The accelerated domain names whose configurations you want to delete. Separate multiple accelerated domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.org
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The names of the features that you want to delete. Separate multiple feature names with commas (,). For more information about feature names, see [Parameters for configuring features for domain names](https://help.aliyun.com/document_detail/388460.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// referer_white_list_set,https_force
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchDeleteCdnDomainConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteCdnDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteCdnDomainConfigRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchDeleteCdnDomainConfigRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *BatchDeleteCdnDomainConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BatchDeleteCdnDomainConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchDeleteCdnDomainConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchDeleteCdnDomainConfigRequest) SetDomainNames(v string) *BatchDeleteCdnDomainConfigRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchDeleteCdnDomainConfigRequest) SetFunctionNames(v string) *BatchDeleteCdnDomainConfigRequest {
	s.FunctionNames = &v
	return s
}

func (s *BatchDeleteCdnDomainConfigRequest) SetOwnerAccount(v string) *BatchDeleteCdnDomainConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchDeleteCdnDomainConfigRequest) SetOwnerId(v int64) *BatchDeleteCdnDomainConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchDeleteCdnDomainConfigRequest) SetSecurityToken(v string) *BatchDeleteCdnDomainConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchDeleteCdnDomainConfigRequest) Validate() error {
	return dara.Validate(s)
}
