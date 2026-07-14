// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMarketingFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddMarketingFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddMarketingFlowResponseBody
	GetCode() *string
	SetData(v string) *AddMarketingFlowResponseBody
	GetData() *string
	SetMessage(v string) *AddMarketingFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddMarketingFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddMarketingFlowResponseBody
	GetSuccess() *bool
}

type AddMarketingFlowResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddMarketingFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMarketingFlowResponseBody) GoString() string {
	return s.String()
}

func (s *AddMarketingFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddMarketingFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddMarketingFlowResponseBody) GetData() *string {
	return s.Data
}

func (s *AddMarketingFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddMarketingFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMarketingFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddMarketingFlowResponseBody) SetAccessDeniedDetail(v string) *AddMarketingFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddMarketingFlowResponseBody) SetCode(v string) *AddMarketingFlowResponseBody {
	s.Code = &v
	return s
}

func (s *AddMarketingFlowResponseBody) SetData(v string) *AddMarketingFlowResponseBody {
	s.Data = &v
	return s
}

func (s *AddMarketingFlowResponseBody) SetMessage(v string) *AddMarketingFlowResponseBody {
	s.Message = &v
	return s
}

func (s *AddMarketingFlowResponseBody) SetRequestId(v string) *AddMarketingFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMarketingFlowResponseBody) SetSuccess(v bool) *AddMarketingFlowResponseBody {
	s.Success = &v
	return s
}

func (s *AddMarketingFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
