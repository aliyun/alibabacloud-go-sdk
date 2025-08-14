// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeByPassShuntEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeByPassShuntEventRequest
	GetLang() *string
	SetEventId(v int64) *DescribeByPassShuntEventRequest
	GetEventId() *int64
	SetRegId(v string) *DescribeByPassShuntEventRequest
	GetRegId() *string
}

type DescribeByPassShuntEventRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Event ID.
	//
	// example:
	//
	// 25
	EventId *int64 `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeByPassShuntEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeByPassShuntEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeByPassShuntEventRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeByPassShuntEventRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *DescribeByPassShuntEventRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeByPassShuntEventRequest) SetLang(v string) *DescribeByPassShuntEventRequest {
	s.Lang = &v
	return s
}

func (s *DescribeByPassShuntEventRequest) SetEventId(v int64) *DescribeByPassShuntEventRequest {
	s.EventId = &v
	return s
}

func (s *DescribeByPassShuntEventRequest) SetRegId(v string) *DescribeByPassShuntEventRequest {
	s.RegId = &v
	return s
}

func (s *DescribeByPassShuntEventRequest) Validate() error {
	return dara.Validate(s)
}
