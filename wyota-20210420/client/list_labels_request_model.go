// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelContent(v string) *ListLabelsRequest
	GetLabelContent() *string
	SetLabelId(v string) *ListLabelsRequest
	GetLabelId() *string
	SetMaxResults(v int32) *ListLabelsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListLabelsRequest
	GetNextToken() *string
}

type ListLabelsRequest struct {
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLabelsRequest) GoString() string {
	return s.String()
}

func (s *ListLabelsRequest) GetLabelContent() *string {
	return s.LabelContent
}

func (s *ListLabelsRequest) GetLabelId() *string {
	return s.LabelId
}

func (s *ListLabelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLabelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLabelsRequest) SetLabelContent(v string) *ListLabelsRequest {
	s.LabelContent = &v
	return s
}

func (s *ListLabelsRequest) SetLabelId(v string) *ListLabelsRequest {
	s.LabelId = &v
	return s
}

func (s *ListLabelsRequest) SetMaxResults(v int32) *ListLabelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLabelsRequest) SetNextToken(v string) *ListLabelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListLabelsRequest) Validate() error {
	return dara.Validate(s)
}
