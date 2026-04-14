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
	if s.AppApiRelationInfos != nil {
		if err := s.AppApiRelationInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo struct {
	ApiId               *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName             *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthVaildTime       *string `json:"AuthVaildTime,omitempty" xml:"AuthVaildTime,omitempty"`
	AuthorizationSource *string `json:"AuthorizationSource,omitempty" xml:"AuthorizationSource,omitempty"`
	CreatedTime         *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Method              *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Operator            *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Path                *string `json:"Path,omitempty" xml:"Path,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageAlias          *string `json:"StageAlias,omitempty" xml:"StageAlias,omitempty"`
	StageName           *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
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
