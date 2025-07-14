// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAppStaticsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetLength(v int32) *DescribeWebAppStaticsOutput
	GetLength() *int32
	SetWebAppStatics(v []*WebStaticsInfo) *DescribeWebAppStaticsOutput
	GetWebAppStatics() []*WebStaticsInfo
}

type DescribeWebAppStaticsOutput struct {
	Length        *int32            `json:"Length,omitempty" xml:"Length,omitempty"`
	WebAppStatics []*WebStaticsInfo `json:"WebAppStatics,omitempty" xml:"WebAppStatics,omitempty" type:"Repeated"`
}

func (s DescribeWebAppStaticsOutput) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAppStaticsOutput) GoString() string {
	return s.String()
}

func (s *DescribeWebAppStaticsOutput) GetLength() *int32 {
	return s.Length
}

func (s *DescribeWebAppStaticsOutput) GetWebAppStatics() []*WebStaticsInfo {
	return s.WebAppStatics
}

func (s *DescribeWebAppStaticsOutput) SetLength(v int32) *DescribeWebAppStaticsOutput {
	s.Length = &v
	return s
}

func (s *DescribeWebAppStaticsOutput) SetWebAppStatics(v []*WebStaticsInfo) *DescribeWebAppStaticsOutput {
	s.WebAppStatics = v
	return s
}

func (s *DescribeWebAppStaticsOutput) Validate() error {
	return dara.Validate(s)
}
