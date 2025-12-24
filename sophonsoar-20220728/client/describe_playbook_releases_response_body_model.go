// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookReleasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v *DescribePlaybookReleasesResponseBodyPage) *DescribePlaybookReleasesResponseBody
	GetPage() *DescribePlaybookReleasesResponseBodyPage
	SetRecords(v []*DescribePlaybookReleasesResponseBodyRecords) *DescribePlaybookReleasesResponseBody
	GetRecords() []*DescribePlaybookReleasesResponseBodyRecords
	SetRequestId(v string) *DescribePlaybookReleasesResponseBody
	GetRequestId() *string
}

type DescribePlaybookReleasesResponseBody struct {
	// The pagination information.
	Page *DescribePlaybookReleasesResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The information about the playbook version.
	Records []*DescribePlaybookReleasesResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3DFBE11C-6EB6-5166-92D6-3397796AFE1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookReleasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookReleasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesResponseBody) GetPage() *DescribePlaybookReleasesResponseBodyPage {
	return s.Page
}

func (s *DescribePlaybookReleasesResponseBody) GetRecords() []*DescribePlaybookReleasesResponseBodyRecords {
	return s.Records
}

func (s *DescribePlaybookReleasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlaybookReleasesResponseBody) SetPage(v *DescribePlaybookReleasesResponseBodyPage) *DescribePlaybookReleasesResponseBody {
	s.Page = v
	return s
}

func (s *DescribePlaybookReleasesResponseBody) SetRecords(v []*DescribePlaybookReleasesResponseBodyRecords) *DescribePlaybookReleasesResponseBody {
	s.Records = v
	return s
}

func (s *DescribePlaybookReleasesResponseBody) SetRequestId(v string) *DescribePlaybookReleasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePlaybookReleasesResponseBodyPage struct {
	// The page number.
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
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePlaybookReleasesResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookReleasesResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePlaybookReleasesResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePlaybookReleasesResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePlaybookReleasesResponseBodyPage) SetPageNumber(v int32) *DescribePlaybookReleasesResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyPage) SetPageSize(v int32) *DescribePlaybookReleasesResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyPage) SetTotalCount(v int32) *DescribePlaybookReleasesResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyPage) Validate() error {
	return dara.Validate(s)
}

type DescribePlaybookReleasesResponseBodyRecords struct {
	// The ID of the Alibaba Cloud account that is used to publish the version.
	//
	// example:
	//
	// 145xxxx985
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The description of the layer version.
	//
	// example:
	//
	// This is a new version
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the version was created. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1655277397000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the version was modified. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1691460804000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The record ID.
	//
	// example:
	//
	// 80xxx
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The MD5 value configured for the published version of the playbook.
	//
	// example:
	//
	// be0a4ef084dd174abe47xxxxx
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
}

func (s DescribePlaybookReleasesResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookReleasesResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesResponseBodyRecords) GetCreator() *string {
	return s.Creator
}

func (s *DescribePlaybookReleasesResponseBodyRecords) GetDescription() *string {
	return s.Description
}

func (s *DescribePlaybookReleasesResponseBodyRecords) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribePlaybookReleasesResponseBodyRecords) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribePlaybookReleasesResponseBodyRecords) GetId() *int32 {
	return s.Id
}

func (s *DescribePlaybookReleasesResponseBodyRecords) GetTaskflowMd5() *string {
	return s.TaskflowMd5
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetCreator(v string) *DescribePlaybookReleasesResponseBodyRecords {
	s.Creator = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetDescription(v string) *DescribePlaybookReleasesResponseBodyRecords {
	s.Description = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetGmtCreate(v int64) *DescribePlaybookReleasesResponseBodyRecords {
	s.GmtCreate = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetGmtModified(v int64) *DescribePlaybookReleasesResponseBodyRecords {
	s.GmtModified = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetId(v int32) *DescribePlaybookReleasesResponseBodyRecords {
	s.Id = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) SetTaskflowMd5(v string) *DescribePlaybookReleasesResponseBodyRecords {
	s.TaskflowMd5 = &v
	return s
}

func (s *DescribePlaybookReleasesResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
