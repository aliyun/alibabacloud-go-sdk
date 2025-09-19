// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetsPropertyDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *GetAssetsPropertyDetailResponseBodyPageInfo) *GetAssetsPropertyDetailResponseBody
	GetPageInfo() *GetAssetsPropertyDetailResponseBodyPageInfo
	SetPropertys(v []*GetAssetsPropertyDetailResponseBodyPropertys) *GetAssetsPropertyDetailResponseBody
	GetPropertys() []*GetAssetsPropertyDetailResponseBodyPropertys
	SetRequestId(v string) *GetAssetsPropertyDetailResponseBody
	GetRequestId() *string
}

type GetAssetsPropertyDetailResponseBody struct {
	// The pagination information.
	PageInfo *GetAssetsPropertyDetailResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// An array that consists of the details about the asset fingerprints.
	Propertys []*GetAssetsPropertyDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C1AE3F3-18FA-4108-BBB9-AFA1A032****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAssetsPropertyDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAssetsPropertyDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetAssetsPropertyDetailResponseBody) GetPageInfo() *GetAssetsPropertyDetailResponseBodyPageInfo {
	return s.PageInfo
}

func (s *GetAssetsPropertyDetailResponseBody) GetPropertys() []*GetAssetsPropertyDetailResponseBodyPropertys {
	return s.Propertys
}

func (s *GetAssetsPropertyDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAssetsPropertyDetailResponseBody) SetPageInfo(v *GetAssetsPropertyDetailResponseBodyPageInfo) *GetAssetsPropertyDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetAssetsPropertyDetailResponseBody) SetPropertys(v []*GetAssetsPropertyDetailResponseBodyPropertys) *GetAssetsPropertyDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *GetAssetsPropertyDetailResponseBody) SetRequestId(v string) *GetAssetsPropertyDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAssetsPropertyDetailResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAssetsPropertyDetailResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAssetsPropertyDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetAssetsPropertyDetailResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *GetAssetsPropertyDetailResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAssetsPropertyDetailResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAssetsPropertyDetailResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetAssetsPropertyDetailResponseBodyPageInfo) SetCount(v int32) *GetAssetsPropertyDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPageInfo) SetCurrentPage(v int32) *GetAssetsPropertyDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPageInfo) SetPageSize(v int32) *GetAssetsPropertyDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPageInfo) SetTotalCount(v int32) *GetAssetsPropertyDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type GetAssetsPropertyDetailResponseBodyPropertys struct {
	// The name of the container.
	//
	// example:
	//
	// 5-rce_web_1
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The timestamp of the last fingerprint collection. Unit: milliseconds.
	//
	// example:
	//
	// 1649149566000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The domain name of the website.
	//
	// > This parameter is returned only when **Biz*	- is set to **web_server**.
	//
	// example:
	//
	// localhost
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	File     *string `json:"File,omitempty" xml:"File,omitempty"`
	// The path to the kernel module file.
	//
	// > This parameter is returned only when **Biz*	- is set to **lkm**.
	//
	// example:
	//
	// /lib/modules/4****
	Filepath *string `json:"Filepath,omitempty" xml:"Filepath,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// registry-vpc.cn-beijing.aliyuncs.com/acs/aliyun-ingress-controller****
	ImageName        *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstallationPath *string `json:"InstallationPath,omitempty" xml:"InstallationPath,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// i-hp35tftuh52wbp1g****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// hc-host-****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 47.42.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 100.104.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The IP address of the server.
	//
	// example:
	//
	// 47.42.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The listening protocol that the website uses.
	//
	// > This parameter is returned only when **Biz*	- is set to **web_server**.
	//
	// example:
	//
	// tcp
	ListenProtocol    *string `json:"ListenProtocol,omitempty" xml:"ListenProtocol,omitempty"`
	MiddlewareName    *string `json:"MiddlewareName,omitempty" xml:"MiddlewareName,omitempty"`
	MiddlewareVersion *string `json:"MiddlewareVersion,omitempty" xml:"MiddlewareVersion,omitempty"`
	ModelName         *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The name of the module.
	//
	// > This parameter is returned only when **Biz*	- is set to **lkm**.
	//
	// example:
	//
	// alihids
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The path. The value of this parameter varies based on the value of **Biz**.
	//
	// 	- If **Biz*	- is set to **web_server**, the value of this parameter indicates the path to the website root directory.
	//
	// 	- If **Biz*	- is set to **autorun**, the value of this parameter indicates the path to the startup item.
	//
	// example:
	//
	// /lib/systemd/system****
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The permissions on the root directory of the website.
	//
	// > This parameter is returned only when **Biz*	- is set to **web_server**.
	//
	// example:
	//
	// 755
	PathMode *string `json:"PathMode,omitempty" xml:"PathMode,omitempty"`
	// The process ID (PID) of the process that runs the website service.
	//
	// > This parameter is returned only when **Biz*	- is set to **web_server**.
	//
	// example:
	//
	// 813
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The port of the website.
	//
	// > This parameter is returned only when **Biz*	- is set to **web_server**.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The timestamp generated when the process was started. Unit: milliseconds.
	//
	// > This parameter is returned only when **Biz*	- is set to **web_server**.
	//
	// example:
	//
	// 1671186801000
	ProcessStarted *int64 `json:"ProcessStarted,omitempty" xml:"ProcessStarted,omitempty"`
	// The ID of the region in which the server resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the website.
	//
	// > This parameter is returned only when **Biz*	- is set to **web_server**.
	//
	// example:
	//
	// nginx
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The size of the kernel module.
	//
	// > This parameter is returned only when **Biz*	- is set to **lkm**.
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The number of times that the kernel module is referenced.
	//
	// > This parameter is returned only when **Biz*	- is set to **lkm**.
	//
	// example:
	//
	// 0
	UsedByCount *int32 `json:"UsedByCount,omitempty" xml:"UsedByCount,omitempty"`
	// The user who started the process of the website.
	//
	// > This parameter is returned only when **Biz*	- is set to **web_server**.
	//
	// example:
	//
	// root
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 162eb349-c2d9-4f8b-805c-75b43d4c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The path to the root directory of the website.
	//
	// > This parameter is returned only when **Biz*	- is set to **web_server**.
	//
	// example:
	//
	// /usr/share/nginx/html
	WebPath *string `json:"WebPath,omitempty" xml:"WebPath,omitempty"`
}

func (s GetAssetsPropertyDetailResponseBodyPropertys) String() string {
	return dara.Prettify(s)
}

func (s GetAssetsPropertyDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetContainerName() *string {
	return s.ContainerName
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetDomain() *string {
	return s.Domain
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetEndPoint() *string {
	return s.EndPoint
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetFile() *string {
	return s.File
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetFilepath() *string {
	return s.Filepath
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetImageName() *string {
	return s.ImageName
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetInstallationPath() *string {
	return s.InstallationPath
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetInternetIp() *string {
	return s.InternetIp
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetIp() *string {
	return s.Ip
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetListenProtocol() *string {
	return s.ListenProtocol
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetMiddlewareName() *string {
	return s.MiddlewareName
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetMiddlewareVersion() *string {
	return s.MiddlewareVersion
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetModelName() *string {
	return s.ModelName
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetPath() *string {
	return s.Path
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetPathMode() *string {
	return s.PathMode
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetPid() *string {
	return s.Pid
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetPort() *string {
	return s.Port
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetProcessStarted() *int64 {
	return s.ProcessStarted
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetServerType() *string {
	return s.ServerType
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetSize() *int32 {
	return s.Size
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetUsedByCount() *int32 {
	return s.UsedByCount
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetUser() *string {
	return s.User
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetUuid() *string {
	return s.Uuid
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) GetWebPath() *string {
	return s.WebPath
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetContainerName(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.ContainerName = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetDomain(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.Domain = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetEndPoint(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.EndPoint = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetFile(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.File = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetFilepath(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.Filepath = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetImageName(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.ImageName = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetInstallationPath(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.InstallationPath = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetInstanceId(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetInstanceName(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetInternetIp(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetIntranetIp(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetIp(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.Ip = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetListenProtocol(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.ListenProtocol = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetMiddlewareName(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.MiddlewareName = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetMiddlewareVersion(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.MiddlewareVersion = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetModelName(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.ModelName = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetModuleName(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.ModuleName = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetPath(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.Path = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetPathMode(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.PathMode = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetPid(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.Pid = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetPort(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.Port = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetProcessStarted(v int64) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.ProcessStarted = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetRegionId(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.RegionId = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetServerType(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.ServerType = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetSize(v int32) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.Size = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetUsedByCount(v int32) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.UsedByCount = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetUser(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.User = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetUuid(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) SetWebPath(v string) *GetAssetsPropertyDetailResponseBodyPropertys {
	s.WebPath = &v
	return s
}

func (s *GetAssetsPropertyDetailResponseBodyPropertys) Validate() error {
	return dara.Validate(s)
}
