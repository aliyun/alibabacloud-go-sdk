// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListHostGroupsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListHostGroupsResponseBody
	GetErrorMessage() *string
	SetHostGroups(v []*ListHostGroupsResponseBodyHostGroups) *ListHostGroupsResponseBody
	GetHostGroups() []*ListHostGroupsResponseBodyHostGroups
	SetNextToken(v string) *ListHostGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListHostGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHostGroupsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListHostGroupsResponseBody
	GetTotalCount() *int64
}

type ListHostGroupsResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string                                 `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HostGroups   []*ListHostGroupsResponseBodyHostGroups `json:"hostGroups,omitempty" xml:"hostGroups,omitempty" type:"Repeated"`
	// example:
	//
	// asassasassa
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 50
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListHostGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListHostGroupsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListHostGroupsResponseBody) GetHostGroups() []*ListHostGroupsResponseBodyHostGroups {
	return s.HostGroups
}

func (s *ListHostGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHostGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHostGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListHostGroupsResponseBody) SetErrorCode(v string) *ListHostGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetErrorMessage(v string) *ListHostGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetHostGroups(v []*ListHostGroupsResponseBodyHostGroups) *ListHostGroupsResponseBody {
	s.HostGroups = v
	return s
}

func (s *ListHostGroupsResponseBody) SetNextToken(v string) *ListHostGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetRequestId(v string) *ListHostGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetSuccess(v bool) *ListHostGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetTotalCount(v int64) *ListHostGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostGroupsResponseBody) Validate() error {
	if s.HostGroups != nil {
		for _, item := range s.HostGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHostGroupsResponseBodyHostGroups struct {
	// example:
	//
	// cn-beijing
	AliyunRegion *string `json:"aliyunRegion,omitempty" xml:"aliyunRegion,omitempty"`
	// example:
	//
	// 1586863220000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 2222222222222
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// example:
	//
	// 主机组
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// ecs
	EcsLabelKey *string `json:"ecsLabelKey,omitempty" xml:"ecsLabelKey,omitempty"`
	// example:
	//
	// value
	EcsLabelValue *string `json:"ecsLabelValue,omitempty" xml:"ecsLabelValue,omitempty"`
	// example:
	//
	// ECS_ALIYUN
	EcsType *string `json:"ecsType,omitempty" xml:"ecsType,omitempty"`
	// example:
	//
	// 3
	HostNum *int64 `json:"hostNum,omitempty" xml:"hostNum,omitempty"`
	// 323232
	//
	// example:
	//
	// 部署组Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 211111111
	ModifierAccountId *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	// example:
	//
	// 部署组
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1212122
	ServiceConnectionId *int64 `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
	// example:
	//
	// ECS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1586863220000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListHostGroupsResponseBodyHostGroups) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsResponseBodyHostGroups) GoString() string {
	return s.String()
}

func (s *ListHostGroupsResponseBodyHostGroups) GetAliyunRegion() *string {
	return s.AliyunRegion
}

func (s *ListHostGroupsResponseBodyHostGroups) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListHostGroupsResponseBodyHostGroups) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *ListHostGroupsResponseBodyHostGroups) GetDescription() *string {
	return s.Description
}

func (s *ListHostGroupsResponseBodyHostGroups) GetEcsLabelKey() *string {
	return s.EcsLabelKey
}

func (s *ListHostGroupsResponseBodyHostGroups) GetEcsLabelValue() *string {
	return s.EcsLabelValue
}

func (s *ListHostGroupsResponseBodyHostGroups) GetEcsType() *string {
	return s.EcsType
}

func (s *ListHostGroupsResponseBodyHostGroups) GetHostNum() *int64 {
	return s.HostNum
}

func (s *ListHostGroupsResponseBodyHostGroups) GetId() *int64 {
	return s.Id
}

func (s *ListHostGroupsResponseBodyHostGroups) GetModifierAccountId() *string {
	return s.ModifierAccountId
}

func (s *ListHostGroupsResponseBodyHostGroups) GetName() *string {
	return s.Name
}

func (s *ListHostGroupsResponseBodyHostGroups) GetServiceConnectionId() *int64 {
	return s.ServiceConnectionId
}

func (s *ListHostGroupsResponseBodyHostGroups) GetType() *string {
	return s.Type
}

func (s *ListHostGroupsResponseBodyHostGroups) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListHostGroupsResponseBodyHostGroups) SetAliyunRegion(v string) *ListHostGroupsResponseBodyHostGroups {
	s.AliyunRegion = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetCreateTime(v int64) *ListHostGroupsResponseBodyHostGroups {
	s.CreateTime = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetCreatorAccountId(v string) *ListHostGroupsResponseBodyHostGroups {
	s.CreatorAccountId = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetDescription(v string) *ListHostGroupsResponseBodyHostGroups {
	s.Description = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetEcsLabelKey(v string) *ListHostGroupsResponseBodyHostGroups {
	s.EcsLabelKey = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetEcsLabelValue(v string) *ListHostGroupsResponseBodyHostGroups {
	s.EcsLabelValue = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetEcsType(v string) *ListHostGroupsResponseBodyHostGroups {
	s.EcsType = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetHostNum(v int64) *ListHostGroupsResponseBodyHostGroups {
	s.HostNum = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetId(v int64) *ListHostGroupsResponseBodyHostGroups {
	s.Id = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetModifierAccountId(v string) *ListHostGroupsResponseBodyHostGroups {
	s.ModifierAccountId = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetName(v string) *ListHostGroupsResponseBodyHostGroups {
	s.Name = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetServiceConnectionId(v int64) *ListHostGroupsResponseBodyHostGroups {
	s.ServiceConnectionId = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetType(v string) *ListHostGroupsResponseBodyHostGroups {
	s.Type = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetUpdateTime(v int64) *ListHostGroupsResponseBodyHostGroups {
	s.UpdateTime = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) Validate() error {
	return dara.Validate(s)
}
