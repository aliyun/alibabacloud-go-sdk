// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSampleDataByTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *AddSampleDataByTextRequest
	GetLang() *string
	SetDataValue(v string) *AddSampleDataByTextRequest
	GetDataValue() *string
	SetRegId(v string) *AddSampleDataByTextRequest
	GetRegId() *string
	SetSampleBatchUuid(v string) *AddSampleDataByTextRequest
	GetSampleBatchUuid() *string
}

type AddSampleDataByTextRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// List data.
	//
	// example:
	//
	// testA \\n testB
	DataValue *string `json:"dataValue,omitempty" xml:"dataValue,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Sample UUID.
	//
	// example:
	//
	// 1806507582222226
	//
	// _saf_sample_split_coupon_abuse_saf_sample_split_mobile_saf_sample_split_pass
	SampleBatchUuid *string `json:"sampleBatchUuid,omitempty" xml:"sampleBatchUuid,omitempty"`
}

func (s AddSampleDataByTextRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSampleDataByTextRequest) GoString() string {
	return s.String()
}

func (s *AddSampleDataByTextRequest) GetLang() *string {
	return s.Lang
}

func (s *AddSampleDataByTextRequest) GetDataValue() *string {
	return s.DataValue
}

func (s *AddSampleDataByTextRequest) GetRegId() *string {
	return s.RegId
}

func (s *AddSampleDataByTextRequest) GetSampleBatchUuid() *string {
	return s.SampleBatchUuid
}

func (s *AddSampleDataByTextRequest) SetLang(v string) *AddSampleDataByTextRequest {
	s.Lang = &v
	return s
}

func (s *AddSampleDataByTextRequest) SetDataValue(v string) *AddSampleDataByTextRequest {
	s.DataValue = &v
	return s
}

func (s *AddSampleDataByTextRequest) SetRegId(v string) *AddSampleDataByTextRequest {
	s.RegId = &v
	return s
}

func (s *AddSampleDataByTextRequest) SetSampleBatchUuid(v string) *AddSampleDataByTextRequest {
	s.SampleBatchUuid = &v
	return s
}

func (s *AddSampleDataByTextRequest) Validate() error {
	return dara.Validate(s)
}
