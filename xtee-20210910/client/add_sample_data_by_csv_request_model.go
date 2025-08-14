// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSampleDataByCsvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *AddSampleDataByCsvRequest
	GetLang() *string
	SetOssFileName(v string) *AddSampleDataByCsvRequest
	GetOssFileName() *string
	SetRegId(v string) *AddSampleDataByCsvRequest
	GetRegId() *string
	SetSampleBatchUuid(v string) *AddSampleDataByCsvRequest
	GetSampleBatchUuid() *string
}

type AddSampleDataByCsvRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Uploaded OSS address.
	//
	// example:
	//
	// saf/path/test.csv
	OssFileName *string `json:"ossFileName,omitempty" xml:"ossFileName,omitempty"`
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
	// 1806507582222226_saf_sample_split_coupon_abuse_saf_sample_split_mobile_saf_sample_split_pass
	SampleBatchUuid *string `json:"sampleBatchUuid,omitempty" xml:"sampleBatchUuid,omitempty"`
}

func (s AddSampleDataByCsvRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSampleDataByCsvRequest) GoString() string {
	return s.String()
}

func (s *AddSampleDataByCsvRequest) GetLang() *string {
	return s.Lang
}

func (s *AddSampleDataByCsvRequest) GetOssFileName() *string {
	return s.OssFileName
}

func (s *AddSampleDataByCsvRequest) GetRegId() *string {
	return s.RegId
}

func (s *AddSampleDataByCsvRequest) GetSampleBatchUuid() *string {
	return s.SampleBatchUuid
}

func (s *AddSampleDataByCsvRequest) SetLang(v string) *AddSampleDataByCsvRequest {
	s.Lang = &v
	return s
}

func (s *AddSampleDataByCsvRequest) SetOssFileName(v string) *AddSampleDataByCsvRequest {
	s.OssFileName = &v
	return s
}

func (s *AddSampleDataByCsvRequest) SetRegId(v string) *AddSampleDataByCsvRequest {
	s.RegId = &v
	return s
}

func (s *AddSampleDataByCsvRequest) SetSampleBatchUuid(v string) *AddSampleDataByCsvRequest {
	s.SampleBatchUuid = &v
	return s
}

func (s *AddSampleDataByCsvRequest) Validate() error {
	return dara.Validate(s)
}
