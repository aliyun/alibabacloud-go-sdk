// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneratePresignedUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GeneratePresignedUrlResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GeneratePresignedUrlResponseBody
	GetCode() *string
	SetData(v *GeneratePresignedUrlResponseBodyData) *GeneratePresignedUrlResponseBody
	GetData() *GeneratePresignedUrlResponseBodyData
	SetMessage(v string) *GeneratePresignedUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GeneratePresignedUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GeneratePresignedUrlResponseBody
	GetSuccess() *bool
}

type GeneratePresignedUrlResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GeneratePresignedUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. Alibaba Cloud generates a unique identifier for each API request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GeneratePresignedUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GeneratePresignedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GeneratePresignedUrlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GeneratePresignedUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GeneratePresignedUrlResponseBody) GetData() *GeneratePresignedUrlResponseBodyData {
	return s.Data
}

func (s *GeneratePresignedUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GeneratePresignedUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GeneratePresignedUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GeneratePresignedUrlResponseBody) SetAccessDeniedDetail(v string) *GeneratePresignedUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GeneratePresignedUrlResponseBody) SetCode(v string) *GeneratePresignedUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GeneratePresignedUrlResponseBody) SetData(v *GeneratePresignedUrlResponseBodyData) *GeneratePresignedUrlResponseBody {
	s.Data = v
	return s
}

func (s *GeneratePresignedUrlResponseBody) SetMessage(v string) *GeneratePresignedUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GeneratePresignedUrlResponseBody) SetRequestId(v string) *GeneratePresignedUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GeneratePresignedUrlResponseBody) SetSuccess(v bool) *GeneratePresignedUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GeneratePresignedUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GeneratePresignedUrlResponseBodyData struct {
	// The material path.
	//
	// > For Viber, the recommended image size is 800 × 800.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GeneratePresignedUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GeneratePresignedUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GeneratePresignedUrlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GeneratePresignedUrlResponseBodyData) SetUrl(v string) *GeneratePresignedUrlResponseBodyData {
	s.Url = &v
	return s
}

func (s *GeneratePresignedUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
