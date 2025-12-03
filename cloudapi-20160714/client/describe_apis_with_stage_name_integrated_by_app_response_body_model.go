// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisWithStageNameIntegratedByAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppApiRelationInfos(v *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfos) *DescribeApisWithStageNameIntegratedByAppResponseBody
	GetAppApiRelationInfos() *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfos
	SetPageNumber(v int32) *DescribeApisWithStageNameIntegratedByAppResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisWithStageNameIntegratedByAppResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApisWithStageNameIntegratedByAppResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApisWithStageNameIntegratedByAppResponseBody
	GetTotalCount() *int32
}

type DescribeApisWithStageNameIntegratedByAppResponseBody struct {
	// The authorization information of the API.
	AppApiRelationInfos *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfos `json:"AppApiRelationInfos,omitempty" xml:"AppApiRelationInfos,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6C87A26A-6A18-4B8E-8099-705278381A2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisWithStageNameIntegratedByAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisWithStageNameIntegratedByAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBody) GetAppApiRelationInfos() *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfos {
	return s.AppApiRelationInfos
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBody) SetAppApiRelationInfos(v *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfos) *DescribeApisWithStageNameIntegratedByAppResponseBody {
	s.AppApiRelationInfos = v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBody) SetPageNumber(v int32) *DescribeApisWithStageNameIntegratedByAppResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBody) SetPageSize(v int32) *DescribeApisWithStageNameIntegratedByAppResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBody) SetRequestId(v string) *DescribeApisWithStageNameIntegratedByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBody) SetTotalCount(v int32) *DescribeApisWithStageNameIntegratedByAppResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBody) Validate() error {
	if s.AppApiRelationInfos != nil {
		if err := s.AppApiRelationInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfos struct {
	AppApiRelationInfo []*DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo `json:"AppApiRelationInfo,omitempty" xml:"AppApiRelationInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfos) GetAppApiRelationInfo() []*DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	return s.AppApiRelationInfo
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfos) SetAppApiRelationInfo(v []*DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfos {
	s.AppApiRelationInfo = v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfos) Validate() error {
	if s.AppApiRelationInfo != nil {
		for _, item := range s.AppApiRelationInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo struct {
	// The API ID.
	//
	// example:
	//
	// c625795a1e664782a918eaa629e96ab5
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The API name.
	//
	// example:
	//
	// sd2_h5_player_history_competition
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The authorization source.
	//
	// example:
	//
	// SELF
	AuthorizationSource *string `json:"AuthorizationSource,omitempty" xml:"AuthorizationSource,omitempty"`
	// The time when the authorization was created.
	//
	// example:
	//
	// Sun, 02 Jan 2022 16:00:00 GMT
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The API description.
	//
	// example:
	//
	// b1008kz
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 38bed48d664149e4b9b5ea576319673f
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group.
	//
	// example:
	//
	// manager_vehicles
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The request HTTP method of the API.
	//
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The authorizer. Valid values:
	//
	// 	- **PROVIDER:*	- the API owner
	//
	// 	- **CONSUMER:*	- the API caller
	//
	// example:
	//
	// PROVIDER
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The request path of the API.
	//
	// example:
	//
	// /api/offline/cacheData
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The mapping information between environments and authorizations.
	StageNameAndAuth map[string]*string `json:"StageNameAndAuth,omitempty" xml:"StageNameAndAuth,omitempty"`
}

func (s DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetAuthorizationSource() *string {
	return s.AuthorizationSource
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetMethod() *string {
	return s.Method
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetOperator() *string {
	return s.Operator
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetPath() *string {
	return s.Path
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetStageNameAndAuth() map[string]*string {
	return s.StageNameAndAuth
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetApiId(v string) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetApiName(v string) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetAuthorizationSource(v string) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.AuthorizationSource = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetCreatedTime(v string) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetDescription(v string) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetGroupId(v string) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetGroupName(v string) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetMethod(v string) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Method = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetOperator(v string) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Operator = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetPath(v string) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Path = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetRegionId(v string) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetStageNameAndAuth(v map[string]*string) *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.StageNameAndAuth = v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) Validate() error {
	return dara.Validate(s)
}
