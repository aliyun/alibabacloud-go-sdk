// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProhibitedTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListProhibitedTagsRequest
	GetCurrentPage() *int64
	SetName(v string) *ListProhibitedTagsRequest
	GetName() *string
	SetPageSize(v int64) *ListProhibitedTagsRequest
	GetPageSize() *int64
	SetPolicyId(v string) *ListProhibitedTagsRequest
	GetPolicyId() *string
	SetSoftwareId(v *ListProhibitedTagsRequestSoftwareId) *ListProhibitedTagsRequest
	GetSoftwareId() *ListProhibitedTagsRequestSoftwareId
	SetTagIds(v []*ListProhibitedTagsRequestTagIds) *ListProhibitedTagsRequest
	GetTagIds() []*ListProhibitedTagsRequestTagIds
}

type ListProhibitedTagsRequest struct {
	// The page number of the current page in a paged query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the prohibited software tag. Fuzzy match is supported. The name can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page in a paged query. Valid values: 1 to 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the software prohibition policy. You can obtain the value from the following operations:
	//
	// - [ListProhibitedPolicies](~~ListProhibitedPolicies~~): Lists software prohibition policies.
	//
	// - [CreateProhibitedPolicy](~~CreateProhibitedPolicy~~): Creates a software prohibition policy.
	//
	// example:
	//
	// pid-867ef4007c8a****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The unique identifier of the prohibited software.
	SoftwareId *ListProhibitedTagsRequestSoftwareId `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty" type:"Struct"`
	// The collection of prohibited software tag IDs. Duplicate values are not allowed. A maximum of 500 IDs can be specified.
	TagIds []*ListProhibitedTagsRequestTagIds `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListProhibitedTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedTagsRequest) GoString() string {
	return s.String()
}

func (s *ListProhibitedTagsRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListProhibitedTagsRequest) GetName() *string {
	return s.Name
}

func (s *ListProhibitedTagsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListProhibitedTagsRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListProhibitedTagsRequest) GetSoftwareId() *ListProhibitedTagsRequestSoftwareId {
	return s.SoftwareId
}

func (s *ListProhibitedTagsRequest) GetTagIds() []*ListProhibitedTagsRequestTagIds {
	return s.TagIds
}

func (s *ListProhibitedTagsRequest) SetCurrentPage(v int64) *ListProhibitedTagsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListProhibitedTagsRequest) SetName(v string) *ListProhibitedTagsRequest {
	s.Name = &v
	return s
}

func (s *ListProhibitedTagsRequest) SetPageSize(v int64) *ListProhibitedTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProhibitedTagsRequest) SetPolicyId(v string) *ListProhibitedTagsRequest {
	s.PolicyId = &v
	return s
}

func (s *ListProhibitedTagsRequest) SetSoftwareId(v *ListProhibitedTagsRequestSoftwareId) *ListProhibitedTagsRequest {
	s.SoftwareId = v
	return s
}

func (s *ListProhibitedTagsRequest) SetTagIds(v []*ListProhibitedTagsRequestTagIds) *ListProhibitedTagsRequest {
	s.TagIds = v
	return s
}

func (s *ListProhibitedTagsRequest) Validate() error {
	if s.SoftwareId != nil {
		if err := s.SoftwareId.Validate(); err != nil {
			return err
		}
	}
	if s.TagIds != nil {
		for _, item := range s.TagIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProhibitedTagsRequestSoftwareId struct {
	// Indicates whether the prohibited software is a system built-in entry. Valid values:
	//
	// - **true**: A system built-in prohibited software entry that is shared across all Alibaba Cloud accounts and cannot be modified or deleted.
	//
	// - **false**: Custom prohibited software under the current Alibaba Cloud account.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the prohibited software. You can obtain the value from the following operations:
	//
	// - [ListProhibitedSoftware](~~ListProhibitedSoftware~~): Lists prohibited software.
	//
	// - [CreateProhibitedSoftware](~~CreateProhibitedSoftware~~): Creates custom prohibited software.
	//
	// example:
	//
	// swb-83995ff2ae38****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
}

func (s ListProhibitedTagsRequestSoftwareId) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedTagsRequestSoftwareId) GoString() string {
	return s.String()
}

func (s *ListProhibitedTagsRequestSoftwareId) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListProhibitedTagsRequestSoftwareId) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *ListProhibitedTagsRequestSoftwareId) SetIsDefault(v bool) *ListProhibitedTagsRequestSoftwareId {
	s.IsDefault = &v
	return s
}

func (s *ListProhibitedTagsRequestSoftwareId) SetSoftwareId(v string) *ListProhibitedTagsRequestSoftwareId {
	s.SoftwareId = &v
	return s
}

func (s *ListProhibitedTagsRequestSoftwareId) Validate() error {
	return dara.Validate(s)
}

type ListProhibitedTagsRequestTagIds struct {
	// Indicates whether the prohibited software tag is a system built-in tag. Valid values:
	//
	// - **true**: A system built-in tag that is shared across all Alibaba Cloud accounts and cannot be modified or deleted.
	//
	// - **false**: A custom tag under the current Alibaba Cloud account.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the prohibited software tag. You can obtain the value from the following operations:
	//
	// - [ListProhibitedTags](~~ListProhibitedTags~~): Lists prohibited software tags.
	//
	// - [CreateProhibitedTag](~~CreateProhibitedTag~~): Creates a custom prohibited software tag.
	//
	// example:
	//
	// tag-3a5f8e50c396****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListProhibitedTagsRequestTagIds) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedTagsRequestTagIds) GoString() string {
	return s.String()
}

func (s *ListProhibitedTagsRequestTagIds) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListProhibitedTagsRequestTagIds) GetTagId() *string {
	return s.TagId
}

func (s *ListProhibitedTagsRequestTagIds) SetIsDefault(v bool) *ListProhibitedTagsRequestTagIds {
	s.IsDefault = &v
	return s
}

func (s *ListProhibitedTagsRequestTagIds) SetTagId(v string) *ListProhibitedTagsRequestTagIds {
	s.TagId = &v
	return s
}

func (s *ListProhibitedTagsRequestTagIds) Validate() error {
	return dara.Validate(s)
}
