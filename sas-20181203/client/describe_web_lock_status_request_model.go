// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeWebLockStatusRequest
	GetFrom() *string
	SetLang(v string) *DescribeWebLockStatusRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeWebLockStatusRequest
	GetSourceIp() *string
}

type DescribeWebLockStatusRequest struct {
	// The ID of the request source. Set the value to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
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
	// The source IP address of the request.
	//
	// example:
	//
	// 221.214.XXX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeWebLockStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebLockStatusRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeWebLockStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWebLockStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWebLockStatusRequest) SetFrom(v string) *DescribeWebLockStatusRequest {
	s.From = &v
	return s
}

func (s *DescribeWebLockStatusRequest) SetLang(v string) *DescribeWebLockStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebLockStatusRequest) SetSourceIp(v string) *DescribeWebLockStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebLockStatusRequest) Validate() error {
	return dara.Validate(s)
}
