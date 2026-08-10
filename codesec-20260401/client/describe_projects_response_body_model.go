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
	Items      []*DescribeProjectsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	MaxResults *int32                               `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                              `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId  *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int64                               `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	ConfigRevision *int64 `json:"configRevision,omitempty" xml:"configRevision,omitempty"`
	// 扫描项目创建时间（RFC3339）
	CreatedAt         *string                                   `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy         *string                                   `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Description       *string                                   `json:"description,omitempty" xml:"description,omitempty"`
	Engines           *DescribeProjectsResponseBodyItemsEngines `json:"engines,omitempty" xml:"engines,omitempty" type:"Struct"`
	Id                *int64                                    `json:"id,omitempty" xml:"id,omitempty"`
	InstructionPrompt *string                                   `json:"instructionPrompt,omitempty" xml:"instructionPrompt,omitempty"`
	Name              *string                                   `json:"name,omitempty" xml:"name,omitempty"`
	Source            *DescribeProjectsResponseBodyItemsSource  `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	// 扫描项目更新时间（RFC3339）
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
	Sast *bool `json:"sast,omitempty" xml:"sast,omitempty"`
	Sca  *bool `json:"sca,omitempty" xml:"sca,omitempty"`
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
