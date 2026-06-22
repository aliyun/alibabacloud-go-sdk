// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPledgeTemplateAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetPledgeTemplateAddressResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetPledgeTemplateAddressResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *GetPledgeTemplateAddressResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *GetPledgeTemplateAddressResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPledgeTemplateAddressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPledgeTemplateAddressResponseBody
	GetSuccess() *bool
}

type GetPledgeTemplateAddressResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPledgeTemplateAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPledgeTemplateAddressResponseBody) GoString() string {
	return s.String()
}

func (s *GetPledgeTemplateAddressResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetPledgeTemplateAddressResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPledgeTemplateAddressResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetPledgeTemplateAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPledgeTemplateAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPledgeTemplateAddressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPledgeTemplateAddressResponseBody) SetAccessDeniedDetail(v string) *GetPledgeTemplateAddressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetPledgeTemplateAddressResponseBody) SetCode(v string) *GetPledgeTemplateAddressResponseBody {
	s.Code = &v
	return s
}

func (s *GetPledgeTemplateAddressResponseBody) SetData(v map[string]interface{}) *GetPledgeTemplateAddressResponseBody {
	s.Data = v
	return s
}

func (s *GetPledgeTemplateAddressResponseBody) SetMessage(v string) *GetPledgeTemplateAddressResponseBody {
	s.Message = &v
	return s
}

func (s *GetPledgeTemplateAddressResponseBody) SetRequestId(v string) *GetPledgeTemplateAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPledgeTemplateAddressResponseBody) SetSuccess(v bool) *GetPledgeTemplateAddressResponseBody {
	s.Success = &v
	return s
}

func (s *GetPledgeTemplateAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
