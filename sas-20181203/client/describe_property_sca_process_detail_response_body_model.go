// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyScaProcessDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertyScaProcessDetailResponseBodyPageInfo) *DescribePropertyScaProcessDetailResponseBody
	GetPageInfo() *DescribePropertyScaProcessDetailResponseBodyPageInfo
	SetPropertys(v []*DescribePropertyScaProcessDetailResponseBodyPropertys) *DescribePropertyScaProcessDetailResponseBody
	GetPropertys() []*DescribePropertyScaProcessDetailResponseBodyPropertys
	SetRequestId(v string) *DescribePropertyScaProcessDetailResponseBody
	GetRequestId() *string
}

type DescribePropertyScaProcessDetailResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertyScaProcessDetailResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The processes collected by the asset fingerprints feature.
	Propertys []*DescribePropertyScaProcessDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ADE57832-9666-511C-9A80-B87DE2E8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyScaProcessDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaProcessDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaProcessDetailResponseBody) GetPageInfo() *DescribePropertyScaProcessDetailResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertyScaProcessDetailResponseBody) GetPropertys() []*DescribePropertyScaProcessDetailResponseBodyPropertys {
	return s.Propertys
}

func (s *DescribePropertyScaProcessDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyScaProcessDetailResponseBody) SetPageInfo(v *DescribePropertyScaProcessDetailResponseBodyPageInfo) *DescribePropertyScaProcessDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBody) SetPropertys(v []*DescribePropertyScaProcessDetailResponseBodyPropertys) *DescribePropertyScaProcessDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBody) SetRequestId(v string) *DescribePropertyScaProcessDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyScaProcessDetailResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
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
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePropertyScaProcessDetailResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaProcessDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaProcessDetailResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyScaProcessDetailResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyScaProcessDetailResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyScaProcessDetailResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertyScaProcessDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertyScaProcessDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyScaProcessDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyScaProcessDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyScaProcessDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyScaProcessDetailResponseBodyPropertys struct {
	// The command line of the process.
	//
	// example:
	//
	// java -Xms128m -Xmx512m -DNACOS_URL=http://10.184.XX.XX:8848 -DNACOS_NAMESPACE=iam-sit -jar /opt/service/xxl-job/xxl-job-admin-2.3.0.jar
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The timestamp at which the last asset fingerprint collection is performed. Unit: milliseconds.
	//
	// example:
	//
	// 1597987834000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The instance ID of the server.
	//
	// example:
	//
	// i-hp35tftuh52wbp1g****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name of the server.
	//
	// example:
	//
	// hc-host-****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 120.26.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The name of the process.
	//
	// example:
	//
	// java
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the process.
	//
	// example:
	//
	// 522
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 162eb349-c2d9-4f8b-805c-75b43d4c****
	Uuid    *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribePropertyScaProcessDetailResponseBodyPropertys) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaProcessDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) GetCmdline() *string {
	return s.Cmdline
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) GetName() *string {
	return s.Name
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) GetPid() *string {
	return s.Pid
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) GetVersion() *string {
	return s.Version
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) SetCmdline(v string) *DescribePropertyScaProcessDetailResponseBodyPropertys {
	s.Cmdline = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertyScaProcessDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertyScaProcessDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertyScaProcessDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertyScaProcessDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertyScaProcessDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) SetName(v string) *DescribePropertyScaProcessDetailResponseBodyPropertys {
	s.Name = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) SetPid(v string) *DescribePropertyScaProcessDetailResponseBodyPropertys {
	s.Pid = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertyScaProcessDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) SetVersion(v string) *DescribePropertyScaProcessDetailResponseBodyPropertys {
	s.Version = &v
	return s
}

func (s *DescribePropertyScaProcessDetailResponseBodyPropertys) Validate() error {
	return dara.Validate(s)
}
