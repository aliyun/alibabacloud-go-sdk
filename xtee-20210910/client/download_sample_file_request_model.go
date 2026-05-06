// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSampleFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DownloadSampleFileRequest
	GetLang() *string
	SetRegId(v string) *DownloadSampleFileRequest
	GetRegId() *string
	SetSampleId(v int32) *DownloadSampleFileRequest
	GetSampleId() *int32
	SetTab(v string) *DownloadSampleFileRequest
	GetTab() *string
}

type DownloadSampleFileRequest struct {
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
	// example:
	//
	// INTERNET
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
}

func (s DownloadSampleFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadSampleFileRequest) GoString() string {
	return s.String()
}

func (s *DownloadSampleFileRequest) GetLang() *string {
	return s.Lang
}

func (s *DownloadSampleFileRequest) GetRegId() *string {
	return s.RegId
}

func (s *DownloadSampleFileRequest) GetSampleId() *int32 {
	return s.SampleId
}

func (s *DownloadSampleFileRequest) GetTab() *string {
	return s.Tab
}

func (s *DownloadSampleFileRequest) SetLang(v string) *DownloadSampleFileRequest {
	s.Lang = &v
	return s
}

func (s *DownloadSampleFileRequest) SetRegId(v string) *DownloadSampleFileRequest {
	s.RegId = &v
	return s
}

func (s *DownloadSampleFileRequest) SetSampleId(v int32) *DownloadSampleFileRequest {
	s.SampleId = &v
	return s
}

func (s *DownloadSampleFileRequest) SetTab(v string) *DownloadSampleFileRequest {
	s.Tab = &v
	return s
}

func (s *DownloadSampleFileRequest) Validate() error {
	return dara.Validate(s)
}
