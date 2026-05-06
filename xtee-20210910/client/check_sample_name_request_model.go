// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSampleNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CheckSampleNameRequest
	GetLang() *string
	SetRegId(v string) *CheckSampleNameRequest
	GetRegId() *string
	SetSampleName(v string) *CheckSampleNameRequest
	GetSampleName() *string
}

type CheckSampleNameRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// example:
	//
	// SampleTest
	SampleName *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
}

func (s CheckSampleNameRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckSampleNameRequest) GoString() string {
	return s.String()
}

func (s *CheckSampleNameRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckSampleNameRequest) GetRegId() *string {
	return s.RegId
}

func (s *CheckSampleNameRequest) GetSampleName() *string {
	return s.SampleName
}

func (s *CheckSampleNameRequest) SetLang(v string) *CheckSampleNameRequest {
	s.Lang = &v
	return s
}

func (s *CheckSampleNameRequest) SetRegId(v string) *CheckSampleNameRequest {
	s.RegId = &v
	return s
}

func (s *CheckSampleNameRequest) SetSampleName(v string) *CheckSampleNameRequest {
	s.SampleName = &v
	return s
}

func (s *CheckSampleNameRequest) Validate() error {
	return dara.Validate(s)
}
