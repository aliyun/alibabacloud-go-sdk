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
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// Example value example value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// Example value example value.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned with the result.
	//
	// example:
	//
	// Example value.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// Example value.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
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
