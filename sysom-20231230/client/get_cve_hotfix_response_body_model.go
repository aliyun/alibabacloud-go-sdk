// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCveHotfixResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCveHotfixResponseBody
	GetCode() *string
	SetData(v []*GetCveHotfixResponseBodyData) *GetCveHotfixResponseBody
	GetData() []*GetCveHotfixResponseBodyData
	SetMessage(v string) *GetCveHotfixResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCveHotfixResponseBody
	GetRequestId() *string
}

type GetCveHotfixResponseBody struct {
	// The status code.
	//
	// - If `code == Success`, the authorization is successful.
	//
	// - Other status codes indicate that the authorization failed. If the authorization failed, check the `message` field for the detailed fault information.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data []*GetCveHotfixResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// “”
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// example-request-id
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCveHotfixResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCveHotfixResponseBody) GoString() string {
	return s.String()
}

func (s *GetCveHotfixResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCveHotfixResponseBody) GetData() []*GetCveHotfixResponseBodyData {
	return s.Data
}

func (s *GetCveHotfixResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCveHotfixResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCveHotfixResponseBody) SetCode(v string) *GetCveHotfixResponseBody {
	s.Code = &v
	return s
}

func (s *GetCveHotfixResponseBody) SetData(v []*GetCveHotfixResponseBodyData) *GetCveHotfixResponseBody {
	s.Data = v
	return s
}

func (s *GetCveHotfixResponseBody) SetMessage(v string) *GetCveHotfixResponseBody {
	s.Message = &v
	return s
}

func (s *GetCveHotfixResponseBody) SetRequestId(v string) *GetCveHotfixResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCveHotfixResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCveHotfixResponseBodyData struct {
	// The CVE ID.
	//
	// example:
	//
	// CVE-2026-12345
	CveId *string `json:"cveId,omitempty" xml:"cveId,omitempty"`
	// The hotfix information for the CVE.
	//
	// if can be null:
	// true
	Hotfix *GetCveHotfixResponseBodyDataHotfix `json:"hotfix,omitempty" xml:"hotfix,omitempty" type:"Struct"`
}

func (s GetCveHotfixResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCveHotfixResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCveHotfixResponseBodyData) GetCveId() *string {
	return s.CveId
}

func (s *GetCveHotfixResponseBodyData) GetHotfix() *GetCveHotfixResponseBodyDataHotfix {
	return s.Hotfix
}

func (s *GetCveHotfixResponseBodyData) SetCveId(v string) *GetCveHotfixResponseBodyData {
	s.CveId = &v
	return s
}

func (s *GetCveHotfixResponseBodyData) SetHotfix(v *GetCveHotfixResponseBodyDataHotfix) *GetCveHotfixResponseBodyData {
	s.Hotfix = v
	return s
}

func (s *GetCveHotfixResponseBodyData) Validate() error {
	if s.Hotfix != nil {
		if err := s.Hotfix.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCveHotfixResponseBodyDataHotfix struct {
	// The hotfix description.
	//
	// example:
	//
	// Fixes the kernel issue corresponding to the CVE
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The hotfix name.
	//
	// example:
	//
	// 12345678
	HotfixName *string `json:"hotfixName,omitempty" xml:"hotfixName,omitempty"`
	// The hotfix module name.
	//
	// example:
	//
	// kpatch_example
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// The list of published packages that meet the kernel requirements.
	Packages []*GetCveHotfixResponseBodyDataHotfixPackages `json:"packages,omitempty" xml:"packages,omitempty" type:"Repeated"`
	// The hotfix version.
	//
	// example:
	//
	// 1.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetCveHotfixResponseBodyDataHotfix) String() string {
	return dara.Prettify(s)
}

func (s GetCveHotfixResponseBodyDataHotfix) GoString() string {
	return s.String()
}

func (s *GetCveHotfixResponseBodyDataHotfix) GetDescription() *string {
	return s.Description
}

func (s *GetCveHotfixResponseBodyDataHotfix) GetHotfixName() *string {
	return s.HotfixName
}

func (s *GetCveHotfixResponseBodyDataHotfix) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetCveHotfixResponseBodyDataHotfix) GetPackages() []*GetCveHotfixResponseBodyDataHotfixPackages {
	return s.Packages
}

func (s *GetCveHotfixResponseBodyDataHotfix) GetVersion() *string {
	return s.Version
}

func (s *GetCveHotfixResponseBodyDataHotfix) SetDescription(v string) *GetCveHotfixResponseBodyDataHotfix {
	s.Description = &v
	return s
}

func (s *GetCveHotfixResponseBodyDataHotfix) SetHotfixName(v string) *GetCveHotfixResponseBodyDataHotfix {
	s.HotfixName = &v
	return s
}

func (s *GetCveHotfixResponseBodyDataHotfix) SetModuleName(v string) *GetCveHotfixResponseBodyDataHotfix {
	s.ModuleName = &v
	return s
}

func (s *GetCveHotfixResponseBodyDataHotfix) SetPackages(v []*GetCveHotfixResponseBodyDataHotfixPackages) *GetCveHotfixResponseBodyDataHotfix {
	s.Packages = v
	return s
}

func (s *GetCveHotfixResponseBodyDataHotfix) SetVersion(v string) *GetCveHotfixResponseBodyDataHotfix {
	s.Version = &v
	return s
}

func (s *GetCveHotfixResponseBodyDataHotfix) Validate() error {
	if s.Packages != nil {
		for _, item := range s.Packages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCveHotfixResponseBodyDataHotfixPackages struct {
	// The architecture.
	//
	// example:
	//
	// x86_64
	Arch *string `json:"arch,omitempty" xml:"arch,omitempty"`
	// The distribution. This value is null if the distribution is not registered.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// alinux3
	Distro *string `json:"distro,omitempty" xml:"distro,omitempty"`
	// The compatible full kernel version.
	//
	// example:
	//
	// 5.10.134-18.al8.x86_64
	KernelVersion *string `json:"kernelVersion,omitempty" xml:"kernelVersion,omitempty"`
	// The RPM package identifier.
	//
	// example:
	//
	// kernel-hotfix-example-1.0-1.x86_64
	RpmNevra *string `json:"rpmNevra,omitempty" xml:"rpmNevra,omitempty"`
}

func (s GetCveHotfixResponseBodyDataHotfixPackages) String() string {
	return dara.Prettify(s)
}

func (s GetCveHotfixResponseBodyDataHotfixPackages) GoString() string {
	return s.String()
}

func (s *GetCveHotfixResponseBodyDataHotfixPackages) GetArch() *string {
	return s.Arch
}

func (s *GetCveHotfixResponseBodyDataHotfixPackages) GetDistro() *string {
	return s.Distro
}

func (s *GetCveHotfixResponseBodyDataHotfixPackages) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *GetCveHotfixResponseBodyDataHotfixPackages) GetRpmNevra() *string {
	return s.RpmNevra
}

func (s *GetCveHotfixResponseBodyDataHotfixPackages) SetArch(v string) *GetCveHotfixResponseBodyDataHotfixPackages {
	s.Arch = &v
	return s
}

func (s *GetCveHotfixResponseBodyDataHotfixPackages) SetDistro(v string) *GetCveHotfixResponseBodyDataHotfixPackages {
	s.Distro = &v
	return s
}

func (s *GetCveHotfixResponseBodyDataHotfixPackages) SetKernelVersion(v string) *GetCveHotfixResponseBodyDataHotfixPackages {
	s.KernelVersion = &v
	return s
}

func (s *GetCveHotfixResponseBodyDataHotfixPackages) SetRpmNevra(v string) *GetCveHotfixResponseBodyDataHotfixPackages {
	s.RpmNevra = &v
	return s
}

func (s *GetCveHotfixResponseBodyDataHotfixPackages) Validate() error {
	return dara.Validate(s)
}
