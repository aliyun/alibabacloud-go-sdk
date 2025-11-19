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
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// ListResult<InstanceSSL>
	Data *ModifySecurityIPGroupRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 840F51F7-9C01-538D-94F6-AE712905****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	GlobalSecurityIPGroupRel []*ModifySecurityIPGroupRelationResponseBodyDataGlobalSecurityIPGroupRel `json:"GlobalSecurityIPGroupRel,omitempty" xml:"GlobalSecurityIPGroupRel,omitempty" type:"Repeated"`
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
	// example:
	//
	// 192.168.0.1,192.168.100.0/24
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// example:
	//
	// saas_jump
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// example:
	//
	// g-v8kwereyd6u7kx****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
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
