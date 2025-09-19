// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectPluginStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListFileProtectPluginStatusResponseBodyData) *ListFileProtectPluginStatusResponseBody
	GetData() []*ListFileProtectPluginStatusResponseBodyData
	SetPageInfo(v *ListFileProtectPluginStatusResponseBodyPageInfo) *ListFileProtectPluginStatusResponseBody
	GetPageInfo() *ListFileProtectPluginStatusResponseBodyPageInfo
	SetRequestId(v string) *ListFileProtectPluginStatusResponseBody
	GetRequestId() *string
}

type ListFileProtectPluginStatusResponseBody struct {
	// The data returned if the call is successful.
	Data []*ListFileProtectPluginStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListFileProtectPluginStatusResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 60F289EC-BAA3-5DF1-8476-B3F05A14EBC2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFileProtectPluginStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectPluginStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileProtectPluginStatusResponseBody) GetData() []*ListFileProtectPluginStatusResponseBodyData {
	return s.Data
}

func (s *ListFileProtectPluginStatusResponseBody) GetPageInfo() *ListFileProtectPluginStatusResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListFileProtectPluginStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFileProtectPluginStatusResponseBody) SetData(v []*ListFileProtectPluginStatusResponseBodyData) *ListFileProtectPluginStatusResponseBody {
	s.Data = v
	return s
}

func (s *ListFileProtectPluginStatusResponseBody) SetPageInfo(v *ListFileProtectPluginStatusResponseBodyPageInfo) *ListFileProtectPluginStatusResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListFileProtectPluginStatusResponseBody) SetRequestId(v string) *ListFileProtectPluginStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFileProtectPluginStatusResponseBodyData struct {
	// The version of the Security Center agent.
	//
	// example:
	//
	// 00_41
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The returned code after you install the Security Center agent. Valid values:
	//
	// 1.  0: The installation is successful.
	//
	// 2.  \\-2: The kernel does not support the installation.
	//
	// example:
	//
	// -2
	InstallCode *string `json:"InstallCode,omitempty" xml:"InstallCode,omitempty"`
	// The returned message after you install the Security Center agent.
	//
	// example:
	//
	// driver file not exist
	InstallMessage *string `json:"InstallMessage,omitempty" xml:"InstallMessage,omitempty"`
	// Indicates whether the Security Center agent is installed.
	//
	// example:
	//
	// true
	Installed *bool `json:"Installed,omitempty" xml:"Installed,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// i-wz92q7m5hsbgfhdss***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address that is associated with the instance.
	//
	// example:
	//
	// 172.16.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address that is associated with the instance.
	//
	// example:
	//
	// 10.42.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// Indicates whether the Security Center agent is online. Valid value:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- **windows**: Windows
	//
	// 	- **linux**: Linux
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// Indicates whether the core file monitoring file is supported.
	//
	// example:
	//
	// true
	SupportFile *bool `json:"SupportFile,omitempty" xml:"SupportFile,omitempty"`
	// The UUID of the asset.
	//
	// example:
	//
	// 6690a46c-0edb-4663-a641-3629d1a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListFileProtectPluginStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectPluginStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFileProtectPluginStatusResponseBodyData) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *ListFileProtectPluginStatusResponseBodyData) GetInstallCode() *string {
	return s.InstallCode
}

func (s *ListFileProtectPluginStatusResponseBodyData) GetInstallMessage() *string {
	return s.InstallMessage
}

func (s *ListFileProtectPluginStatusResponseBodyData) GetInstalled() *bool {
	return s.Installed
}

func (s *ListFileProtectPluginStatusResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListFileProtectPluginStatusResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListFileProtectPluginStatusResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListFileProtectPluginStatusResponseBodyData) GetOnline() *bool {
	return s.Online
}

func (s *ListFileProtectPluginStatusResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *ListFileProtectPluginStatusResponseBodyData) GetSupportFile() *bool {
	return s.SupportFile
}

func (s *ListFileProtectPluginStatusResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *ListFileProtectPluginStatusResponseBodyData) SetClientVersion(v string) *ListFileProtectPluginStatusResponseBodyData {
	s.ClientVersion = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyData) SetInstallCode(v string) *ListFileProtectPluginStatusResponseBodyData {
	s.InstallCode = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyData) SetInstallMessage(v string) *ListFileProtectPluginStatusResponseBodyData {
	s.InstallMessage = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyData) SetInstalled(v bool) *ListFileProtectPluginStatusResponseBodyData {
	s.Installed = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyData) SetInstanceName(v string) *ListFileProtectPluginStatusResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyData) SetInternetIp(v string) *ListFileProtectPluginStatusResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyData) SetIntranetIp(v string) *ListFileProtectPluginStatusResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyData) SetOnline(v bool) *ListFileProtectPluginStatusResponseBodyData {
	s.Online = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyData) SetPlatform(v string) *ListFileProtectPluginStatusResponseBodyData {
	s.Platform = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyData) SetSupportFile(v bool) *ListFileProtectPluginStatusResponseBodyData {
	s.SupportFile = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyData) SetUuid(v string) *ListFileProtectPluginStatusResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListFileProtectPluginStatusResponseBodyPageInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileProtectPluginStatusResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectPluginStatusResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListFileProtectPluginStatusResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListFileProtectPluginStatusResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileProtectPluginStatusResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFileProtectPluginStatusResponseBodyPageInfo) SetCurrentPage(v int32) *ListFileProtectPluginStatusResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyPageInfo) SetPageSize(v int32) *ListFileProtectPluginStatusResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyPageInfo) SetTotalCount(v int32) *ListFileProtectPluginStatusResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListFileProtectPluginStatusResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
