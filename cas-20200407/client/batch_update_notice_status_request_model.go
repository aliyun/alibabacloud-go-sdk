// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateNoticeStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *BatchUpdateNoticeStatusRequest
	GetIds() *string
	SetLang(v string) *BatchUpdateNoticeStatusRequest
	GetLang() *string
	SetNoticeBiz(v string) *BatchUpdateNoticeStatusRequest
	GetNoticeBiz() *string
	SetNoticeStatus(v string) *BatchUpdateNoticeStatusRequest
	GetNoticeStatus() *string
	SetSourceIp(v string) *BatchUpdateNoticeStatusRequest
	GetSourceIp() *string
}

type BatchUpdateNoticeStatusRequest struct {
	// The list of primary key identifiers to be synchronized to Certificate Management Service. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 888
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The language type for the request and the received message. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The business code of message notification. The value is fixed as ssl.
	//
	// example:
	//
	// ssl
	NoticeBiz *string `json:"NoticeBiz,omitempty" xml:"NoticeBiz,omitempty"`
	// Specifies whether to enable message notification.
	//
	// \\--enable: enables message notification.
	//
	// \\--disable: disables message notification.
	//
	// This parameter is required.
	//
	// example:
	//
	// disable
	NoticeStatus *string `json:"NoticeStatus,omitempty" xml:"NoticeStatus,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 47.98.242.200
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s BatchUpdateNoticeStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateNoticeStatusRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateNoticeStatusRequest) GetIds() *string {
	return s.Ids
}

func (s *BatchUpdateNoticeStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *BatchUpdateNoticeStatusRequest) GetNoticeBiz() *string {
	return s.NoticeBiz
}

func (s *BatchUpdateNoticeStatusRequest) GetNoticeStatus() *string {
	return s.NoticeStatus
}

func (s *BatchUpdateNoticeStatusRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *BatchUpdateNoticeStatusRequest) SetIds(v string) *BatchUpdateNoticeStatusRequest {
	s.Ids = &v
	return s
}

func (s *BatchUpdateNoticeStatusRequest) SetLang(v string) *BatchUpdateNoticeStatusRequest {
	s.Lang = &v
	return s
}

func (s *BatchUpdateNoticeStatusRequest) SetNoticeBiz(v string) *BatchUpdateNoticeStatusRequest {
	s.NoticeBiz = &v
	return s
}

func (s *BatchUpdateNoticeStatusRequest) SetNoticeStatus(v string) *BatchUpdateNoticeStatusRequest {
	s.NoticeStatus = &v
	return s
}

func (s *BatchUpdateNoticeStatusRequest) SetSourceIp(v string) *BatchUpdateNoticeStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *BatchUpdateNoticeStatusRequest) Validate() error {
	return dara.Validate(s)
}
