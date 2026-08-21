// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProhibitedSoftwareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListProhibitedSoftwareRequest
	GetCurrentPage() *int64
	SetDeviceType(v string) *ListProhibitedSoftwareRequest
	GetDeviceType() *string
	SetName(v string) *ListProhibitedSoftwareRequest
	GetName() *string
	SetPageSize(v int64) *ListProhibitedSoftwareRequest
	GetPageSize() *int64
	SetPolicyId(v string) *ListProhibitedSoftwareRequest
	GetPolicyId() *string
	SetProcessName(v string) *ListProhibitedSoftwareRequest
	GetProcessName() *string
	SetSoftwareIds(v []*ListProhibitedSoftwareRequestSoftwareIds) *ListProhibitedSoftwareRequest
	GetSoftwareIds() []*ListProhibitedSoftwareRequestSoftwareIds
	SetTagId(v *ListProhibitedSoftwareRequestTagId) *ListProhibitedSoftwareRequest
	GetTagId() *ListProhibitedSoftwareRequestTagId
}

type ListProhibitedSoftwareRequest struct {
	// The page number of the current page in a paged query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The operating system type for which the prohibited software has configured processes.
	//
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The name of the prohibited software.
	//
	// example:
	//
	// Thunder
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
	// pid-5a1e8c3f7b09****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The process name.
	//
	// example:
	//
	// Thunder.exe
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The collection of prohibited software IDs. Duplicate values are not allowed.
	SoftwareIds []*ListProhibitedSoftwareRequestSoftwareIds `json:"SoftwareIds,omitempty" xml:"SoftwareIds,omitempty" type:"Repeated"`
	// The unique identifier of the prohibited software tag.
	TagId *ListProhibitedSoftwareRequestTagId `json:"TagId,omitempty" xml:"TagId,omitempty" type:"Struct"`
}

func (s ListProhibitedSoftwareRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedSoftwareRequest) GoString() string {
	return s.String()
}

func (s *ListProhibitedSoftwareRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListProhibitedSoftwareRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ListProhibitedSoftwareRequest) GetName() *string {
	return s.Name
}

func (s *ListProhibitedSoftwareRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListProhibitedSoftwareRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListProhibitedSoftwareRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *ListProhibitedSoftwareRequest) GetSoftwareIds() []*ListProhibitedSoftwareRequestSoftwareIds {
	return s.SoftwareIds
}

func (s *ListProhibitedSoftwareRequest) GetTagId() *ListProhibitedSoftwareRequestTagId {
	return s.TagId
}

func (s *ListProhibitedSoftwareRequest) SetCurrentPage(v int64) *ListProhibitedSoftwareRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListProhibitedSoftwareRequest) SetDeviceType(v string) *ListProhibitedSoftwareRequest {
	s.DeviceType = &v
	return s
}

func (s *ListProhibitedSoftwareRequest) SetName(v string) *ListProhibitedSoftwareRequest {
	s.Name = &v
	return s
}

func (s *ListProhibitedSoftwareRequest) SetPageSize(v int64) *ListProhibitedSoftwareRequest {
	s.PageSize = &v
	return s
}

func (s *ListProhibitedSoftwareRequest) SetPolicyId(v string) *ListProhibitedSoftwareRequest {
	s.PolicyId = &v
	return s
}

func (s *ListProhibitedSoftwareRequest) SetProcessName(v string) *ListProhibitedSoftwareRequest {
	s.ProcessName = &v
	return s
}

func (s *ListProhibitedSoftwareRequest) SetSoftwareIds(v []*ListProhibitedSoftwareRequestSoftwareIds) *ListProhibitedSoftwareRequest {
	s.SoftwareIds = v
	return s
}

func (s *ListProhibitedSoftwareRequest) SetTagId(v *ListProhibitedSoftwareRequestTagId) *ListProhibitedSoftwareRequest {
	s.TagId = v
	return s
}

func (s *ListProhibitedSoftwareRequest) Validate() error {
	if s.SoftwareIds != nil {
		for _, item := range s.SoftwareIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TagId != nil {
		if err := s.TagId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProhibitedSoftwareRequestSoftwareIds struct {
	// Indicates whether the prohibited software is a system built-in prohibited software. Valid values:
	//
	// - **true**: A system built-in prohibited software that is shared across all Alibaba Cloud accounts and cannot be modified or deleted.
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
	// swb-d9f669a09746****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
}

func (s ListProhibitedSoftwareRequestSoftwareIds) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedSoftwareRequestSoftwareIds) GoString() string {
	return s.String()
}

func (s *ListProhibitedSoftwareRequestSoftwareIds) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListProhibitedSoftwareRequestSoftwareIds) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *ListProhibitedSoftwareRequestSoftwareIds) SetIsDefault(v bool) *ListProhibitedSoftwareRequestSoftwareIds {
	s.IsDefault = &v
	return s
}

func (s *ListProhibitedSoftwareRequestSoftwareIds) SetSoftwareId(v string) *ListProhibitedSoftwareRequestSoftwareIds {
	s.SoftwareId = &v
	return s
}

func (s *ListProhibitedSoftwareRequestSoftwareIds) Validate() error {
	return dara.Validate(s)
}

type ListProhibitedSoftwareRequestTagId struct {
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
	// tag-7b2c9e4a1d8f****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListProhibitedSoftwareRequestTagId) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedSoftwareRequestTagId) GoString() string {
	return s.String()
}

func (s *ListProhibitedSoftwareRequestTagId) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListProhibitedSoftwareRequestTagId) GetTagId() *string {
	return s.TagId
}

func (s *ListProhibitedSoftwareRequestTagId) SetIsDefault(v bool) *ListProhibitedSoftwareRequestTagId {
	s.IsDefault = &v
	return s
}

func (s *ListProhibitedSoftwareRequestTagId) SetTagId(v string) *ListProhibitedSoftwareRequestTagId {
	s.TagId = &v
	return s
}

func (s *ListProhibitedSoftwareRequestTagId) Validate() error {
	return dara.Validate(s)
}
