// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppApiRelationInfos(v *DescribeApisByAppResponseBodyAppApiRelationInfos) *DescribeApisByAppResponseBody
	GetAppApiRelationInfos() *DescribeApisByAppResponseBodyAppApiRelationInfos
	SetPageNumber(v int32) *DescribeApisByAppResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisByAppResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApisByAppResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApisByAppResponseBody
	GetTotalCount() *int32
}

type DescribeApisByAppResponseBody struct {
	// The API authorizations.
	AppApiRelationInfos *DescribeApisByAppResponseBodyAppApiRelationInfos `json:"AppApiRelationInfos,omitempty" xml:"AppApiRelationInfos,omitempty" type:"Struct"`
	// The number of pages to return the results on.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisByAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisByAppResponseBody) GetAppApiRelationInfos() *DescribeApisByAppResponseBodyAppApiRelationInfos {
	return s.AppApiRelationInfos
}

func (s *DescribeApisByAppResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisByAppResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisByAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisByAppResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApisByAppResponseBody) SetAppApiRelationInfos(v *DescribeApisByAppResponseBodyAppApiRelationInfos) *DescribeApisByAppResponseBody {
	s.AppApiRelationInfos = v
	return s
}

func (s *DescribeApisByAppResponseBody) SetPageNumber(v int32) *DescribeApisByAppResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByAppResponseBody) SetPageSize(v int32) *DescribeApisByAppResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByAppResponseBody) SetRequestId(v string) *DescribeApisByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisByAppResponseBody) SetTotalCount(v int32) *DescribeApisByAppResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisByAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApisByAppResponseBodyAppApiRelationInfos struct {
	AppApiRelationInfo []*DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo `json:"AppApiRelationInfo,omitempty" xml:"AppApiRelationInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisByAppResponseBodyAppApiRelationInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByAppResponseBodyAppApiRelationInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfos) GetAppApiRelationInfo() []*DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	return s.AppApiRelationInfo
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfos) SetAppApiRelationInfo(v []*DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) *DescribeApisByAppResponseBodyAppApiRelationInfos {
	s.AppApiRelationInfo = v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo struct {
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
	// The validity period of the authorization.
	//
	// example:
	//
	// Sun, 02 Jan 2022 16:00:00 GMT
	AuthVaildTime *string `json:"AuthVaildTime,omitempty" xml:"AuthVaildTime,omitempty"`
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
	// The authorization description.
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
	// wb2022012101605749
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The HTTP method of the API.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The authorizer. Valid values:
	//
	// 	- **PROVIDER**: API owner
	//
	// 	- **CONSUMER**: API caller
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
	// The environment alias.
	//
	// example:
	//
	// Online
	StageAlias *string `json:"StageAlias,omitempty" xml:"StageAlias,omitempty"`
	// The environment name.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetAuthVaildTime() *string {
	return s.AuthVaildTime
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetAuthorizationSource() *string {
	return s.AuthorizationSource
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetMethod() *string {
	return s.Method
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetOperator() *string {
	return s.Operator
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetPath() *string {
	return s.Path
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetStageAlias() *string {
	return s.StageAlias
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetApiId(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetApiName(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetAuthVaildTime(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.AuthVaildTime = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetAuthorizationSource(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.AuthorizationSource = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetCreatedTime(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetDescription(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetGroupId(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetGroupName(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetMethod(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Method = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetOperator(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Operator = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetPath(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Path = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetRegionId(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetStageAlias(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.StageAlias = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetStageName(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.StageName = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) Validate() error {
	return dara.Validate(s)
}
