// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertOrUpdateRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InsertOrUpdateRegionResponseBody
	GetCode() *int32
	SetMessage(v string) *InsertOrUpdateRegionResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertOrUpdateRegionResponseBody
	GetRequestId() *string
	SetUserDefineRegionEntity(v *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) *InsertOrUpdateRegionResponseBody
	GetUserDefineRegionEntity() *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity
}

type InsertOrUpdateRegionResponseBody struct {
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
	// 37793352-C568-4A5A-BF69-4DC853******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the custom namespace.
	UserDefineRegionEntity *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity `json:"UserDefineRegionEntity,omitempty" xml:"UserDefineRegionEntity,omitempty" type:"Struct"`
}

func (s InsertOrUpdateRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertOrUpdateRegionResponseBody) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateRegionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InsertOrUpdateRegionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertOrUpdateRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertOrUpdateRegionResponseBody) GetUserDefineRegionEntity() *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	return s.UserDefineRegionEntity
}

func (s *InsertOrUpdateRegionResponseBody) SetCode(v int32) *InsertOrUpdateRegionResponseBody {
	s.Code = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBody) SetMessage(v string) *InsertOrUpdateRegionResponseBody {
	s.Message = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBody) SetRequestId(v string) *InsertOrUpdateRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBody) SetUserDefineRegionEntity(v *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) *InsertOrUpdateRegionResponseBody {
	s.UserDefineRegionEntity = v
	return s
}

func (s *InsertOrUpdateRegionResponseBody) Validate() error {
	if s.UserDefineRegionEntity != nil {
		if err := s.UserDefineRegionEntity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertOrUpdateRegionResponseBodyUserDefineRegionEntity struct {
	// The ID of the region to which the namespace belongs.
	//
	// example:
	//
	// cn-beijing
	BelongRegion *string `json:"BelongRegion,omitempty" xml:"BelongRegion,omitempty"`
	// Indicates whether remote debugging is enabled. Valid values:
	//
	// 	- true: Remote debugging is enabled.
	//
	// 	- false: Remote debugging is disabled.
	//
	// example:
	//
	// false
	DebugEnable *bool `json:"DebugEnable,omitempty" xml:"DebugEnable,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// Logical region
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the namespace is created or modified. If this parameter is left empty or 0 is returned, the namespace is created. Otherwise, the namespace is modified.
	//
	// example:
	//
	// 15160
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the namespace.
	//
	// 	- The ID of a custom namespace is in the `region ID:namespace identifier` format. Example: cn-beijing:tdy218.
	//
	// 	- The ID of the default namespace is in the `region ID` format. Example: cn-beijing.
	//
	// example:
	//
	// cn-beijing:test
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test_region
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The ID of the Alibaba Cloud account to which the custom namespace belongs.
	//
	// example:
	//
	// edas_****_test@aliyun-****.com
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) String() string {
	return dara.Prettify(s)
}

func (s InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) GetBelongRegion() *string {
	return s.BelongRegion
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) GetDebugEnable() *bool {
	return s.DebugEnable
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) GetDescription() *string {
	return s.Description
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) GetId() *int64 {
	return s.Id
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) GetRegionId() *string {
	return s.RegionId
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) GetRegionName() *string {
	return s.RegionName
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) GetUserId() *string {
	return s.UserId
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetBelongRegion(v string) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.BelongRegion = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetDebugEnable(v bool) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.DebugEnable = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetDescription(v string) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.Description = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetId(v int64) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.Id = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetRegionId(v string) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.RegionId = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetRegionName(v string) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.RegionName = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetUserId(v string) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.UserId = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) Validate() error {
	return dara.Validate(s)
}
