// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSoftwaresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalPackages(v *ListSoftwaresResponseBodyAdditionalPackages) *ListSoftwaresResponseBody
	GetAdditionalPackages() *ListSoftwaresResponseBodyAdditionalPackages
	SetPageNumber(v string) *ListSoftwaresResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *ListSoftwaresResponseBody
	GetPageSize() *string
	SetRequestId(v string) *ListSoftwaresResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListSoftwaresResponseBody
	GetTotalCount() *string
}

type ListSoftwaresResponseBody struct {
	// The information about the software that can be installed in the cluster.
	AdditionalPackages *ListSoftwaresResponseBodyAdditionalPackages `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty" type:"Struct"`
	// The page number. Pages start from page 1. Default value: 1.
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
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSoftwaresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwaresResponseBody) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBody) GetAdditionalPackages() *ListSoftwaresResponseBodyAdditionalPackages {
	return s.AdditionalPackages
}

func (s *ListSoftwaresResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListSoftwaresResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *ListSoftwaresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSoftwaresResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListSoftwaresResponseBody) SetAdditionalPackages(v *ListSoftwaresResponseBodyAdditionalPackages) *ListSoftwaresResponseBody {
	s.AdditionalPackages = v
	return s
}

func (s *ListSoftwaresResponseBody) SetPageNumber(v string) *ListSoftwaresResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSoftwaresResponseBody) SetPageSize(v string) *ListSoftwaresResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSoftwaresResponseBody) SetRequestId(v string) *ListSoftwaresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSoftwaresResponseBody) SetTotalCount(v string) *ListSoftwaresResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSoftwaresResponseBody) Validate() error {
	if s.AdditionalPackages != nil {
		if err := s.AdditionalPackages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSoftwaresResponseBodyAdditionalPackages struct {
	AdditionalPackageInfos []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos `json:"AdditionalPackageInfos,omitempty" xml:"AdditionalPackageInfos,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBodyAdditionalPackages) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwaresResponseBodyAdditionalPackages) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodyAdditionalPackages) GetAdditionalPackageInfos() []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	return s.AdditionalPackageInfos
}

func (s *ListSoftwaresResponseBodyAdditionalPackages) SetAdditionalPackageInfos(v []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) *ListSoftwaresResponseBodyAdditionalPackages {
	s.AdditionalPackageInfos = v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackages) Validate() error {
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

type ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos struct {
	// The application category.
	//
	// example:
	//
	// NWP
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
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
	// The information about the software versions that can be installed in the cluster.
	Versions *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Struct"`
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GetCategory() *string {
	return s.Category
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GetDescription() *string {
	return s.Description
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GetIcon() *string {
	return s.Icon
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GetName() *string {
	return s.Name
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GetVersions() *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions {
	return s.Versions
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetCategory(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Category = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetDescription(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Description = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetIcon(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Icon = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetName(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Name = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetVersions(v *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Versions = v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) Validate() error {
	if s.Versions != nil {
		if err := s.Versions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions struct {
	VersionInfos []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos `json:"VersionInfos,omitempty" xml:"VersionInfos,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions) GetVersionInfos() []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos {
	return s.VersionInfos
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions) SetVersionInfos(v []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions {
	s.VersionInfos = v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions) Validate() error {
	if s.VersionInfos != nil {
		for _, item := range s.VersionInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos struct {
	// Indicates whether the version is the latest.
	//
	// example:
	//
	// false
	Latest *string `json:"Latest,omitempty" xml:"Latest,omitempty"`
	// The information about the supported OSs.
	SupportOs *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs `json:"SupportOs,omitempty" xml:"SupportOs,omitempty" type:"Struct"`
	// The software version.
	//
	// example:
	//
	// 2024.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) GetLatest() *string {
	return s.Latest
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) GetSupportOs() *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs {
	return s.SupportOs
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) GetVersion() *string {
	return s.Version
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) SetLatest(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos {
	s.Latest = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) SetSupportOs(v *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos {
	s.SupportOs = v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) SetVersion(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos {
	s.Version = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) Validate() error {
	if s.SupportOs != nil {
		if err := s.SupportOs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs struct {
	SupportOsInfos []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos `json:"SupportOsInfos,omitempty" xml:"SupportOsInfos,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs) GetSupportOsInfos() []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos {
	return s.SupportOsInfos
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs) SetSupportOsInfos(v []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs {
	s.SupportOsInfos = v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs) Validate() error {
	if s.SupportOsInfos != nil {
		for _, item := range s.SupportOsInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos struct {
	// The OS architecture. Valid values:
	//
	// 	- x86_64
	//
	// 	- arm64
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The image tag.
	//
	// example:
	//
	// Alibaba Cloud Linux  3.2104 LTS 64 bit ARM Edition
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) GetArchitecture() *string {
	return s.Architecture
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) GetOsTag() *string {
	return s.OsTag
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) SetArchitecture(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos {
	s.Architecture = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) SetOsTag(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos {
	s.OsTag = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) Validate() error {
	return dara.Validate(s)
}
