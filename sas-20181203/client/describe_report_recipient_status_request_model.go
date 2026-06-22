// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReportRecipientStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeReportRecipientStatusRequest
	GetLang() *string
	SetRecipients(v string) *DescribeReportRecipientStatusRequest
	GetRecipients() *string
	SetSourceIp(v string) *DescribeReportRecipientStatusRequest
	GetSourceIp() *string
}

type DescribeReportRecipientStatusRequest struct {
	// The language type of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The email addresses of the contacts. Separate multiple email addresses with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com,username@example.com
	Recipients *string `json:"Recipients,omitempty" xml:"Recipients,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 39.174.xxx.xxx
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeReportRecipientStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeReportRecipientStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeReportRecipientStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeReportRecipientStatusRequest) GetRecipients() *string {
	return s.Recipients
}

func (s *DescribeReportRecipientStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeReportRecipientStatusRequest) SetLang(v string) *DescribeReportRecipientStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeReportRecipientStatusRequest) SetRecipients(v string) *DescribeReportRecipientStatusRequest {
	s.Recipients = &v
	return s
}

func (s *DescribeReportRecipientStatusRequest) SetSourceIp(v string) *DescribeReportRecipientStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeReportRecipientStatusRequest) Validate() error {
	return dara.Validate(s)
}
