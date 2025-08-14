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
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// Scenario.
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
