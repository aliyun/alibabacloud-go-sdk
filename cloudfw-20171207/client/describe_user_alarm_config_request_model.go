// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAlarmConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeUserAlarmConfigRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeUserAlarmConfigRequest
	GetSourceIp() *string
}

type DescribeUserAlarmConfigRequest struct {
	// The language of the response. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 59.82.135.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeUserAlarmConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAlarmConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserAlarmConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeUserAlarmConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeUserAlarmConfigRequest) SetLang(v string) *DescribeUserAlarmConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserAlarmConfigRequest) SetSourceIp(v string) *DescribeUserAlarmConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserAlarmConfigRequest) Validate() error {
	return dara.Validate(s)
}
