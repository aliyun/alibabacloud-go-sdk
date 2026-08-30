// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDlpOutboundLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*ListDlpOutboundLogsResponseBodyLogs) *ListDlpOutboundLogsResponseBody
	GetLogs() []*ListDlpOutboundLogsResponseBodyLogs
	SetRequestId(v string) *ListDlpOutboundLogsResponseBody
	GetRequestId() *string
	SetTotalNumber(v int32) *ListDlpOutboundLogsResponseBody
	GetTotalNumber() *int32
}

type ListDlpOutboundLogsResponseBody struct {
	// The log objects.
	Logs []*ListDlpOutboundLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9350F06D-80FB-5B8B-983A-0E25CBD08479
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that match the query conditions.
	//
	// example:
	//
	// 120
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s ListDlpOutboundLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDlpOutboundLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDlpOutboundLogsResponseBody) GetLogs() []*ListDlpOutboundLogsResponseBodyLogs {
	return s.Logs
}

func (s *ListDlpOutboundLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDlpOutboundLogsResponseBody) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *ListDlpOutboundLogsResponseBody) SetLogs(v []*ListDlpOutboundLogsResponseBodyLogs) *ListDlpOutboundLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListDlpOutboundLogsResponseBody) SetRequestId(v string) *ListDlpOutboundLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBody) SetTotalNumber(v int32) *ListDlpOutboundLogsResponseBody {
	s.TotalNumber = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBody) Validate() error {
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDlpOutboundLogsResponseBodyLogs struct {
	// The file category.
	//
	// example:
	//
	// Document
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The primary channel ID.
	//
	// example:
	//
	// 3
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The primary channel.
	//
	// example:
	//
	// Web Upload
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The department.
	//
	// example:
	//
	// R&D Department
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The local path of the file on the device.
	//
	// example:
	//
	// C:\\Users\\a\\Desktop\\x.xlsx
	DevFilePath *string `json:"DevFilePath,omitempty" xml:"DevFilePath,omitempty"`
	// The unique identifier of the device.
	//
	// example:
	//
	// 69FFB8EB-3E5E-5F11-A9D3-59ED8CBB345D
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// The device type.
	//
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The storage policy ID.
	//
	// example:
	//
	// 12345
	DlpStorageConfigId *string `json:"DlpStorageConfigId,omitempty" xml:"DlpStorageConfigId,omitempty"`
	// The storage type.
	//
	// example:
	//
	// OSS
	DlpStorageType *string `json:"DlpStorageType,omitempty" xml:"DlpStorageType,omitempty"`
	// The outbound destination address or URL.
	//
	// example:
	//
	// https://mail.example.com
	DstAddr *string `json:"DstAddr,omitempty" xml:"DstAddr,omitempty"`
	// The file source.
	//
	// example:
	//
	// download
	FileOrigin *string `json:"FileOrigin,omitempty" xml:"FileOrigin,omitempty"`
	// The host of the file source.
	//
	// example:
	//
	// drive.example.com
	FileOriginHost *string `json:"FileOriginHost,omitempty" xml:"FileOriginHost,omitempty"`
	// The referrer of the file source.
	//
	// example:
	//
	// https://drive.example.com/
	FileOriginReferrer *string `json:"FileOriginReferrer,omitempty" xml:"FileOriginReferrer,omitempty"`
	// The file content preview snippet.
	//
	// example:
	//
	// preview
	FilePreview *string `json:"FilePreview,omitempty" xml:"FilePreview,omitempty"`
	// The file type.
	//
	// example:
	//
	// xlsx
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The organizational structure path.
	//
	// example:
	//
	// R&D Center-Data Platform Team
	GroupInfo *string `json:"GroupInfo,omitempty" xml:"GroupInfo,omitempty"`
	// The hostname of the device.
	//
	// example:
	//
	// MacBook-Pro
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// LogId
	//
	// example:
	//
	// da817m4mfrcs6xxxx3hg
	LogId *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// The matched dictionary statistics.
	MatchedDicts *ListDlpOutboundLogsResponseBodyLogsMatchedDicts `json:"MatchedDicts,omitempty" xml:"MatchedDicts,omitempty" type:"Struct"`
	// The list of matched policy details.
	MatchedPolicies []*ListDlpOutboundLogsResponseBodyLogsMatchedPolicies `json:"MatchedPolicies,omitempty" xml:"MatchedPolicies,omitempty" type:"Repeated"`
	// The object path of the sensitive file in the storage bucket.
	//
	// example:
	//
	// dlp/2026/08/10/xxx.xlsx
	OssFileName *string `json:"OssFileName,omitempty" xml:"OssFileName,omitempty"`
	// The policy action.
	//
	// example:
	//
	// block_and_hint
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// The list of matched policy names.
	PolicyName []*string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty" type:"Repeated"`
	// The outbound process name.
	//
	// example:
	//
	// chrome.exe
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The secondary channel description.
	//
	// example:
	//
	// WeCom-FileSend
	ProcessNameDesc *string `json:"ProcessNameDesc,omitempty" xml:"ProcessNameDesc,omitempty"`
	// The risk level.
	//
	// example:
	//
	// low
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The matched risk scenario.
	//
	// example:
	//
	// file
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The storage path of the screenshot file.
	//
	// example:
	//
	// dlp/screen/xxx.png
	ScreenFilePath *string `json:"ScreenFilePath,omitempty" xml:"ScreenFilePath,omitempty"`
	// The file size.
	//
	// example:
	//
	// 1048576
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The original file name.
	//
	// example:
	//
	// 2026Q2FinancialReport.xlsx
	SrcFileName *string `json:"SrcFileName,omitempty" xml:"SrcFileName,omitempty"`
	// The source IP address of the device.
	//
	// example:
	//
	// 10.0.0.5
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// The time when the event occurred.
	//
	// example:
	//
	// 2026-08-10 14:03:22
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time when the log was reported.
	//
	// example:
	//
	// 2026-08-10 14:03:25
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
	// The username.
	//
	// example:
	//
	// zhangsan
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ListDlpOutboundLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s ListDlpOutboundLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetCategory() *string {
	return s.Category
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetChannelId() *string {
	return s.ChannelId
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetDepartment() *string {
	return s.Department
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetDevFilePath() *string {
	return s.DevFilePath
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetDeviceTag() *string {
	return s.DeviceTag
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetDlpStorageConfigId() *string {
	return s.DlpStorageConfigId
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetDlpStorageType() *string {
	return s.DlpStorageType
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetDstAddr() *string {
	return s.DstAddr
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetFileOrigin() *string {
	return s.FileOrigin
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetFileOriginHost() *string {
	return s.FileOriginHost
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetFileOriginReferrer() *string {
	return s.FileOriginReferrer
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetFilePreview() *string {
	return s.FilePreview
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetFormat() *string {
	return s.Format
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetGroupInfo() *string {
	return s.GroupInfo
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetHostName() *string {
	return s.HostName
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetLogId() *string {
	return s.LogId
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetMatchedDicts() *ListDlpOutboundLogsResponseBodyLogsMatchedDicts {
	return s.MatchedDicts
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetMatchedPolicies() []*ListDlpOutboundLogsResponseBodyLogsMatchedPolicies {
	return s.MatchedPolicies
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetOssFileName() *string {
	return s.OssFileName
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetPolicyAction() *string {
	return s.PolicyAction
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetPolicyName() []*string {
	return s.PolicyName
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetProcessName() *string {
	return s.ProcessName
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetProcessNameDesc() *string {
	return s.ProcessNameDesc
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetScene() *string {
	return s.Scene
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetScreenFilePath() *string {
	return s.ScreenFilePath
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetSize() *string {
	return s.Size
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetSrcFileName() *string {
	return s.SrcFileName
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetSrcIp() *string {
	return s.SrcIp
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetUploadTime() *string {
	return s.UploadTime
}

func (s *ListDlpOutboundLogsResponseBodyLogs) GetUser() *string {
	return s.User
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetCategory(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.Category = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetChannelId(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.ChannelId = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetChannelType(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.ChannelType = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetDepartment(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.Department = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetDevFilePath(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.DevFilePath = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetDeviceTag(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.DeviceTag = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetDeviceType(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.DeviceType = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetDlpStorageConfigId(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.DlpStorageConfigId = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetDlpStorageType(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.DlpStorageType = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetDstAddr(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.DstAddr = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetFileOrigin(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.FileOrigin = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetFileOriginHost(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.FileOriginHost = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetFileOriginReferrer(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.FileOriginReferrer = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetFilePreview(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.FilePreview = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetFormat(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.Format = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetGroupInfo(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.GroupInfo = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetHostName(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.HostName = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetLogId(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.LogId = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetMatchedDicts(v *ListDlpOutboundLogsResponseBodyLogsMatchedDicts) *ListDlpOutboundLogsResponseBodyLogs {
	s.MatchedDicts = v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetMatchedPolicies(v []*ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) *ListDlpOutboundLogsResponseBodyLogs {
	s.MatchedPolicies = v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetOssFileName(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.OssFileName = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetPolicyAction(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.PolicyAction = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetPolicyName(v []*string) *ListDlpOutboundLogsResponseBodyLogs {
	s.PolicyName = v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetProcessName(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.ProcessName = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetProcessNameDesc(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.ProcessNameDesc = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetRiskLevel(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.RiskLevel = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetScene(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.Scene = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetScreenFilePath(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.ScreenFilePath = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetSize(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.Size = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetSrcFileName(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.SrcFileName = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetSrcIp(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.SrcIp = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetStartTime(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.StartTime = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetUploadTime(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.UploadTime = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) SetUser(v string) *ListDlpOutboundLogsResponseBodyLogs {
	s.User = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogs) Validate() error {
	if s.MatchedDicts != nil {
		if err := s.MatchedDicts.Validate(); err != nil {
			return err
		}
	}
	if s.MatchedPolicies != nil {
		for _, item := range s.MatchedPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDlpOutboundLogsResponseBodyLogsMatchedDicts struct {
	// The matched built-in dictionaries.
	InnerDicts []*ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts `json:"InnerDicts,omitempty" xml:"InnerDicts,omitempty" type:"Repeated"`
	// The matched keywords.
	Keywords []*ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	// The matched built-in dictionaries.
	UserDicts []*ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts `json:"UserDicts,omitempty" xml:"UserDicts,omitempty" type:"Repeated"`
}

func (s ListDlpOutboundLogsResponseBodyLogsMatchedDicts) String() string {
	return dara.Prettify(s)
}

func (s ListDlpOutboundLogsResponseBodyLogsMatchedDicts) GoString() string {
	return s.String()
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDicts) GetInnerDicts() []*ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts {
	return s.InnerDicts
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDicts) GetKeywords() []*ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords {
	return s.Keywords
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDicts) GetUserDicts() []*ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts {
	return s.UserDicts
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDicts) SetInnerDicts(v []*ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts) *ListDlpOutboundLogsResponseBodyLogsMatchedDicts {
	s.InnerDicts = v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDicts) SetKeywords(v []*ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords) *ListDlpOutboundLogsResponseBodyLogsMatchedDicts {
	s.Keywords = v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDicts) SetUserDicts(v []*ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts) *ListDlpOutboundLogsResponseBodyLogsMatchedDicts {
	s.UserDicts = v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDicts) Validate() error {
	if s.InnerDicts != nil {
		for _, item := range s.InnerDicts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Keywords != nil {
		for _, item := range s.Keywords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserDicts != nil {
		for _, item := range s.UserDicts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts struct {
	// The number of matches.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The dictionary name.
	//
	// example:
	//
	// Identify Number
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts) String() string {
	return dara.Prettify(s)
}

func (s ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts) GoString() string {
	return s.String()
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts) GetCount() *int64 {
	return s.Count
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts) GetName() *string {
	return s.Name
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts) SetCount(v int64) *ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts {
	s.Count = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts) SetName(v string) *ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts {
	s.Name = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsInnerDicts) Validate() error {
	return dara.Validate(s)
}

type ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords struct {
	// The number of matches.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The keyword name.
	//
	// example:
	//
	// Report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords) String() string {
	return dara.Prettify(s)
}

func (s ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords) GoString() string {
	return s.String()
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords) GetCount() *int64 {
	return s.Count
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords) GetName() *string {
	return s.Name
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords) SetCount(v int64) *ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords {
	s.Count = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords) SetName(v string) *ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords {
	s.Name = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsKeywords) Validate() error {
	return dara.Validate(s)
}

type ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts struct {
	// The number of matches.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The dictionary name.
	//
	// example:
	//
	// Work Number
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts) String() string {
	return dara.Prettify(s)
}

func (s ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts) GoString() string {
	return s.String()
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts) GetCount() *int64 {
	return s.Count
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts) GetName() *string {
	return s.Name
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts) SetCount(v int64) *ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts {
	s.Count = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts) SetName(v string) *ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts {
	s.Name = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedDictsUserDicts) Validate() error {
	return dara.Validate(s)
}

type ListDlpOutboundLogsResponseBodyLogsMatchedPolicies struct {
	// The corresponding data level.
	//
	// example:
	//
	// L3
	EngineDataLevel *string `json:"EngineDataLevel,omitempty" xml:"EngineDataLevel,omitempty"`
	// The corresponding data type.
	//
	// example:
	//
	// Company Data
	EngineDataType *string `json:"EngineDataType,omitempty" xml:"EngineDataType,omitempty"`
	// The matched detection rule name.
	//
	// example:
	//
	// ID No
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The policy description.
	//
	// example:
	//
	// gitlab-tag-policy
	PolicyDesc *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	// The matched policy name.
	//
	// example:
	//
	// default
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) GoString() string {
	return s.String()
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) GetEngineDataLevel() *string {
	return s.EngineDataLevel
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) GetEngineDataType() *string {
	return s.EngineDataType
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) GetEngineName() *string {
	return s.EngineName
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) GetPolicyDesc() *string {
	return s.PolicyDesc
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) SetEngineDataLevel(v string) *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies {
	s.EngineDataLevel = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) SetEngineDataType(v string) *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies {
	s.EngineDataType = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) SetEngineName(v string) *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies {
	s.EngineName = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) SetPolicyDesc(v string) *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies {
	s.PolicyDesc = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) SetPolicyName(v string) *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies {
	s.PolicyName = &v
	return s
}

func (s *ListDlpOutboundLogsResponseBodyLogsMatchedPolicies) Validate() error {
	return dara.Validate(s)
}
