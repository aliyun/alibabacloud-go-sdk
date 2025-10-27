// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdcProbeScanResultListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeIdcProbeScanResultListResponseBodyInstances) *DescribeIdcProbeScanResultListResponseBody
	GetInstances() []*DescribeIdcProbeScanResultListResponseBodyInstances
	SetPageInfo(v *DescribeIdcProbeScanResultListResponseBodyPageInfo) *DescribeIdcProbeScanResultListResponseBody
	GetPageInfo() *DescribeIdcProbeScanResultListResponseBodyPageInfo
	SetRequestId(v string) *DescribeIdcProbeScanResultListResponseBody
	GetRequestId() *string
}

type DescribeIdcProbeScanResultListResponseBody struct {
	// The instances.
	Instances []*DescribeIdcProbeScanResultListResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeIdcProbeScanResultListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A3D7C47D-3F11-57BB-90E8-E5C20C61****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIdcProbeScanResultListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdcProbeScanResultListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIdcProbeScanResultListResponseBody) GetInstances() []*DescribeIdcProbeScanResultListResponseBodyInstances {
	return s.Instances
}

func (s *DescribeIdcProbeScanResultListResponseBody) GetPageInfo() *DescribeIdcProbeScanResultListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeIdcProbeScanResultListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIdcProbeScanResultListResponseBody) SetInstances(v []*DescribeIdcProbeScanResultListResponseBodyInstances) *DescribeIdcProbeScanResultListResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBody) SetPageInfo(v *DescribeIdcProbeScanResultListResponseBodyPageInfo) *DescribeIdcProbeScanResultListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBody) SetRequestId(v string) *DescribeIdcProbeScanResultListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIdcProbeScanResultListResponseBodyInstances struct {
	// The status of the client of the instance on which the probe is installed. Valid values:
	//
	// 	- **online**: The Security Center agent on the asset is **enabled**.
	//
	// 	- **offline**: The Security Center agent on the asset is **disabled**.
	//
	// example:
	//
	// online
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	// The name of the IDC.
	//
	// example:
	//
	// Test
	IdcName *string `json:"IdcName,omitempty" xml:"IdcName,omitempty"`
	// The CIDR blocks.
	//
	// example:
	//
	// 192.168.2.0/24
	IpSegment *string `json:"IpSegment,omitempty" xml:"IpSegment,omitempty"`
	// The timestamp when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1720006818000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The operating system type of the asset. Valid values:
	//
	// 	- **windows**
	//
	// 	- **linux**
	//
	// example:
	//
	// Linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The private IP address of the associated instance.
	//
	// example:
	//
	// 42.121.*.*
	ProbeInternetIp *string `json:"ProbeInternetIp,omitempty" xml:"ProbeInternetIp,omitempty"`
	// The private IP address of the associated instance.
	//
	// example:
	//
	// 192.168.*.*
	ProbeIntranetIp *string `json:"ProbeIntranetIp,omitempty" xml:"ProbeIntranetIp,omitempty"`
	// The name of the associated instance.
	//
	// example:
	//
	// i-xxxx
	ProbeMachineName *string `json:"ProbeMachineName,omitempty" xml:"ProbeMachineName,omitempty"`
	// The UUID of the associated instance.
	//
	// example:
	//
	// 11C96623-E106-59C9-866D-A6C82911****
	ProbeUuid *string `json:"ProbeUuid,omitempty" xml:"ProbeUuid,omitempty"`
	// The ID of the scan result.
	//
	// example:
	//
	// 1231
	ScanResultId *int64 `json:"ScanResultId,omitempty" xml:"ScanResultId,omitempty"`
	// The IP address that is scanned.
	//
	// example:
	//
	// 192.168.*.*
	ScannedIp *string `json:"ScannedIp,omitempty" xml:"ScannedIp,omitempty"`
	// The port that is scanned.
	//
	// example:
	//
	// 22
	ValidPort *string `json:"ValidPort,omitempty" xml:"ValidPort,omitempty"`
}

func (s DescribeIdcProbeScanResultListResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdcProbeScanResultListResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) GetClientStatus() *string {
	return s.ClientStatus
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) GetIdcName() *string {
	return s.IdcName
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) GetIpSegment() *string {
	return s.IpSegment
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) GetOs() *string {
	return s.Os
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) GetProbeInternetIp() *string {
	return s.ProbeInternetIp
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) GetProbeIntranetIp() *string {
	return s.ProbeIntranetIp
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) GetProbeMachineName() *string {
	return s.ProbeMachineName
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) GetProbeUuid() *string {
	return s.ProbeUuid
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) GetScanResultId() *int64 {
	return s.ScanResultId
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) GetScannedIp() *string {
	return s.ScannedIp
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) GetValidPort() *string {
	return s.ValidPort
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) SetClientStatus(v string) *DescribeIdcProbeScanResultListResponseBodyInstances {
	s.ClientStatus = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) SetIdcName(v string) *DescribeIdcProbeScanResultListResponseBodyInstances {
	s.IdcName = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) SetIpSegment(v string) *DescribeIdcProbeScanResultListResponseBodyInstances {
	s.IpSegment = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) SetLastScanTime(v int64) *DescribeIdcProbeScanResultListResponseBodyInstances {
	s.LastScanTime = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) SetOs(v string) *DescribeIdcProbeScanResultListResponseBodyInstances {
	s.Os = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) SetProbeInternetIp(v string) *DescribeIdcProbeScanResultListResponseBodyInstances {
	s.ProbeInternetIp = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) SetProbeIntranetIp(v string) *DescribeIdcProbeScanResultListResponseBodyInstances {
	s.ProbeIntranetIp = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) SetProbeMachineName(v string) *DescribeIdcProbeScanResultListResponseBodyInstances {
	s.ProbeMachineName = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) SetProbeUuid(v string) *DescribeIdcProbeScanResultListResponseBodyInstances {
	s.ProbeUuid = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) SetScanResultId(v int64) *DescribeIdcProbeScanResultListResponseBodyInstances {
	s.ScanResultId = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) SetScannedIp(v string) *DescribeIdcProbeScanResultListResponseBodyInstances {
	s.ScannedIp = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) SetValidPort(v string) *DescribeIdcProbeScanResultListResponseBodyInstances {
	s.ValidPort = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeIdcProbeScanResultListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 9
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
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
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIdcProbeScanResultListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdcProbeScanResultListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeIdcProbeScanResultListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeIdcProbeScanResultListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeIdcProbeScanResultListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIdcProbeScanResultListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIdcProbeScanResultListResponseBodyPageInfo) SetCount(v int32) *DescribeIdcProbeScanResultListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeIdcProbeScanResultListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyPageInfo) SetPageSize(v int32) *DescribeIdcProbeScanResultListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeIdcProbeScanResultListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeIdcProbeScanResultListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
