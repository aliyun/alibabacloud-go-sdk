// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSampleDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetSampleDetailRequest
	GetLang() *string
	SetRegId(v string) *GetSampleDetailRequest
	GetRegId() *string
	SetSampleId(v int32) *GetSampleDetailRequest
	GetSampleId() *int32
}

type GetSampleDetailRequest struct {
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
	// 1
	SampleId *int32 `json:"SampleId,omitempty" xml:"SampleId,omitempty"`
}

func (s GetSampleDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSampleDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSampleDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *GetSampleDetailRequest) GetRegId() *string {
	return s.RegId
}

func (s *GetSampleDetailRequest) GetSampleId() *int32 {
	return s.SampleId
}

func (s *GetSampleDetailRequest) SetLang(v string) *GetSampleDetailRequest {
	s.Lang = &v
	return s
}

func (s *GetSampleDetailRequest) SetRegId(v string) *GetSampleDetailRequest {
	s.RegId = &v
	return s
}

func (s *GetSampleDetailRequest) SetSampleId(v int32) *GetSampleDetailRequest {
	s.SampleId = &v
	return s
}

func (s *GetSampleDetailRequest) Validate() error {
	return dara.Validate(s)
}
