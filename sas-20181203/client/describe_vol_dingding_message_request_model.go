// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVolDingdingMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVolDingdingMessageRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeVolDingdingMessageRequest
	GetSourceIp() *string
}

type DescribeVolDingdingMessageRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeVolDingdingMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVolDingdingMessageRequest) GoString() string {
	return s.String()
}

func (s *DescribeVolDingdingMessageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVolDingdingMessageRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeVolDingdingMessageRequest) SetLang(v string) *DescribeVolDingdingMessageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVolDingdingMessageRequest) SetSourceIp(v string) *DescribeVolDingdingMessageRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVolDingdingMessageRequest) Validate() error {
	return dara.Validate(s)
}
