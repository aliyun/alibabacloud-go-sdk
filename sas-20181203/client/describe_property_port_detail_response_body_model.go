// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyPortDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertyPortDetailResponseBodyPageInfo) *DescribePropertyPortDetailResponseBody
	GetPageInfo() *DescribePropertyPortDetailResponseBodyPageInfo
	SetPropertys(v []*DescribePropertyPortDetailResponseBodyPropertys) *DescribePropertyPortDetailResponseBody
	GetPropertys() []*DescribePropertyPortDetailResponseBodyPropertys
	SetRequestId(v string) *DescribePropertyPortDetailResponseBody
	GetRequestId() *string
}

type DescribePropertyPortDetailResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertyPortDetailResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The fingerprints of the ports.
	Propertys []*DescribePropertyPortDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0FA7F1F4-488D-52CA-9BFC-3E47793B49D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyPortDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyPortDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortDetailResponseBody) GetPageInfo() *DescribePropertyPortDetailResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertyPortDetailResponseBody) GetPropertys() []*DescribePropertyPortDetailResponseBodyPropertys {
	return s.Propertys
}

func (s *DescribePropertyPortDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyPortDetailResponseBody) SetPageInfo(v *DescribePropertyPortDetailResponseBodyPageInfo) *DescribePropertyPortDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyPortDetailResponseBody) SetPropertys(v []*DescribePropertyPortDetailResponseBodyPropertys) *DescribePropertyPortDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertyPortDetailResponseBody) SetRequestId(v string) *DescribePropertyPortDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBody) Validate() error {
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

type DescribePropertyPortDetailResponseBodyPageInfo struct {
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
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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

func (s DescribePropertyPortDetailResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyPortDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertyPortDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyPortDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) SetNextToken(v string) *DescribePropertyPortDetailResponseBodyPageInfo {
	s.NextToken = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyPortDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyPortDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyPortDetailResponseBodyPropertys struct {
	// The IP address bound to the port.
	//
	// example:
	//
	// 0.0.X.X
	BindIp *string `json:"BindIp,omitempty" xml:"BindIp,omitempty"`
	// The timestamp of the last fingerprint collection. Unit: milliseconds.
	//
	// example:
	//
	// 1649149566000
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
	// 192.168.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 100.104.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The IP address of the network interface controller (NIC) that is bound to the listening port.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the server process that listens on the port.
	//
	// example:
	//
	// 522
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The listener port.
	//
	// example:
	//
	// 22
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the server process.
	//
	// example:
	//
	// sshd
	ProcName *string `json:"ProcName,omitempty" xml:"ProcName,omitempty"`
	// The network protocol that is used by the listening port.
	//
	// example:
	//
	// tcp
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 162eb349-c2d9-4f8b-805c-75b43d4c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribePropertyPortDetailResponseBodyPropertys) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyPortDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) GetBindIp() *string {
	return s.BindIp
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) GetIp() *string {
	return s.Ip
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) GetPid() *string {
	return s.Pid
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) GetPort() *string {
	return s.Port
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) GetProcName() *string {
	return s.ProcName
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) GetProto() *string {
	return s.Proto
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetBindIp(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.BindIp = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertyPortDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetIp(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.Ip = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetPid(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.Pid = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetPort(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.Port = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetProcName(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.ProcName = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetProto(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.Proto = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) Validate() error {
	return dara.Validate(s)
}
