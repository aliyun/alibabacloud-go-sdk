// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRouteMapRequest
	GetClientToken() *string
	SetDescription(v string) *CreateRouteMapRequest
	GetDescription() *string
	SetDirection(v string) *CreateRouteMapRequest
	GetDirection() *string
	SetDryRun(v bool) *CreateRouteMapRequest
	GetDryRun() *bool
	SetName(v string) *CreateRouteMapRequest
	GetName() *string
	SetPriority(v int32) *CreateRouteMapRequest
	GetPriority() *int32
	SetResourceGroupId(v string) *CreateRouteMapRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateRouteMapRequestTag) *CreateRouteMapRequest
	GetTag() []*CreateRouteMapRequestTag
}

type CreateRouteMapRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the parameter value on your client. Make sure that the token is unique among different requests. The client token only supports ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the route map.
	//
	// The description can be empty or 1 to 256 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// descname
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The direction of the route map. Valid values:
	//
	// - **Inbound**: inbound. The direction in which the instance route table receives routes from the associated network instance. This affects the routes that the current instance can learn from the associated network instance.
	//
	// - **Outbound**: outbound. The direction in which the instance route table sends routes to the associated network instance. This affects the routes that the current instance advertises to the associated network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Inbound
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// Specifies whether to perform a dry run on the request. The dry run includes checks on permissions and instance status. Valid values:
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, the route map is created.
	//
	// - **true**: sends a dry run request. The system only checks the request and does not create the route map. The check items include whether required parameters are specified and whether the request format is valid. If the request fails the check, the corresponding error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the route map.
	//
	// The name can be empty or 1 to 128 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Policy priority of the route map. Valid values: **1*	- to **100**. A smaller value indicates a higher priority.
	//
	// > Policy priority of each route map under an Alibaba Cloud account must be unique. When a route map is applied, the system starts matching route map rules from the route map with the smallest priority value based on policy priority of the route map rules. Therefore, when you specify policy priority of a route map, make sure that policy priority meets the expected matching order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm3unpnuw****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags at a time.
	Tag []*CreateRouteMapRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateRouteMapRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteMapRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRouteMapRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRouteMapRequest) GetDirection() *string {
	return s.Direction
}

func (s *CreateRouteMapRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateRouteMapRequest) GetName() *string {
	return s.Name
}

func (s *CreateRouteMapRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateRouteMapRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRouteMapRequest) GetTag() []*CreateRouteMapRequestTag {
	return s.Tag
}

func (s *CreateRouteMapRequest) SetClientToken(v string) *CreateRouteMapRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRouteMapRequest) SetDescription(v string) *CreateRouteMapRequest {
	s.Description = &v
	return s
}

func (s *CreateRouteMapRequest) SetDirection(v string) *CreateRouteMapRequest {
	s.Direction = &v
	return s
}

func (s *CreateRouteMapRequest) SetDryRun(v bool) *CreateRouteMapRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRouteMapRequest) SetName(v string) *CreateRouteMapRequest {
	s.Name = &v
	return s
}

func (s *CreateRouteMapRequest) SetPriority(v int32) *CreateRouteMapRequest {
	s.Priority = &v
	return s
}

func (s *CreateRouteMapRequest) SetResourceGroupId(v string) *CreateRouteMapRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRouteMapRequest) SetTag(v []*CreateRouteMapRequestTag) *CreateRouteMapRequest {
	s.Tag = v
	return s
}

func (s *CreateRouteMapRequest) Validate() error {
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

type CreateRouteMapRequestTag struct {
	// The tag key of the resource.
	//
	// The tag key cannot be an empty string. It can be up to 64 characters in length, and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// The tag value cannot be empty. It can be up to 128 characters in length, and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// Each tag key corresponds to a tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// tagtest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRouteMapRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapRequestTag) GoString() string {
	return s.String()
}

func (s *CreateRouteMapRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateRouteMapRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateRouteMapRequestTag) SetKey(v string) *CreateRouteMapRequestTag {
	s.Key = &v
	return s
}

func (s *CreateRouteMapRequestTag) SetValue(v string) *CreateRouteMapRequestTag {
	s.Value = &v
	return s
}

func (s *CreateRouteMapRequestTag) Validate() error {
	return dara.Validate(s)
}
