// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseMarketingFLowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *PauseMarketingFLowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *PauseMarketingFLowResponseBody
	GetCode() *string
	SetData(v string) *PauseMarketingFLowResponseBody
	GetData() *string
	SetMessage(v string) *PauseMarketingFLowResponseBody
	GetMessage() *string
	SetRequestId(v string) *PauseMarketingFLowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PauseMarketingFLowResponseBody
	GetSuccess() *bool
}

type PauseMarketingFLowResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
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

func (s PauseMarketingFLowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseMarketingFLowResponseBody) GoString() string {
	return s.String()
}

func (s *PauseMarketingFLowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *PauseMarketingFLowResponseBody) GetCode() *string {
	return s.Code
}

func (s *PauseMarketingFLowResponseBody) GetData() *string {
	return s.Data
}

func (s *PauseMarketingFLowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PauseMarketingFLowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseMarketingFLowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PauseMarketingFLowResponseBody) SetAccessDeniedDetail(v string) *PauseMarketingFLowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *PauseMarketingFLowResponseBody) SetCode(v string) *PauseMarketingFLowResponseBody {
	s.Code = &v
	return s
}

func (s *PauseMarketingFLowResponseBody) SetData(v string) *PauseMarketingFLowResponseBody {
	s.Data = &v
	return s
}

func (s *PauseMarketingFLowResponseBody) SetMessage(v string) *PauseMarketingFLowResponseBody {
	s.Message = &v
	return s
}

func (s *PauseMarketingFLowResponseBody) SetRequestId(v string) *PauseMarketingFLowResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseMarketingFLowResponseBody) SetSuccess(v bool) *PauseMarketingFLowResponseBody {
	s.Success = &v
	return s
}

func (s *PauseMarketingFLowResponseBody) Validate() error {
	return dara.Validate(s)
}
