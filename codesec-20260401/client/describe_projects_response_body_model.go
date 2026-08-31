// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeProjectsResponseBodyItems) *DescribeProjectsResponseBody
	GetItems() []*DescribeProjectsResponseBodyItems
	SetMaxResults(v int32) *DescribeProjectsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeProjectsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeProjectsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeProjectsResponseBody
	GetTotalCount() *int64
}

type DescribeProjectsResponseBody struct {
	// The list of projects.
	Items []*DescribeProjectsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page size.
	//
	// > If not specified, all projects are displayed.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. An empty value indicates the last page.
	//
	// example:
	//
	// eyJ0IjoiMjAyNi0wNy0xNlQwNzo1MzozOC4wMjFaIiwiaSI6MTAwMDQ0OH0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9A1F403F-0A85-5578-8B7C-55E3E9408659
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s DescribeProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectsResponseBody) GetItems() []*DescribeProjectsResponseBodyItems {
	return s.Items
}

func (s *DescribeProjectsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeProjectsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProjectsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeProjectsResponseBody) SetItems(v []*DescribeProjectsResponseBodyItems) *DescribeProjectsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeProjectsResponseBody) SetMaxResults(v int32) *DescribeProjectsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeProjectsResponseBody) SetNextToken(v string) *DescribeProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeProjectsResponseBody) SetRequestId(v string) *DescribeProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectsResponseBody) SetTotalCount(v int64) *DescribeProjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeProjectsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeProjectsResponseBodyItems struct {
	// The project configuration version number.
	//
	// example:
	//
	// 1
	ConfigRevision *int64 `json:"configRevision,omitempty" xml:"configRevision,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-07-28T03:36:31.573Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The user ID of the project creator.
	//
	// example:
	//
	// 11111
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The description.
	//
	// example:
	//
	// 111
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The engine switches for the project or scan snapshot (SAST and SCA only).
	Engines *DescribeProjectsResponseBodyItemsEngines `json:"engines,omitempty" xml:"engines,omitempty" type:"Struct"`
	// The project ID.
	//
	// example:
	//
	// 934
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The natural language prompt provided by the user that describes scanning or result processing preferences, such as ignoring low-risk vulnerabilities.
	//
	// example:
	//
	// 1111
	InstructionPrompt *string `json:"instructionPrompt,omitempty" xml:"instructionPrompt,omitempty"`
	// The project name.
	//
	// example:
	//
	// manual-hDecBn
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The project source.
	Source *DescribeProjectsResponseBodyItemsSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	// The update time.
	//
	// example:
	//
	// 2026-07-28T03:36:31.573Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s DescribeProjectsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeProjectsResponseBodyItems) GetConfigRevision() *int64 {
	return s.ConfigRevision
}

func (s *DescribeProjectsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeProjectsResponseBodyItems) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *DescribeProjectsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeProjectsResponseBodyItems) GetEngines() *DescribeProjectsResponseBodyItemsEngines {
	return s.Engines
}

func (s *DescribeProjectsResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeProjectsResponseBodyItems) GetInstructionPrompt() *string {
	return s.InstructionPrompt
}

func (s *DescribeProjectsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeProjectsResponseBodyItems) GetSource() *DescribeProjectsResponseBodyItemsSource {
	return s.Source
}

func (s *DescribeProjectsResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DescribeProjectsResponseBodyItems) SetConfigRevision(v int64) *DescribeProjectsResponseBodyItems {
	s.ConfigRevision = &v
	return s
}

func (s *DescribeProjectsResponseBodyItems) SetCreatedAt(v string) *DescribeProjectsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *DescribeProjectsResponseBodyItems) SetCreatedBy(v string) *DescribeProjectsResponseBodyItems {
	s.CreatedBy = &v
	return s
}

func (s *DescribeProjectsResponseBodyItems) SetDescription(v string) *DescribeProjectsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeProjectsResponseBodyItems) SetEngines(v *DescribeProjectsResponseBodyItemsEngines) *DescribeProjectsResponseBodyItems {
	s.Engines = v
	return s
}

func (s *DescribeProjectsResponseBodyItems) SetId(v int64) *DescribeProjectsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeProjectsResponseBodyItems) SetInstructionPrompt(v string) *DescribeProjectsResponseBodyItems {
	s.InstructionPrompt = &v
	return s
}

func (s *DescribeProjectsResponseBodyItems) SetName(v string) *DescribeProjectsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeProjectsResponseBodyItems) SetSource(v *DescribeProjectsResponseBodyItemsSource) *DescribeProjectsResponseBodyItems {
	s.Source = v
	return s
}

func (s *DescribeProjectsResponseBodyItems) SetUpdatedAt(v string) *DescribeProjectsResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeProjectsResponseBodyItems) Validate() error {
	if s.Engines != nil {
		if err := s.Engines.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeProjectsResponseBodyItemsEngines struct {
	// Indicates whether SAST is supported.
	//
	// example:
	//
	// true
	Sast *bool `json:"sast,omitempty" xml:"sast,omitempty"`
	// Indicates whether SCA is supported.
	//
	// example:
	//
	// true
	Sca *bool `json:"sca,omitempty" xml:"sca,omitempty"`
}

func (s DescribeProjectsResponseBodyItemsEngines) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectsResponseBodyItemsEngines) GoString() string {
	return s.String()
}

func (s *DescribeProjectsResponseBodyItemsEngines) GetSast() *bool {
	return s.Sast
}

func (s *DescribeProjectsResponseBodyItemsEngines) GetSca() *bool {
	return s.Sca
}

func (s *DescribeProjectsResponseBodyItemsEngines) SetSast(v bool) *DescribeProjectsResponseBodyItemsEngines {
	s.Sast = &v
	return s
}

func (s *DescribeProjectsResponseBodyItemsEngines) SetSca(v bool) *DescribeProjectsResponseBodyItemsEngines {
	s.Sca = &v
	return s
}

func (s *DescribeProjectsResponseBodyItemsEngines) Validate() error {
	return dara.Validate(s)
}

type DescribeProjectsResponseBodyItemsSource struct {
	// The project type.
	//
	// example:
	//
	// manual_upload
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeProjectsResponseBodyItemsSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectsResponseBodyItemsSource) GoString() string {
	return s.String()
}

func (s *DescribeProjectsResponseBodyItemsSource) GetType() *string {
	return s.Type
}

func (s *DescribeProjectsResponseBodyItemsSource) SetType(v string) *DescribeProjectsResponseBodyItemsSource {
	s.Type = &v
	return s
}

func (s *DescribeProjectsResponseBodyItemsSource) Validate() error {
	return dara.Validate(s)
}
