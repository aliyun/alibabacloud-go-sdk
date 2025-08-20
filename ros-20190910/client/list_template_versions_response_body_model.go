// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTemplateVersionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTemplateVersionsResponseBody
	GetRequestId() *string
	SetVersions(v []*ListTemplateVersionsResponseBodyVersions) *ListTemplateVersionsResponseBody
	GetVersions() []*ListTemplateVersionsResponseBodyVersions
}

type ListTemplateVersionsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The versions.
	Versions []*ListTemplateVersionsResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListTemplateVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplateVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplateVersionsResponseBody) GetVersions() []*ListTemplateVersionsResponseBodyVersions {
	return s.Versions
}

func (s *ListTemplateVersionsResponseBody) SetNextToken(v string) *ListTemplateVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetRequestId(v string) *ListTemplateVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetVersions(v []*ListTemplateVersionsResponseBodyVersions) *ListTemplateVersionsResponseBody {
	s.Versions = v
	return s
}

func (s *ListTemplateVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTemplateVersionsResponseBodyVersions struct {
	// The time when the version was created.
	//
	// example:
	//
	// 2020-02-27T07:47:47
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the version.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The template ID. This parameter applies to shared and private templates. For a shared template, the template ID is the same as the Alibaba Cloud Resource Name (ARN) of the template.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name that corresponds to the specified version.
	//
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version number.
	//
	// For a shared template, this parameter is returned only if VersionOption is set to AllVersions.
	//
	// Valid values: v1 to v100.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the version was last updated.
	//
	// example:
	//
	// 2020-02-27T07:47:47
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTemplateVersionsResponseBodyVersions) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateVersionsResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponseBodyVersions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTemplateVersionsResponseBodyVersions) GetDescription() *string {
	return s.Description
}

func (s *ListTemplateVersionsResponseBodyVersions) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTemplateVersionsResponseBodyVersions) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListTemplateVersionsResponseBodyVersions) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListTemplateVersionsResponseBodyVersions) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListTemplateVersionsResponseBodyVersions) SetCreateTime(v string) *ListTemplateVersionsResponseBodyVersions {
	s.CreateTime = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetDescription(v string) *ListTemplateVersionsResponseBodyVersions {
	s.Description = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetTemplateId(v string) *ListTemplateVersionsResponseBodyVersions {
	s.TemplateId = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetTemplateName(v string) *ListTemplateVersionsResponseBodyVersions {
	s.TemplateName = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetTemplateVersion(v string) *ListTemplateVersionsResponseBodyVersions {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetUpdateTime(v string) *ListTemplateVersionsResponseBodyVersions {
	s.UpdateTime = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) Validate() error {
	return dara.Validate(s)
}
