// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadCdsFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadFileModel(v *DownloadCdsFileResponseBodyDownloadFileModel) *DownloadCdsFileResponseBody
	GetDownloadFileModel() *DownloadCdsFileResponseBodyDownloadFileModel
	SetMessage(v string) *DownloadCdsFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *DownloadCdsFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DownloadCdsFileResponseBody
	GetSuccess() *bool
}

type DownloadCdsFileResponseBody struct {
	// The download URL of the file.
	DownloadFileModel *DownloadCdsFileResponseBodyDownloadFileModel `json:"DownloadFileModel,omitempty" xml:"DownloadFileModel,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E3ED9519-DD73-5C86-9C0A-43C9281C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DownloadCdsFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadCdsFileResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadCdsFileResponseBody) GetDownloadFileModel() *DownloadCdsFileResponseBodyDownloadFileModel {
	return s.DownloadFileModel
}

func (s *DownloadCdsFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DownloadCdsFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadCdsFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DownloadCdsFileResponseBody) SetDownloadFileModel(v *DownloadCdsFileResponseBodyDownloadFileModel) *DownloadCdsFileResponseBody {
	s.DownloadFileModel = v
	return s
}

func (s *DownloadCdsFileResponseBody) SetMessage(v string) *DownloadCdsFileResponseBody {
	s.Message = &v
	return s
}

func (s *DownloadCdsFileResponseBody) SetRequestId(v string) *DownloadCdsFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadCdsFileResponseBody) SetSuccess(v bool) *DownloadCdsFileResponseBody {
	s.Success = &v
	return s
}

func (s *DownloadCdsFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type DownloadCdsFileResponseBodyDownloadFileModel struct {
	// This parameter is deprecated.
	DownloadType *string `json:"DownloadType,omitempty" xml:"DownloadType,omitempty"`
	// The download URL.
	//
	// example:
	//
	// https://pds-XXXX-bj-1693807057.oss-cn-beijing.aliyuncs.com/A0SKfLOp%2F2%2F6662612e0570fb2bdd5549759716d433439f0572%2F6662612ee3804e4901794928b14f9a7477640ee7?di=XXXX&dr=1030&f=667d5a322ebf7409e91c485d808fb3bd8a73efbb&response-content-disposition=attachment%3B%20
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// This parameter is deprecated.
	ExpirationSecond *string `json:"ExpirationSecond,omitempty" xml:"ExpirationSecond,omitempty"`
	// The validity period of the download URL.
	//
	// example:
	//
	// 2024-07-18T02:55:49.795Z
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 667d5a322ebf7409e91c485d808fb3bd8a73efbb
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// example:
	//
	// 1594642
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// This parameter is deprecated.
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
}

func (s DownloadCdsFileResponseBodyDownloadFileModel) String() string {
	return dara.Prettify(s)
}

func (s DownloadCdsFileResponseBodyDownloadFileModel) GoString() string {
	return s.String()
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) GetDownloadType() *string {
	return s.DownloadType
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) GetExpirationSecond() *string {
	return s.ExpirationSecond
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) GetFileId() *string {
	return s.FileId
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) GetSize() *int64 {
	return s.Size
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) SetDownloadType(v string) *DownloadCdsFileResponseBodyDownloadFileModel {
	s.DownloadType = &v
	return s
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) SetDownloadUrl(v string) *DownloadCdsFileResponseBodyDownloadFileModel {
	s.DownloadUrl = &v
	return s
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) SetExpirationSecond(v string) *DownloadCdsFileResponseBodyDownloadFileModel {
	s.ExpirationSecond = &v
	return s
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) SetExpirationTime(v string) *DownloadCdsFileResponseBodyDownloadFileModel {
	s.ExpirationTime = &v
	return s
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) SetFileId(v string) *DownloadCdsFileResponseBodyDownloadFileModel {
	s.FileId = &v
	return s
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) SetSize(v int64) *DownloadCdsFileResponseBodyDownloadFileModel {
	s.Size = &v
	return s
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) SetStreamUrl(v string) *DownloadCdsFileResponseBodyDownloadFileModel {
	s.StreamUrl = &v
	return s
}

func (s *DownloadCdsFileResponseBodyDownloadFileModel) Validate() error {
	return dara.Validate(s)
}
