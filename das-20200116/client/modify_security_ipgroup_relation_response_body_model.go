// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIPGroupRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifySecurityIPGroupRelationResponseBody
	GetCode() *string
	SetData(v *ModifySecurityIPGroupRelationResponseBodyData) *ModifySecurityIPGroupRelationResponseBody
	GetData() *ModifySecurityIPGroupRelationResponseBodyData
	SetMessage(v string) *ModifySecurityIPGroupRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySecurityIPGroupRelationResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifySecurityIPGroupRelationResponseBody
	GetSuccess() *string
}

type ModifySecurityIPGroupRelationResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ModifySecurityIPGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// > If the request is successful, Successful is returned. If the request fails, an error message, such as an error code, is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 840F51F7-9C01-538D-94F6-AE712905****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySecurityIPGroupRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPGroupRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifySecurityIPGroupRelationResponseBody) GetData() *ModifySecurityIPGroupRelationResponseBodyData {
	return s.Data
}

func (s *ModifySecurityIPGroupRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySecurityIPGroupRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityIPGroupRelationResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifySecurityIPGroupRelationResponseBody) SetCode(v string) *ModifySecurityIPGroupRelationResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySecurityIPGroupRelationResponseBody) SetData(v *ModifySecurityIPGroupRelationResponseBodyData) *ModifySecurityIPGroupRelationResponseBody {
	s.Data = v
	return s
}

func (s *ModifySecurityIPGroupRelationResponseBody) SetMessage(v string) *ModifySecurityIPGroupRelationResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySecurityIPGroupRelationResponseBody) SetRequestId(v string) *ModifySecurityIPGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityIPGroupRelationResponseBody) SetSuccess(v string) *ModifySecurityIPGroupRelationResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySecurityIPGroupRelationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifySecurityIPGroupRelationResponseBodyData struct {
	// The information about the cross-product IP address whitelist template that is attached to the instance.
	GlobalSecurityIPGroupRel []*ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel `json:"GlobalSecurityIPGroupRel,omitempty" xml:"GlobalSecurityIPGroupRel,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifySecurityIPGroupRelationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPGroupRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPGroupRelationResponseBodyData) GetGlobalSecurityIPGroupRel() []*ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel {
	return s.GlobalSecurityIPGroupRel
}

func (s *ModifySecurityIPGroupRelationResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifySecurityIPGroupRelationResponseBodyData) SetGlobalSecurityIPGroupRel(v []*ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) *ModifySecurityIPGroupRelationResponseBodyData {
	s.GlobalSecurityIPGroupRel = v
	return s
}

func (s *ModifySecurityIPGroupRelationResponseBodyData) SetInstanceId(v string) *ModifySecurityIPGroupRelationResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ModifySecurityIPGroupRelationResponseBodyData) Validate() error {
	if s.GlobalSecurityIPGroupRel != nil {
		for _, item := range s.GlobalSecurityIPGroupRel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel struct {
	// The IP addresses in the whitelist template.
	//
	// > Separate multiple IP addresses with commas (,). A maximum of 1,000 IP addresses or CIDR blocks can be added to all IP address whitelists.
	//
	// example:
	//
	// 192.168.0.1,192.168.100.0/24
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the IP address whitelist template. The name must meet the following requirements:
	//
	// - It must consist of lowercase letters, digits, and underscores (_).
	//
	// - It must start with a letter and end with a letter or a digit.
	//
	// - It must be 2 to 120 characters in length.
	//
	// example:
	//
	// saas_jump
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The ID of the IP address whitelist template.
	//
	// example:
	//
	// g-v8kwereyd6u7kx****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) GetGIpList() *string {
	return s.GIpList
}

func (s *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) SetGIpList(v string) *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel {
	s.GIpList = &v
	return s
}

func (s *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) SetGlobalIgName(v string) *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel {
	s.GlobalIgName = &v
	return s
}

func (s *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) SetGlobalSecurityGroupId(v string) *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) SetRegionId(v string) *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel) Validate() error {
	return dara.Validate(s)
}
