// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBruteForceRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMachineList(v []*DescribeBruteForceRecordsResponseBodyMachineList) *DescribeBruteForceRecordsResponseBody
	GetMachineList() []*DescribeBruteForceRecordsResponseBodyMachineList
	SetPageInfo(v *DescribeBruteForceRecordsResponseBodyPageInfo) *DescribeBruteForceRecordsResponseBody
	GetPageInfo() *DescribeBruteForceRecordsResponseBodyPageInfo
	SetRequestId(v string) *DescribeBruteForceRecordsResponseBody
	GetRequestId() *string
}

type DescribeBruteForceRecordsResponseBody struct {
	// The IP addresses.
	MachineList []*DescribeBruteForceRecordsResponseBodyMachineList `json:"MachineList,omitempty" xml:"MachineList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeBruteForceRecordsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBruteForceRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBruteForceRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceRecordsResponseBody) GetMachineList() []*DescribeBruteForceRecordsResponseBodyMachineList {
	return s.MachineList
}

func (s *DescribeBruteForceRecordsResponseBody) GetPageInfo() *DescribeBruteForceRecordsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeBruteForceRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBruteForceRecordsResponseBody) SetMachineList(v []*DescribeBruteForceRecordsResponseBodyMachineList) *DescribeBruteForceRecordsResponseBody {
	s.MachineList = v
	return s
}

func (s *DescribeBruteForceRecordsResponseBody) SetPageInfo(v *DescribeBruteForceRecordsResponseBodyPageInfo) *DescribeBruteForceRecordsResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeBruteForceRecordsResponseBody) SetRequestId(v string) *DescribeBruteForceRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBruteForceRecordsResponseBodyMachineList struct {
	// The status of the host network extension. Valid values:
	//
	// 	- **true**: online
	//
	// 	- **false**: offline
	//
	// example:
	//
	// false
	AliNetOnline *bool `json:"AliNetOnline,omitempty" xml:"AliNetOnline,omitempty"`
	// The timestamp when the block action on the IP address becomes invalid.
	//
	// example:
	//
	// 1671506882063
	BlockExpireDate *int64 `json:"BlockExpireDate,omitempty" xml:"BlockExpireDate,omitempty"`
	// The IP address that is blocked.
	//
	// example:
	//
	// 10.12.XX.XX
	BlockIp *string `json:"BlockIp,omitempty" xml:"BlockIp,omitempty"`
	// The blocking type. Valid values:
	//
	// 	- **group**: security group
	//
	// 	- **alinet**: host network extension
	//
	// example:
	//
	// alinet
	BlockType *string `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The error code returned when the defense rule fails to block the IP address.
	//
	// example:
	//
	// InstanceSecurityGroupLimitExceeded
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the primary key that is recorded in the defense rule.
	//
	// example:
	//
	// 112XX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance name of the server.
	//
	// example:
	//
	// record-test-***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 120.79.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The port that is attacked.
	//
	// example:
	//
	// 22/22
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the defense rule.
	//
	// example:
	//
	// AntiRuleName
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the defense rule. Valid values:
	//
	// 	- **userRule**: custom rule
	//
	// 	- **blinkRule**: system rule
	//
	// example:
	//
	// userRule
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the defense rule. Valid values:
	//
	// 	- **0**: invalid
	//
	// 	- **1**: enabled
	//
	// 	- **2**: failed
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server on which the defense rule takes effect.
	//
	// example:
	//
	// 6d5b361f-958d-48a8-a9d2-d6e82c1****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeBruteForceRecordsResponseBodyMachineList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBruteForceRecordsResponseBodyMachineList) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetAliNetOnline() *bool {
	return s.AliNetOnline
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetBlockExpireDate() *int64 {
	return s.BlockExpireDate
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetBlockIp() *string {
	return s.BlockIp
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetBlockType() *string {
	return s.BlockType
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetId() *int64 {
	return s.Id
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetPort() *string {
	return s.Port
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetSource() *string {
	return s.Source
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetAliNetOnline(v bool) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.AliNetOnline = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetBlockExpireDate(v int64) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.BlockExpireDate = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetBlockIp(v string) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.BlockIp = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetBlockType(v string) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.BlockType = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetErrorCode(v string) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.ErrorCode = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetId(v int64) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.Id = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetInstanceName(v string) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.InstanceName = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetInternetIp(v string) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.InternetIp = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetIntranetIp(v string) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetPort(v string) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.Port = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetRuleName(v string) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.RuleName = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetSource(v string) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.Source = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetStatus(v int32) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.Status = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) SetUuid(v string) *DescribeBruteForceRecordsResponseBodyMachineList {
	s.Uuid = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyMachineList) Validate() error {
	return dara.Validate(s)
}

type DescribeBruteForceRecordsResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 263
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBruteForceRecordsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeBruteForceRecordsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceRecordsResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeBruteForceRecordsResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeBruteForceRecordsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBruteForceRecordsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBruteForceRecordsResponseBodyPageInfo) SetCount(v int32) *DescribeBruteForceRecordsResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeBruteForceRecordsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyPageInfo) SetPageSize(v int32) *DescribeBruteForceRecordsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyPageInfo) SetTotalCount(v int32) *DescribeBruteForceRecordsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeBruteForceRecordsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
