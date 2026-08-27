// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstagramPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteInstagramPageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteInstagramPageResponseBody
	GetCode() *string
	SetData(v string) *DeleteInstagramPageResponseBody
	GetData() *string
	SetMessage(v string) *DeleteInstagramPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstagramPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstagramPageResponseBody
	GetSuccess() *bool
}

type DeleteInstagramPageResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, refer to [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response message.
	//
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// gfdg435t-hf544**
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

func (s DeleteInstagramPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstagramPageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstagramPageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteInstagramPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstagramPageResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteInstagramPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstagramPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstagramPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstagramPageResponseBody) SetAccessDeniedDetail(v string) *DeleteInstagramPageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteInstagramPageResponseBody) SetCode(v string) *DeleteInstagramPageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstagramPageResponseBody) SetData(v string) *DeleteInstagramPageResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstagramPageResponseBody) SetMessage(v string) *DeleteInstagramPageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstagramPageResponseBody) SetRequestId(v string) *DeleteInstagramPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstagramPageResponseBody) SetSuccess(v bool) *DeleteInstagramPageResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstagramPageResponseBody) Validate() error {
	return dara.Validate(s)
}
