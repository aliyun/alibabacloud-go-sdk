// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppOtaVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAppOtaVersionResponseBody
	GetCode() *string
	SetData(v *DescribeAppOtaVersionResponseBodyData) *DescribeAppOtaVersionResponseBody
	GetData() *DescribeAppOtaVersionResponseBodyData
	SetMessage(v string) *DescribeAppOtaVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAppOtaVersionResponseBody
	GetRequestId() *string
}

type DescribeAppOtaVersionResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeAppOtaVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppOtaVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppOtaVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppOtaVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAppOtaVersionResponseBody) GetData() *DescribeAppOtaVersionResponseBodyData {
	return s.Data
}

func (s *DescribeAppOtaVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAppOtaVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppOtaVersionResponseBody) SetCode(v string) *DescribeAppOtaVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBody) SetData(v *DescribeAppOtaVersionResponseBodyData) *DescribeAppOtaVersionResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAppOtaVersionResponseBody) SetMessage(v string) *DescribeAppOtaVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBody) SetRequestId(v string) *DescribeAppOtaVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAppOtaVersionResponseBodyData struct {
	AppOtaInfoDTOList []*DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList `json:"AppOtaInfoDTOList,omitempty" xml:"AppOtaInfoDTOList,omitempty" type:"Repeated"`
}

func (s DescribeAppOtaVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppOtaVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAppOtaVersionResponseBodyData) GetAppOtaInfoDTOList() []*DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	return s.AppOtaInfoDTOList
}

func (s *DescribeAppOtaVersionResponseBodyData) SetAppOtaInfoDTOList(v []*DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) *DescribeAppOtaVersionResponseBodyData {
	s.AppOtaInfoDTOList = v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList struct {
	AppVersion      *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	Channel         *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	DownloadUrl     *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	FullDownloadUrl *string `json:"FullDownloadUrl,omitempty" xml:"FullDownloadUrl,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	OsType          *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OtaType         *int32  `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ProtocolType    *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	ReleaseNote     *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseNoteEn   *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	SessionType     *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	VersionCode     *int64  `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionType     *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	VersionUid      *string `json:"VersionUid,omitempty" xml:"VersionUid,omitempty"`
}

func (s DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GoString() string {
	return s.String()
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetAppVersion() *string {
	return s.AppVersion
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetChannel() *string {
	return s.Channel
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetFullDownloadUrl() *string {
	return s.FullDownloadUrl
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetMd5() *string {
	return s.Md5
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetOsType() *string {
	return s.OsType
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetOtaType() *int32 {
	return s.OtaType
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetProject() *string {
	return s.Project
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetReleaseNoteEn() *string {
	return s.ReleaseNoteEn
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetSessionType() *string {
	return s.SessionType
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetSize() *int64 {
	return s.Size
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetVersionCode() *int64 {
	return s.VersionCode
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetVersionType() *string {
	return s.VersionType
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) GetVersionUid() *string {
	return s.VersionUid
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetAppVersion(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.AppVersion = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetChannel(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.Channel = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetDownloadUrl(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetFullDownloadUrl(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.FullDownloadUrl = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetGmtCreate(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetMd5(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.Md5 = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetOsType(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.OsType = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetOtaType(v int32) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.OtaType = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetProject(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.Project = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetProtocolType(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.ProtocolType = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetReleaseNote(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetReleaseNoteEn(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.ReleaseNoteEn = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetSessionType(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.SessionType = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetSize(v int64) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.Size = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetStatus(v int32) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.Status = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetVersionCode(v int64) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.VersionCode = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetVersionType(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.VersionType = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) SetVersionUid(v string) *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList {
	s.VersionUid = &v
	return s
}

func (s *DescribeAppOtaVersionResponseBodyDataAppOtaInfoDTOList) Validate() error {
	return dara.Validate(s)
}
