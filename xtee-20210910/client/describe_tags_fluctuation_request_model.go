// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsFluctuationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeTagsFluctuationRequest
	GetLang() *string
	SetEventCodes(v string) *DescribeTagsFluctuationRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeTagsFluctuationRequest
	GetRegId() *string
}

type DescribeTagsFluctuationRequest struct {
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
	// de_afghcf6411
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeTagsFluctuationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsFluctuationRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsFluctuationRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTagsFluctuationRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeTagsFluctuationRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeTagsFluctuationRequest) SetLang(v string) *DescribeTagsFluctuationRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTagsFluctuationRequest) SetEventCodes(v string) *DescribeTagsFluctuationRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeTagsFluctuationRequest) SetRegId(v string) *DescribeTagsFluctuationRequest {
	s.RegId = &v
	return s
}

func (s *DescribeTagsFluctuationRequest) Validate() error {
	return dara.Validate(s)
}
