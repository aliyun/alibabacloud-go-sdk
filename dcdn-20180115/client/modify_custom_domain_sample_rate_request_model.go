// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomDomainSampleRateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseConfigID(v string) *ModifyCustomDomainSampleRateRequest
	GetBaseConfigID() *string
	SetDomainNames(v string) *ModifyCustomDomainSampleRateRequest
	GetDomainNames() *string
	SetSampleRate(v float32) *ModifyCustomDomainSampleRateRequest
	GetSampleRate() *float32
	SetSinkID(v int64) *ModifyCustomDomainSampleRateRequest
	GetSinkID() *int64
}

type ModifyCustomDomainSampleRateRequest struct {
	BaseConfigID *string  `json:"BaseConfigID,omitempty" xml:"BaseConfigID,omitempty"`
	DomainNames  *string  `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	SampleRate   *float32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	SinkID       *int64   `json:"SinkID,omitempty" xml:"SinkID,omitempty"`
}

func (s ModifyCustomDomainSampleRateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomDomainSampleRateRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomDomainSampleRateRequest) GetBaseConfigID() *string {
	return s.BaseConfigID
}

func (s *ModifyCustomDomainSampleRateRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *ModifyCustomDomainSampleRateRequest) GetSampleRate() *float32 {
	return s.SampleRate
}

func (s *ModifyCustomDomainSampleRateRequest) GetSinkID() *int64 {
	return s.SinkID
}

func (s *ModifyCustomDomainSampleRateRequest) SetBaseConfigID(v string) *ModifyCustomDomainSampleRateRequest {
	s.BaseConfigID = &v
	return s
}

func (s *ModifyCustomDomainSampleRateRequest) SetDomainNames(v string) *ModifyCustomDomainSampleRateRequest {
	s.DomainNames = &v
	return s
}

func (s *ModifyCustomDomainSampleRateRequest) SetSampleRate(v float32) *ModifyCustomDomainSampleRateRequest {
	s.SampleRate = &v
	return s
}

func (s *ModifyCustomDomainSampleRateRequest) SetSinkID(v int64) *ModifyCustomDomainSampleRateRequest {
	s.SinkID = &v
	return s
}

func (s *ModifyCustomDomainSampleRateRequest) Validate() error {
	return dara.Validate(s)
}
