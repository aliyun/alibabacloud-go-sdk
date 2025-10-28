// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDefineRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListUserDefineRegionResponseBody
	GetCode() *int32
	SetMessage(v string) *ListUserDefineRegionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListUserDefineRegionResponseBody
	GetRequestId() *string
	SetUserDefineRegionList(v *ListUserDefineRegionResponseBodyUserDefineRegionList) *ListUserDefineRegionResponseBody
	GetUserDefineRegionList() *ListUserDefineRegionResponseBodyUserDefineRegionList
}

type ListUserDefineRegionResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// b197-40ab-9155-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The namespaces.
	UserDefineRegionList *ListUserDefineRegionResponseBodyUserDefineRegionList `json:"UserDefineRegionList,omitempty" xml:"UserDefineRegionList,omitempty" type:"Struct"`
}

func (s ListUserDefineRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefineRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDefineRegionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListUserDefineRegionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserDefineRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserDefineRegionResponseBody) GetUserDefineRegionList() *ListUserDefineRegionResponseBodyUserDefineRegionList {
	return s.UserDefineRegionList
}

func (s *ListUserDefineRegionResponseBody) SetCode(v int32) *ListUserDefineRegionResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserDefineRegionResponseBody) SetMessage(v string) *ListUserDefineRegionResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserDefineRegionResponseBody) SetRequestId(v string) *ListUserDefineRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDefineRegionResponseBody) SetUserDefineRegionList(v *ListUserDefineRegionResponseBodyUserDefineRegionList) *ListUserDefineRegionResponseBody {
	s.UserDefineRegionList = v
	return s
}

func (s *ListUserDefineRegionResponseBody) Validate() error {
	if s.UserDefineRegionList != nil {
		if err := s.UserDefineRegionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserDefineRegionResponseBodyUserDefineRegionList struct {
	UserDefineRegionEntity []*ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity `json:"UserDefineRegionEntity,omitempty" xml:"UserDefineRegionEntity,omitempty" type:"Repeated"`
}

func (s ListUserDefineRegionResponseBodyUserDefineRegionList) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefineRegionResponseBodyUserDefineRegionList) GoString() string {
	return s.String()
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionList) GetUserDefineRegionEntity() []*ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	return s.UserDefineRegionEntity
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionList) SetUserDefineRegionEntity(v []*ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) *ListUserDefineRegionResponseBodyUserDefineRegionList {
	s.UserDefineRegionEntity = v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionList) Validate() error {
	if s.UserDefineRegionEntity != nil {
		for _, item := range s.UserDefineRegionEntity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity struct {
	// The ID of the region to which the namespace belongs.
	//
	// example:
	//
	// cn-shenzhen
	BelongRegion *string `json:"BelongRegion,omitempty" xml:"BelongRegion,omitempty"`
	// Indicates whether remote debugging is allowed.
	//
	// example:
	//
	// false
	DebugEnable *bool `json:"DebugEnable,omitempty" xml:"DebugEnable,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// betaappManager
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identifier of the namespace.
	//
	// example:
	//
	// 1330
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the registry. Valid values:
	//
	// 	- default: shared service registry of Enterprise Distributed Application Service (EDAS)
	//
	// 	- exclusive_mse: Microservices Engine (MSE) Nacos registry
	//
	// example:
	//
	// default: EDAS
	MseInstanceId *string `json:"MseInstanceId,omitempty" xml:"MseInstanceId,omitempty"`
	// The ID of the namespace.
	//
	// > The ID cannot be changed after the namespace is created. The ID is in the `Physical region ID:Logical region identifier` format .
	//
	// example:
	//
	// cn-shenzhen:betaappManager
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// betaappManager
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The ID of the MSE instance.
	//
	// example:
	//
	// mse_prepaid_public_cn-tl32n******
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The ID of the Alibaba Cloud account to which the namespace belongs.
	//
	// example:
	//
	// edas_****_test@aliyun-****.com
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) GoString() string {
	return s.String()
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) GetBelongRegion() *string {
	return s.BelongRegion
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) GetDebugEnable() *bool {
	return s.DebugEnable
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) GetDescription() *string {
	return s.Description
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) GetId() *int64 {
	return s.Id
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) GetMseInstanceId() *string {
	return s.MseInstanceId
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) GetRegionName() *string {
	return s.RegionName
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) GetRegistryType() *string {
	return s.RegistryType
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) GetUserId() *string {
	return s.UserId
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetBelongRegion(v string) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.BelongRegion = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetDebugEnable(v bool) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.DebugEnable = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetDescription(v string) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.Description = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetId(v int64) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.Id = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetMseInstanceId(v string) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.MseInstanceId = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetRegionId(v string) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.RegionId = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetRegionName(v string) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.RegionName = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetRegistryType(v string) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.RegistryType = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetUserId(v string) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.UserId = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) Validate() error {
	return dara.Validate(s)
}
