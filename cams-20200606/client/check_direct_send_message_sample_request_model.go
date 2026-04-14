// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDirectSendMessageSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *CheckDirectSendMessageSampleRequest
	GetCustSpaceId() *string
	SetInteractive(v map[string]interface{}) *CheckDirectSendMessageSampleRequest
	GetInteractive() map[string]interface{}
	SetText(v map[string]interface{}) *CheckDirectSendMessageSampleRequest
	GetText() map[string]interface{}
	SetType(v string) *CheckDirectSendMessageSampleRequest
	GetType() *string
}

type CheckDirectSendMessageSampleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cams-xx**
	CustSpaceId *string                `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	Interactive map[string]interface{} `json:"Interactive,omitempty" xml:"Interactive,omitempty"`
	// example:
	//
	// {"text": "This is a direct send message"}
	Text map[string]interface{} `json:"Text,omitempty" xml:"Text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CheckDirectSendMessageSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDirectSendMessageSampleRequest) GoString() string {
	return s.String()
}

func (s *CheckDirectSendMessageSampleRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CheckDirectSendMessageSampleRequest) GetInteractive() map[string]interface{} {
	return s.Interactive
}

func (s *CheckDirectSendMessageSampleRequest) GetText() map[string]interface{} {
	return s.Text
}

func (s *CheckDirectSendMessageSampleRequest) GetType() *string {
	return s.Type
}

func (s *CheckDirectSendMessageSampleRequest) SetCustSpaceId(v string) *CheckDirectSendMessageSampleRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CheckDirectSendMessageSampleRequest) SetInteractive(v map[string]interface{}) *CheckDirectSendMessageSampleRequest {
	s.Interactive = v
	return s
}

func (s *CheckDirectSendMessageSampleRequest) SetText(v map[string]interface{}) *CheckDirectSendMessageSampleRequest {
	s.Text = v
	return s
}

func (s *CheckDirectSendMessageSampleRequest) SetType(v string) *CheckDirectSendMessageSampleRequest {
	s.Type = &v
	return s
}

func (s *CheckDirectSendMessageSampleRequest) Validate() error {
	return dara.Validate(s)
}
