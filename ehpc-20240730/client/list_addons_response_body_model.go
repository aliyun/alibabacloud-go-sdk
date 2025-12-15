// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*ListAddonsResponseBodyAddons) *ListAddonsResponseBody
	GetAddons() []*ListAddonsResponseBodyAddons
	SetPageNumber(v int32) *ListAddonsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAddonsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAddonsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAddonsResponseBody
	GetTotalCount() *int32
}

type ListAddonsResponseBody struct {
	// The information about the addons.
	Addons []*ListAddonsResponseBodyAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number of the returned page. Default value: 1
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBody) GetAddons() []*ListAddonsResponseBodyAddons {
	return s.Addons
}

func (s *ListAddonsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAddonsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAddonsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAddonsResponseBody) SetAddons(v []*ListAddonsResponseBodyAddons) *ListAddonsResponseBody {
	s.Addons = v
	return s
}

func (s *ListAddonsResponseBody) SetPageNumber(v int32) *ListAddonsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAddonsResponseBody) SetPageSize(v int32) *ListAddonsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAddonsResponseBody) SetRequestId(v string) *ListAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAddonsResponseBody) SetTotalCount(v int32) *ListAddonsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAddonsResponseBody) Validate() error {
	if s.Addons != nil {
		for _, item := range s.Addons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAddonsResponseBodyAddons struct {
	// The addon ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login-1.0-W4g****
	AddonId *string `json:"AddonId,omitempty" xml:"AddonId,omitempty"`
	// The addon description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the addon was installed.
	//
	// example:
	//
	// 2024-08-22 18:11:17
	InstallTime *string `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	// The addon label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The addon state.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAddonsResponseBodyAddons) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddons) GetAddonId() *string {
	return s.AddonId
}

func (s *ListAddonsResponseBodyAddons) GetDescription() *string {
	return s.Description
}

func (s *ListAddonsResponseBodyAddons) GetInstallTime() *string {
	return s.InstallTime
}

func (s *ListAddonsResponseBodyAddons) GetLabel() *string {
	return s.Label
}

func (s *ListAddonsResponseBodyAddons) GetName() *string {
	return s.Name
}

func (s *ListAddonsResponseBodyAddons) GetStatus() *string {
	return s.Status
}

func (s *ListAddonsResponseBodyAddons) GetVersion() *string {
	return s.Version
}

func (s *ListAddonsResponseBodyAddons) SetAddonId(v string) *ListAddonsResponseBodyAddons {
	s.AddonId = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetDescription(v string) *ListAddonsResponseBodyAddons {
	s.Description = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetInstallTime(v string) *ListAddonsResponseBodyAddons {
	s.InstallTime = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetLabel(v string) *ListAddonsResponseBodyAddons {
	s.Label = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetName(v string) *ListAddonsResponseBodyAddons {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetStatus(v string) *ListAddonsResponseBodyAddons {
	s.Status = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetVersion(v string) *ListAddonsResponseBodyAddons {
	s.Version = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) Validate() error {
	return dara.Validate(s)
}
