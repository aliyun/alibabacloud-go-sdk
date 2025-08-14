// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventDetailByRequestIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventDetailByRequestIdRequest
	GetLang() *string
	SetEventCode(v string) *DescribeEventDetailByRequestIdRequest
	GetEventCode() *string
	SetEventTime(v int64) *DescribeEventDetailByRequestIdRequest
	GetEventTime() *int64
	SetRegId(v string) *DescribeEventDetailByRequestIdRequest
	GetRegId() *string
	SetSRequestId(v string) *DescribeEventDetailByRequestIdRequest
	GetSRequestId() *string
}

type DescribeEventDetailByRequestIdRequest struct {
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
	// Event code
	//
	// example:
	//
	// de_azywkh7523
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event execution time
	//
	// example:
	//
	// 1752571330000
	EventTime *int64 `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// BD6B08EC-1B44-5378-8838-C76A36415C55
	SRequestId *string `json:"sRequestId,omitempty" xml:"sRequestId,omitempty"`
}

func (s DescribeEventDetailByRequestIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailByRequestIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailByRequestIdRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventDetailByRequestIdRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeEventDetailByRequestIdRequest) GetEventTime() *int64 {
	return s.EventTime
}

func (s *DescribeEventDetailByRequestIdRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventDetailByRequestIdRequest) GetSRequestId() *string {
	return s.SRequestId
}

func (s *DescribeEventDetailByRequestIdRequest) SetLang(v string) *DescribeEventDetailByRequestIdRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventDetailByRequestIdRequest) SetEventCode(v string) *DescribeEventDetailByRequestIdRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeEventDetailByRequestIdRequest) SetEventTime(v int64) *DescribeEventDetailByRequestIdRequest {
	s.EventTime = &v
	return s
}

func (s *DescribeEventDetailByRequestIdRequest) SetRegId(v string) *DescribeEventDetailByRequestIdRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventDetailByRequestIdRequest) SetSRequestId(v string) *DescribeEventDetailByRequestIdRequest {
	s.SRequestId = &v
	return s
}

func (s *DescribeEventDetailByRequestIdRequest) Validate() error {
	return dara.Validate(s)
}
