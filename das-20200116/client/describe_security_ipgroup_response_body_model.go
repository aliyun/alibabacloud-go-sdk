// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIPGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSecurityIPGroupResponseBody
	GetCode() *string
	SetData(v *DescribeSecurityIPGroupResponseBodyData) *DescribeSecurityIPGroupResponseBody
	GetData() *DescribeSecurityIPGroupResponseBodyData
	SetMessage(v string) *DescribeSecurityIPGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSecurityIPGroupResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSecurityIPGroupResponseBody
	GetSuccess() *string
}

type DescribeSecurityIPGroupResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// ListResult<InstanceSSL>
	Data *DescribeSecurityIPGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned for the request.
	//
	// > If the request is successful, **Successful*	- is returned. If the request fails, an exception message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CAC553F1-C669-53F1-A295-2CF050E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSecurityIPGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSecurityIPGroupResponseBody) GetData() *DescribeSecurityIPGroupResponseBodyData {
	return s.Data
}

func (s *DescribeSecurityIPGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSecurityIPGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityIPGroupResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSecurityIPGroupResponseBody) SetCode(v string) *DescribeSecurityIPGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBody) SetData(v *DescribeSecurityIPGroupResponseBodyData) *DescribeSecurityIPGroupResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSecurityIPGroupResponseBody) SetMessage(v string) *DescribeSecurityIPGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBody) SetRequestId(v string) *DescribeSecurityIPGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBody) SetSuccess(v string) *DescribeSecurityIPGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityIPGroupResponseBodyData struct {
	// The list of all cross-engine whitelist templates for the user in the specified region.
	SecurityIpGroups []*DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups `json:"SecurityIpGroups,omitempty" xml:"SecurityIpGroups,omitempty" type:"Repeated"`
}

func (s DescribeSecurityIPGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPGroupResponseBodyData) GetSecurityIpGroups() []*DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups {
	return s.SecurityIpGroups
}

func (s *DescribeSecurityIPGroupResponseBodyData) SetSecurityIpGroups(v []*DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) *DescribeSecurityIPGroupResponseBodyData {
	s.SecurityIpGroups = v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyData) Validate() error {
	if s.SecurityIpGroups != nil {
		for _, item := range s.SecurityIpGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups struct {
	// The list of database instances bound to the cross-engine whitelist template.
	DbInstances []*string `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Repeated"`
	// The instance information for each product bound to the template.
	EngineInfoList []*DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList `json:"EngineInfoList,omitempty" xml:"EngineInfoList,omitempty" type:"Repeated"`
	// The ECS security group ID. This field is invalid and contains redundant data that will be deprecated.
	//
	// example:
	//
	// null
	GEcsSgIdList *string `json:"GEcsSgIdList,omitempty" xml:"GEcsSgIdList,omitempty"`
	// The IP addresses in the whitelist template.
	//
	// > Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 192.168.1.28/32
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The IP whitelist template name. The name must meet the following requirements:
	//
	// - Contains only lowercase letters, digits, and underscores (_).
	//
	// - Starts with a letter and ends with a letter or digit.
	//
	// - Contains 2 to 120 characters in length.
	//
	// example:
	//
	// test_123
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The IP whitelist template ID.
	//
	// example:
	//
	// g-1no2rzybnqcv****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP type.
	//
	// example:
	//
	// ipv4
	SecurityIpType *string `json:"SecurityIpType,omitempty" xml:"SecurityIpType,omitempty"`
	// The account ID. You can obtain the ID of the **logon account*	- on the **Security Settings*	- page in **Account Management*	- of the Alibaba Cloud console.
	//
	// example:
	//
	// 160-79abe3f4****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 641***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The network type of the whitelist.
	//
	// example:
	//
	// mix
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) GetDbInstances() []*string {
	return s.DbInstances
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) GetEngineInfoList() []*DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList {
	return s.EngineInfoList
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) GetGEcsSgIdList() *string {
	return s.GEcsSgIdList
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) GetGIpList() *string {
	return s.GIpList
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) GetSecurityIpType() *string {
	return s.SecurityIpType
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) GetUid() *string {
	return s.Uid
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) GetUserId() *string {
	return s.UserId
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) GetWhitelistNetType() *string {
	return s.WhitelistNetType
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) SetDbInstances(v []*string) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups {
	s.DbInstances = v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) SetEngineInfoList(v []*DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups {
	s.EngineInfoList = v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) SetGEcsSgIdList(v string) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups {
	s.GEcsSgIdList = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) SetGIpList(v string) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups {
	s.GIpList = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) SetGlobalIgName(v string) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups {
	s.GlobalIgName = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) SetGlobalSecurityGroupId(v string) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) SetRegionId(v string) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) SetSecurityIpType(v string) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups {
	s.SecurityIpType = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) SetUid(v string) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups {
	s.Uid = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) SetUserId(v string) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups {
	s.UserId = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) SetWhitelistNetType(v string) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups {
	s.WhitelistNetType = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroups) Validate() error {
	if s.EngineInfoList != nil {
		for _, item := range s.EngineInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList struct {
	// The database engine type of the target instance.
	//
	// example:
	//
	// PolarDBMySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The list of database instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The number of instances under the current logon account.
	//
	// example:
	//
	// 10
	InstanceNum *int32 `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
}

func (s DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList) GetInstanceNum() *int32 {
	return s.InstanceNum
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList) SetEngineName(v string) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList {
	s.EngineName = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList) SetInstanceIds(v []*string) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList {
	s.InstanceIds = v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList) SetInstanceNum(v int32) *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList {
	s.InstanceNum = &v
	return s
}

func (s *DescribeSecurityIPGroupResponseBodyDataSecurityIpGroupsEngineInfoList) Validate() error {
	return dara.Validate(s)
}
