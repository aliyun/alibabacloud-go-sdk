// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSubTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DownloadSubTaskResultRequest
	GetLang() *string
	SetRegId(v string) *DownloadSubTaskResultRequest
	GetRegId() *string
	SetSubTaskId(v int32) *DownloadSubTaskResultRequest
	GetSubTaskId() *int32
}

type DownloadSubTaskResultRequest struct {
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
	// 2
	SubTaskId *int32 `json:"SubTaskId,omitempty" xml:"SubTaskId,omitempty"`
}

func (s DownloadSubTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadSubTaskResultRequest) GoString() string {
	return s.String()
}

func (s *DownloadSubTaskResultRequest) GetLang() *string {
	return s.Lang
}

func (s *DownloadSubTaskResultRequest) GetRegId() *string {
	return s.RegId
}

func (s *DownloadSubTaskResultRequest) GetSubTaskId() *int32 {
	return s.SubTaskId
}

func (s *DownloadSubTaskResultRequest) SetLang(v string) *DownloadSubTaskResultRequest {
	s.Lang = &v
	return s
}

func (s *DownloadSubTaskResultRequest) SetRegId(v string) *DownloadSubTaskResultRequest {
	s.RegId = &v
	return s
}

func (s *DownloadSubTaskResultRequest) SetSubTaskId(v int32) *DownloadSubTaskResultRequest {
	s.SubTaskId = &v
	return s
}

func (s *DownloadSubTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
