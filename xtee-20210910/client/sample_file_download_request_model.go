// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSampleFileDownloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SampleFileDownloadRequest
	GetLang() *string
	SetRegId(v string) *SampleFileDownloadRequest
	GetRegId() *string
	SetTab(v string) *SampleFileDownloadRequest
	GetTab() *string
}

type SampleFileDownloadRequest struct {
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// The scenario.
	//
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
}

func (s SampleFileDownloadRequest) String() string {
	return dara.Prettify(s)
}

func (s SampleFileDownloadRequest) GoString() string {
	return s.String()
}

func (s *SampleFileDownloadRequest) GetLang() *string {
	return s.Lang
}

func (s *SampleFileDownloadRequest) GetRegId() *string {
	return s.RegId
}

func (s *SampleFileDownloadRequest) GetTab() *string {
	return s.Tab
}

func (s *SampleFileDownloadRequest) SetLang(v string) *SampleFileDownloadRequest {
	s.Lang = &v
	return s
}

func (s *SampleFileDownloadRequest) SetRegId(v string) *SampleFileDownloadRequest {
	s.RegId = &v
	return s
}

func (s *SampleFileDownloadRequest) SetTab(v string) *SampleFileDownloadRequest {
	s.Tab = &v
	return s
}

func (s *SampleFileDownloadRequest) Validate() error {
	return dara.Validate(s)
}
