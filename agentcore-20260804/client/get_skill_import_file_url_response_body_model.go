// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillImportFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSkillImportFileUrlResponseBodyData) *GetSkillImportFileUrlResponseBody
	GetData() *GetSkillImportFileUrlResponseBodyData
	SetRequestId(v string) *GetSkillImportFileUrlResponseBody
	GetRequestId() *string
}

type GetSkillImportFileUrlResponseBody struct {
	// The response data.
	Data *GetSkillImportFileUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSkillImportFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillImportFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillImportFileUrlResponseBody) GetData() *GetSkillImportFileUrlResponseBodyData {
	return s.Data
}

func (s *GetSkillImportFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillImportFileUrlResponseBody) SetData(v *GetSkillImportFileUrlResponseBodyData) *GetSkillImportFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillImportFileUrlResponseBody) SetRequestId(v string) *GetSkillImportFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillImportFileUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillImportFileUrlResponseBodyData struct {
	// The Content-Type of the upload file.
	//
	// example:
	//
	// application/zip
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The maximum file size allowed for upload, in bytes.
	//
	// example:
	//
	// 10485760
	MaxSize *string `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	// The OSS object name.
	//
	// example:
	//
	// imports/example.zip
	OssObjectName *string `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
	// The OSS pre-signed upload URL.
	//
	// example:
	//
	// https://example.com/artifacts/example.zip
	UploadUrl *string `json:"uploadUrl,omitempty" xml:"uploadUrl,omitempty"`
}

func (s GetSkillImportFileUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSkillImportFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillImportFileUrlResponseBodyData) GetContentType() *string {
	return s.ContentType
}

func (s *GetSkillImportFileUrlResponseBodyData) GetMaxSize() *string {
	return s.MaxSize
}

func (s *GetSkillImportFileUrlResponseBodyData) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *GetSkillImportFileUrlResponseBodyData) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *GetSkillImportFileUrlResponseBodyData) SetContentType(v string) *GetSkillImportFileUrlResponseBodyData {
	s.ContentType = &v
	return s
}

func (s *GetSkillImportFileUrlResponseBodyData) SetMaxSize(v string) *GetSkillImportFileUrlResponseBodyData {
	s.MaxSize = &v
	return s
}

func (s *GetSkillImportFileUrlResponseBodyData) SetOssObjectName(v string) *GetSkillImportFileUrlResponseBodyData {
	s.OssObjectName = &v
	return s
}

func (s *GetSkillImportFileUrlResponseBodyData) SetUploadUrl(v string) *GetSkillImportFileUrlResponseBodyData {
	s.UploadUrl = &v
	return s
}

func (s *GetSkillImportFileUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
