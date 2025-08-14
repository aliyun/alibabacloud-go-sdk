// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDecisionResultFluctuationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDecisionResultFluctuationRequest
	GetLang() *string
	SetEventCodes(v string) *DescribeDecisionResultFluctuationRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeDecisionResultFluctuationRequest
	GetRegId() *string
}

type DescribeDecisionResultFluctuationRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Event code.
	//
	// example:
	//
	// de_ahqhsw7665,de_ahqhsw7622
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeDecisionResultFluctuationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDecisionResultFluctuationRequest) GoString() string {
	return s.String()
}

func (s *DescribeDecisionResultFluctuationRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDecisionResultFluctuationRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeDecisionResultFluctuationRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeDecisionResultFluctuationRequest) SetLang(v string) *DescribeDecisionResultFluctuationRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDecisionResultFluctuationRequest) SetEventCodes(v string) *DescribeDecisionResultFluctuationRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeDecisionResultFluctuationRequest) SetRegId(v string) *DescribeDecisionResultFluctuationRequest {
	s.RegId = &v
	return s
}

func (s *DescribeDecisionResultFluctuationRequest) Validate() error {
	return dara.Validate(s)
}
