// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOtaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeviceOtaInfoResponseBody
	GetCode() *string
	SetData(v *GetDeviceOtaInfoResponseBodyData) *GetDeviceOtaInfoResponseBody
	GetData() *GetDeviceOtaInfoResponseBodyData
	SetMessage(v string) *GetDeviceOtaInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeviceOtaInfoResponseBody
	GetRequestId() *string
}

type GetDeviceOtaInfoResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDeviceOtaInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceOtaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeviceOtaInfoResponseBody) GetData() *GetDeviceOtaInfoResponseBodyData {
	return s.Data
}

func (s *GetDeviceOtaInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeviceOtaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceOtaInfoResponseBody) SetCode(v string) *GetDeviceOtaInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBody) SetData(v *GetDeviceOtaInfoResponseBodyData) *GetDeviceOtaInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceOtaInfoResponseBody) SetMessage(v string) *GetDeviceOtaInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBody) SetRequestId(v string) *GetDeviceOtaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeviceOtaInfoResponseBodyData struct {
	Version *GetDeviceOtaInfoResponseBodyDataVersion `json:"Version,omitempty" xml:"Version,omitempty" type:"Struct"`
}

func (s GetDeviceOtaInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoResponseBodyData) GetVersion() *GetDeviceOtaInfoResponseBodyDataVersion {
	return s.Version
}

func (s *GetDeviceOtaInfoResponseBodyData) SetVersion(v *GetDeviceOtaInfoResponseBodyDataVersion) *GetDeviceOtaInfoResponseBodyData {
	s.Version = v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDeviceOtaInfoResponseBodyDataVersion struct {
	AndroidHorizontalMultiCnImageDownloadUrl *string `json:"AndroidHorizontalMultiCnImageDownloadUrl,omitempty" xml:"AndroidHorizontalMultiCnImageDownloadUrl,omitempty"`
	AndroidHorizontalMultiEnImageDownloadUrl *string `json:"AndroidHorizontalMultiEnImageDownloadUrl,omitempty" xml:"AndroidHorizontalMultiEnImageDownloadUrl,omitempty"`
	AndroidVerticalMultiCnImageDownloadUrl   *string `json:"AndroidVerticalMultiCnImageDownloadUrl,omitempty" xml:"AndroidVerticalMultiCnImageDownloadUrl,omitempty"`
	AndroidVerticalMultiEnImageDownloadUrl   *string `json:"AndroidVerticalMultiEnImageDownloadUrl,omitempty" xml:"AndroidVerticalMultiEnImageDownloadUrl,omitempty"`
	CnImageDownloadUrl                       *string `json:"CnImageDownloadUrl,omitempty" xml:"CnImageDownloadUrl,omitempty"`
	Creator                                  *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CustomForceUpgrade                       *bool   `json:"CustomForceUpgrade,omitempty" xml:"CustomForceUpgrade,omitempty"`
	DownloadUrl                              *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	EnImageDownloadUrl                       *string `json:"EnImageDownloadUrl,omitempty" xml:"EnImageDownloadUrl,omitempty"`
	ForceUpgrade                             *int32  `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	IsAppDownloadUrl                         *bool   `json:"IsAppDownloadUrl,omitempty" xml:"IsAppDownloadUrl,omitempty"`
	LocalDownloadUrl                         *string `json:"LocalDownloadUrl,omitempty" xml:"LocalDownloadUrl,omitempty"`
	Md5                                      *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Model                                    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	MultiCnImageDownloadUrl                  *string `json:"MultiCnImageDownloadUrl,omitempty" xml:"MultiCnImageDownloadUrl,omitempty"`
	MultiEnImageDownloadUrl                  *string `json:"MultiEnImageDownloadUrl,omitempty" xml:"MultiEnImageDownloadUrl,omitempty"`
	ReleaseNote                              *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseNoteEn                            *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	Size                                     *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Version                                  *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionCode                              *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionType                              *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
	WyForceUpgrade                           *bool   `json:"WyForceUpgrade,omitempty" xml:"WyForceUpgrade,omitempty"`
}

func (s GetDeviceOtaInfoResponseBodyDataVersion) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaInfoResponseBodyDataVersion) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetAndroidHorizontalMultiCnImageDownloadUrl() *string {
	return s.AndroidHorizontalMultiCnImageDownloadUrl
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetAndroidHorizontalMultiEnImageDownloadUrl() *string {
	return s.AndroidHorizontalMultiEnImageDownloadUrl
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetAndroidVerticalMultiCnImageDownloadUrl() *string {
	return s.AndroidVerticalMultiCnImageDownloadUrl
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetAndroidVerticalMultiEnImageDownloadUrl() *string {
	return s.AndroidVerticalMultiEnImageDownloadUrl
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetCnImageDownloadUrl() *string {
	return s.CnImageDownloadUrl
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetCreator() *string {
	return s.Creator
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetCustomForceUpgrade() *bool {
	return s.CustomForceUpgrade
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetEnImageDownloadUrl() *string {
	return s.EnImageDownloadUrl
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetForceUpgrade() *int32 {
	return s.ForceUpgrade
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetIsAppDownloadUrl() *bool {
	return s.IsAppDownloadUrl
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetLocalDownloadUrl() *string {
	return s.LocalDownloadUrl
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetMd5() *string {
	return s.Md5
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetModel() *string {
	return s.Model
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetMultiCnImageDownloadUrl() *string {
	return s.MultiCnImageDownloadUrl
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetMultiEnImageDownloadUrl() *string {
	return s.MultiEnImageDownloadUrl
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetReleaseNoteEn() *string {
	return s.ReleaseNoteEn
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetSize() *int64 {
	return s.Size
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetVersion() *string {
	return s.Version
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetVersionCode() *string {
	return s.VersionCode
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetVersionType() *string {
	return s.VersionType
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) GetWyForceUpgrade() *bool {
	return s.WyForceUpgrade
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetAndroidHorizontalMultiCnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.AndroidHorizontalMultiCnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetAndroidHorizontalMultiEnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.AndroidHorizontalMultiEnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetAndroidVerticalMultiCnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.AndroidVerticalMultiCnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetAndroidVerticalMultiEnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.AndroidVerticalMultiEnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetCnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.CnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetCreator(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.Creator = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetCustomForceUpgrade(v bool) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.CustomForceUpgrade = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.DownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetEnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.EnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetForceUpgrade(v int32) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.ForceUpgrade = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetIsAppDownloadUrl(v bool) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.IsAppDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetLocalDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.LocalDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetMd5(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.Md5 = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetModel(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.Model = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetMultiCnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.MultiCnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetMultiEnImageDownloadUrl(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.MultiEnImageDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetReleaseNote(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.ReleaseNote = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetReleaseNoteEn(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.ReleaseNoteEn = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetSize(v int64) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.Size = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetVersion(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.Version = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetVersionCode(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.VersionCode = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetVersionType(v string) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.VersionType = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) SetWyForceUpgrade(v bool) *GetDeviceOtaInfoResponseBodyDataVersion {
	s.WyForceUpgrade = &v
	return s
}

func (s *GetDeviceOtaInfoResponseBodyDataVersion) Validate() error {
	return dara.Validate(s)
}
