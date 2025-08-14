// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExcuteNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeExcuteNumRequest
	GetCode() *string
	SetDegree(v string) *DescribeExcuteNumRequest
	GetDegree() *string
	SetEndDate(v string) *DescribeExcuteNumRequest
	GetEndDate() *string
	SetLang(v string) *DescribeExcuteNumRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeExcuteNumRequest
	GetSourceIp() *string
	SetStartDate(v string) *DescribeExcuteNumRequest
	GetStartDate() *string
}

type DescribeExcuteNumRequest struct {
	// This parameter is required.
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Degree    *string `json:"Degree,omitempty" xml:"Degree,omitempty"`
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeExcuteNumRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcuteNumRequest) GoString() string {
	return s.String()
}

func (s *DescribeExcuteNumRequest) GetCode() *string {
	return s.Code
}

func (s *DescribeExcuteNumRequest) GetDegree() *string {
	return s.Degree
}

func (s *DescribeExcuteNumRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeExcuteNumRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeExcuteNumRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeExcuteNumRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeExcuteNumRequest) SetCode(v string) *DescribeExcuteNumRequest {
	s.Code = &v
	return s
}

func (s *DescribeExcuteNumRequest) SetDegree(v string) *DescribeExcuteNumRequest {
	s.Degree = &v
	return s
}

func (s *DescribeExcuteNumRequest) SetEndDate(v string) *DescribeExcuteNumRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeExcuteNumRequest) SetLang(v string) *DescribeExcuteNumRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExcuteNumRequest) SetSourceIp(v string) *DescribeExcuteNumRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeExcuteNumRequest) SetStartDate(v string) *DescribeExcuteNumRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeExcuteNumRequest) Validate() error {
	return dara.Validate(s)
}
