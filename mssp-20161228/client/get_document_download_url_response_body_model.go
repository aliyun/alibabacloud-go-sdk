// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocumentDownloadUrlResponseBody
	GetCode() *string
	SetData(v string) *GetDocumentDownloadUrlResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetDocumentDownloadUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDocumentDownloadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocumentDownloadUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDocumentDownloadUrlResponseBody
	GetSuccess() *bool
}

type GetDocumentDownloadUrlResponseBody struct {
	// API status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// OSS file access URL.
	//
	// example:
	//
	// https://oos-cn.ctyunapi.cn/example-bucket/test/1.jpg
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Message of the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// C7BE80B4-7692-54FA-AB22-2A7DF08C4754
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful: - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocumentDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentDownloadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocumentDownloadUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *GetDocumentDownloadUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDocumentDownloadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentDownloadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDocumentDownloadUrlResponseBody) SetCode(v string) *GetDocumentDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetData(v string) *GetDocumentDownloadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetHttpStatusCode(v int32) *GetDocumentDownloadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetMessage(v string) *GetDocumentDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetRequestId(v string) *GetDocumentDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetSuccess(v bool) *GetDocumentDownloadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
