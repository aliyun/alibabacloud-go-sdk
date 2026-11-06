// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillsMarketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribePolarClawSkillsMarketResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribePolarClawSkillsMarketResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribePolarClawSkillsMarketResponseBody
	GetRequestId() *string
	SetSkillList(v []*DescribePolarClawSkillsMarketResponseBodySkillList) *DescribePolarClawSkillsMarketResponseBody
	GetSkillList() []*DescribePolarClawSkillsMarketResponseBodySkillList
	SetTotalRecordCount(v int32) *DescribePolarClawSkillsMarketResponseBody
	GetTotalRecordCount() *int32
}

type DescribePolarClawSkillsMarketResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The skill list.
	SkillList []*DescribePolarClawSkillsMarketResponseBodySkillList `json:"SkillList,omitempty" xml:"SkillList,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 9
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribePolarClawSkillsMarketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillsMarketResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillsMarketResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolarClawSkillsMarketResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribePolarClawSkillsMarketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawSkillsMarketResponseBody) GetSkillList() []*DescribePolarClawSkillsMarketResponseBodySkillList {
	return s.SkillList
}

func (s *DescribePolarClawSkillsMarketResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribePolarClawSkillsMarketResponseBody) SetPageNumber(v int32) *DescribePolarClawSkillsMarketResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarClawSkillsMarketResponseBody) SetPageRecordCount(v int32) *DescribePolarClawSkillsMarketResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribePolarClawSkillsMarketResponseBody) SetRequestId(v string) *DescribePolarClawSkillsMarketResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawSkillsMarketResponseBody) SetSkillList(v []*DescribePolarClawSkillsMarketResponseBodySkillList) *DescribePolarClawSkillsMarketResponseBody {
	s.SkillList = v
	return s
}

func (s *DescribePolarClawSkillsMarketResponseBody) SetTotalRecordCount(v int32) *DescribePolarClawSkillsMarketResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribePolarClawSkillsMarketResponseBody) Validate() error {
	if s.SkillList != nil {
		for _, item := range s.SkillList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarClawSkillsMarketResponseBodySkillList struct {
	// The assessment status.
	//
	// example:
	//
	// 2
	AssessmentStatus *string `json:"AssessmentStatus,omitempty" xml:"AssessmentStatus,omitempty"`
	// The number of downloads.
	//
	// example:
	//
	// 3
	DownloadCount *string `json:"DownloadCount,omitempty" xml:"DownloadCount,omitempty"`
	// The skill name.
	//
	// example:
	//
	// alibabacloud-polardbx-sql
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The skill version.
	//
	// example:
	//
	// 0.0.1-beta.1
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
	// The OpenAPI summary.
	//
	// example:
	//
	// Automatically created repository via Image-manage
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The tag list.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The title.
	//
	// example:
	//
	// Alcohol
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribePolarClawSkillsMarketResponseBodySkillList) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillsMarketResponseBodySkillList) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) GetAssessmentStatus() *string {
	return s.AssessmentStatus
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) GetDownloadCount() *string {
	return s.DownloadCount
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) GetSkillName() *string {
	return s.SkillName
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) GetSummary() *string {
	return s.Summary
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) GetTagList() []*string {
	return s.TagList
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) GetTitle() *string {
	return s.Title
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) SetAssessmentStatus(v string) *DescribePolarClawSkillsMarketResponseBodySkillList {
	s.AssessmentStatus = &v
	return s
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) SetDownloadCount(v string) *DescribePolarClawSkillsMarketResponseBodySkillList {
	s.DownloadCount = &v
	return s
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) SetSkillName(v string) *DescribePolarClawSkillsMarketResponseBodySkillList {
	s.SkillName = &v
	return s
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) SetSkillVersion(v string) *DescribePolarClawSkillsMarketResponseBodySkillList {
	s.SkillVersion = &v
	return s
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) SetSummary(v string) *DescribePolarClawSkillsMarketResponseBodySkillList {
	s.Summary = &v
	return s
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) SetTagList(v []*string) *DescribePolarClawSkillsMarketResponseBodySkillList {
	s.TagList = v
	return s
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) SetTitle(v string) *DescribePolarClawSkillsMarketResponseBodySkillList {
	s.Title = &v
	return s
}

func (s *DescribePolarClawSkillsMarketResponseBodySkillList) Validate() error {
	return dara.Validate(s)
}
