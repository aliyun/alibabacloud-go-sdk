// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessRelateMaliciousResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListAgentlessRelateMaliciousResponseBodyList) *ListAgentlessRelateMaliciousResponseBody
	GetList() []*ListAgentlessRelateMaliciousResponseBodyList
	SetPageInfo(v *ListAgentlessRelateMaliciousResponseBodyPageInfo) *ListAgentlessRelateMaliciousResponseBody
	GetPageInfo() *ListAgentlessRelateMaliciousResponseBodyPageInfo
	SetRequestId(v string) *ListAgentlessRelateMaliciousResponseBody
	GetRequestId() *string
}

type ListAgentlessRelateMaliciousResponseBody struct {
	// The list of hosts that are associated with the risk.
	List []*ListAgentlessRelateMaliciousResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListAgentlessRelateMaliciousResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FCA270CC-24C8-5DA1-9184-20F2605C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAgentlessRelateMaliciousResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRelateMaliciousResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentlessRelateMaliciousResponseBody) GetList() []*ListAgentlessRelateMaliciousResponseBodyList {
	return s.List
}

func (s *ListAgentlessRelateMaliciousResponseBody) GetPageInfo() *ListAgentlessRelateMaliciousResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListAgentlessRelateMaliciousResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentlessRelateMaliciousResponseBody) SetList(v []*ListAgentlessRelateMaliciousResponseBodyList) *ListAgentlessRelateMaliciousResponseBody {
	s.List = v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBody) SetPageInfo(v *ListAgentlessRelateMaliciousResponseBodyPageInfo) *ListAgentlessRelateMaliciousResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBody) SetRequestId(v string) *ListAgentlessRelateMaliciousResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAgentlessRelateMaliciousResponseBodyList struct {
	// The details of the alert events.
	Details []*ListAgentlessRelateMaliciousResponseBodyListDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The URL to download the malicious image sample.
	//
	// example:
	//
	// https://upgrade-rule-pkg.oss-cn-beijing.aliyuncs.com/totalpackage/***
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The file path.
	//
	// example:
	//
	// /root/logs/arms/tracelogs/log****
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The timestamp when the first scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 168257753****
	FirstScanTimestamp *int64 `json:"FirstScanTimestamp,omitempty" xml:"FirstScanTimestamp,omitempty"`
	// The highlighted JSON string.
	//
	// example:
	//
	// {\\"highlight\\":{\\"ruleVersion\\":\\"20230223\\",\\"ruleId\\":600139,\\"events\\":[[207,284]]}}
	HighLight *string `json:"HighLight,omitempty" xml:"HighLight,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 8012
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the asset.
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
	// The timestamp when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 168257753****
	LatestScanTimestamp *int64 `json:"LatestScanTimestamp,omitempty" xml:"LatestScanTimestamp,omitempty"`
	// The severity of the malicious file. Valid values:
	//
	// 	- serious
	//
	// 	- suspicious
	//
	// 	- remind
	//
	// example:
	//
	// serious
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The MD5 hash value of the malicious file.
	//
	// example:
	//
	// 1f2e13a7c51ee89316ae50066515****
	MaliciousMd5 *string `json:"MaliciousMd5,omitempty" xml:"MaliciousMd5,omitempty"`
	// The name of the malicious file.
	//
	// example:
	//
	// WebShell
	MaliciousName *string `json:"MaliciousName,omitempty" xml:"MaliciousName,omitempty"`
	// The type of the virus.
	//
	// example:
	//
	// WebShell
	MaliciousType *string `json:"MaliciousType,omitempty" xml:"MaliciousType,omitempty"`
	// The handling result of the alert.
	//
	// example:
	//
	// addWhitelist.USER.Success
	OperateResult *string `json:"OperateResult,omitempty" xml:"OperateResult,omitempty"`
	// The timestamp when the alert is handled. Unit: milliseconds.
	//
	// example:
	//
	// 168257753****
	OperateTimestamp *string `json:"OperateTimestamp,omitempty" xml:"OperateTimestamp,omitempty"`
	// The partition of the disk.
	//
	// example:
	//
	// /dev/xvda1
	Partition *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The ID of the task object.
	//
	// example:
	//
	// m-****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	//
	// example:
	//
	// image_***
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the task object. Valid values:
	//
	// 	- **1**: snapshot.
	//
	// 	- **2**: image.
	//
	// example:
	//
	// 2
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 49e25e0f-bb51-4a5a-a1b3-13a4ddaa****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListAgentlessRelateMaliciousResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRelateMaliciousResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetDetails() []*ListAgentlessRelateMaliciousResponseBodyListDetails {
	return s.Details
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetFilePath() *string {
	return s.FilePath
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetFirstScanTimestamp() *int64 {
	return s.FirstScanTimestamp
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetHighLight() *string {
	return s.HighLight
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetId() *int64 {
	return s.Id
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetLatestScanTimestamp() *int64 {
	return s.LatestScanTimestamp
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetLevel() *string {
	return s.Level
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetMaliciousMd5() *string {
	return s.MaliciousMd5
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetMaliciousName() *string {
	return s.MaliciousName
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetMaliciousType() *string {
	return s.MaliciousType
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetOperateResult() *string {
	return s.OperateResult
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetOperateTimestamp() *string {
	return s.OperateTimestamp
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetPartition() *string {
	return s.Partition
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetTargetId() *string {
	return s.TargetId
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetTargetType() *string {
	return s.TargetType
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) GetUuid() *string {
	return s.Uuid
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetDetails(v []*ListAgentlessRelateMaliciousResponseBodyListDetails) *ListAgentlessRelateMaliciousResponseBodyList {
	s.Details = v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetDownloadUrl(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.DownloadUrl = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetFilePath(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.FilePath = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetFirstScanTimestamp(v int64) *ListAgentlessRelateMaliciousResponseBodyList {
	s.FirstScanTimestamp = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetHighLight(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.HighLight = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetId(v int64) *ListAgentlessRelateMaliciousResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetInstanceName(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.InstanceName = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetInternetIp(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.InternetIp = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetIntranetIp(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.IntranetIp = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetLatestScanTimestamp(v int64) *ListAgentlessRelateMaliciousResponseBodyList {
	s.LatestScanTimestamp = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetLevel(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.Level = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetMaliciousMd5(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.MaliciousMd5 = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetMaliciousName(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.MaliciousName = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetMaliciousType(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.MaliciousType = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetOperateResult(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.OperateResult = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetOperateTimestamp(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.OperateTimestamp = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetPartition(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.Partition = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetTargetId(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.TargetId = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetTargetName(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.TargetName = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetTargetType(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.TargetType = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) SetUuid(v string) *ListAgentlessRelateMaliciousResponseBodyList {
	s.Uuid = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListAgentlessRelateMaliciousResponseBodyListDetails struct {
	// The name of the detailed item.
	//
	// example:
	//
	// MD5
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name key of the detailed item.
	//
	// example:
	//
	// ${suspicious.mp.db.maliciousfilemd5}
	NameKey *string `json:"NameKey,omitempty" xml:"NameKey,omitempty"`
	// The type of the detailed item.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the detailed item.
	//
	// example:
	//
	// 1f2e13a7c51ee89316ae50066515****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAgentlessRelateMaliciousResponseBodyListDetails) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRelateMaliciousResponseBodyListDetails) GoString() string {
	return s.String()
}

func (s *ListAgentlessRelateMaliciousResponseBodyListDetails) GetName() *string {
	return s.Name
}

func (s *ListAgentlessRelateMaliciousResponseBodyListDetails) GetNameKey() *string {
	return s.NameKey
}

func (s *ListAgentlessRelateMaliciousResponseBodyListDetails) GetType() *string {
	return s.Type
}

func (s *ListAgentlessRelateMaliciousResponseBodyListDetails) GetValue() *string {
	return s.Value
}

func (s *ListAgentlessRelateMaliciousResponseBodyListDetails) SetName(v string) *ListAgentlessRelateMaliciousResponseBodyListDetails {
	s.Name = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyListDetails) SetNameKey(v string) *ListAgentlessRelateMaliciousResponseBodyListDetails {
	s.NameKey = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyListDetails) SetType(v string) *ListAgentlessRelateMaliciousResponseBodyListDetails {
	s.Type = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyListDetails) SetValue(v string) *ListAgentlessRelateMaliciousResponseBodyListDetails {
	s.Value = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyListDetails) Validate() error {
	return dara.Validate(s)
}

type ListAgentlessRelateMaliciousResponseBodyPageInfo struct {
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
	// 197
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentlessRelateMaliciousResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRelateMaliciousResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListAgentlessRelateMaliciousResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListAgentlessRelateMaliciousResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAgentlessRelateMaliciousResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentlessRelateMaliciousResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentlessRelateMaliciousResponseBodyPageInfo) SetCount(v int32) *ListAgentlessRelateMaliciousResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyPageInfo) SetCurrentPage(v int32) *ListAgentlessRelateMaliciousResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyPageInfo) SetPageSize(v int32) *ListAgentlessRelateMaliciousResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyPageInfo) SetTotalCount(v int32) *ListAgentlessRelateMaliciousResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListAgentlessRelateMaliciousResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
