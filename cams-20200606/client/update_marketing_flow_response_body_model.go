// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMarketingFLowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateMarketingFLowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateMarketingFLowResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *UpdateMarketingFLowResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *UpdateMarketingFLowResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMarketingFLowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMarketingFLowResponseBody
	GetSuccess() *bool
}

type UpdateMarketingFLowResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// true
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ewtrew-fghdfg43564ZZ
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMarketingFLowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMarketingFLowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMarketingFLowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateMarketingFLowResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMarketingFLowResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *UpdateMarketingFLowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMarketingFLowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMarketingFLowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMarketingFLowResponseBody) SetAccessDeniedDetail(v string) *UpdateMarketingFLowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateMarketingFLowResponseBody) SetCode(v string) *UpdateMarketingFLowResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMarketingFLowResponseBody) SetData(v map[string]interface{}) *UpdateMarketingFLowResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMarketingFLowResponseBody) SetMessage(v string) *UpdateMarketingFLowResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMarketingFLowResponseBody) SetRequestId(v string) *UpdateMarketingFLowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMarketingFLowResponseBody) SetSuccess(v bool) *UpdateMarketingFLowResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMarketingFLowResponseBody) Validate() error {
	return dara.Validate(s)
}
