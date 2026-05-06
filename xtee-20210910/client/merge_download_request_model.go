// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeDownloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *MergeDownloadRequest
	GetLang() *string
	SetRegId(v string) *MergeDownloadRequest
	GetRegId() *string
	SetSubTaskIds(v string) *MergeDownloadRequest
	GetSubTaskIds() *string
}

type MergeDownloadRequest struct {
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
	// 1,2
	SubTaskIds *string `json:"SubTaskIds,omitempty" xml:"SubTaskIds,omitempty"`
}

func (s MergeDownloadRequest) String() string {
	return dara.Prettify(s)
}

func (s MergeDownloadRequest) GoString() string {
	return s.String()
}

func (s *MergeDownloadRequest) GetLang() *string {
	return s.Lang
}

func (s *MergeDownloadRequest) GetRegId() *string {
	return s.RegId
}

func (s *MergeDownloadRequest) GetSubTaskIds() *string {
	return s.SubTaskIds
}

func (s *MergeDownloadRequest) SetLang(v string) *MergeDownloadRequest {
	s.Lang = &v
	return s
}

func (s *MergeDownloadRequest) SetRegId(v string) *MergeDownloadRequest {
	s.RegId = &v
	return s
}

func (s *MergeDownloadRequest) SetSubTaskIds(v string) *MergeDownloadRequest {
	s.SubTaskIds = &v
	return s
}

func (s *MergeDownloadRequest) Validate() error {
	return dara.Validate(s)
}
