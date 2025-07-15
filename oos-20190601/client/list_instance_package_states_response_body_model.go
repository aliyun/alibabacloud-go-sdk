// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancePackageStatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *ListInstancePackageStatesResponseBody
	GetMaxResults() *string
	SetNextToken(v string) *ListInstancePackageStatesResponseBody
	GetNextToken() *string
	SetPackageStates(v []*ListInstancePackageStatesResponseBodyPackageStates) *ListInstancePackageStatesResponseBody
	GetPackageStates() []*ListInstancePackageStatesResponseBodyPackageStates
	SetRequestId(v string) *ListInstancePackageStatesResponseBody
	GetRequestId() *string
}

type ListInstancePackageStatesResponseBody struct {
	// Page size.
	//
	// example:
	//
	// 50
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token string for pagination.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctzxxxxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// List of extensions
	PackageStates []*ListInstancePackageStatesResponseBodyPackageStates `json:"PackageStates,omitempty" xml:"PackageStates,omitempty" type:"Repeated"`
	// ID of the request
	//
	// example:
	//
	// 1306108F-610C-40FD-AAD5-XXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancePackageStatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePackageStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancePackageStatesResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListInstancePackageStatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancePackageStatesResponseBody) GetPackageStates() []*ListInstancePackageStatesResponseBodyPackageStates {
	return s.PackageStates
}

func (s *ListInstancePackageStatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancePackageStatesResponseBody) SetMaxResults(v string) *ListInstancePackageStatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePackageStatesResponseBody) SetNextToken(v string) *ListInstancePackageStatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancePackageStatesResponseBody) SetPackageStates(v []*ListInstancePackageStatesResponseBodyPackageStates) *ListInstancePackageStatesResponseBody {
	s.PackageStates = v
	return s
}

func (s *ListInstancePackageStatesResponseBody) SetRequestId(v string) *ListInstancePackageStatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancePackageStatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstancePackageStatesResponseBodyPackageStates struct {
	// Description
	//
	// example:
	//
	// template description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Parameters
	//
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Publisher
	//
	// example:
	//
	// Alibaba Cloud
	Publisher *string `json:"Publisher,omitempty" xml:"Publisher,omitempty"`
	// Template type
	//
	// example:
	//
	// Package
	TemplateCategory *string `json:"TemplateCategory,omitempty" xml:"TemplateCategory,omitempty"`
	// Template ID
	//
	// example:
	//
	// 087b1e11072a40259f6fxxxxxxxxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template name.
	//
	// example:
	//
	// ACS-ECS-Docker
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template version number
	//
	// example:
	//
	// v3
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// Template version name
	//
	// example:
	//
	// fix bug
	TemplateVersionName *string `json:"TemplateVersionName,omitempty" xml:"TemplateVersionName,omitempty"`
	// Update time.
	//
	// example:
	//
	// 2024-05-04T11:17:28
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListInstancePackageStatesResponseBodyPackageStates) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePackageStatesResponseBodyPackageStates) GoString() string {
	return s.String()
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) GetDescription() *string {
	return s.Description
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) GetParameters() *string {
	return s.Parameters
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) GetPublisher() *string {
	return s.Publisher
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) GetTemplateCategory() *string {
	return s.TemplateCategory
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) GetTemplateVersionName() *string {
	return s.TemplateVersionName
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetDescription(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.Description = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetParameters(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.Parameters = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetPublisher(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.Publisher = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetTemplateCategory(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.TemplateCategory = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetTemplateId(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.TemplateId = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetTemplateName(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.TemplateName = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetTemplateVersion(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.TemplateVersion = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetTemplateVersionName(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.TemplateVersionName = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetUpdateTime(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.UpdateTime = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) Validate() error {
	return dara.Validate(s)
}
