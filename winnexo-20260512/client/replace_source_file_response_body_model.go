// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceSourceFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReplaceSourceFileResponseBody
	GetCode() *string
	SetFilePath(v string) *ReplaceSourceFileResponseBody
	GetFilePath() *string
	SetFilePublicUrl(v string) *ReplaceSourceFileResponseBody
	GetFilePublicUrl() *string
	SetFileRecordId(v string) *ReplaceSourceFileResponseBody
	GetFileRecordId() *string
	SetMessage(v string) *ReplaceSourceFileResponseBody
	GetMessage() *string
	SetName(v string) *ReplaceSourceFileResponseBody
	GetName() *string
	SetRequestId(v string) *ReplaceSourceFileResponseBody
	GetRequestId() *string
	SetSourceId(v string) *ReplaceSourceFileResponseBody
	GetSourceId() *string
	SetSourceType(v string) *ReplaceSourceFileResponseBody
	GetSourceType() *string
	SetStatus(v string) *ReplaceSourceFileResponseBody
	GetStatus() *string
}

type ReplaceSourceFileResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The OSS persistent address of the file after replacement.
	//
	// example:
	//
	// string_value
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// The OSS persistent address of the file after replacement.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The file record ID after replacement.
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The data source type. The value is fixed to FILE.
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The data source status after re-parsing.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ReplaceSourceFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceSourceFileResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceSourceFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReplaceSourceFileResponseBody) GetFilePath() *string {
	return s.FilePath
}

func (s *ReplaceSourceFileResponseBody) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *ReplaceSourceFileResponseBody) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *ReplaceSourceFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReplaceSourceFileResponseBody) GetName() *string {
	return s.Name
}

func (s *ReplaceSourceFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceSourceFileResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceSourceFileResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *ReplaceSourceFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ReplaceSourceFileResponseBody) SetCode(v string) *ReplaceSourceFileResponseBody {
	s.Code = &v
	return s
}

func (s *ReplaceSourceFileResponseBody) SetFilePath(v string) *ReplaceSourceFileResponseBody {
	s.FilePath = &v
	return s
}

func (s *ReplaceSourceFileResponseBody) SetFilePublicUrl(v string) *ReplaceSourceFileResponseBody {
	s.FilePublicUrl = &v
	return s
}

func (s *ReplaceSourceFileResponseBody) SetFileRecordId(v string) *ReplaceSourceFileResponseBody {
	s.FileRecordId = &v
	return s
}

func (s *ReplaceSourceFileResponseBody) SetMessage(v string) *ReplaceSourceFileResponseBody {
	s.Message = &v
	return s
}

func (s *ReplaceSourceFileResponseBody) SetName(v string) *ReplaceSourceFileResponseBody {
	s.Name = &v
	return s
}

func (s *ReplaceSourceFileResponseBody) SetRequestId(v string) *ReplaceSourceFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceSourceFileResponseBody) SetSourceId(v string) *ReplaceSourceFileResponseBody {
	s.SourceId = &v
	return s
}

func (s *ReplaceSourceFileResponseBody) SetSourceType(v string) *ReplaceSourceFileResponseBody {
	s.SourceType = &v
	return s
}

func (s *ReplaceSourceFileResponseBody) SetStatus(v string) *ReplaceSourceFileResponseBody {
	s.Status = &v
	return s
}

func (s *ReplaceSourceFileResponseBody) Validate() error {
	return dara.Validate(s)
}
