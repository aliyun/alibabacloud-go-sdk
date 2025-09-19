// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertySoftwareDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertySoftwareDetailResponseBodyPageInfo) *DescribePropertySoftwareDetailResponseBody
	GetPageInfo() *DescribePropertySoftwareDetailResponseBodyPageInfo
	SetPropertys(v []*DescribePropertySoftwareDetailResponseBodyPropertys) *DescribePropertySoftwareDetailResponseBody
	GetPropertys() []*DescribePropertySoftwareDetailResponseBodyPropertys
	SetRequestId(v string) *DescribePropertySoftwareDetailResponseBody
	GetRequestId() *string
}

type DescribePropertySoftwareDetailResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertySoftwareDetailResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The details of the software asset.
	Propertys []*DescribePropertySoftwareDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6AEE7412-0065-1135-B790-AE2C38BA68FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertySoftwareDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertySoftwareDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareDetailResponseBody) GetPageInfo() *DescribePropertySoftwareDetailResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertySoftwareDetailResponseBody) GetPropertys() []*DescribePropertySoftwareDetailResponseBodyPropertys {
	return s.Propertys
}

func (s *DescribePropertySoftwareDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertySoftwareDetailResponseBody) SetPageInfo(v *DescribePropertySoftwareDetailResponseBodyPageInfo) *DescribePropertySoftwareDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBody) SetPropertys(v []*DescribePropertySoftwareDetailResponseBodyPropertys) *DescribePropertySoftwareDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBody) SetRequestId(v string) *DescribePropertySoftwareDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePropertySoftwareDetailResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Default value: **10**.
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

func (s DescribePropertySoftwareDetailResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertySoftwareDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertySoftwareDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertySoftwareDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertySoftwareDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertySoftwareDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertySoftwareDetailResponseBodyPropertys struct {
	// The timestamp generated when the last asset fingerprint collection is performed. Unit: milliseconds.
	//
	// example:
	//
	// 1649149566000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The time at which the software is installed.
	//
	// example:
	//
	// 2022-04-07 10:54:49
	InstallTime *string `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	// The timestamp generated when the software is installed. Unit: milliseconds.
	//
	// example:
	//
	// 1649066826000
	InstallTimeDt *int64 `json:"InstallTimeDt,omitempty" xml:"InstallTimeDt,omitempty"`
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
	// 192.168.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 100.104.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The IP addresses of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The name of the software.
	//
	// example:
	//
	// aaa_base
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The installation path of the software.
	//
	// example:
	//
	// /etc/test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 162eb349-c2d9-4f8b-805c-75b43d4c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The version of the software.
	//
	// example:
	//
	// 3.10.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribePropertySoftwareDetailResponseBodyPropertys) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertySoftwareDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) GetInstallTime() *string {
	return s.InstallTime
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) GetInstallTimeDt() *int64 {
	return s.InstallTimeDt
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) GetIp() *string {
	return s.Ip
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) GetName() *string {
	return s.Name
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) GetPath() *string {
	return s.Path
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) GetVersion() *string {
	return s.Version
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetInstallTime(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.InstallTime = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetInstallTimeDt(v int64) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.InstallTimeDt = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetIp(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.Ip = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetName(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.Name = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetPath(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.Path = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetVersion(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.Version = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) Validate() error {
	return dara.Validate(s)
}
