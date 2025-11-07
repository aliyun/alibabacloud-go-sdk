// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTemplateVersionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTemplateVersionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTemplateVersionsResponseBody
	GetRequestId() *string
	SetTemplateVersions(v []*ListTemplateVersionsResponseBodyTemplateVersions) *ListTemplateVersionsResponseBody
	GetTemplateVersions() []*ListTemplateVersionsResponseBodyTemplateVersions
}

type ListTemplateVersionsResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// NJSNDKLJS-SJKJDO090k30-JSDK232
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E6CD612B-5889-4F1A-823F-8A4029E46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The versions of the template.
	TemplateVersions []*ListTemplateVersionsResponseBodyTemplateVersions `json:"TemplateVersions,omitempty" xml:"TemplateVersions,omitempty" type:"Repeated"`
}

func (s ListTemplateVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplateVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplateVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplateVersionsResponseBody) GetTemplateVersions() []*ListTemplateVersionsResponseBodyTemplateVersions {
	return s.TemplateVersions
}

func (s *ListTemplateVersionsResponseBody) SetMaxResults(v int32) *ListTemplateVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetNextToken(v string) *ListTemplateVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetRequestId(v string) *ListTemplateVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetTemplateVersions(v []*ListTemplateVersionsResponseBodyTemplateVersions) *ListTemplateVersionsResponseBody {
	s.TemplateVersions = v
	return s
}

func (s *ListTemplateVersionsResponseBody) Validate() error {
	if s.TemplateVersions != nil {
		for _, item := range s.TemplateVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTemplateVersionsResponseBodyTemplateVersions struct {
	// The description of the version.
	//
	// example:
	//
	// Detach the eip from the special instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The format of the template content. Valid values: YAML and JSON.
	//
	// example:
	//
	// YAML
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The number of the version.
	//
	// example:
	//
	// v2
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The user who last updated the version.
	//
	// example:
	//
	// foo
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the version was last updated.
	//
	// example:
	//
	// 2020-05-19T06:05:41Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The name of the version.
	//
	// example:
	//
	// baz
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListTemplateVersionsResponseBodyTemplateVersions) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateVersionsResponseBodyTemplateVersions) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) GetDescription() *string {
	return s.Description
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) GetTemplateFormat() *string {
	return s.TemplateFormat
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) GetVersionName() *string {
	return s.VersionName
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetDescription(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.Description = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetTemplateFormat(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.TemplateFormat = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetTemplateVersion(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetUpdatedBy(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.UpdatedBy = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetUpdatedDate(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.UpdatedDate = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetVersionName(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.VersionName = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) Validate() error {
	return dara.Validate(s)
}
