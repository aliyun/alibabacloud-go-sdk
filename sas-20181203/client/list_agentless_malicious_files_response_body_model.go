// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessMaliciousFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListAgentlessMaliciousFilesResponseBodyList) *ListAgentlessMaliciousFilesResponseBody
	GetList() []*ListAgentlessMaliciousFilesResponseBodyList
	SetPageInfo(v *ListAgentlessMaliciousFilesResponseBodyPageInfo) *ListAgentlessMaliciousFilesResponseBody
	GetPageInfo() *ListAgentlessMaliciousFilesResponseBodyPageInfo
	SetRequestId(v string) *ListAgentlessMaliciousFilesResponseBody
	GetRequestId() *string
}

type ListAgentlessMaliciousFilesResponseBody struct {
	// The malicious files.
	List []*ListAgentlessMaliciousFilesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListAgentlessMaliciousFilesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CE343162-35BF-565E-B85E-8E3B7A2B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAgentlessMaliciousFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessMaliciousFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentlessMaliciousFilesResponseBody) GetList() []*ListAgentlessMaliciousFilesResponseBodyList {
	return s.List
}

func (s *ListAgentlessMaliciousFilesResponseBody) GetPageInfo() *ListAgentlessMaliciousFilesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListAgentlessMaliciousFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentlessMaliciousFilesResponseBody) SetList(v []*ListAgentlessMaliciousFilesResponseBodyList) *ListAgentlessMaliciousFilesResponseBody {
	s.List = v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBody) SetPageInfo(v *ListAgentlessMaliciousFilesResponseBodyPageInfo) *ListAgentlessMaliciousFilesResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBody) SetRequestId(v string) *ListAgentlessMaliciousFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBody) Validate() error {
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

type ListAgentlessMaliciousFilesResponseBodyList struct {
	// The details of the alert event.
	Details []*ListAgentlessMaliciousFilesResponseBodyListDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The URL to download the malicious image sample.
	//
	// example:
	//
	// https://upgrade-rule-pkg.oss-cn-beijing.aliyuncs.com/totalpackage/rule/V3.166.2/totalpackage_rule_V3.16_6.2_tip_20230220110****
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The file path.
	//
	// example:
	//
	// /root/etc/g****
	FilePath *int64 `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The timestamp generated when the first scan was performed. Unit: milliseconds.
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
	// 1****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// sql-test-001****
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
	// The timestamp generated when the last scan was performed. Unit: milliseconds.
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
	// The remarks.
	Notes []*ListAgentlessMaliciousFilesResponseBodyListNotes `json:"Notes,omitempty" xml:"Notes,omitempty" type:"Repeated"`
	// The handling result of the alert.
	//
	// example:
	//
	// addWhitelist.USER.Success
	OperateResult *string `json:"OperateResult,omitempty" xml:"OperateResult,omitempty"`
	// The timestamp generated when the alert is handled. Unit: milliseconds.
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
	// The ID of the asset that is scanned.
	//
	// example:
	//
	// m-****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the asset that is scanned.
	//
	// example:
	//
	// jenkins****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the asset that is scanned. Valid values:
	//
	// 	- 2: The scanned asset is an image.
	//
	// example:
	//
	// 2
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The UUID of the asset.
	//
	// example:
	//
	// 49e25e0f-bb51-4a5a-a1b3-13a4ddaa****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListAgentlessMaliciousFilesResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessMaliciousFilesResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetDetails() []*ListAgentlessMaliciousFilesResponseBodyListDetails {
	return s.Details
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetFilePath() *int64 {
	return s.FilePath
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetFirstScanTimestamp() *int64 {
	return s.FirstScanTimestamp
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetHighLight() *string {
	return s.HighLight
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetId() *int64 {
	return s.Id
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetLatestScanTimestamp() *int64 {
	return s.LatestScanTimestamp
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetLevel() *string {
	return s.Level
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetMaliciousMd5() *string {
	return s.MaliciousMd5
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetMaliciousName() *string {
	return s.MaliciousName
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetMaliciousType() *string {
	return s.MaliciousType
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetNotes() []*ListAgentlessMaliciousFilesResponseBodyListNotes {
	return s.Notes
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetOperateResult() *string {
	return s.OperateResult
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetOperateTimestamp() *string {
	return s.OperateTimestamp
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetPartition() *string {
	return s.Partition
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetTargetId() *string {
	return s.TargetId
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetTargetType() *string {
	return s.TargetType
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) GetUuid() *string {
	return s.Uuid
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetDetails(v []*ListAgentlessMaliciousFilesResponseBodyListDetails) *ListAgentlessMaliciousFilesResponseBodyList {
	s.Details = v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetDownloadUrl(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.DownloadUrl = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetFilePath(v int64) *ListAgentlessMaliciousFilesResponseBodyList {
	s.FilePath = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetFirstScanTimestamp(v int64) *ListAgentlessMaliciousFilesResponseBodyList {
	s.FirstScanTimestamp = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetHighLight(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.HighLight = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetId(v int64) *ListAgentlessMaliciousFilesResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetInstanceName(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.InstanceName = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetInternetIp(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.InternetIp = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetIntranetIp(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.IntranetIp = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetLatestScanTimestamp(v int64) *ListAgentlessMaliciousFilesResponseBodyList {
	s.LatestScanTimestamp = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetLevel(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.Level = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetMaliciousMd5(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.MaliciousMd5 = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetMaliciousName(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.MaliciousName = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetMaliciousType(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.MaliciousType = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetNotes(v []*ListAgentlessMaliciousFilesResponseBodyListNotes) *ListAgentlessMaliciousFilesResponseBodyList {
	s.Notes = v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetOperateResult(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.OperateResult = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetOperateTimestamp(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.OperateTimestamp = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetPartition(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.Partition = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetTargetId(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.TargetId = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetTargetName(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.TargetName = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetTargetType(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.TargetType = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) SetUuid(v string) *ListAgentlessMaliciousFilesResponseBodyList {
	s.Uuid = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyList) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Notes != nil {
		for _, item := range s.Notes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentlessMaliciousFilesResponseBodyListDetails struct {
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
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAgentlessMaliciousFilesResponseBodyListDetails) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessMaliciousFilesResponseBodyListDetails) GoString() string {
	return s.String()
}

func (s *ListAgentlessMaliciousFilesResponseBodyListDetails) GetName() *string {
	return s.Name
}

func (s *ListAgentlessMaliciousFilesResponseBodyListDetails) GetNameKey() *string {
	return s.NameKey
}

func (s *ListAgentlessMaliciousFilesResponseBodyListDetails) GetType() *string {
	return s.Type
}

func (s *ListAgentlessMaliciousFilesResponseBodyListDetails) GetValue() *string {
	return s.Value
}

func (s *ListAgentlessMaliciousFilesResponseBodyListDetails) SetName(v string) *ListAgentlessMaliciousFilesResponseBodyListDetails {
	s.Name = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyListDetails) SetNameKey(v string) *ListAgentlessMaliciousFilesResponseBodyListDetails {
	s.NameKey = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyListDetails) SetType(v string) *ListAgentlessMaliciousFilesResponseBodyListDetails {
	s.Type = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyListDetails) SetValue(v string) *ListAgentlessMaliciousFilesResponseBodyListDetails {
	s.Value = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyListDetails) Validate() error {
	return dara.Validate(s)
}

type ListAgentlessMaliciousFilesResponseBodyListNotes struct {
	// The remark.
	//
	// example:
	//
	// Serious***
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The ID of the remark.
	//
	// example:
	//
	// 50****
	NoteId *string `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// The time when the remark was created.
	//
	// example:
	//
	// 2023-04-27 14:49:21
	NoteTime *string `json:"NoteTime,omitempty" xml:"NoteTime,omitempty"`
}

func (s ListAgentlessMaliciousFilesResponseBodyListNotes) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessMaliciousFilesResponseBodyListNotes) GoString() string {
	return s.String()
}

func (s *ListAgentlessMaliciousFilesResponseBodyListNotes) GetNote() *string {
	return s.Note
}

func (s *ListAgentlessMaliciousFilesResponseBodyListNotes) GetNoteId() *string {
	return s.NoteId
}

func (s *ListAgentlessMaliciousFilesResponseBodyListNotes) GetNoteTime() *string {
	return s.NoteTime
}

func (s *ListAgentlessMaliciousFilesResponseBodyListNotes) SetNote(v string) *ListAgentlessMaliciousFilesResponseBodyListNotes {
	s.Note = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyListNotes) SetNoteId(v string) *ListAgentlessMaliciousFilesResponseBodyListNotes {
	s.NoteId = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyListNotes) SetNoteTime(v string) *ListAgentlessMaliciousFilesResponseBodyListNotes {
	s.NoteTime = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyListNotes) Validate() error {
	return dara.Validate(s)
}

type ListAgentlessMaliciousFilesResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
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
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentlessMaliciousFilesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessMaliciousFilesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListAgentlessMaliciousFilesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListAgentlessMaliciousFilesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAgentlessMaliciousFilesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentlessMaliciousFilesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentlessMaliciousFilesResponseBodyPageInfo) SetCount(v int32) *ListAgentlessMaliciousFilesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyPageInfo) SetCurrentPage(v int32) *ListAgentlessMaliciousFilesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyPageInfo) SetPageSize(v int32) *ListAgentlessMaliciousFilesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyPageInfo) SetTotalCount(v int32) *ListAgentlessMaliciousFilesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListAgentlessMaliciousFilesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
