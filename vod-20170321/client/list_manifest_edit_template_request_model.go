// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManifestEditTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListManifestEditTemplateRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListManifestEditTemplateRequest
	GetNextToken() *string
}

type ListManifestEditTemplateRequest struct {
	// The maximum number of entries to return per page. Default value: **30**. Valid values: **1*	- to **200**.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination cursor. You do not need to specify this parameter for the first call. The value is a string in the format of page_N, where N is the page number. Pass the NextToken value returned in the previous response to retrieve the next page. If no NextToken is returned in the response, all data has been retrieved.
	//
	// example:
	//
	// page_2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListManifestEditTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListManifestEditTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListManifestEditTemplateRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListManifestEditTemplateRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListManifestEditTemplateRequest) SetMaxResults(v int32) *ListManifestEditTemplateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListManifestEditTemplateRequest) SetNextToken(v string) *ListManifestEditTemplateRequest {
	s.NextToken = &v
	return s
}

func (s *ListManifestEditTemplateRequest) Validate() error {
	return dara.Validate(s)
}
