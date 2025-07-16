// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomDomainSampleRateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ModifyCustomDomainSampleRateResponseBodyContent) *ModifyCustomDomainSampleRateResponseBody
	GetContent() *ModifyCustomDomainSampleRateResponseBodyContent
	SetRequestId(v string) *ModifyCustomDomainSampleRateResponseBody
	GetRequestId() *string
}

type ModifyCustomDomainSampleRateResponseBody struct {
	Content   *ModifyCustomDomainSampleRateResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCustomDomainSampleRateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomDomainSampleRateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustomDomainSampleRateResponseBody) GetContent() *ModifyCustomDomainSampleRateResponseBodyContent {
	return s.Content
}

func (s *ModifyCustomDomainSampleRateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustomDomainSampleRateResponseBody) SetContent(v *ModifyCustomDomainSampleRateResponseBodyContent) *ModifyCustomDomainSampleRateResponseBody {
	s.Content = v
	return s
}

func (s *ModifyCustomDomainSampleRateResponseBody) SetRequestId(v string) *ModifyCustomDomainSampleRateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCustomDomainSampleRateResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyCustomDomainSampleRateResponseBodyContent struct {
	Content []*ModifyCustomDomainSampleRateResponseBodyContentContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s ModifyCustomDomainSampleRateResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomDomainSampleRateResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ModifyCustomDomainSampleRateResponseBodyContent) GetContent() []*ModifyCustomDomainSampleRateResponseBodyContentContent {
	return s.Content
}

func (s *ModifyCustomDomainSampleRateResponseBodyContent) SetContent(v []*ModifyCustomDomainSampleRateResponseBodyContentContent) *ModifyCustomDomainSampleRateResponseBodyContent {
	s.Content = v
	return s
}

func (s *ModifyCustomDomainSampleRateResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type ModifyCustomDomainSampleRateResponseBodyContentContent struct {
	DomainName *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ErrMessage *string  `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	SampleRate *float32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
}

func (s ModifyCustomDomainSampleRateResponseBodyContentContent) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomDomainSampleRateResponseBodyContentContent) GoString() string {
	return s.String()
}

func (s *ModifyCustomDomainSampleRateResponseBodyContentContent) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyCustomDomainSampleRateResponseBodyContentContent) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyCustomDomainSampleRateResponseBodyContentContent) GetSampleRate() *float32 {
	return s.SampleRate
}

func (s *ModifyCustomDomainSampleRateResponseBodyContentContent) SetDomainName(v string) *ModifyCustomDomainSampleRateResponseBodyContentContent {
	s.DomainName = &v
	return s
}

func (s *ModifyCustomDomainSampleRateResponseBodyContentContent) SetErrMessage(v string) *ModifyCustomDomainSampleRateResponseBodyContentContent {
	s.ErrMessage = &v
	return s
}

func (s *ModifyCustomDomainSampleRateResponseBodyContentContent) SetSampleRate(v float32) *ModifyCustomDomainSampleRateResponseBodyContentContent {
	s.SampleRate = &v
	return s
}

func (s *ModifyCustomDomainSampleRateResponseBodyContentContent) Validate() error {
	return dara.Validate(s)
}
