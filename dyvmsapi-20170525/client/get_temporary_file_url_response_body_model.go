// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemporaryFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetTemporaryFileUrlResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetTemporaryFileUrlResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *GetTemporaryFileUrlResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *GetTemporaryFileUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTemporaryFileUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTemporaryFileUrlResponseBody
	GetSuccess() *bool
}

type GetTemporaryFileUrlResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTemporaryFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemporaryFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemporaryFileUrlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetTemporaryFileUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTemporaryFileUrlResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetTemporaryFileUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTemporaryFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemporaryFileUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTemporaryFileUrlResponseBody) SetAccessDeniedDetail(v string) *GetTemporaryFileUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetTemporaryFileUrlResponseBody) SetCode(v string) *GetTemporaryFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemporaryFileUrlResponseBody) SetData(v map[string]interface{}) *GetTemporaryFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetTemporaryFileUrlResponseBody) SetMessage(v string) *GetTemporaryFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetTemporaryFileUrlResponseBody) SetRequestId(v string) *GetTemporaryFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemporaryFileUrlResponseBody) SetSuccess(v bool) *GetTemporaryFileUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetTemporaryFileUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
