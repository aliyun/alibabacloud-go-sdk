// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendFeedbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReason(v string) *SendFeedbackRequest
	GetReason() *string
	SetRiskLabel(v string) *SendFeedbackRequest
	GetRiskLabel() *string
	SetSampleType(v string) *SendFeedbackRequest
	GetSampleType() *string
	SetValue(v string) *SendFeedbackRequest
	GetValue() *string
}

type SendFeedbackRequest struct {
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Sample labels. User-defined, separated by commas.
	//
	// example:
	//
	// OTHERS
	RiskLabel *string `json:"RiskLabel,omitempty" xml:"RiskLabel,omitempty"`
	// Sample type. For phone number type samples, input PHONE; for email type samples, input EMAIL; for account type samples, input ACCOUNT.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHONE
	SampleType *string `json:"SampleType,omitempty" xml:"SampleType,omitempty"`
	// Sample value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000000000
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SendFeedbackRequest) String() string {
	return dara.Prettify(s)
}

func (s SendFeedbackRequest) GoString() string {
	return s.String()
}

func (s *SendFeedbackRequest) GetReason() *string {
	return s.Reason
}

func (s *SendFeedbackRequest) GetRiskLabel() *string {
	return s.RiskLabel
}

func (s *SendFeedbackRequest) GetSampleType() *string {
	return s.SampleType
}

func (s *SendFeedbackRequest) GetValue() *string {
	return s.Value
}

func (s *SendFeedbackRequest) SetReason(v string) *SendFeedbackRequest {
	s.Reason = &v
	return s
}

func (s *SendFeedbackRequest) SetRiskLabel(v string) *SendFeedbackRequest {
	s.RiskLabel = &v
	return s
}

func (s *SendFeedbackRequest) SetSampleType(v string) *SendFeedbackRequest {
	s.SampleType = &v
	return s
}

func (s *SendFeedbackRequest) SetValue(v string) *SendFeedbackRequest {
	s.Value = &v
	return s
}

func (s *SendFeedbackRequest) Validate() error {
	return dara.Validate(s)
}
