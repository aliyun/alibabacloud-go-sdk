// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotAttackerPortraitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListHoneypotAttackerPortraitRequest
	GetCurrentPage() *int32
	SetEndTimeStamp(v int64) *ListHoneypotAttackerPortraitRequest
	GetEndTimeStamp() *int64
	SetLang(v string) *ListHoneypotAttackerPortraitRequest
	GetLang() *string
	SetPageSize(v int32) *ListHoneypotAttackerPortraitRequest
	GetPageSize() *int32
	SetSrcIp(v string) *ListHoneypotAttackerPortraitRequest
	GetSrcIp() *string
	SetStartTimeStamp(v int64) *ListHoneypotAttackerPortraitRequest
	GetStartTimeStamp() *int64
}

type ListHoneypotAttackerPortraitRequest struct {
	// The page number. Default value: **1**. Pages start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1672285044000
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source IP address of the attack.
	//
	// example:
	//
	// 101.133.155.***
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1672249044000
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
}

func (s ListHoneypotAttackerPortraitRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotAttackerPortraitRequest) GoString() string {
	return s.String()
}

func (s *ListHoneypotAttackerPortraitRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotAttackerPortraitRequest) GetEndTimeStamp() *int64 {
	return s.EndTimeStamp
}

func (s *ListHoneypotAttackerPortraitRequest) GetLang() *string {
	return s.Lang
}

func (s *ListHoneypotAttackerPortraitRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotAttackerPortraitRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *ListHoneypotAttackerPortraitRequest) GetStartTimeStamp() *int64 {
	return s.StartTimeStamp
}

func (s *ListHoneypotAttackerPortraitRequest) SetCurrentPage(v int32) *ListHoneypotAttackerPortraitRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotAttackerPortraitRequest) SetEndTimeStamp(v int64) *ListHoneypotAttackerPortraitRequest {
	s.EndTimeStamp = &v
	return s
}

func (s *ListHoneypotAttackerPortraitRequest) SetLang(v string) *ListHoneypotAttackerPortraitRequest {
	s.Lang = &v
	return s
}

func (s *ListHoneypotAttackerPortraitRequest) SetPageSize(v int32) *ListHoneypotAttackerPortraitRequest {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotAttackerPortraitRequest) SetSrcIp(v string) *ListHoneypotAttackerPortraitRequest {
	s.SrcIp = &v
	return s
}

func (s *ListHoneypotAttackerPortraitRequest) SetStartTimeStamp(v int64) *ListHoneypotAttackerPortraitRequest {
	s.StartTimeStamp = &v
	return s
}

func (s *ListHoneypotAttackerPortraitRequest) Validate() error {
	return dara.Validate(s)
}
