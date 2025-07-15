// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTemplateVersionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTemplateVersionsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTemplateVersionsRequest
	GetRegionId() *string
	SetShareType(v string) *ListTemplateVersionsRequest
	GetShareType() *string
	SetTemplateName(v string) *ListTemplateVersionsRequest
	GetTemplateName() *string
}

type ListTemplateVersionsRequest struct {
	// The number of entries per page. Valid values: 10 to 100
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// H8xj9c-398djs9-39ajd9asdjjJ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the template. Valid values: Private and Public.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// describeinstances
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListTemplateVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplateVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplateVersionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTemplateVersionsRequest) GetShareType() *string {
	return s.ShareType
}

func (s *ListTemplateVersionsRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListTemplateVersionsRequest) SetMaxResults(v int32) *ListTemplateVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetNextToken(v string) *ListTemplateVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetRegionId(v string) *ListTemplateVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetShareType(v string) *ListTemplateVersionsRequest {
	s.ShareType = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetTemplateName(v string) *ListTemplateVersionsRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplateVersionsRequest) Validate() error {
	return dara.Validate(s)
}
