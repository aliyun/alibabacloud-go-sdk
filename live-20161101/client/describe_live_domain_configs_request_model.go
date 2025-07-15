// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainConfigsRequest
	GetDomainName() *string
	SetFunctionNames(v string) *DescribeLiveDomainConfigsRequest
	GetFunctionNames() *string
	SetOwnerId(v int64) *DescribeLiveDomainConfigsRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeLiveDomainConfigsRequest
	GetSecurityToken() *string
}

type DescribeLiveDomainConfigsRequest struct {
	// The ingest domain or streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The names of the features. Separate multiple features with commas (,). For more information, see the **Features specified by the Functions parameter*	- section in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// set_req_host_header,set_hashkey_args
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeLiveDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainConfigsRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *DescribeLiveDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveDomainConfigsRequest) SetDomainName(v string) *DescribeLiveDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainConfigsRequest) SetFunctionNames(v string) *DescribeLiveDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeLiveDomainConfigsRequest) SetOwnerId(v int64) *DescribeLiveDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainConfigsRequest) SetSecurityToken(v string) *DescribeLiveDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
