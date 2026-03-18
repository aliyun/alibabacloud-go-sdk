// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListPackagesResponseBodyData) *ListPackagesResponseBody
	GetData() *ListPackagesResponseBodyData
	SetRequestId(v string) *ListPackagesResponseBody
	GetRequestId() *string
}

type ListPackagesResponseBody struct {
	Data      *ListPackagesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPackagesResponseBody) GetData() *ListPackagesResponseBodyData {
	return s.Data
}

func (s *ListPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPackagesResponseBody) SetData(v *ListPackagesResponseBodyData) *ListPackagesResponseBody {
	s.Data = v
	return s
}

func (s *ListPackagesResponseBody) SetRequestId(v string) *ListPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPackagesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPackagesResponseBodyData struct {
	CreatedPackages   []*ListPackagesResponseBodyDataCreatedPackages   `json:"createdPackages,omitempty" xml:"createdPackages,omitempty" type:"Repeated"`
	InstalledPackages []*ListPackagesResponseBodyDataInstalledPackages `json:"installedPackages,omitempty" xml:"installedPackages,omitempty" type:"Repeated"`
}

func (s ListPackagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPackagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPackagesResponseBodyData) GetCreatedPackages() []*ListPackagesResponseBodyDataCreatedPackages {
	return s.CreatedPackages
}

func (s *ListPackagesResponseBodyData) GetInstalledPackages() []*ListPackagesResponseBodyDataInstalledPackages {
	return s.InstalledPackages
}

func (s *ListPackagesResponseBodyData) SetCreatedPackages(v []*ListPackagesResponseBodyDataCreatedPackages) *ListPackagesResponseBodyData {
	s.CreatedPackages = v
	return s
}

func (s *ListPackagesResponseBodyData) SetInstalledPackages(v []*ListPackagesResponseBodyDataInstalledPackages) *ListPackagesResponseBodyData {
	s.InstalledPackages = v
	return s
}

func (s *ListPackagesResponseBodyData) Validate() error {
	if s.CreatedPackages != nil {
		for _, item := range s.CreatedPackages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InstalledPackages != nil {
		for _, item := range s.InstalledPackages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPackagesResponseBodyDataCreatedPackages struct {
	CreateTime *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPackagesResponseBodyDataCreatedPackages) String() string {
	return dara.Prettify(s)
}

func (s ListPackagesResponseBodyDataCreatedPackages) GoString() string {
	return s.String()
}

func (s *ListPackagesResponseBodyDataCreatedPackages) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListPackagesResponseBodyDataCreatedPackages) GetName() *string {
	return s.Name
}

func (s *ListPackagesResponseBodyDataCreatedPackages) SetCreateTime(v int64) *ListPackagesResponseBodyDataCreatedPackages {
	s.CreateTime = &v
	return s
}

func (s *ListPackagesResponseBodyDataCreatedPackages) SetName(v string) *ListPackagesResponseBodyDataCreatedPackages {
	s.Name = &v
	return s
}

func (s *ListPackagesResponseBodyDataCreatedPackages) Validate() error {
	return dara.Validate(s)
}

type ListPackagesResponseBodyDataInstalledPackages struct {
	InstallTime   *int64  `json:"installTime,omitempty" xml:"installTime,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceProject *string `json:"sourceProject,omitempty" xml:"sourceProject,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListPackagesResponseBodyDataInstalledPackages) String() string {
	return dara.Prettify(s)
}

func (s ListPackagesResponseBodyDataInstalledPackages) GoString() string {
	return s.String()
}

func (s *ListPackagesResponseBodyDataInstalledPackages) GetInstallTime() *int64 {
	return s.InstallTime
}

func (s *ListPackagesResponseBodyDataInstalledPackages) GetName() *string {
	return s.Name
}

func (s *ListPackagesResponseBodyDataInstalledPackages) GetSourceProject() *string {
	return s.SourceProject
}

func (s *ListPackagesResponseBodyDataInstalledPackages) GetStatus() *string {
	return s.Status
}

func (s *ListPackagesResponseBodyDataInstalledPackages) SetInstallTime(v int64) *ListPackagesResponseBodyDataInstalledPackages {
	s.InstallTime = &v
	return s
}

func (s *ListPackagesResponseBodyDataInstalledPackages) SetName(v string) *ListPackagesResponseBodyDataInstalledPackages {
	s.Name = &v
	return s
}

func (s *ListPackagesResponseBodyDataInstalledPackages) SetSourceProject(v string) *ListPackagesResponseBodyDataInstalledPackages {
	s.SourceProject = &v
	return s
}

func (s *ListPackagesResponseBodyDataInstalledPackages) SetStatus(v string) *ListPackagesResponseBodyDataInstalledPackages {
	s.Status = &v
	return s
}

func (s *ListPackagesResponseBodyDataInstalledPackages) Validate() error {
	return dara.Validate(s)
}
