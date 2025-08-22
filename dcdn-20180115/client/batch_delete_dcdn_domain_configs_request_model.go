// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchDeleteDcdnDomainConfigsRequest
	GetDomainNames() *string
	SetFunctionNames(v string) *BatchDeleteDcdnDomainConfigsRequest
	GetFunctionNames() *string
	SetOwnerAccount(v string) *BatchDeleteDcdnDomainConfigsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BatchDeleteDcdnDomainConfigsRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchDeleteDcdnDomainConfigsRequest
	GetSecurityToken() *string
}

type BatchDeleteDcdnDomainConfigsRequest struct {
	// The accelerated domain names whose configurations you want to delete. Separate multiple accelerated domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The names of the features that you want to delete. Separate multiple feature names with commas (,). For more information about feature names, see [Feature settings for a domain name](https://help.aliyun.com/document_detail/410622.html).
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

func (s BatchDeleteDcdnDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnDomainConfigsRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchDeleteDcdnDomainConfigsRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *BatchDeleteDcdnDomainConfigsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BatchDeleteDcdnDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchDeleteDcdnDomainConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchDeleteDcdnDomainConfigsRequest) SetDomainNames(v string) *BatchDeleteDcdnDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsRequest) SetFunctionNames(v string) *BatchDeleteDcdnDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsRequest) SetOwnerAccount(v string) *BatchDeleteDcdnDomainConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsRequest) SetOwnerId(v int64) *BatchDeleteDcdnDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsRequest) SetSecurityToken(v string) *BatchDeleteDcdnDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchDeleteDcdnDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
