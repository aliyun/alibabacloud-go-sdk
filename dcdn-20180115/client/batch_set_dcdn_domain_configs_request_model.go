// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDcdnDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchSetDcdnDomainConfigsRequest
	GetDomainNames() *string
	SetFunctions(v string) *BatchSetDcdnDomainConfigsRequest
	GetFunctions() *string
	SetOwnerAccount(v string) *BatchSetDcdnDomainConfigsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BatchSetDcdnDomainConfigsRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchSetDcdnDomainConfigsRequest
	GetSecurityToken() *string
}

type BatchSetDcdnDomainConfigsRequest struct {
	// The accelerated domain names. Specify multiple accelerated domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.org
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The features that you want to configure. Format:
	//
	// 	- **functionName**: The name of the feature. Separate multiple values with commas (,). For more information, see [A list of features](https://help.aliyun.com/document_detail/410622.html).
	//
	// 	- **argName**: The feature parameters for **functionName**.
	//
	// 	- **argValue**: The parameter values set for **functionName**.
	//
	// 	- **parentid**: the rule ID. This parameter is optional. You can use the **condition*	- rules engine to create a rule. For information, see [Feature settings for domain names](https://help.aliyun.com/document_detail/388460.html). A rule can identify parameters that are included in requests and filter requests based on the identified parameters. After you create a rule, a [configid](https://help.aliyun.com/document_detail/388994.html) is generated. A configid can be used as parentId that is referenced by other features. This way, you can combine rules and features for flexible configurations.
	//
	// If the **parentId*	- parameter is **-1**, the existing rules in the configurations are deleted.
	//
	// ````[
	//
	//  {
	//
	//    "functionArgs": [
	//
	//     {
	//
	//      "argName": "Parameter A",
	//
	//      "argValue": Value of parameter A"
	//
	//     },
	//
	//   {
	//
	//     "argName": "Parameter B",
	//
	//     "argValue": "Value of Parameter B"
	//
	//      }
	//
	//  ],
	//
	//  "functionName": "Feature name"
	//
	//     }
	//
	// ]```
	//
	// ````
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"functionArgs":[{"argName":"switch","argValue":"on"},{"argName":"region","argValue":"*"}],"functionName":"ipv6"}]
	Functions     *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchSetDcdnDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDcdnDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnDomainConfigsRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchSetDcdnDomainConfigsRequest) GetFunctions() *string {
	return s.Functions
}

func (s *BatchSetDcdnDomainConfigsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BatchSetDcdnDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchSetDcdnDomainConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchSetDcdnDomainConfigsRequest) SetDomainNames(v string) *BatchSetDcdnDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsRequest) SetFunctions(v string) *BatchSetDcdnDomainConfigsRequest {
	s.Functions = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsRequest) SetOwnerAccount(v string) *BatchSetDcdnDomainConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsRequest) SetOwnerId(v int64) *BatchSetDcdnDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsRequest) SetSecurityToken(v string) *BatchSetDcdnDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchSetDcdnDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
