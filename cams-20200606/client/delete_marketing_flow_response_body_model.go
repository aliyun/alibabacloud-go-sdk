// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMarketingFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteMarketingFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteMarketingFlowResponseBody
	GetCode() *string
	SetData(v string) *DeleteMarketingFlowResponseBody
	GetData() *string
	SetMessage(v string) *DeleteMarketingFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMarketingFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMarketingFlowResponseBody
	GetSuccess() *bool
}

type DeleteMarketingFlowResponseBody struct {
	// The access denied details.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data object.
	//
	// example:
	//
	// {}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. Alibaba Cloud generates a unique identifier for each request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
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

func (s DeleteMarketingFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMarketingFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMarketingFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteMarketingFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMarketingFlowResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteMarketingFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMarketingFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMarketingFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMarketingFlowResponseBody) SetAccessDeniedDetail(v string) *DeleteMarketingFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteMarketingFlowResponseBody) SetCode(v string) *DeleteMarketingFlowResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMarketingFlowResponseBody) SetData(v string) *DeleteMarketingFlowResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteMarketingFlowResponseBody) SetMessage(v string) *DeleteMarketingFlowResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMarketingFlowResponseBody) SetRequestId(v string) *DeleteMarketingFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMarketingFlowResponseBody) SetSuccess(v bool) *DeleteMarketingFlowResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMarketingFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
