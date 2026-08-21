// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProhibitedTagsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListProhibitedTagsShrinkRequest
	GetCurrentPage() *int64
	SetName(v string) *ListProhibitedTagsShrinkRequest
	GetName() *string
	SetPageSize(v int64) *ListProhibitedTagsShrinkRequest
	GetPageSize() *int64
	SetPolicyId(v string) *ListProhibitedTagsShrinkRequest
	GetPolicyId() *string
	SetSoftwareIdShrink(v string) *ListProhibitedTagsShrinkRequest
	GetSoftwareIdShrink() *string
	SetTagIds(v []*ListProhibitedTagsShrinkRequestTagIds) *ListProhibitedTagsShrinkRequest
	GetTagIds() []*ListProhibitedTagsShrinkRequestTagIds
}

type ListProhibitedTagsShrinkRequest struct {
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
	SoftwareIdShrink *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	// The collection of prohibited software tag IDs. Duplicate values are not allowed. A maximum of 500 IDs can be specified.
	TagIds []*ListProhibitedTagsShrinkRequestTagIds `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListProhibitedTagsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedTagsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProhibitedTagsShrinkRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListProhibitedTagsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListProhibitedTagsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListProhibitedTagsShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListProhibitedTagsShrinkRequest) GetSoftwareIdShrink() *string {
	return s.SoftwareIdShrink
}

func (s *ListProhibitedTagsShrinkRequest) GetTagIds() []*ListProhibitedTagsShrinkRequestTagIds {
	return s.TagIds
}

func (s *ListProhibitedTagsShrinkRequest) SetCurrentPage(v int64) *ListProhibitedTagsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListProhibitedTagsShrinkRequest) SetName(v string) *ListProhibitedTagsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListProhibitedTagsShrinkRequest) SetPageSize(v int64) *ListProhibitedTagsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProhibitedTagsShrinkRequest) SetPolicyId(v string) *ListProhibitedTagsShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *ListProhibitedTagsShrinkRequest) SetSoftwareIdShrink(v string) *ListProhibitedTagsShrinkRequest {
	s.SoftwareIdShrink = &v
	return s
}

func (s *ListProhibitedTagsShrinkRequest) SetTagIds(v []*ListProhibitedTagsShrinkRequestTagIds) *ListProhibitedTagsShrinkRequest {
	s.TagIds = v
	return s
}

func (s *ListProhibitedTagsShrinkRequest) Validate() error {
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

type ListProhibitedTagsShrinkRequestTagIds struct {
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

func (s ListProhibitedTagsShrinkRequestTagIds) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedTagsShrinkRequestTagIds) GoString() string {
	return s.String()
}

func (s *ListProhibitedTagsShrinkRequestTagIds) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListProhibitedTagsShrinkRequestTagIds) GetTagId() *string {
	return s.TagId
}

func (s *ListProhibitedTagsShrinkRequestTagIds) SetIsDefault(v bool) *ListProhibitedTagsShrinkRequestTagIds {
	s.IsDefault = &v
	return s
}

func (s *ListProhibitedTagsShrinkRequestTagIds) SetTagId(v string) *ListProhibitedTagsShrinkRequestTagIds {
	s.TagId = &v
	return s
}

func (s *ListProhibitedTagsShrinkRequestTagIds) Validate() error {
	return dara.Validate(s)
}
