// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeToRCSSignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpgradeToRCSSignatureResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpgradeToRCSSignatureResponseBody
	GetCode() *string
	SetData(v *UpgradeToRCSSignatureResponseBodyData) *UpgradeToRCSSignatureResponseBody
	GetData() *UpgradeToRCSSignatureResponseBodyData
	SetMessage(v string) *UpgradeToRCSSignatureResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeToRCSSignatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeToRCSSignatureResponseBody
	GetSuccess() *bool
}

type UpgradeToRCSSignatureResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpgradeToRCSSignatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UpgradeToRCSSignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeToRCSSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeToRCSSignatureResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpgradeToRCSSignatureResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpgradeToRCSSignatureResponseBody) GetData() *UpgradeToRCSSignatureResponseBodyData {
	return s.Data
}

func (s *UpgradeToRCSSignatureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeToRCSSignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeToRCSSignatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeToRCSSignatureResponseBody) SetAccessDeniedDetail(v string) *UpgradeToRCSSignatureResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpgradeToRCSSignatureResponseBody) SetCode(v string) *UpgradeToRCSSignatureResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeToRCSSignatureResponseBody) SetData(v *UpgradeToRCSSignatureResponseBodyData) *UpgradeToRCSSignatureResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeToRCSSignatureResponseBody) SetMessage(v string) *UpgradeToRCSSignatureResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeToRCSSignatureResponseBody) SetRequestId(v string) *UpgradeToRCSSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeToRCSSignatureResponseBody) SetSuccess(v bool) *UpgradeToRCSSignatureResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeToRCSSignatureResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeToRCSSignatureResponseBodyData struct {
	// example:
	//
	// 示例值示例值
	ChatbotCode *string `json:"ChatbotCode,omitempty" xml:"ChatbotCode,omitempty"`
	// example:
	//
	// 14
	SignId *int64 `json:"SignId,omitempty" xml:"SignId,omitempty"`
}

func (s UpgradeToRCSSignatureResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpgradeToRCSSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeToRCSSignatureResponseBodyData) GetChatbotCode() *string {
	return s.ChatbotCode
}

func (s *UpgradeToRCSSignatureResponseBodyData) GetSignId() *int64 {
	return s.SignId
}

func (s *UpgradeToRCSSignatureResponseBodyData) SetChatbotCode(v string) *UpgradeToRCSSignatureResponseBodyData {
	s.ChatbotCode = &v
	return s
}

func (s *UpgradeToRCSSignatureResponseBodyData) SetSignId(v int64) *UpgradeToRCSSignatureResponseBodyData {
	s.SignId = &v
	return s
}

func (s *UpgradeToRCSSignatureResponseBodyData) Validate() error {
	return dara.Validate(s)
}
