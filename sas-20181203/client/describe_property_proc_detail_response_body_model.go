// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyProcDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertyProcDetailResponseBodyPageInfo) *DescribePropertyProcDetailResponseBody
	GetPageInfo() *DescribePropertyProcDetailResponseBodyPageInfo
	SetPropertys(v []*DescribePropertyProcDetailResponseBodyPropertys) *DescribePropertyProcDetailResponseBody
	GetPropertys() []*DescribePropertyProcDetailResponseBodyPropertys
	SetRequestId(v string) *DescribePropertyProcDetailResponseBody
	GetRequestId() *string
}

type DescribePropertyProcDetailResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertyProcDetailResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// An array that consists of the fingerprints of the processes.
	Propertys []*DescribePropertyProcDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// AA47D46F-10DE-138C-BBB4-8A0003F75CD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyProcDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyProcDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcDetailResponseBody) GetPageInfo() *DescribePropertyProcDetailResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertyProcDetailResponseBody) GetPropertys() []*DescribePropertyProcDetailResponseBodyPropertys {
	return s.Propertys
}

func (s *DescribePropertyProcDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyProcDetailResponseBody) SetPageInfo(v *DescribePropertyProcDetailResponseBodyPageInfo) *DescribePropertyProcDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyProcDetailResponseBody) SetPropertys(v []*DescribePropertyProcDetailResponseBodyPropertys) *DescribePropertyProcDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertyProcDetailResponseBody) SetRequestId(v string) *DescribePropertyProcDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Propertys != nil {
		for _, item := range s.Propertys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePropertyProcDetailResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 1
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePropertyProcDetailResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyProcDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertyProcDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyProcDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyProcDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyProcDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyProcDetailResponseBodyPropertys struct {
	// The startup parameter of the process.
	//
	// example:
	//
	// ./8888
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The timestamp of last data collection. Unit: milliseconds.
	//
	// example:
	//
	// 1565686951000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The permission that is required to run the process.
	//
	// example:
	//
	// root
	EuidName *string `json:"EuidName,omitempty" xml:"EuidName,omitempty"`
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// The ID of the server that is associated with the process.
	//
	// example:
	//
	// i-hp35tftuh52wbp1g****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server that is associated with the process.
	//
	// example:
	//
	// hc-host-****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// Indicates whether the process is a package installation process. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	IsPackage *int32 `json:"IsPackage,omitempty" xml:"IsPackage,omitempty"`
	// The MD5 hash value of the process file.
	//
	// example:
	//
	// 842644ea3d88bd7f7e14c1c089ef****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the process.
	//
	// example:
	//
	// agetty
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The path of the process.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The ID of the process.
	//
	// example:
	//
	// 12826
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The name of the parent process to which the process belongs.
	//
	// example:
	//
	// start***.s
	Pname *string `json:"Pname,omitempty" xml:"Pname,omitempty"`
	// The time when the process starts.
	//
	// example:
	//
	// 2019-08-07 10:09:05
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The timestamp when the process starts. Unit: milliseconds.
	//
	// example:
	//
	// 1648783107000
	StartTimeDt *int64 `json:"StartTimeDt,omitempty" xml:"StartTimeDt,omitempty"`
	// The status of the process.
	//
	// example:
	//
	// sleeping
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The user who runs the process.
	//
	// example:
	//
	// root
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The UUID of the server that is associated with the process.
	//
	// example:
	//
	// 162eb349-c2d9-4f8b-805c-75b43d4c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribePropertyProcDetailResponseBodyPropertys) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyProcDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetCmdline() *string {
	return s.Cmdline
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetEuidName() *string {
	return s.EuidName
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetFileHash() *string {
	return s.FileHash
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetIsPackage() *int32 {
	return s.IsPackage
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetMd5() *string {
	return s.Md5
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetName() *string {
	return s.Name
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetPath() *string {
	return s.Path
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetPid() *string {
	return s.Pid
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetPname() *string {
	return s.Pname
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetStartTimeDt() *int64 {
	return s.StartTimeDt
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetState() *string {
	return s.State
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetUser() *string {
	return s.User
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetCmdline(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Cmdline = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertyProcDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetEuidName(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.EuidName = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetFileHash(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.FileHash = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetIsPackage(v int32) *DescribePropertyProcDetailResponseBodyPropertys {
	s.IsPackage = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetMd5(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Md5 = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetName(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Name = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetPath(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Path = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetPid(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Pid = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetPname(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Pname = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetStartTime(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.StartTime = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetStartTimeDt(v int64) *DescribePropertyProcDetailResponseBodyPropertys {
	s.StartTimeDt = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetState(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.State = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetUser(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.User = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) Validate() error {
	return dara.Validate(s)
}
