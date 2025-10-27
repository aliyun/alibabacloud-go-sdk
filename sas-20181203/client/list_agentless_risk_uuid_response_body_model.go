// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessRiskUuidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListAgentlessRiskUuidResponseBodyList) *ListAgentlessRiskUuidResponseBody
	GetList() []*ListAgentlessRiskUuidResponseBodyList
	SetPageInfo(v *ListAgentlessRiskUuidResponseBodyPageInfo) *ListAgentlessRiskUuidResponseBody
	GetPageInfo() *ListAgentlessRiskUuidResponseBodyPageInfo
	SetRequestId(v string) *ListAgentlessRiskUuidResponseBody
	GetRequestId() *string
}

type ListAgentlessRiskUuidResponseBody struct {
	// The information about the hosts.
	List []*ListAgentlessRiskUuidResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListAgentlessRiskUuidResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05AE6B16-741A-5A75-9BF0-BC747DD9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAgentlessRiskUuidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRiskUuidResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentlessRiskUuidResponseBody) GetList() []*ListAgentlessRiskUuidResponseBodyList {
	return s.List
}

func (s *ListAgentlessRiskUuidResponseBody) GetPageInfo() *ListAgentlessRiskUuidResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListAgentlessRiskUuidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentlessRiskUuidResponseBody) SetList(v []*ListAgentlessRiskUuidResponseBodyList) *ListAgentlessRiskUuidResponseBody {
	s.List = v
	return s
}

func (s *ListAgentlessRiskUuidResponseBody) SetPageInfo(v *ListAgentlessRiskUuidResponseBodyPageInfo) *ListAgentlessRiskUuidResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListAgentlessRiskUuidResponseBody) SetRequestId(v string) *ListAgentlessRiskUuidResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
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

type ListAgentlessRiskUuidResponseBodyList struct {
	// The number of baseline risks.
	//
	// example:
	//
	// 1
	BaselineCount *int32 `json:"BaselineCount,omitempty" xml:"BaselineCount,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// s-bp1g6wxdwps7s9dz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name of the asset.
	//
	// example:
	//
	// sql-test-0****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 8.210.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.25.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The number of malicious samples.
	//
	// example:
	//
	// 1
	MaliciousCount *int32 `json:"MaliciousCount,omitempty" xml:"MaliciousCount,omitempty"`
	// The timestamp of the detection. Unit: milliseconds.
	//
	// example:
	//
	// 168257753****
	ScanTime *int64 `json:"ScanTime,omitempty" xml:"ScanTime,omitempty"`
	// The ID of the asset that is detected.
	//
	// example:
	//
	// 30****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the asset that is detected.
	//
	// example:
	//
	// test****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 49e25e0f-bb51-4a5a-a1b3-13a4ddaa****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The number of detected vulnerabilities.
	//
	// example:
	//
	// 1
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
}

func (s ListAgentlessRiskUuidResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRiskUuidResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListAgentlessRiskUuidResponseBodyList) GetBaselineCount() *int32 {
	return s.BaselineCount
}

func (s *ListAgentlessRiskUuidResponseBodyList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentlessRiskUuidResponseBodyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListAgentlessRiskUuidResponseBodyList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListAgentlessRiskUuidResponseBodyList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListAgentlessRiskUuidResponseBodyList) GetMaliciousCount() *int32 {
	return s.MaliciousCount
}

func (s *ListAgentlessRiskUuidResponseBodyList) GetScanTime() *int64 {
	return s.ScanTime
}

func (s *ListAgentlessRiskUuidResponseBodyList) GetTargetId() *string {
	return s.TargetId
}

func (s *ListAgentlessRiskUuidResponseBodyList) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAgentlessRiskUuidResponseBodyList) GetUuid() *string {
	return s.Uuid
}

func (s *ListAgentlessRiskUuidResponseBodyList) GetVulCount() *int32 {
	return s.VulCount
}

func (s *ListAgentlessRiskUuidResponseBodyList) SetBaselineCount(v int32) *ListAgentlessRiskUuidResponseBodyList {
	s.BaselineCount = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyList) SetInstanceId(v string) *ListAgentlessRiskUuidResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyList) SetInstanceName(v string) *ListAgentlessRiskUuidResponseBodyList {
	s.InstanceName = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyList) SetInternetIp(v string) *ListAgentlessRiskUuidResponseBodyList {
	s.InternetIp = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyList) SetIntranetIp(v string) *ListAgentlessRiskUuidResponseBodyList {
	s.IntranetIp = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyList) SetMaliciousCount(v int32) *ListAgentlessRiskUuidResponseBodyList {
	s.MaliciousCount = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyList) SetScanTime(v int64) *ListAgentlessRiskUuidResponseBodyList {
	s.ScanTime = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyList) SetTargetId(v string) *ListAgentlessRiskUuidResponseBodyList {
	s.TargetId = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyList) SetTargetName(v string) *ListAgentlessRiskUuidResponseBodyList {
	s.TargetName = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyList) SetUuid(v string) *ListAgentlessRiskUuidResponseBodyList {
	s.Uuid = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyList) SetVulCount(v int32) *ListAgentlessRiskUuidResponseBodyList {
	s.VulCount = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListAgentlessRiskUuidResponseBodyPageInfo struct {
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
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentlessRiskUuidResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRiskUuidResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListAgentlessRiskUuidResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAgentlessRiskUuidResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentlessRiskUuidResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentlessRiskUuidResponseBodyPageInfo) SetCurrentPage(v int32) *ListAgentlessRiskUuidResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyPageInfo) SetPageSize(v int32) *ListAgentlessRiskUuidResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyPageInfo) SetTotalCount(v int32) *ListAgentlessRiskUuidResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListAgentlessRiskUuidResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
