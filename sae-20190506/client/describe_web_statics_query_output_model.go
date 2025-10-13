// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebStaticsQueryOutput interface {
	dara.Model
	String() string
	GoString() string
	SetLength(v int32) *DescribeWebStaticsQueryOutput
	GetLength() *int32
	SetWebStatics(v []*WebStaticsInfo) *DescribeWebStaticsQueryOutput
	GetWebStatics() []*WebStaticsInfo
}

type DescribeWebStaticsQueryOutput struct {
	Length     *int32            `json:"Length,omitempty" xml:"Length,omitempty"`
	WebStatics []*WebStaticsInfo `json:"WebStatics,omitempty" xml:"WebStatics,omitempty" type:"Repeated"`
}

func (s DescribeWebStaticsQueryOutput) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebStaticsQueryOutput) GoString() string {
	return s.String()
}

func (s *DescribeWebStaticsQueryOutput) GetLength() *int32 {
	return s.Length
}

func (s *DescribeWebStaticsQueryOutput) GetWebStatics() []*WebStaticsInfo {
	return s.WebStatics
}

func (s *DescribeWebStaticsQueryOutput) SetLength(v int32) *DescribeWebStaticsQueryOutput {
	s.Length = &v
	return s
}

func (s *DescribeWebStaticsQueryOutput) SetWebStatics(v []*WebStaticsInfo) *DescribeWebStaticsQueryOutput {
	s.WebStatics = v
	return s
}

func (s *DescribeWebStaticsQueryOutput) Validate() error {
	if s.WebStatics != nil {
		for _, item := range s.WebStatics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
