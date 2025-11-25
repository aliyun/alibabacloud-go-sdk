// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlrGrantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSlrGrantRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeSlrGrantRequest
	GetSourceIp() *string
}

type DescribeSlrGrantRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 113.225.22.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSlrGrantRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlrGrantRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlrGrantRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSlrGrantRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSlrGrantRequest) SetLang(v string) *DescribeSlrGrantRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlrGrantRequest) SetSourceIp(v string) *DescribeSlrGrantRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlrGrantRequest) Validate() error {
	return dara.Validate(s)
}
