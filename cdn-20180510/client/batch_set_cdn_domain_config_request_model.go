// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetCdnDomainConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchSetCdnDomainConfigRequest
	GetDomainNames() *string
	SetFunctions(v string) *BatchSetCdnDomainConfigRequest
	GetFunctions() *string
	SetOwnerAccount(v string) *BatchSetCdnDomainConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BatchSetCdnDomainConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchSetCdnDomainConfigRequest
	GetSecurityToken() *string
}

type BatchSetCdnDomainConfigRequest struct {
	// The accelerated domain names. You can specify multiple accelerated domain names and separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The features that you want to configure. Format:
	//
	// 	- **functionName**: the name of the feature. This parameter is required. Separate multiple values with commas (,). For more information, see [Parameters for configuring features for domain names](https://help.aliyun.com/document_detail/388460.html).
	//
	// 	- **argName**: the feature parameter for **functionName**. This parameter is required. You can specify multiple feature parameters.
	//
	// 	- **argValue**: the parameter value that is specified for **functionName**. This parameter is required.
	//
	// 	- **parentid**: the rule condition ID. This parameter is optional. You can use the **condition*	- rule engine to create a rule condition. For information, see [Parameters for configuring features for domain names](https://help.aliyun.com/document_detail/388460.html). A rule condition can identify parameters that are included in requests and filter requests based on the identified parameters. After you create a rule condition, a [configid](https://help.aliyun.com/document_detail/388994.html) is generated. A configid can be used as parentId that is referenced by other features. This way, you can combine rule conditions and features for flexible configurations.
	//
	// If the **ParentId*	- parameter is \\*\\*-1\\*\\*, the existing rule conditions in the configurations are deleted.
	//
	// ```[{
	//
	//    "functionArgs": [{
	//
	//      "argName": "Parameter A",
	//
	//      "argValue": "Value of parameter A"
	//
	//     },
	//
	//   {
	//
	//     "argName": "Parameter B",
	//
	//     "argValue": "Value of parameter B"
	//
	//      }],
	//
	//  "functionName": "Feature name"
	//
	//  "parentId": Optional. parentId corresponds to configid of the referenced rule condition
	//
	// }]
	//
	// ```
	//
	// The following code provides a sample configuration if **parentId*	- is not used. In this example, the **origin_request_header*	- feature is used to add back-to-origin HTTP headers, and the rule condition whose configuration ID is **configid=222728944812032*	- is referenced.
	//
	// ```[{
	//
	//         "functionArgs": [{
	//
	//             "argName": "header_operation_type",
	//
	//             "argValue": "add"
	//
	//         }, {
	//
	//             "argName": "header_name",
	//
	//             "argValue": "Accept-Encoding"
	//
	//         }, {
	//
	//             "argName": "header_value",
	//
	//             "argValue": "gzip"
	//
	//         }, {
	//
	//             "argName": "duplicate",
	//
	//             "argValue": "off"
	//
	//         }],
	//
	//         "functionName": "origin_request_header"
	//
	// }]
	//
	// ```
	//
	// The following code shows a sample configuration if **parentId*	- is used. In this example, the **origin_request_header*	- feature is used to add back-to-origin HTTP headers, and the rule condition whose configuration ID is **222728944812032*	- is referenced.
	//
	// ```[{
	//
	//         "functionArgs": [{
	//
	//             "argName": "header_operation_type",
	//
	//             "argValue": "add"
	//
	//         }, {
	//
	//             "argName": "header_name",
	//
	//             "argValue": "Accept-Encoding"
	//
	//         }, {
	//
	//             "argName": "header_value",
	//
	//             "argValue": "gzip"
	//
	//         }, {
	//
	//             "argName": "duplicate",
	//
	//             "argValue": "off"
	//
	//         }],
	//
	//         "functionName": "origin_request_header",
	//
	//         "parentId": 222728944812032
	//
	// }]
	//
	// ```
	//
	// The following code provides a sample configuration that deletes the reference to **parentId*	- for a feature that uses **parentId**. This example shows how to delete the rule condition that has a configuration ID of **222728944812032*	- and is referenced when **origin_request_header*	- feature is used to add back-to-origin HTTP headers.
	//
	// ```[{
	//
	//         "functionArgs": [{
	//
	//             "argName": "header_operation_type",
	//
	//             "argValue": "add"
	//
	//         }, {
	//
	//             "argName": "header_name",
	//
	//             "argValue": "Accept-Encoding"
	//
	//         }, {
	//
	//             "argName": "header_value",
	//
	//             "argValue": "gzip"
	//
	//         }, {
	//
	//             "argName": "duplicate",
	//
	//             "argValue": "off"
	//
	//         }],
	//
	//         "functionName": "origin_request_header",
	//
	//         "parentId": -1
	//
	// }]
	//
	// ```
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"functionArgs": [{"argName": "key","argValue": "Content-Encoding"},{"argName": "value","argValue": "gzip"}],"functionName": "set_resp_header"} ]
	Functions     *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchSetCdnDomainConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSetCdnDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *BatchSetCdnDomainConfigRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchSetCdnDomainConfigRequest) GetFunctions() *string {
	return s.Functions
}

func (s *BatchSetCdnDomainConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BatchSetCdnDomainConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchSetCdnDomainConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchSetCdnDomainConfigRequest) SetDomainNames(v string) *BatchSetCdnDomainConfigRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetCdnDomainConfigRequest) SetFunctions(v string) *BatchSetCdnDomainConfigRequest {
	s.Functions = &v
	return s
}

func (s *BatchSetCdnDomainConfigRequest) SetOwnerAccount(v string) *BatchSetCdnDomainConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchSetCdnDomainConfigRequest) SetOwnerId(v int64) *BatchSetCdnDomainConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetCdnDomainConfigRequest) SetSecurityToken(v string) *BatchSetCdnDomainConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchSetCdnDomainConfigRequest) Validate() error {
	return dara.Validate(s)
}
