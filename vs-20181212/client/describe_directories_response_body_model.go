// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody
	GetDirectories() []*DescribeDirectoriesResponseBodyDirectories
	SetPageCount(v int64) *DescribeDirectoriesResponseBody
	GetPageCount() *int64
	SetPageNum(v int64) *DescribeDirectoriesResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeDirectoriesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDirectoriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDirectoriesResponseBody
	GetTotalCount() *int64
}

type DescribeDirectoriesResponseBody struct {
	Directories []*DescribeDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBody) GetDirectories() []*DescribeDirectoriesResponseBodyDirectories {
	return s.Directories
}

func (s *DescribeDirectoriesResponseBody) GetPageCount() *int64 {
	return s.PageCount
}

func (s *DescribeDirectoriesResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeDirectoriesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDirectoriesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDirectoriesResponseBody) SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetPageCount(v int64) *DescribeDirectoriesResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetPageNum(v int64) *DescribeDirectoriesResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetPageSize(v int64) *DescribeDirectoriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetRequestId(v string) *DescribeDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetTotalCount(v int64) *DescribeDirectoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) Validate() error {
	if s.Directories != nil {
		for _, item := range s.Directories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDirectoriesResponseBodyDirectories struct {
	// example:
	//
	// 2021-09-10T10:00:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 399*****488-cn-qingdao
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 399*****774-cn-qingdao
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectories) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectories) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetDescription() *string {
	return s.Description
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetId() *string {
	return s.Id
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetName() *string {
	return s.Name
}

func (s *DescribeDirectoriesResponseBodyDirectories) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetCreatedTime(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDescription(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Description = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetGroupId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.GroupId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Id = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Name = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetParentId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.ParentId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) Validate() error {
	return dara.Validate(s)
}
