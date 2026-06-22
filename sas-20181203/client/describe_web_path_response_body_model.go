// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigList(v []*DescribeWebPathResponseBodyConfigList) *DescribeWebPathResponseBody
	GetConfigList() []*DescribeWebPathResponseBodyConfigList
	SetCount(v int32) *DescribeWebPathResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeWebPathResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeWebPathResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeWebPathResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWebPathResponseBody
	GetTotalCount() *int32
}

type DescribeWebPathResponseBody struct {
	// The list of web paths.
	ConfigList []*DescribeWebPathResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the current page in a paged query. Paging is used to display results.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries per page in a paged query. Paging is used to display results.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request. The China Chinese Cloud generates a unique identifier for each request, which can be used for troubleshooting and diagnostics.
	//
	// example:
	//
	// B37C9052-A73E-4707-A024-92477028****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebPathResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebPathResponseBody) GetConfigList() []*DescribeWebPathResponseBodyConfigList {
	return s.ConfigList
}

func (s *DescribeWebPathResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeWebPathResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWebPathResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebPathResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWebPathResponseBody) SetConfigList(v []*DescribeWebPathResponseBodyConfigList) *DescribeWebPathResponseBody {
	s.ConfigList = v
	return s
}

func (s *DescribeWebPathResponseBody) SetCount(v int32) *DescribeWebPathResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeWebPathResponseBody) SetCurrentPage(v int32) *DescribeWebPathResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebPathResponseBody) SetPageSize(v int32) *DescribeWebPathResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWebPathResponseBody) SetRequestId(v string) *DescribeWebPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebPathResponseBody) SetTotalCount(v int32) *DescribeWebPathResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebPathResponseBody) Validate() error {
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebPathResponseBodyConfigList struct {
	// The list of servers on which the web directory takes effect.
	TargetList []*DescribeWebPathResponseBodyConfigListTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
	// The web directory.
	//
	// example:
	//
	// /root/www****
	WebPath *string `json:"WebPath,omitempty" xml:"WebPath,omitempty"`
	// The type of the web path. Valid values:
	//
	// - **def**: automatically identified by the system.
	//
	// - **customize**: manually added.
	//
	// example:
	//
	// def
	WebPathType *string `json:"WebPathType,omitempty" xml:"WebPathType,omitempty"`
}

func (s DescribeWebPathResponseBodyConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebPathResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeWebPathResponseBodyConfigList) GetTargetList() []*DescribeWebPathResponseBodyConfigListTargetList {
	return s.TargetList
}

func (s *DescribeWebPathResponseBodyConfigList) GetWebPath() *string {
	return s.WebPath
}

func (s *DescribeWebPathResponseBodyConfigList) GetWebPathType() *string {
	return s.WebPathType
}

func (s *DescribeWebPathResponseBodyConfigList) SetTargetList(v []*DescribeWebPathResponseBodyConfigListTargetList) *DescribeWebPathResponseBodyConfigList {
	s.TargetList = v
	return s
}

func (s *DescribeWebPathResponseBodyConfigList) SetWebPath(v string) *DescribeWebPathResponseBodyConfigList {
	s.WebPath = &v
	return s
}

func (s *DescribeWebPathResponseBodyConfigList) SetWebPathType(v string) *DescribeWebPathResponseBodyConfigList {
	s.WebPathType = &v
	return s
}

func (s *DescribeWebPathResponseBodyConfigList) Validate() error {
	if s.TargetList != nil {
		for _, item := range s.TargetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebPathResponseBodyConfigListTargetList struct {
	// The target object.
	//
	// example:
	//
	// 82048187-bb9b-4e19-8320-7b4ddb97****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The target type. Valid values:
	//
	// - **uuid**.
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeWebPathResponseBodyConfigListTargetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebPathResponseBodyConfigListTargetList) GoString() string {
	return s.String()
}

func (s *DescribeWebPathResponseBodyConfigListTargetList) GetTarget() *string {
	return s.Target
}

func (s *DescribeWebPathResponseBodyConfigListTargetList) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeWebPathResponseBodyConfigListTargetList) SetTarget(v string) *DescribeWebPathResponseBodyConfigListTargetList {
	s.Target = &v
	return s
}

func (s *DescribeWebPathResponseBodyConfigListTargetList) SetTargetType(v string) *DescribeWebPathResponseBodyConfigListTargetList {
	s.TargetType = &v
	return s
}

func (s *DescribeWebPathResponseBodyConfigListTargetList) Validate() error {
	return dara.Validate(s)
}
