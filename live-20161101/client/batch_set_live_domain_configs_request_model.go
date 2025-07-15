// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetLiveDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchSetLiveDomainConfigsRequest
	GetDomainNames() *string
	SetFunctions(v string) *BatchSetLiveDomainConfigsRequest
	GetFunctions() *string
	SetOwnerAccount(v string) *BatchSetLiveDomainConfigsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BatchSetLiveDomainConfigsRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchSetLiveDomainConfigsRequest
	GetSecurityToken() *string
}

type BatchSetLiveDomainConfigsRequest struct {
	// The domain names that you want to batch configure. Supported domain names include ingest domains, main streaming domains, and sub-streaming domains. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com,example.aliyundoc.com,example.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The list of features.
	//
	// Some features, such as `filetype_based_ttl_set`, support multiple configuration records. To update one of the configuration records, use `configId` to identify the record. For more information, see **Format of the Functions parameter*	- and **Features specified by the Functions parameter**.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"functionArgs":[{"argName":"file_type","argValue":"jpg"},{"argName":"ttl","argValue":"18"},{"argName":"weight","argValue":"30"}],"functionName":"filetype_based_ttl_set","configId":506***}]
	Functions     *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchSetLiveDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSetLiveDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchSetLiveDomainConfigsRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchSetLiveDomainConfigsRequest) GetFunctions() *string {
	return s.Functions
}

func (s *BatchSetLiveDomainConfigsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BatchSetLiveDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchSetLiveDomainConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchSetLiveDomainConfigsRequest) SetDomainNames(v string) *BatchSetLiveDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetLiveDomainConfigsRequest) SetFunctions(v string) *BatchSetLiveDomainConfigsRequest {
	s.Functions = &v
	return s
}

func (s *BatchSetLiveDomainConfigsRequest) SetOwnerAccount(v string) *BatchSetLiveDomainConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchSetLiveDomainConfigsRequest) SetOwnerId(v int64) *BatchSetLiveDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetLiveDomainConfigsRequest) SetSecurityToken(v string) *BatchSetLiveDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchSetLiveDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
