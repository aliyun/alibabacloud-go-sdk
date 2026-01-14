// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmsTemplateCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *SmsTemplateCreateResponseBody
	GetCode() *int64
	SetMessage(v string) *SmsTemplateCreateResponseBody
	GetMessage() *string
	SetModel(v string) *SmsTemplateCreateResponseBody
	GetModel() *string
	SetRequestId(v string) *SmsTemplateCreateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SmsTemplateCreateResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *SmsTemplateCreateResponseBody
	GetTimestamp() *int64
}

type SmsTemplateCreateResponseBody struct {
	// example:
	//
	// 23
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s SmsTemplateCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SmsTemplateCreateResponseBody) GoString() string {
	return s.String()
}

func (s *SmsTemplateCreateResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SmsTemplateCreateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SmsTemplateCreateResponseBody) GetModel() *string {
	return s.Model
}

func (s *SmsTemplateCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SmsTemplateCreateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SmsTemplateCreateResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *SmsTemplateCreateResponseBody) SetCode(v int64) *SmsTemplateCreateResponseBody {
	s.Code = &v
	return s
}

func (s *SmsTemplateCreateResponseBody) SetMessage(v string) *SmsTemplateCreateResponseBody {
	s.Message = &v
	return s
}

func (s *SmsTemplateCreateResponseBody) SetModel(v string) *SmsTemplateCreateResponseBody {
	s.Model = &v
	return s
}

func (s *SmsTemplateCreateResponseBody) SetRequestId(v string) *SmsTemplateCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmsTemplateCreateResponseBody) SetSuccess(v bool) *SmsTemplateCreateResponseBody {
	s.Success = &v
	return s
}

func (s *SmsTemplateCreateResponseBody) SetTimestamp(v int64) *SmsTemplateCreateResponseBody {
	s.Timestamp = &v
	return s
}

func (s *SmsTemplateCreateResponseBody) Validate() error {
	return dara.Validate(s)
}
