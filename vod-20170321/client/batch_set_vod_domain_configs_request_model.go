// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetVodDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchSetVodDomainConfigsRequest
	GetDomainNames() *string
	SetFunctions(v string) *BatchSetVodDomainConfigsRequest
	GetFunctions() *string
	SetOwnerAccount(v string) *BatchSetVodDomainConfigsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BatchSetVodDomainConfigsRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchSetVodDomainConfigsRequest
	GetSecurityToken() *string
}

type BatchSetVodDomainConfigsRequest struct {
	// The accelerated domain names for ApsaraVideo VOD. Separate multiple domain names with commas (,). You can configure up to 50 domain names at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The list of features.
	//
	// - functionName (feature name, required): For the features that can be configured and their feature name parameters, see [Domain name configuration features](https://help.aliyun.com/document_detail/2411639.html).
	//
	// - argName (parameter name, required): The configuration items of functionName. You can configure multiple configuration items.
	//
	// - argValue (parameter value, required): The values of the configuration items of functionName.
	//
	// For detailed information about the features that can be configured for accelerated domain names, including feature names and parameter names, see [Domain name configuration features](https://help.aliyun.com/document_detail/2411639.html).
	//
	// > Some features, such as filetype_based_ttl_set (file expiration time), support multiple configuration rules. To update a specific configuration rule, specify the configId of that rule. Example:
	//
	// `[{"functionArgs":[{"argName":"file_type","argValue":"jpg"},{"argName":"ttl","argValue":"18"},{"argName":"weight","argValue":"30"}],"functionName":"filetype_based_ttl_set","configId":5068995}]`
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"functionArgs":[{"argName":"domain_name","argValue":"www.example.com"}],"functionName":"set_req_host_header"}]
	Functions     *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchSetVodDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSetVodDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchSetVodDomainConfigsRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchSetVodDomainConfigsRequest) GetFunctions() *string {
	return s.Functions
}

func (s *BatchSetVodDomainConfigsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BatchSetVodDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchSetVodDomainConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchSetVodDomainConfigsRequest) SetDomainNames(v string) *BatchSetVodDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetVodDomainConfigsRequest) SetFunctions(v string) *BatchSetVodDomainConfigsRequest {
	s.Functions = &v
	return s
}

func (s *BatchSetVodDomainConfigsRequest) SetOwnerAccount(v string) *BatchSetVodDomainConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchSetVodDomainConfigsRequest) SetOwnerId(v int64) *BatchSetVodDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetVodDomainConfigsRequest) SetSecurityToken(v string) *BatchSetVodDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchSetVodDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
