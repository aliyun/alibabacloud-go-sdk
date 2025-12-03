// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiSummarys(v *DescribeApisResponseBodyApiSummarys) *DescribeApisResponseBody
	GetApiSummarys() *DescribeApisResponseBodyApiSummarys
	SetPageNumber(v int32) *DescribeApisResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApisResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApisResponseBody
	GetTotalCount() *int32
}

type DescribeApisResponseBody struct {
	// The queried API definitions.
	ApiSummarys *DescribeApisResponseBodyApiSummarys `json:"ApiSummarys,omitempty" xml:"ApiSummarys,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 3
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 15
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ002
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 32
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBody) GetApiSummarys() *DescribeApisResponseBodyApiSummarys {
	return s.ApiSummarys
}

func (s *DescribeApisResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApisResponseBody) SetApiSummarys(v *DescribeApisResponseBodyApiSummarys) *DescribeApisResponseBody {
	s.ApiSummarys = v
	return s
}

func (s *DescribeApisResponseBody) SetPageNumber(v int32) *DescribeApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisResponseBody) SetPageSize(v int32) *DescribeApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisResponseBody) SetRequestId(v string) *DescribeApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisResponseBody) SetTotalCount(v int32) *DescribeApisResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisResponseBody) Validate() error {
	if s.ApiSummarys != nil {
		if err := s.ApiSummarys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApisResponseBodyApiSummarys struct {
	ApiSummary []*DescribeApisResponseBodyApiSummarysApiSummary `json:"ApiSummary,omitempty" xml:"ApiSummary,omitempty" type:"Repeated"`
}

func (s DescribeApisResponseBodyApiSummarys) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisResponseBodyApiSummarys) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBodyApiSummarys) GetApiSummary() []*DescribeApisResponseBodyApiSummarysApiSummary {
	return s.ApiSummary
}

func (s *DescribeApisResponseBodyApiSummarys) SetApiSummary(v []*DescribeApisResponseBodyApiSummarysApiSummary) *DescribeApisResponseBodyApiSummarys {
	s.ApiSummary = v
	return s
}

func (s *DescribeApisResponseBodyApiSummarys) Validate() error {
	if s.ApiSummary != nil {
		for _, item := range s.ApiSummary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisResponseBodyApiSummarysApiSummary struct {
	// The API ID.
	//
	// example:
	//
	// 5af418828f0344a3b588c0cc1331a3bc
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The HTTP method of the API request.
	//
	// example:
	//
	// GET
	ApiMethod *string `json:"ApiMethod,omitempty" xml:"ApiMethod,omitempty"`
	// The API name.
	//
	// example:
	//
	// CreateObject
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The request path of the API.
	//
	// example:
	//
	// /sqb/
	ApiPath *string `json:"ApiPath,omitempty" xml:"ApiPath,omitempty"`
	// The time when the API was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-05-20T18:18:25Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The API publishing statuses.
	DeployedInfos *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfos `json:"DeployedInfos,omitempty" xml:"DeployedInfos,omitempty" type:"Struct"`
	// The API description.
	//
	// example:
	//
	// Creates an object
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The API group ID.
	//
	// example:
	//
	// 1084f9034c744137901057206b39d2b6
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group to which the API belongs.
	//
	// example:
	//
	// myGroup2
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when the API was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-07-23T08:28:48Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the region to which the API belongs.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags that are added to the APIs.
	TagList *DescribeApisResponseBodyApiSummarysApiSummaryTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Struct"`
	// Indicates whether the API is public. Valid values:
	//
	// 	- **PUBLIC**: The API is public.
	//
	// 	- **PRIVATE**: The API is private.
	//
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeApisResponseBodyApiSummarysApiSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisResponseBodyApiSummarysApiSummary) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetApiMethod() *string {
	return s.ApiMethod
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetApiPath() *string {
	return s.ApiPath
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetDeployedInfos() *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfos {
	return s.DeployedInfos
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetDescription() *string {
	return s.Description
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetTagList() *DescribeApisResponseBodyApiSummarysApiSummaryTagList {
	return s.TagList
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) GetVisibility() *string {
	return s.Visibility
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetApiId(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.ApiId = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetApiMethod(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.ApiMethod = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetApiName(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.ApiName = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetApiPath(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.ApiPath = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetCreatedTime(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetDeployedInfos(v *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfos) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.DeployedInfos = v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetDescription(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.Description = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetGroupId(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.GroupId = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetGroupName(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.GroupName = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetModifiedTime(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetRegionId(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.RegionId = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetTagList(v *DescribeApisResponseBodyApiSummarysApiSummaryTagList) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.TagList = v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetVisibility(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.Visibility = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) Validate() error {
	if s.DeployedInfos != nil {
		if err := s.DeployedInfos.Validate(); err != nil {
			return err
		}
	}
	if s.TagList != nil {
		if err := s.TagList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfos struct {
	DeployedInfo []*DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo `json:"DeployedInfo,omitempty" xml:"DeployedInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfos) GetDeployedInfo() []*DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo {
	return s.DeployedInfo
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfos) SetDeployedInfo(v []*DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo) *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfos {
	s.DeployedInfo = v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfos) Validate() error {
	if s.DeployedInfo != nil {
		for _, item := range s.DeployedInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo struct {
	// The deployment status. Valid values: DEPLOYED and NONDEPLOYED.
	//
	// example:
	//
	// DEPLOYED
	DeployedStatus *string `json:"DeployedStatus,omitempty" xml:"DeployedStatus,omitempty"`
	// The deployed version.
	//
	// example:
	//
	// 20220103170737313
	EffectiveVersion *string `json:"EffectiveVersion,omitempty" xml:"EffectiveVersion,omitempty"`
	// Stage Name:
	//
	// 	- **RELEASE**: production environment
	//
	// 	- **PRE**: staging environment
	//
	// 	- **TEST**: test environment
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo) GetDeployedStatus() *string {
	return s.DeployedStatus
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo) GetEffectiveVersion() *string {
	return s.EffectiveVersion
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo) SetDeployedStatus(v string) *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo {
	s.DeployedStatus = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo) SetEffectiveVersion(v string) *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo {
	s.EffectiveVersion = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo) SetStageName(v string) *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo {
	s.StageName = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryDeployedInfosDeployedInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeApisResponseBodyApiSummarysApiSummaryTagList struct {
	Tag []*DescribeApisResponseBodyApiSummarysApiSummaryTagListTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeApisResponseBodyApiSummarysApiSummaryTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisResponseBodyApiSummarysApiSummaryTagList) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryTagList) GetTag() []*DescribeApisResponseBodyApiSummarysApiSummaryTagListTag {
	return s.Tag
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryTagList) SetTag(v []*DescribeApisResponseBodyApiSummarysApiSummaryTagListTag) *DescribeApisResponseBodyApiSummarysApiSummaryTagList {
	s.Tag = v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryTagList) Validate() error {
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

type DescribeApisResponseBodyApiSummarysApiSummaryTagListTag struct {
	// The tag key.
	//
	// example:
	//
	// tagKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tagValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeApisResponseBodyApiSummarysApiSummaryTagListTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisResponseBodyApiSummarysApiSummaryTagListTag) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryTagListTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryTagListTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryTagListTag) SetTagKey(v string) *DescribeApisResponseBodyApiSummarysApiSummaryTagListTag {
	s.TagKey = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryTagListTag) SetTagValue(v string) *DescribeApisResponseBodyApiSummarysApiSummaryTagListTag {
	s.TagValue = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummaryTagListTag) Validate() error {
	return dara.Validate(s)
}
