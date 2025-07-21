// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOtaInfoTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeviceOtaInfoTestResponseBody
	GetCode() *string
	SetData(v *GetDeviceOtaInfoTestResponseBodyData) *GetDeviceOtaInfoTestResponseBody
	GetData() *GetDeviceOtaInfoTestResponseBodyData
	SetMessage(v string) *GetDeviceOtaInfoTestResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeviceOtaInfoTestResponseBody
	GetRequestId() *string
}

type GetDeviceOtaInfoTestResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDeviceOtaInfoTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceOtaInfoTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaInfoTestResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoTestResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeviceOtaInfoTestResponseBody) GetData() *GetDeviceOtaInfoTestResponseBodyData {
	return s.Data
}

func (s *GetDeviceOtaInfoTestResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeviceOtaInfoTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceOtaInfoTestResponseBody) SetCode(v string) *GetDeviceOtaInfoTestResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBody) SetData(v *GetDeviceOtaInfoTestResponseBodyData) *GetDeviceOtaInfoTestResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBody) SetMessage(v string) *GetDeviceOtaInfoTestResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBody) SetRequestId(v string) *GetDeviceOtaInfoTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeviceOtaInfoTestResponseBodyData struct {
	Version *GetDeviceOtaInfoTestResponseBodyDataVersion `json:"Version,omitempty" xml:"Version,omitempty" type:"Struct"`
}

func (s GetDeviceOtaInfoTestResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaInfoTestResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoTestResponseBodyData) GetVersion() *GetDeviceOtaInfoTestResponseBodyDataVersion {
	return s.Version
}

func (s *GetDeviceOtaInfoTestResponseBodyData) SetVersion(v *GetDeviceOtaInfoTestResponseBodyDataVersion) *GetDeviceOtaInfoTestResponseBodyData {
	s.Version = v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDeviceOtaInfoTestResponseBodyDataVersion struct {
	Creator          *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DownloadUrl      *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ForceUpgrade     *int32  `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	LocalDownloadUrl *string `json:"LocalDownloadUrl,omitempty" xml:"LocalDownloadUrl,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Model            *string `json:"Model,omitempty" xml:"Model,omitempty"`
	ReleaseNote      *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Version          *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionCode      *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionType      *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s GetDeviceOtaInfoTestResponseBodyDataVersion) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaInfoTestResponseBodyDataVersion) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) GetCreator() *string {
	return s.Creator
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) GetForceUpgrade() *int32 {
	return s.ForceUpgrade
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) GetLocalDownloadUrl() *string {
	return s.LocalDownloadUrl
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) GetMd5() *string {
	return s.Md5
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) GetModel() *string {
	return s.Model
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) GetSize() *int64 {
	return s.Size
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) GetVersion() *string {
	return s.Version
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) GetVersionCode() *string {
	return s.VersionCode
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) GetVersionType() *string {
	return s.VersionType
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetCreator(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.Creator = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetDownloadUrl(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.DownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetForceUpgrade(v int32) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.ForceUpgrade = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetLocalDownloadUrl(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.LocalDownloadUrl = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetMd5(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.Md5 = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetModel(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.Model = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetReleaseNote(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.ReleaseNote = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetSize(v int64) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.Size = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetVersion(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.Version = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetVersionCode(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.VersionCode = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) SetVersionType(v string) *GetDeviceOtaInfoTestResponseBodyDataVersion {
	s.VersionType = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponseBodyDataVersion) Validate() error {
	return dara.Validate(s)
}
