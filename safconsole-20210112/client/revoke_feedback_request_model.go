// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeFeedbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSampleType(v string) *RevokeFeedbackRequest
	GetSampleType() *string
	SetValue(v string) *RevokeFeedbackRequest
	GetValue() *string
}

type RevokeFeedbackRequest struct {
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

func (s RevokeFeedbackRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeFeedbackRequest) GoString() string {
	return s.String()
}

func (s *RevokeFeedbackRequest) GetSampleType() *string {
	return s.SampleType
}

func (s *RevokeFeedbackRequest) GetValue() *string {
	return s.Value
}

func (s *RevokeFeedbackRequest) SetSampleType(v string) *RevokeFeedbackRequest {
	s.SampleType = &v
	return s
}

func (s *RevokeFeedbackRequest) SetValue(v string) *RevokeFeedbackRequest {
	s.Value = &v
	return s
}

func (s *RevokeFeedbackRequest) Validate() error {
	return dara.Validate(s)
}
