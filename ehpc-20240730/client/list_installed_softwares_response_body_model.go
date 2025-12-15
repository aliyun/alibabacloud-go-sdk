// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstalledSoftwaresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalPackages(v *ListInstalledSoftwaresResponseBodyAdditionalPackages) *ListInstalledSoftwaresResponseBody
	GetAdditionalPackages() *ListInstalledSoftwaresResponseBodyAdditionalPackages
	SetPageNumber(v string) *ListInstalledSoftwaresResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *ListInstalledSoftwaresResponseBody
	GetPageSize() *string
	SetRequestId(v string) *ListInstalledSoftwaresResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListInstalledSoftwaresResponseBody
	GetTotalCount() *string
}

type ListInstalledSoftwaresResponseBody struct {
	// The list of installed software.
	AdditionalPackages *ListInstalledSoftwaresResponseBodyAdditionalPackages `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstalledSoftwaresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstalledSoftwaresResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwaresResponseBody) GetAdditionalPackages() *ListInstalledSoftwaresResponseBodyAdditionalPackages {
	return s.AdditionalPackages
}

func (s *ListInstalledSoftwaresResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListInstalledSoftwaresResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *ListInstalledSoftwaresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstalledSoftwaresResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListInstalledSoftwaresResponseBody) SetAdditionalPackages(v *ListInstalledSoftwaresResponseBodyAdditionalPackages) *ListInstalledSoftwaresResponseBody {
	s.AdditionalPackages = v
	return s
}

func (s *ListInstalledSoftwaresResponseBody) SetPageNumber(v string) *ListInstalledSoftwaresResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBody) SetPageSize(v string) *ListInstalledSoftwaresResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBody) SetRequestId(v string) *ListInstalledSoftwaresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBody) SetTotalCount(v string) *ListInstalledSoftwaresResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBody) Validate() error {
	if s.AdditionalPackages != nil {
		if err := s.AdditionalPackages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstalledSoftwaresResponseBodyAdditionalPackages struct {
	AdditionalPackageInfos []*ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos `json:"AdditionalPackageInfos,omitempty" xml:"AdditionalPackageInfos,omitempty" type:"Repeated"`
}

func (s ListInstalledSoftwaresResponseBodyAdditionalPackages) String() string {
	return dara.Prettify(s)
}

func (s ListInstalledSoftwaresResponseBodyAdditionalPackages) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackages) GetAdditionalPackageInfos() []*ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	return s.AdditionalPackageInfos
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackages) SetAdditionalPackageInfos(v []*ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) *ListInstalledSoftwaresResponseBodyAdditionalPackages {
	s.AdditionalPackageInfos = v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackages) Validate() error {
	if s.AdditionalPackageInfos != nil {
		for _, item := range s.AdditionalPackageInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos struct {
	// The category into which the software falls.
	//
	// example:
	//
	// NWP
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the software was installed.
	//
	// example:
	//
	// 2024-03-05 18:24:08
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The software description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the software icon.
	//
	// example:
	//
	// https://gw.alicdn.com/imgextra/i2/O1CN01FIkxZ81LmE0fvrAyR_!!6000000001341-55-tps-6349-1603.svg
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The software name.
	//
	// example:
	//
	// gromacs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The installation status of the software.
	//
	// Valid values:
	//
	// 	- Installed
	//
	// 	- Uninstalled
	//
	// 	- Installing
	//
	// 	- Exception
	//
	// example:
	//
	// Installed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The software version.
	//
	// example:
	//
	// 2024.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) String() string {
	return dara.Prettify(s)
}

func (s ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GetCategory() *string {
	return s.Category
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GetDescription() *string {
	return s.Description
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GetIcon() *string {
	return s.Icon
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GetName() *string {
	return s.Name
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GetStatus() *string {
	return s.Status
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GetVersion() *string {
	return s.Version
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetCategory(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Category = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetCreateTime(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.CreateTime = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetDescription(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Description = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetIcon(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Icon = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetName(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Name = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetStatus(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Status = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetVersion(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Version = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) Validate() error {
	return dara.Validate(s)
}
