// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventLogDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventLogDetailRequest
	GetLang() *string
	SetRegId(v string) *DescribeEventLogDetailRequest
	GetRegId() *string
	SetReqIdByLog(v string) *DescribeEventLogDetailRequest
	GetReqIdByLog() *string
}

type DescribeEventLogDetailRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Log details.
	//
	// This parameter is required.
	//
	// example:
	//
	// BD6B08EC-1B44-5378-8838-C76A36415C55
	ReqIdByLog *string `json:"reqIdByLog,omitempty" xml:"reqIdByLog,omitempty"`
}

func (s DescribeEventLogDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLogDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventLogDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventLogDetailRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventLogDetailRequest) GetReqIdByLog() *string {
	return s.ReqIdByLog
}

func (s *DescribeEventLogDetailRequest) SetLang(v string) *DescribeEventLogDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventLogDetailRequest) SetRegId(v string) *DescribeEventLogDetailRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventLogDetailRequest) SetReqIdByLog(v string) *DescribeEventLogDetailRequest {
	s.ReqIdByLog = &v
	return s
}

func (s *DescribeEventLogDetailRequest) Validate() error {
	return dara.Validate(s)
}
