// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteSampleRequest
	GetLang() *string
	SetRegId(v string) *DeleteSampleRequest
	GetRegId() *string
	SetSampleId(v int32) *DeleteSampleRequest
	GetSampleId() *int32
}

type DeleteSampleRequest struct {
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

func (s DeleteSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleRequest) GoString() string {
	return s.String()
}

func (s *DeleteSampleRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteSampleRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteSampleRequest) GetSampleId() *int32 {
	return s.SampleId
}

func (s *DeleteSampleRequest) SetLang(v string) *DeleteSampleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteSampleRequest) SetRegId(v string) *DeleteSampleRequest {
	s.RegId = &v
	return s
}

func (s *DeleteSampleRequest) SetSampleId(v int32) *DeleteSampleRequest {
	s.SampleId = &v
	return s
}

func (s *DeleteSampleRequest) Validate() error {
	return dara.Validate(s)
}
