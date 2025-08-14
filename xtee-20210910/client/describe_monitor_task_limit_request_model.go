// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorTaskLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeMonitorTaskLimitRequest
	GetLang() *string
	SetRegId(v string) *DescribeMonitorTaskLimitRequest
	GetRegId() *string
}

type DescribeMonitorTaskLimitRequest struct {
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
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeMonitorTaskLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorTaskLimitRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorTaskLimitRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeMonitorTaskLimitRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeMonitorTaskLimitRequest) SetLang(v string) *DescribeMonitorTaskLimitRequest {
	s.Lang = &v
	return s
}

func (s *DescribeMonitorTaskLimitRequest) SetRegId(v string) *DescribeMonitorTaskLimitRequest {
	s.RegId = &v
	return s
}

func (s *DescribeMonitorTaskLimitRequest) Validate() error {
	return dara.Validate(s)
}
