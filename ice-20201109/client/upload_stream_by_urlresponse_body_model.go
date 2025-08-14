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
	SetJobId(v string) *UploadStreamByURLResponseBody
	GetJobId() *string
	SetMediaId(v string) *UploadStreamByURLResponseBody
	GetMediaId() *string
	SetRequestId(v string) *UploadStreamByURLResponseBody
	GetRequestId() *string
	SetSourceURL(v string) *UploadStreamByURLResponseBody
	GetSourceURL() *string
}

type UploadStreamByURLResponseBody struct {
	// The OSS URL of the file.
	//
	// example:
	//
	// http://outin-***.oss-cn-shanghai.aliyuncs.com/stream/48555e8b-181dd5a8c07/48555e8b-181dd5a8c07.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The ID of the upload job.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 411bed50018971edb60b0764a0ec6***
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******89-C21D-4B78-AE24-3788B8******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The URL of the source file that is uploaded in the upload job.
	//
	// example:
	//
	// https://example.com/sample-stream.mp4
	SourceURL *string `json:"SourceURL,omitempty" xml:"SourceURL,omitempty"`
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

func (s *UploadStreamByURLResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UploadStreamByURLResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *UploadStreamByURLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadStreamByURLResponseBody) GetSourceURL() *string {
	return s.SourceURL
}

func (s *UploadStreamByURLResponseBody) SetFileURL(v string) *UploadStreamByURLResponseBody {
	s.FileURL = &v
	return s
}

func (s *UploadStreamByURLResponseBody) SetJobId(v string) *UploadStreamByURLResponseBody {
	s.JobId = &v
	return s
}

func (s *UploadStreamByURLResponseBody) SetMediaId(v string) *UploadStreamByURLResponseBody {
	s.MediaId = &v
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

func (s *UploadStreamByURLResponseBody) Validate() error {
	return dara.Validate(s)
}
