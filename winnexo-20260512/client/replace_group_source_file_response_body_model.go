// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceGroupSourceFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReplaceGroupSourceFileResponseBody
	GetCode() *string
	SetFilePath(v string) *ReplaceGroupSourceFileResponseBody
	GetFilePath() *string
	SetFilePublicUrl(v string) *ReplaceGroupSourceFileResponseBody
	GetFilePublicUrl() *string
	SetFileRecordId(v string) *ReplaceGroupSourceFileResponseBody
	GetFileRecordId() *string
	SetMessage(v string) *ReplaceGroupSourceFileResponseBody
	GetMessage() *string
	SetName(v string) *ReplaceGroupSourceFileResponseBody
	GetName() *string
	SetRequestId(v string) *ReplaceGroupSourceFileResponseBody
	GetRequestId() *string
	SetSourceId(v string) *ReplaceGroupSourceFileResponseBody
	GetSourceId() *string
	SetSourceType(v string) *ReplaceGroupSourceFileResponseBody
	GetSourceType() *string
	SetStatus(v string) *ReplaceGroupSourceFileResponseBody
	GetStatus() *string
}

type ReplaceGroupSourceFileResponseBody struct {
	// The business status code. A value of 200 indicates success. A failure returns a backend error code (ERR.	- / InvalidParameter.*).
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The OSS persistent storage path of the replacement file.
	//
	// example:
	//
	// oss://example/new.txt
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// The OSS persistent storage path of the replacement file.
	//
	// example:
	//
	// https://example.com/new.txt
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The file record ID of the replacement file.
	//
	// example:
	//
	// file_example
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The image name.
	//
	// example:
	//
	// Project resources
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// E68654BD-F7BA-5837-8686-5645D739A47C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// source_example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The data source type. The value is fixed as FILE.
	//
	// example:
	//
	// example
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The data source status. Valid values:
	//
	// - **1**: Online.
	//
	// - **0**: Offline.
	//
	// example:
	//
	// example
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ReplaceGroupSourceFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceGroupSourceFileResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceGroupSourceFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReplaceGroupSourceFileResponseBody) GetFilePath() *string {
	return s.FilePath
}

func (s *ReplaceGroupSourceFileResponseBody) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *ReplaceGroupSourceFileResponseBody) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *ReplaceGroupSourceFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReplaceGroupSourceFileResponseBody) GetName() *string {
	return s.Name
}

func (s *ReplaceGroupSourceFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceGroupSourceFileResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceGroupSourceFileResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *ReplaceGroupSourceFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ReplaceGroupSourceFileResponseBody) SetCode(v string) *ReplaceGroupSourceFileResponseBody {
	s.Code = &v
	return s
}

func (s *ReplaceGroupSourceFileResponseBody) SetFilePath(v string) *ReplaceGroupSourceFileResponseBody {
	s.FilePath = &v
	return s
}

func (s *ReplaceGroupSourceFileResponseBody) SetFilePublicUrl(v string) *ReplaceGroupSourceFileResponseBody {
	s.FilePublicUrl = &v
	return s
}

func (s *ReplaceGroupSourceFileResponseBody) SetFileRecordId(v string) *ReplaceGroupSourceFileResponseBody {
	s.FileRecordId = &v
	return s
}

func (s *ReplaceGroupSourceFileResponseBody) SetMessage(v string) *ReplaceGroupSourceFileResponseBody {
	s.Message = &v
	return s
}

func (s *ReplaceGroupSourceFileResponseBody) SetName(v string) *ReplaceGroupSourceFileResponseBody {
	s.Name = &v
	return s
}

func (s *ReplaceGroupSourceFileResponseBody) SetRequestId(v string) *ReplaceGroupSourceFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceGroupSourceFileResponseBody) SetSourceId(v string) *ReplaceGroupSourceFileResponseBody {
	s.SourceId = &v
	return s
}

func (s *ReplaceGroupSourceFileResponseBody) SetSourceType(v string) *ReplaceGroupSourceFileResponseBody {
	s.SourceType = &v
	return s
}

func (s *ReplaceGroupSourceFileResponseBody) SetStatus(v string) *ReplaceGroupSourceFileResponseBody {
	s.Status = &v
	return s
}

func (s *ReplaceGroupSourceFileResponseBody) Validate() error {
	return dara.Validate(s)
}
