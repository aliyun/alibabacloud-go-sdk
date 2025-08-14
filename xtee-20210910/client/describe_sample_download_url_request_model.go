// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSampleDownloadUrlRequest
	GetLang() *string
	SetRegId(v string) *DescribeSampleDownloadUrlRequest
	GetRegId() *string
	SetSampleId(v int64) *DescribeSampleDownloadUrlRequest
	GetSampleId() *int64
}

type DescribeSampleDownloadUrlRequest struct {
	// Set the language type for request and response, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Sample ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5467
	SampleId *int64 `json:"sampleId,omitempty" xml:"sampleId,omitempty"`
}

func (s DescribeSampleDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleDownloadUrlRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSampleDownloadUrlRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSampleDownloadUrlRequest) GetSampleId() *int64 {
	return s.SampleId
}

func (s *DescribeSampleDownloadUrlRequest) SetLang(v string) *DescribeSampleDownloadUrlRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSampleDownloadUrlRequest) SetRegId(v string) *DescribeSampleDownloadUrlRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSampleDownloadUrlRequest) SetSampleId(v int64) *DescribeSampleDownloadUrlRequest {
	s.SampleId = &v
	return s
}

func (s *DescribeSampleDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
