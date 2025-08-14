// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventTaskHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventTaskHistoryRequest
	GetLang() *string
	SetRegId(v string) *DescribeEventTaskHistoryRequest
	GetRegId() *string
}

type DescribeEventTaskHistoryRequest struct {
	// Set the language type for request and response messages, default value is **zh**. Values:
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
}

func (s DescribeEventTaskHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTaskHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventTaskHistoryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventTaskHistoryRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventTaskHistoryRequest) SetLang(v string) *DescribeEventTaskHistoryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventTaskHistoryRequest) SetRegId(v string) *DescribeEventTaskHistoryRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventTaskHistoryRequest) Validate() error {
	return dara.Validate(s)
}
