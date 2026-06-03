// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDownloadUrlResponseBody
	GetCode() *string
	SetFileHttpUrl(v string) *CreateDownloadUrlResponseBody
	GetFileHttpUrl() *string
	SetHttpStatusCode(v int32) *CreateDownloadUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDownloadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDownloadUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDownloadUrlResponseBody
	GetSuccess() *bool
}

type CreateDownloadUrlResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// http://www.xxx.com/xxx
	FileHttpUrl *string `json:"FileHttpUrl,omitempty" xml:"FileHttpUrl,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EC08CC41-6870-5594-939A-F758F057898F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDownloadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDownloadUrlResponseBody) GetFileHttpUrl() *string {
	return s.FileHttpUrl
}

func (s *CreateDownloadUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDownloadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDownloadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDownloadUrlResponseBody) SetCode(v string) *CreateDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDownloadUrlResponseBody) SetFileHttpUrl(v string) *CreateDownloadUrlResponseBody {
	s.FileHttpUrl = &v
	return s
}

func (s *CreateDownloadUrlResponseBody) SetHttpStatusCode(v int32) *CreateDownloadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDownloadUrlResponseBody) SetMessage(v string) *CreateDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDownloadUrlResponseBody) SetRequestId(v string) *CreateDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDownloadUrlResponseBody) SetSuccess(v bool) *CreateDownloadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
