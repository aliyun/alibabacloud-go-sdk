// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlibabaSideAsn(v int64) *CreateExpressConnectRouterRequest
	GetAlibabaSideAsn() *int64
	SetClientToken(v string) *CreateExpressConnectRouterRequest
	GetClientToken() *string
	SetDescription(v string) *CreateExpressConnectRouterRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateExpressConnectRouterRequest
	GetDryRun() *bool
	SetName(v string) *CreateExpressConnectRouterRequest
	GetName() *string
	SetResourceGroupId(v string) *CreateExpressConnectRouterRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateExpressConnectRouterRequestTag) *CreateExpressConnectRouterRequest
	GetTag() []*CreateExpressConnectRouterRequestTag
}

type CreateExpressConnectRouterRequest struct {
	// The autonomous system number (ASN) of the ECR. Valid values: 45104, 64512 to 65534, and 4200000000 to 4294967294. Default value: 45104. The value 65025 is reserved by Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 45104
	AlibabaSideAsn *int64 `json:"AlibabaSideAsn,omitempty" xml:"AlibabaSideAsn,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the ECR.
	//
	// >  The description can be empty or 0 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the ECR.
	//
	// >  The name must be 0 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the ECR belongs.
	//
	// example:
	//
	// rg-acfmvvajg5q****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The information about the tags.
	//
	// You can specify at most 20 tags in each call.
	Tag []*CreateExpressConnectRouterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateExpressConnectRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterRequest) GetAlibabaSideAsn() *int64 {
	return s.AlibabaSideAsn
}

func (s *CreateExpressConnectRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateExpressConnectRouterRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateExpressConnectRouterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateExpressConnectRouterRequest) GetName() *string {
	return s.Name
}

func (s *CreateExpressConnectRouterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateExpressConnectRouterRequest) GetTag() []*CreateExpressConnectRouterRequestTag {
	return s.Tag
}

func (s *CreateExpressConnectRouterRequest) SetAlibabaSideAsn(v int64) *CreateExpressConnectRouterRequest {
	s.AlibabaSideAsn = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetClientToken(v string) *CreateExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetDescription(v string) *CreateExpressConnectRouterRequest {
	s.Description = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetDryRun(v bool) *CreateExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetName(v string) *CreateExpressConnectRouterRequest {
	s.Name = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetResourceGroupId(v string) *CreateExpressConnectRouterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetTag(v []*CreateExpressConnectRouterRequestTag) *CreateExpressConnectRouterRequest {
	s.Tag = v
	return s
}

func (s *CreateExpressConnectRouterRequest) Validate() error {
	return dara.Validate(s)
}

type CreateExpressConnectRouterRequestTag struct {
	// The tag key. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateExpressConnectRouterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectRouterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateExpressConnectRouterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateExpressConnectRouterRequestTag) SetKey(v string) *CreateExpressConnectRouterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateExpressConnectRouterRequestTag) SetValue(v string) *CreateExpressConnectRouterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateExpressConnectRouterRequestTag) Validate() error {
	return dara.Validate(s)
}
