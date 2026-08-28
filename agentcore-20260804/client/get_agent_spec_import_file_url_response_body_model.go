// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpecImportFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAgentSpecImportFileUrlResponseBodyData) *GetAgentSpecImportFileUrlResponseBody
	GetData() *GetAgentSpecImportFileUrlResponseBodyData
	SetRequestId(v string) *GetAgentSpecImportFileUrlResponseBody
	GetRequestId() *string
}

type GetAgentSpecImportFileUrlResponseBody struct {
	// The response data.
	Data *GetAgentSpecImportFileUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAgentSpecImportFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecImportFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentSpecImportFileUrlResponseBody) GetData() *GetAgentSpecImportFileUrlResponseBodyData {
	return s.Data
}

func (s *GetAgentSpecImportFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentSpecImportFileUrlResponseBody) SetData(v *GetAgentSpecImportFileUrlResponseBodyData) *GetAgentSpecImportFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentSpecImportFileUrlResponseBody) SetRequestId(v string) *GetAgentSpecImportFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentSpecImportFileUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentSpecImportFileUrlResponseBodyData struct {
	// The Content-Type used for the OSS PUT request.
	//
	// example:
	//
	// application/zip
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The maximum number of bytes allowed for the upload.
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

func (s GetAgentSpecImportFileUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecImportFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentSpecImportFileUrlResponseBodyData) GetContentType() *string {
	return s.ContentType
}

func (s *GetAgentSpecImportFileUrlResponseBodyData) GetMaxSize() *string {
	return s.MaxSize
}

func (s *GetAgentSpecImportFileUrlResponseBodyData) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *GetAgentSpecImportFileUrlResponseBodyData) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *GetAgentSpecImportFileUrlResponseBodyData) SetContentType(v string) *GetAgentSpecImportFileUrlResponseBodyData {
	s.ContentType = &v
	return s
}

func (s *GetAgentSpecImportFileUrlResponseBodyData) SetMaxSize(v string) *GetAgentSpecImportFileUrlResponseBodyData {
	s.MaxSize = &v
	return s
}

func (s *GetAgentSpecImportFileUrlResponseBodyData) SetOssObjectName(v string) *GetAgentSpecImportFileUrlResponseBodyData {
	s.OssObjectName = &v
	return s
}

func (s *GetAgentSpecImportFileUrlResponseBodyData) SetUploadUrl(v string) *GetAgentSpecImportFileUrlResponseBodyData {
	s.UploadUrl = &v
	return s
}

func (s *GetAgentSpecImportFileUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
