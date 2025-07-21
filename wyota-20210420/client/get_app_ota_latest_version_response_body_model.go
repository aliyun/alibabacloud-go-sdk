// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppOtaLatestVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAppOtaLatestVersionResponseBody
	GetCode() *string
	SetData(v *GetAppOtaLatestVersionResponseBodyData) *GetAppOtaLatestVersionResponseBody
	GetData() *GetAppOtaLatestVersionResponseBodyData
	SetMessage(v string) *GetAppOtaLatestVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAppOtaLatestVersionResponseBody
	GetRequestId() *string
}

type GetAppOtaLatestVersionResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAppOtaLatestVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppOtaLatestVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppOtaLatestVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppOtaLatestVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAppOtaLatestVersionResponseBody) GetData() *GetAppOtaLatestVersionResponseBodyData {
	return s.Data
}

func (s *GetAppOtaLatestVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAppOtaLatestVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppOtaLatestVersionResponseBody) SetCode(v string) *GetAppOtaLatestVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBody) SetData(v *GetAppOtaLatestVersionResponseBodyData) *GetAppOtaLatestVersionResponseBody {
	s.Data = v
	return s
}

func (s *GetAppOtaLatestVersionResponseBody) SetMessage(v string) *GetAppOtaLatestVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBody) SetRequestId(v string) *GetAppOtaLatestVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAppOtaLatestVersionResponseBodyData struct {
	AppVersion   *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	DownloadUrl  *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ForceUpgrade *int32  `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	Md5          *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	ReleaseNote  *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	TaskUid      *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
	VersionCode  *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionType  *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s GetAppOtaLatestVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAppOtaLatestVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppOtaLatestVersionResponseBodyData) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetAppOtaLatestVersionResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetAppOtaLatestVersionResponseBodyData) GetForceUpgrade() *int32 {
	return s.ForceUpgrade
}

func (s *GetAppOtaLatestVersionResponseBodyData) GetMd5() *string {
	return s.Md5
}

func (s *GetAppOtaLatestVersionResponseBodyData) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *GetAppOtaLatestVersionResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *GetAppOtaLatestVersionResponseBodyData) GetTaskUid() *string {
	return s.TaskUid
}

func (s *GetAppOtaLatestVersionResponseBodyData) GetVersionCode() *string {
	return s.VersionCode
}

func (s *GetAppOtaLatestVersionResponseBodyData) GetVersionType() *string {
	return s.VersionType
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetAppVersion(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetDownloadUrl(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetForceUpgrade(v int32) *GetAppOtaLatestVersionResponseBodyData {
	s.ForceUpgrade = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetMd5(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetReleaseNote(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.ReleaseNote = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetSize(v int64) *GetAppOtaLatestVersionResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetTaskUid(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.TaskUid = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetVersionCode(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.VersionCode = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) SetVersionType(v string) *GetAppOtaLatestVersionResponseBodyData {
	s.VersionType = &v
	return s
}

func (s *GetAppOtaLatestVersionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
