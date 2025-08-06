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
	// The domain name for CDN. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The features to configure.
	//
	// 	- Set this parameter in the following format: `[{"functionArgs":[{"argName":"domain_name","argValue":"www.example.com"}],"functionName":"set_req_host_header"}]`.
	//
	// 	- Specific features, such as filetype_based_ttl_set, support more than one configuration record. To update one of the configuration records, use the configId field to specify the record. `[{"functionArgs":[{"argName":"file_type","argValue":"jpg"},{"argName":"ttl","argValue":"18"},{"argName":"weight","argValue":"30"}],"functionName":"filetype_based_ttl_set","configId":5068995}]`
	//
	// 	- For more information, see the **Feature description*	- section.
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
