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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
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
