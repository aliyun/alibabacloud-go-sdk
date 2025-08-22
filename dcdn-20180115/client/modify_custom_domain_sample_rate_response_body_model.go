// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomDomainSampleRateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*ModifyCustomDomainSampleRateResponseBodyContent) *ModifyCustomDomainSampleRateResponseBody
	GetContent() []*ModifyCustomDomainSampleRateResponseBodyContent
	SetRequestId(v string) *ModifyCustomDomainSampleRateResponseBody
	GetRequestId() *string
}

type ModifyCustomDomainSampleRateResponseBody struct {
	Content   []*ModifyCustomDomainSampleRateResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCustomDomainSampleRateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomDomainSampleRateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustomDomainSampleRateResponseBody) GetContent() []*ModifyCustomDomainSampleRateResponseBodyContent {
	return s.Content
}

func (s *ModifyCustomDomainSampleRateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustomDomainSampleRateResponseBody) SetContent(v []*ModifyCustomDomainSampleRateResponseBodyContent) *ModifyCustomDomainSampleRateResponseBody {
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
	DomainName *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ErrMessage *string  `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	SampleRate *float32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
}

func (s ModifyCustomDomainSampleRateResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomDomainSampleRateResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ModifyCustomDomainSampleRateResponseBodyContent) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyCustomDomainSampleRateResponseBodyContent) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyCustomDomainSampleRateResponseBodyContent) GetSampleRate() *float32 {
	return s.SampleRate
}

func (s *ModifyCustomDomainSampleRateResponseBodyContent) SetDomainName(v string) *ModifyCustomDomainSampleRateResponseBodyContent {
	s.DomainName = &v
	return s
}

func (s *ModifyCustomDomainSampleRateResponseBodyContent) SetErrMessage(v string) *ModifyCustomDomainSampleRateResponseBodyContent {
	s.ErrMessage = &v
	return s
}

func (s *ModifyCustomDomainSampleRateResponseBodyContent) SetSampleRate(v float32) *ModifyCustomDomainSampleRateResponseBodyContent {
	s.SampleRate = &v
	return s
}

func (s *ModifyCustomDomainSampleRateResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
