// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDirectSendMessageSampleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *CheckDirectSendMessageSampleShrinkRequest
	GetCustSpaceId() *string
	SetInteractiveShrink(v string) *CheckDirectSendMessageSampleShrinkRequest
	GetInteractiveShrink() *string
	SetTextShrink(v string) *CheckDirectSendMessageSampleShrinkRequest
	GetTextShrink() *string
	SetType(v string) *CheckDirectSendMessageSampleShrinkRequest
	GetType() *string
}

type CheckDirectSendMessageSampleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cams-xx**
	CustSpaceId       *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	InteractiveShrink *string `json:"Interactive,omitempty" xml:"Interactive,omitempty"`
	// example:
	//
	// {"text": "This is a direct send message"}
	TextShrink *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CheckDirectSendMessageSampleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDirectSendMessageSampleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckDirectSendMessageSampleShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CheckDirectSendMessageSampleShrinkRequest) GetInteractiveShrink() *string {
	return s.InteractiveShrink
}

func (s *CheckDirectSendMessageSampleShrinkRequest) GetTextShrink() *string {
	return s.TextShrink
}

func (s *CheckDirectSendMessageSampleShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CheckDirectSendMessageSampleShrinkRequest) SetCustSpaceId(v string) *CheckDirectSendMessageSampleShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CheckDirectSendMessageSampleShrinkRequest) SetInteractiveShrink(v string) *CheckDirectSendMessageSampleShrinkRequest {
	s.InteractiveShrink = &v
	return s
}

func (s *CheckDirectSendMessageSampleShrinkRequest) SetTextShrink(v string) *CheckDirectSendMessageSampleShrinkRequest {
	s.TextShrink = &v
	return s
}

func (s *CheckDirectSendMessageSampleShrinkRequest) SetType(v string) *CheckDirectSendMessageSampleShrinkRequest {
	s.Type = &v
	return s
}

func (s *CheckDirectSendMessageSampleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
