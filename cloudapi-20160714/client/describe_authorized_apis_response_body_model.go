// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthorizedApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedApis(v *DescribeAuthorizedApisResponseBodyAuthorizedApis) *DescribeAuthorizedApisResponseBody
	GetAuthorizedApis() *DescribeAuthorizedApisResponseBodyAuthorizedApis
	SetPageNumber(v int32) *DescribeAuthorizedApisResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAuthorizedApisResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAuthorizedApisResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAuthorizedApisResponseBody
	GetTotalCount() *int32
}

type DescribeAuthorizedApisResponseBody struct {
	AuthorizedApis *DescribeAuthorizedApisResponseBodyAuthorizedApis `json:"AuthorizedApis,omitempty" xml:"AuthorizedApis,omitempty" type:"Struct"`
	// The page number of the returned page.
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
	// D6E46F10-F26C-4AA0-BB69-FE2743D9AE62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuthorizedApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthorizedApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedApisResponseBody) GetAuthorizedApis() *DescribeAuthorizedApisResponseBodyAuthorizedApis {
	return s.AuthorizedApis
}

func (s *DescribeAuthorizedApisResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAuthorizedApisResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAuthorizedApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuthorizedApisResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAuthorizedApisResponseBody) SetAuthorizedApis(v *DescribeAuthorizedApisResponseBodyAuthorizedApis) *DescribeAuthorizedApisResponseBody {
	s.AuthorizedApis = v
	return s
}

func (s *DescribeAuthorizedApisResponseBody) SetPageNumber(v int32) *DescribeAuthorizedApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBody) SetPageSize(v int32) *DescribeAuthorizedApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBody) SetRequestId(v string) *DescribeAuthorizedApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBody) SetTotalCount(v int32) *DescribeAuthorizedApisResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBody) Validate() error {
	if s.AuthorizedApis != nil {
		if err := s.AuthorizedApis.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAuthorizedApisResponseBodyAuthorizedApis struct {
	AuthorizedApi []*DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi `json:"AuthorizedApi,omitempty" xml:"AuthorizedApi,omitempty" type:"Repeated"`
}

func (s DescribeAuthorizedApisResponseBodyAuthorizedApis) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthorizedApisResponseBodyAuthorizedApis) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApis) GetAuthorizedApi() []*DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	return s.AuthorizedApi
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApis) SetAuthorizedApi(v []*DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) *DescribeAuthorizedApisResponseBodyAuthorizedApis {
	s.AuthorizedApi = v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApis) Validate() error {
	if s.AuthorizedApi != nil {
		for _, item := range s.AuthorizedApi {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi struct {
	ApiId               *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName             *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthVaildTime       *string `json:"AuthVaildTime,omitempty" xml:"AuthVaildTime,omitempty"`
	AuthorizationSource *string `json:"AuthorizationSource,omitempty" xml:"AuthorizationSource,omitempty"`
	AuthorizedTime      *string `json:"AuthorizedTime,omitempty" xml:"AuthorizedTime,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Operator            *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageName           *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GetAuthVaildTime() *string {
	return s.AuthVaildTime
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GetAuthorizationSource() *string {
	return s.AuthorizationSource
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GetAuthorizedTime() *string {
	return s.AuthorizedTime
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GetDescription() *string {
	return s.Description
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GetOperator() *string {
	return s.Operator
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GetStageName() *string {
	return s.StageName
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetApiId(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.ApiId = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetApiName(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.ApiName = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetAuthVaildTime(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.AuthVaildTime = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetAuthorizationSource(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.AuthorizationSource = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetAuthorizedTime(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.AuthorizedTime = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetDescription(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.Description = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetGroupId(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.GroupId = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetGroupName(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.GroupName = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetOperator(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.Operator = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetRegionId(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.RegionId = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetStageName(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.StageName = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) Validate() error {
	return dara.Validate(s)
}
