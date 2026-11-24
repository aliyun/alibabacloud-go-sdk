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
	SetVersion(v string) *CreateExpressConnectRouterRequest
	GetVersion() *string
}

type CreateExpressConnectRouterRequest struct {
	// The Autonomous System Number (ASN) of the Express Connect Router (ECR) instance. Valid values: 45104 (default), 64512 to 65534, and 4200000000 to 4294967294. The value 65025 is reserved by Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 45104
	AlibabaSideAsn *int64 `json:"AlibabaSideAsn,omitempty" xml:"AlibabaSideAsn,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the token from your client to ensure that the token is unique across different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the Express Connect Router (ECR) instance.
	//
	// > The description can be empty or 0 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether it is only a dry run for this request. Valid values:
	//
	// - **true**: Sends a check request without modifying service resources.
	//
	// - **false*	- (default): Sends a normal request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the Express Connect Router (ECR) instance.
	//
	// > The name can be empty or 0 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the Express Connect Router (ECR) instance belongs.
	//
	// example:
	//
	// rg-acfmvvajg5q****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	//
	// You can add up to 20 tags.
	Tag     []*CreateExpressConnectRouterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Version *string                                 `json:"Version,omitempty" xml:"Version,omitempty"`
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

func (s *CreateExpressConnectRouterRequest) GetVersion() *string {
	return s.Version
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

func (s *CreateExpressConnectRouterRequest) SetVersion(v string) *CreateExpressConnectRouterRequest {
	s.Version = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateExpressConnectRouterRequestTag struct {
	// The tag key of the resource. If you specify this parameter, it cannot be an empty string. The tag key can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the listener. Valid values of N: **1*	- to **20**. If specified, the value can be an empty string. The value can be up to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
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
