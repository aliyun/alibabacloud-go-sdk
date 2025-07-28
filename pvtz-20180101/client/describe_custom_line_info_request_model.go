// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLineInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeCustomLineInfoRequest
	GetLang() *string
	SetLineId(v string) *DescribeCustomLineInfoRequest
	GetLineId() *string
}

type DescribeCustomLineInfoRequest struct {
	// The language of the response.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The unique ID of the custom line.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11271
	LineId *string `json:"LineId,omitempty" xml:"LineId,omitempty"`
}

func (s DescribeCustomLineInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLineInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomLineInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCustomLineInfoRequest) GetLineId() *string {
	return s.LineId
}

func (s *DescribeCustomLineInfoRequest) SetLang(v string) *DescribeCustomLineInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCustomLineInfoRequest) SetLineId(v string) *DescribeCustomLineInfoRequest {
	s.LineId = &v
	return s
}

func (s *DescribeCustomLineInfoRequest) Validate() error {
	return dara.Validate(s)
}
