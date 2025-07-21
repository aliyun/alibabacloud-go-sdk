// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceVersionDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDeviceVersionDetailResponseBody
	GetCode() *string
	SetData(v []*DescribeDeviceVersionDetailResponseBodyData) *DescribeDeviceVersionDetailResponseBody
	GetData() []*DescribeDeviceVersionDetailResponseBodyData
	SetMessage(v string) *DescribeDeviceVersionDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDeviceVersionDetailResponseBody
	GetRequestId() *string
}

type DescribeDeviceVersionDetailResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeDeviceVersionDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeviceVersionDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceVersionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceVersionDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDeviceVersionDetailResponseBody) GetData() []*DescribeDeviceVersionDetailResponseBodyData {
	return s.Data
}

func (s *DescribeDeviceVersionDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDeviceVersionDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeviceVersionDetailResponseBody) SetCode(v string) *DescribeDeviceVersionDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBody) SetData(v []*DescribeDeviceVersionDetailResponseBodyData) *DescribeDeviceVersionDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBody) SetMessage(v string) *DescribeDeviceVersionDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBody) SetRequestId(v string) *DescribeDeviceVersionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDeviceVersionDetailResponseBodyData struct {
	AndroidHorizontalMultiCnImageDownloadUrl *string `json:"AndroidHorizontalMultiCnImageDownloadUrl,omitempty" xml:"AndroidHorizontalMultiCnImageDownloadUrl,omitempty"`
	AndroidHorizontalMultiEnImageDownloadUrl *string `json:"AndroidHorizontalMultiEnImageDownloadUrl,omitempty" xml:"AndroidHorizontalMultiEnImageDownloadUrl,omitempty"`
	AndroidVerticalMultiCnImageDownloadUrl   *string `json:"AndroidVerticalMultiCnImageDownloadUrl,omitempty" xml:"AndroidVerticalMultiCnImageDownloadUrl,omitempty"`
	AndroidVerticalMultiEnImageDownloadUrl   *string `json:"AndroidVerticalMultiEnImageDownloadUrl,omitempty" xml:"AndroidVerticalMultiEnImageDownloadUrl,omitempty"`
	Channel                                  *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ClientType                               *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	CnImageDownloadUrl                       *string `json:"CnImageDownloadUrl,omitempty" xml:"CnImageDownloadUrl,omitempty"`
	Creator                                  *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DownloadUrl                              *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	EnImageDownloadUrl                       *string `json:"EnImageDownloadUrl,omitempty" xml:"EnImageDownloadUrl,omitempty"`
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
}

func (s DescribeDeviceVersionDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceVersionDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetAndroidHorizontalMultiCnImageDownloadUrl() *string {
	return s.AndroidHorizontalMultiCnImageDownloadUrl
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetAndroidHorizontalMultiEnImageDownloadUrl() *string {
	return s.AndroidHorizontalMultiEnImageDownloadUrl
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetAndroidVerticalMultiCnImageDownloadUrl() *string {
	return s.AndroidVerticalMultiCnImageDownloadUrl
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetAndroidVerticalMultiEnImageDownloadUrl() *string {
	return s.AndroidVerticalMultiEnImageDownloadUrl
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetChannel() *string {
	return s.Channel
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetClientType() *int32 {
	return s.ClientType
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetCnImageDownloadUrl() *string {
	return s.CnImageDownloadUrl
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetEnImageDownloadUrl() *string {
	return s.EnImageDownloadUrl
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetMd5() *string {
	return s.Md5
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetMultiCnImageDownloadUrl() *string {
	return s.MultiCnImageDownloadUrl
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetMultiEnImageDownloadUrl() *string {
	return s.MultiEnImageDownloadUrl
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetReleaseNoteEn() *string {
	return s.ReleaseNoteEn
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeDeviceVersionDetailResponseBodyData) GetVersionType() *string {
	return s.VersionType
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetAndroidHorizontalMultiCnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.AndroidHorizontalMultiCnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetAndroidHorizontalMultiEnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.AndroidHorizontalMultiEnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetAndroidVerticalMultiCnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.AndroidVerticalMultiCnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetAndroidVerticalMultiEnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.AndroidVerticalMultiEnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetChannel(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.Channel = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetClientType(v int32) *DescribeDeviceVersionDetailResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetCnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.CnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetCreator(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.Creator = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetEnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.EnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetMd5(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetModel(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.Model = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetMultiCnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.MultiCnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetMultiEnImageDownloadUrl(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.MultiEnImageDownloadUrl = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetReleaseNote(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetReleaseNoteEn(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.ReleaseNoteEn = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetSize(v int64) *DescribeDeviceVersionDetailResponseBodyData {
	s.Size = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetVersion(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetVersionCode(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.VersionCode = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) SetVersionType(v string) *DescribeDeviceVersionDetailResponseBodyData {
	s.VersionType = &v
	return s
}

func (s *DescribeDeviceVersionDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
