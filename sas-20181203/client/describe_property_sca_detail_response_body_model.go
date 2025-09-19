// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyScaDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertyScaDetailResponseBodyPageInfo) *DescribePropertyScaDetailResponseBody
	GetPageInfo() *DescribePropertyScaDetailResponseBodyPageInfo
	SetPropertys(v []*DescribePropertyScaDetailResponseBodyPropertys) *DescribePropertyScaDetailResponseBody
	GetPropertys() []*DescribePropertyScaDetailResponseBodyPropertys
	SetRequestId(v string) *DescribePropertyScaDetailResponseBody
	GetRequestId() *string
}

type DescribePropertyScaDetailResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertyScaDetailResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The details about the asset fingerprints returned.
	Propertys []*DescribePropertyScaDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F9146867-16C8-4AAB-BB4FB8C2A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyScaDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaDetailResponseBody) GetPageInfo() *DescribePropertyScaDetailResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertyScaDetailResponseBody) GetPropertys() []*DescribePropertyScaDetailResponseBodyPropertys {
	return s.Propertys
}

func (s *DescribePropertyScaDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyScaDetailResponseBody) SetPageInfo(v *DescribePropertyScaDetailResponseBodyPageInfo) *DescribePropertyScaDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyScaDetailResponseBody) SetPropertys(v []*DescribePropertyScaDetailResponseBodyPropertys) *DescribePropertyScaDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertyScaDetailResponseBody) SetRequestId(v string) *DescribePropertyScaDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyScaDetailResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **10**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePropertyScaDetailResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertyScaDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyScaDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyScaDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyScaDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyScaDetailResponseBodyPropertys struct {
	// The type of the middleware, database, or web service. Valid values:
	//
	// 	- **system_service**: system service
	//
	// 	- **software_library**: software library
	//
	// 	- **docker_component**: container component
	//
	// 	- **database**: database
	//
	// 	- **web_container**: web container
	//
	// 	- **jar**: JAR package
	//
	// 	- **web_framework**: web framework
	//
	// example:
	//
	// software_library
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The display name of the type of the middleware, database, or web service . Valid values:
	//
	// 	- System service
	//
	// 	- Software library
	//
	// 	- Container component
	//
	// 	- Database
	//
	// 	- Web container
	//
	// 	- JAR package
	//
	// 	- Web framework
	//
	// example:
	//
	// System Service
	BizTypeDispaly *string `json:"BizTypeDispaly,omitempty" xml:"BizTypeDispaly,omitempty"`
	// The command line of the process.
	//
	// example:
	//
	// /sbin/dhclient -H iz2zeflhhbtk8gtxzt087az -1 -q -lf /var/lib/dhclient/dhclient--eth0.lease -pf /var/run/dhclient-eth0.pid eth0
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The path to the configuration file.
	//
	// example:
	//
	// /etc/my.cnf
	ConfigPath *string `json:"ConfigPath,omitempty" xml:"ConfigPath,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// 5-rce_web_1
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The latest collection timestamp, which indicates the last timestamp when Security Center collected the information about the middleware, database, or web service. Unit: milliseconds.
	//
	// example:
	//
	// 1597987834000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// registry-vpc.cn-beijing.aliyuncs.com/acs/aliyun-ingress-controller****
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the server on which the middleware, database, or web service is run.
	//
	// example:
	//
	// i-2zeclqj7ti****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server on which the middleware, database, or web service is run.
	//
	// example:
	//
	// Test01
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server on which the middleware, database, or web service is run.
	//
	// example:
	//
	// 47.42.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server on which the middleware, database, or web service is run.
	//
	// example:
	//
	// 192.210.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The public IP address of the server on which the middleware, database, or web service is run.
	//
	// example:
	//
	// 47.42.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The IP address that the process monitors.
	//
	// example:
	//
	// 0.0.XX.XX
	ListenIp *string `json:"ListenIp,omitempty" xml:"ListenIp,omitempty"`
	// The protocol of the traffic on which the process listens. Valid values:
	//
	// 	- **UDP**
	//
	// 	- **TCP**
	//
	// example:
	//
	// UDP
	ListenProtocol *string `json:"ListenProtocol,omitempty" xml:"ListenProtocol,omitempty"`
	// The listening status of the process. Valid values:
	//
	// 	- **NONE**: not listening
	//
	// 	- **LISTEN**: listening
	//
	// example:
	//
	// NONE
	ListenStatus *string `json:"ListenStatus,omitempty" xml:"ListenStatus,omitempty"`
	// The name of the middleware, database, or web service.
	//
	// example:
	//
	// openssl
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The path of the middleware, database, or web service.
	//
	// example:
	//
	// /usr/lib64/libssl.so.1.0.2k
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The PID.
	//
	// example:
	//
	// 756
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The name of the Kubernetes pod.
	//
	// example:
	//
	// myapp-pod
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// The port of the middleware, database, or web service.
	//
	// example:
	//
	// 68
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the parent process.
	//
	// example:
	//
	// 1
	Ppid *string `json:"Ppid,omitempty" xml:"Ppid,omitempty"`
	// The timestamp when the process starts. Unit: milliseconds.
	//
	// example:
	//
	// 1596539788
	ProcessStarted *int64 `json:"ProcessStarted,omitempty" xml:"ProcessStarted,omitempty"`
	// The name of the user who runs the process.
	//
	// example:
	//
	// root
	ProcessUser *string `json:"ProcessUser,omitempty" xml:"ProcessUser,omitempty"`
	// The version verification information about the middleware, database, or web service.
	//
	// example:
	//
	// /usr/lib64/libssl.so.1.0.2k
	Proof *string `json:"Proof,omitempty" xml:"Proof,omitempty"`
	// The version of the runtime environment.
	//
	// >  The value of this parameter can be the Java Development Kit (JDK) version of the runtime environment for a Java process.
	//
	// example:
	//
	// 1.8.0_144
	RuntimeEnvVersion *string `json:"RuntimeEnvVersion,omitempty" xml:"RuntimeEnvVersion,omitempty"`
	// The type of the middleware, database, or web service.
	//
	// example:
	//
	// library
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the server on which the middleware, database, or web service is run.
	//
	// example:
	//
	// uuid-02ebabe7-1c19-a****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The version of the middleware, database, or web service.
	//
	// example:
	//
	// 1.0.2k
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The web directory.
	//
	// example:
	//
	// /usr/share/nginx/html
	WebPath *string `json:"WebPath,omitempty" xml:"WebPath,omitempty"`
}

func (s DescribePropertyScaDetailResponseBodyPropertys) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetBizType() *string {
	return s.BizType
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetBizTypeDispaly() *string {
	return s.BizTypeDispaly
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetCmdline() *string {
	return s.Cmdline
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetConfigPath() *string {
	return s.ConfigPath
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetContainerName() *string {
	return s.ContainerName
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetImageName() *string {
	return s.ImageName
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetIp() *string {
	return s.Ip
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetListenIp() *string {
	return s.ListenIp
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetListenProtocol() *string {
	return s.ListenProtocol
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetListenStatus() *string {
	return s.ListenStatus
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetName() *string {
	return s.Name
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetPath() *string {
	return s.Path
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetPid() *string {
	return s.Pid
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetPodName() *string {
	return s.PodName
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetPort() *string {
	return s.Port
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetPpid() *string {
	return s.Ppid
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetProcessStarted() *int64 {
	return s.ProcessStarted
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetProcessUser() *string {
	return s.ProcessUser
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetProof() *string {
	return s.Proof
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetRuntimeEnvVersion() *string {
	return s.RuntimeEnvVersion
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetType() *string {
	return s.Type
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetVersion() *string {
	return s.Version
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) GetWebPath() *string {
	return s.WebPath
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetBizType(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.BizType = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetBizTypeDispaly(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.BizTypeDispaly = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetCmdline(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Cmdline = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetConfigPath(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ConfigPath = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetContainerName(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ContainerName = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertyScaDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetImageName(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ImageName = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetIp(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Ip = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetListenIp(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ListenIp = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetListenProtocol(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ListenProtocol = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetListenStatus(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ListenStatus = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetName(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Name = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetPath(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Path = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetPid(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Pid = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetPodName(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.PodName = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetPort(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Port = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetPpid(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Ppid = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetProcessStarted(v int64) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ProcessStarted = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetProcessUser(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ProcessUser = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetProof(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Proof = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetRuntimeEnvVersion(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.RuntimeEnvVersion = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetType(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Type = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetVersion(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Version = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetWebPath(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.WebPath = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) Validate() error {
	return dara.Validate(s)
}
