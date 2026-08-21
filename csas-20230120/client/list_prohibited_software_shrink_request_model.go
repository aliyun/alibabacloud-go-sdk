// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProhibitedSoftwareShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListProhibitedSoftwareShrinkRequest
	GetCurrentPage() *int64
	SetDeviceType(v string) *ListProhibitedSoftwareShrinkRequest
	GetDeviceType() *string
	SetName(v string) *ListProhibitedSoftwareShrinkRequest
	GetName() *string
	SetPageSize(v int64) *ListProhibitedSoftwareShrinkRequest
	GetPageSize() *int64
	SetPolicyId(v string) *ListProhibitedSoftwareShrinkRequest
	GetPolicyId() *string
	SetProcessName(v string) *ListProhibitedSoftwareShrinkRequest
	GetProcessName() *string
	SetSoftwareIds(v []*ListProhibitedSoftwareShrinkRequestSoftwareIds) *ListProhibitedSoftwareShrinkRequest
	GetSoftwareIds() []*ListProhibitedSoftwareShrinkRequestSoftwareIds
	SetTagIdShrink(v string) *ListProhibitedSoftwareShrinkRequest
	GetTagIdShrink() *string
}

type ListProhibitedSoftwareShrinkRequest struct {
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
	SoftwareIds []*ListProhibitedSoftwareShrinkRequestSoftwareIds `json:"SoftwareIds,omitempty" xml:"SoftwareIds,omitempty" type:"Repeated"`
	// The unique identifier of the prohibited software tag.
	TagIdShrink *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListProhibitedSoftwareShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedSoftwareShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProhibitedSoftwareShrinkRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListProhibitedSoftwareShrinkRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ListProhibitedSoftwareShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListProhibitedSoftwareShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListProhibitedSoftwareShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListProhibitedSoftwareShrinkRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *ListProhibitedSoftwareShrinkRequest) GetSoftwareIds() []*ListProhibitedSoftwareShrinkRequestSoftwareIds {
	return s.SoftwareIds
}

func (s *ListProhibitedSoftwareShrinkRequest) GetTagIdShrink() *string {
	return s.TagIdShrink
}

func (s *ListProhibitedSoftwareShrinkRequest) SetCurrentPage(v int64) *ListProhibitedSoftwareShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListProhibitedSoftwareShrinkRequest) SetDeviceType(v string) *ListProhibitedSoftwareShrinkRequest {
	s.DeviceType = &v
	return s
}

func (s *ListProhibitedSoftwareShrinkRequest) SetName(v string) *ListProhibitedSoftwareShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListProhibitedSoftwareShrinkRequest) SetPageSize(v int64) *ListProhibitedSoftwareShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProhibitedSoftwareShrinkRequest) SetPolicyId(v string) *ListProhibitedSoftwareShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *ListProhibitedSoftwareShrinkRequest) SetProcessName(v string) *ListProhibitedSoftwareShrinkRequest {
	s.ProcessName = &v
	return s
}

func (s *ListProhibitedSoftwareShrinkRequest) SetSoftwareIds(v []*ListProhibitedSoftwareShrinkRequestSoftwareIds) *ListProhibitedSoftwareShrinkRequest {
	s.SoftwareIds = v
	return s
}

func (s *ListProhibitedSoftwareShrinkRequest) SetTagIdShrink(v string) *ListProhibitedSoftwareShrinkRequest {
	s.TagIdShrink = &v
	return s
}

func (s *ListProhibitedSoftwareShrinkRequest) Validate() error {
	if s.SoftwareIds != nil {
		for _, item := range s.SoftwareIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProhibitedSoftwareShrinkRequestSoftwareIds struct {
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

func (s ListProhibitedSoftwareShrinkRequestSoftwareIds) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedSoftwareShrinkRequestSoftwareIds) GoString() string {
	return s.String()
}

func (s *ListProhibitedSoftwareShrinkRequestSoftwareIds) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListProhibitedSoftwareShrinkRequestSoftwareIds) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *ListProhibitedSoftwareShrinkRequestSoftwareIds) SetIsDefault(v bool) *ListProhibitedSoftwareShrinkRequestSoftwareIds {
	s.IsDefault = &v
	return s
}

func (s *ListProhibitedSoftwareShrinkRequestSoftwareIds) SetSoftwareId(v string) *ListProhibitedSoftwareShrinkRequestSoftwareIds {
	s.SoftwareId = &v
	return s
}

func (s *ListProhibitedSoftwareShrinkRequestSoftwareIds) Validate() error {
	return dara.Validate(s)
}
