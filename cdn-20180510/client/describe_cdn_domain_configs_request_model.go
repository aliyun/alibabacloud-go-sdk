// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DescribeCdnDomainConfigsRequest
	GetConfigId() *string
	SetDomainName(v string) *DescribeCdnDomainConfigsRequest
	GetDomainName() *string
	SetFunctionNames(v string) *DescribeCdnDomainConfigsRequest
	GetFunctionNames() *string
	SetOwnerId(v int64) *DescribeCdnDomainConfigsRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnDomainConfigsRequest
	GetSecurityToken() *string
}

type DescribeCdnDomainConfigsRequest struct {
	// The ID of the configuration. For more information about ConfigId, see [Usage notes on ConfigId](https://help.aliyun.com/document_detail/388994.html).
	//
	// example:
	//
	// 6295
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The names of the features. Separate multiple feature names with commas (,). For more information, see [Parameters for configuring features for domain names](https://help.aliyun.com/document_detail/388460.html).
	//
	// example:
	//
	// aliauth
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCdnDomainConfigsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDomainConfigsRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *DescribeCdnDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnDomainConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnDomainConfigsRequest) SetConfigId(v string) *DescribeCdnDomainConfigsRequest {
	s.ConfigId = &v
	return s
}

func (s *DescribeCdnDomainConfigsRequest) SetDomainName(v string) *DescribeCdnDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainConfigsRequest) SetFunctionNames(v string) *DescribeCdnDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeCdnDomainConfigsRequest) SetOwnerId(v int64) *DescribeCdnDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnDomainConfigsRequest) SetSecurityToken(v string) *DescribeCdnDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
