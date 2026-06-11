// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNumberDistrictInfoTemplateDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetNumberDistrictInfoTemplateDownloadUrlResponseBody
	GetCode() *string
	SetFileHttpUrl(v string) *GetNumberDistrictInfoTemplateDownloadUrlResponseBody
	GetFileHttpUrl() *string
	SetHttpStatusCode(v int32) *GetNumberDistrictInfoTemplateDownloadUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetNumberDistrictInfoTemplateDownloadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNumberDistrictInfoTemplateDownloadUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNumberDistrictInfoTemplateDownloadUrlResponseBody
	GetSuccess() *bool
}

type GetNumberDistrictInfoTemplateDownloadUrlResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// File URL
	//
	// example:
	//
	// http://xxx.xx.com/xx
	FileHttpUrl *string `json:"FileHttpUrl,omitempty" xml:"FileHttpUrl,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNumberDistrictInfoTemplateDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNumberDistrictInfoTemplateDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) GetFileHttpUrl() *string {
	return s.FileHttpUrl
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) SetCode(v string) *GetNumberDistrictInfoTemplateDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) SetFileHttpUrl(v string) *GetNumberDistrictInfoTemplateDownloadUrlResponseBody {
	s.FileHttpUrl = &v
	return s
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) SetHttpStatusCode(v int32) *GetNumberDistrictInfoTemplateDownloadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) SetMessage(v string) *GetNumberDistrictInfoTemplateDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) SetRequestId(v string) *GetNumberDistrictInfoTemplateDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) SetSuccess(v bool) *GetNumberDistrictInfoTemplateDownloadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetNumberDistrictInfoTemplateDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
