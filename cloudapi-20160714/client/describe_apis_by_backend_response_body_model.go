// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByBackendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiInfoList(v *DescribeApisByBackendResponseBodyApiInfoList) *DescribeApisByBackendResponseBody
	GetApiInfoList() *DescribeApisByBackendResponseBodyApiInfoList
	SetPageNumber(v int32) *DescribeApisByBackendResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisByBackendResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApisByBackendResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApisByBackendResponseBody
	GetTotalCount() *int32
}

type DescribeApisByBackendResponseBody struct {
	// The information about the returned API list.
	ApiInfoList *DescribeApisByBackendResponseBodyApiInfoList `json:"ApiInfoList,omitempty" xml:"ApiInfoList,omitempty" type:"Struct"`
	// The number of the current page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEB6EC62-B6C7-5082-A45A-45A204724AC2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisByBackendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByBackendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisByBackendResponseBody) GetApiInfoList() *DescribeApisByBackendResponseBodyApiInfoList {
	return s.ApiInfoList
}

func (s *DescribeApisByBackendResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisByBackendResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisByBackendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisByBackendResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApisByBackendResponseBody) SetApiInfoList(v *DescribeApisByBackendResponseBodyApiInfoList) *DescribeApisByBackendResponseBody {
	s.ApiInfoList = v
	return s
}

func (s *DescribeApisByBackendResponseBody) SetPageNumber(v int32) *DescribeApisByBackendResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByBackendResponseBody) SetPageSize(v int32) *DescribeApisByBackendResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByBackendResponseBody) SetRequestId(v string) *DescribeApisByBackendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisByBackendResponseBody) SetTotalCount(v int32) *DescribeApisByBackendResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisByBackendResponseBody) Validate() error {
	if s.ApiInfoList != nil {
		if err := s.ApiInfoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApisByBackendResponseBodyApiInfoList struct {
	ApiInfo []*DescribeApisByBackendResponseBodyApiInfoListApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisByBackendResponseBodyApiInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByBackendResponseBodyApiInfoList) GoString() string {
	return s.String()
}

func (s *DescribeApisByBackendResponseBodyApiInfoList) GetApiInfo() []*DescribeApisByBackendResponseBodyApiInfoListApiInfo {
	return s.ApiInfo
}

func (s *DescribeApisByBackendResponseBodyApiInfoList) SetApiInfo(v []*DescribeApisByBackendResponseBodyApiInfoListApiInfo) *DescribeApisByBackendResponseBodyApiInfoList {
	s.ApiInfo = v
	return s
}

func (s *DescribeApisByBackendResponseBodyApiInfoList) Validate() error {
	if s.ApiInfo != nil {
		for _, item := range s.ApiInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisByBackendResponseBodyApiInfoListApiInfo struct {
	// The ID of the API.
	//
	// example:
	//
	// 20bcdc9453524b78a8beb1f6de21edb7
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API.
	//
	// example:
	//
	// testApi
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The description of the API.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 15bcdc9453524b7gs8beb1f6de21edb7
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group.
	//
	// example:
	//
	// testApiGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The request method of the API.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The request path of the API.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeApisByBackendResponseBodyApiInfoListApiInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByBackendResponseBodyApiInfoListApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) GetMethod() *string {
	return s.Method
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) GetPath() *string {
	return s.Path
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) SetApiId(v string) *DescribeApisByBackendResponseBodyApiInfoListApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) SetApiName(v string) *DescribeApisByBackendResponseBodyApiInfoListApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) SetDescription(v string) *DescribeApisByBackendResponseBodyApiInfoListApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) SetGroupId(v string) *DescribeApisByBackendResponseBodyApiInfoListApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) SetGroupName(v string) *DescribeApisByBackendResponseBodyApiInfoListApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) SetMethod(v string) *DescribeApisByBackendResponseBodyApiInfoListApiInfo {
	s.Method = &v
	return s
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) SetPath(v string) *DescribeApisByBackendResponseBodyApiInfoListApiInfo {
	s.Path = &v
	return s
}

func (s *DescribeApisByBackendResponseBodyApiInfoListApiInfo) Validate() error {
	return dara.Validate(s)
}
