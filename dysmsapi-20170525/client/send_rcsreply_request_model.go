// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendRCSReplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInReplyToRcsID(v string) *SendRCSReplyRequest
	GetInReplyToRcsID() *string
	SetOutId(v string) *SendRCSReplyRequest
	GetOutId() *string
	SetPhoneNumbers(v string) *SendRCSReplyRequest
	GetPhoneNumbers() *string
	SetSignName(v string) *SendRCSReplyRequest
	GetSignName() *string
	SetTemplateCode(v string) *SendRCSReplyRequest
	GetTemplateCode() *string
	SetTemplateParam(v string) *SendRCSReplyRequest
	GetTemplateParam() *string
}

type SendRCSReplyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	InReplyToRcsID *string `json:"InReplyToRcsID,omitempty" xml:"InReplyToRcsID,omitempty"`
	// example:
	//
	// 示例值
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	PhoneNumbers *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 示例值示例值
	TemplateParam *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
}

func (s SendRCSReplyRequest) String() string {
	return dara.Prettify(s)
}

func (s SendRCSReplyRequest) GoString() string {
	return s.String()
}

func (s *SendRCSReplyRequest) GetInReplyToRcsID() *string {
	return s.InReplyToRcsID
}

func (s *SendRCSReplyRequest) GetOutId() *string {
	return s.OutId
}

func (s *SendRCSReplyRequest) GetPhoneNumbers() *string {
	return s.PhoneNumbers
}

func (s *SendRCSReplyRequest) GetSignName() *string {
	return s.SignName
}

func (s *SendRCSReplyRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendRCSReplyRequest) GetTemplateParam() *string {
	return s.TemplateParam
}

func (s *SendRCSReplyRequest) SetInReplyToRcsID(v string) *SendRCSReplyRequest {
	s.InReplyToRcsID = &v
	return s
}

func (s *SendRCSReplyRequest) SetOutId(v string) *SendRCSReplyRequest {
	s.OutId = &v
	return s
}

func (s *SendRCSReplyRequest) SetPhoneNumbers(v string) *SendRCSReplyRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *SendRCSReplyRequest) SetSignName(v string) *SendRCSReplyRequest {
	s.SignName = &v
	return s
}

func (s *SendRCSReplyRequest) SetTemplateCode(v string) *SendRCSReplyRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendRCSReplyRequest) SetTemplateParam(v string) *SendRCSReplyRequest {
	s.TemplateParam = &v
	return s
}

func (s *SendRCSReplyRequest) Validate() error {
	return dara.Validate(s)
}
