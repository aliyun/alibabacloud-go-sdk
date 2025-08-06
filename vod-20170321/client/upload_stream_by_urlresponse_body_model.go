// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadStreamByURLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileURL(v string) *UploadStreamByURLResponseBody
	GetFileURL() *string
	SetRequestId(v string) *UploadStreamByURLResponseBody
	GetRequestId() *string
	SetSourceURL(v string) *UploadStreamByURLResponseBody
	GetSourceURL() *string
	SetStreamJobId(v string) *UploadStreamByURLResponseBody
	GetStreamJobId() *string
}

type UploadStreamByURLResponseBody struct {
	// The URL of the OSS object.
	//
	// example:
	//
	// http://outin-31059bcee7810a200163e1c8dba****.oss-cn-shanghai.aliyuncs.com/lesson-01.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7AE96389-DF1E-598D-816B-7B40F13B4620
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The URL of the input stream. This parameter is used when you call the [GetURLUploadInfos](https://help.aliyun.com/document_detail/106830.html) operation.
	//
	// example:
	//
	// https://example.com/lesson-01.mp4
	SourceURL *string `json:"SourceURL,omitempty" xml:"SourceURL,omitempty"`
	// The ID of the stream upload job. This parameter is used when you call the [GetURLUploadInfos](https://help.aliyun.com/document_detail/106830.html) operation.
	//
	// In ApsaraVideo VOD, you can upload only one transcoded stream in an upload job. For more information, see the PlayInfo: the playback information about a video stream section in [Basic structures](https://help.aliyun.com/document_detail/52839.html).
	//
	// example:
	//
	// e304b34fb3d959f92baef97b6496****
	StreamJobId *string `json:"StreamJobId,omitempty" xml:"StreamJobId,omitempty"`
}

func (s UploadStreamByURLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadStreamByURLResponseBody) GoString() string {
	return s.String()
}

func (s *UploadStreamByURLResponseBody) GetFileURL() *string {
	return s.FileURL
}

func (s *UploadStreamByURLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadStreamByURLResponseBody) GetSourceURL() *string {
	return s.SourceURL
}

func (s *UploadStreamByURLResponseBody) GetStreamJobId() *string {
	return s.StreamJobId
}

func (s *UploadStreamByURLResponseBody) SetFileURL(v string) *UploadStreamByURLResponseBody {
	s.FileURL = &v
	return s
}

func (s *UploadStreamByURLResponseBody) SetRequestId(v string) *UploadStreamByURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadStreamByURLResponseBody) SetSourceURL(v string) *UploadStreamByURLResponseBody {
	s.SourceURL = &v
	return s
}

func (s *UploadStreamByURLResponseBody) SetStreamJobId(v string) *UploadStreamByURLResponseBody {
	s.StreamJobId = &v
	return s
}

func (s *UploadStreamByURLResponseBody) Validate() error {
	return dara.Validate(s)
}
