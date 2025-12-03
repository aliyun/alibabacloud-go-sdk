// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTestcaseListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v string) *GetTestcaseListRequest
	GetConditions() *string
	SetDirectoryIdentifier(v string) *GetTestcaseListRequest
	GetDirectoryIdentifier() *string
	SetMaxResult(v string) *GetTestcaseListRequest
	GetMaxResult() *string
	SetNextToken(v string) *GetTestcaseListRequest
	GetNextToken() *string
	SetSpaceIdentifier(v string) *GetTestcaseListRequest
	GetSpaceIdentifier() *string
}

type GetTestcaseListRequest struct {
	// example:
	//
	// {\\"conditionGroups\\":[[{\\"fieldIdentifier\\":\\"status\\",\\"operator\\":\\"CONTAINS\\",\\"value\\":[\\"cc961a18adf770c6e423ccc5\\"],\\"toValue\\":null,,\\"className\\":\\"status\\",\\"format\\":\\"list\\"}]]}
	Conditions *string `json:"conditions,omitempty" xml:"conditions,omitempty"`
	// example:
	//
	// e27b8eace6501ce51cf5d56784
	DirectoryIdentifier *string `json:"directoryIdentifier,omitempty" xml:"directoryIdentifier,omitempty"`
	// example:
	//
	// 20
	MaxResult *string `json:"maxResult,omitempty" xml:"maxResult,omitempty"`
	// example:
	//
	// 2591861102250c4522380b33a6
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6d5984c7d92b23fa53d4743c37
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
}

func (s GetTestcaseListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTestcaseListRequest) GoString() string {
	return s.String()
}

func (s *GetTestcaseListRequest) GetConditions() *string {
	return s.Conditions
}

func (s *GetTestcaseListRequest) GetDirectoryIdentifier() *string {
	return s.DirectoryIdentifier
}

func (s *GetTestcaseListRequest) GetMaxResult() *string {
	return s.MaxResult
}

func (s *GetTestcaseListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetTestcaseListRequest) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *GetTestcaseListRequest) SetConditions(v string) *GetTestcaseListRequest {
	s.Conditions = &v
	return s
}

func (s *GetTestcaseListRequest) SetDirectoryIdentifier(v string) *GetTestcaseListRequest {
	s.DirectoryIdentifier = &v
	return s
}

func (s *GetTestcaseListRequest) SetMaxResult(v string) *GetTestcaseListRequest {
	s.MaxResult = &v
	return s
}

func (s *GetTestcaseListRequest) SetNextToken(v string) *GetTestcaseListRequest {
	s.NextToken = &v
	return s
}

func (s *GetTestcaseListRequest) SetSpaceIdentifier(v string) *GetTestcaseListRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *GetTestcaseListRequest) Validate() error {
	return dara.Validate(s)
}
