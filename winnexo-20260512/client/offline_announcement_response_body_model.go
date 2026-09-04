// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineAnnouncementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnnouncementId(v int64) *OfflineAnnouncementResponseBody
	GetAnnouncementId() *int64
	SetChanged(v bool) *OfflineAnnouncementResponseBody
	GetChanged() *bool
	SetCode(v string) *OfflineAnnouncementResponseBody
	GetCode() *string
	SetGmtModified(v string) *OfflineAnnouncementResponseBody
	GetGmtModified() *string
	SetMessage(v string) *OfflineAnnouncementResponseBody
	GetMessage() *string
	SetRequestId(v string) *OfflineAnnouncementResponseBody
	GetRequestId() *string
	SetStatus(v string) *OfflineAnnouncementResponseBody
	GetStatus() *string
	SetUpdatedBy(v int64) *OfflineAnnouncementResponseBody
	GetUpdatedBy() *int64
}

type OfflineAnnouncementResponseBody struct {
	// The business ID of the announcement.
	//
	// example:
	//
	// 1001
	AnnouncementId *int64 `json:"announcementId,omitempty" xml:"announcementId,omitempty"`
	// Indicates whether the status was changed.
	//
	// example:
	//
	// true
	Changed *bool `json:"changed,omitempty" xml:"changed,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The last update time.
	//
	// example:
	//
	// 2026-08-20T14:00:00+08:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The data source status after re-parsing.
	//
	// example:
	//
	// OFFLINE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The user who performed the update.
	//
	// example:
	//
	// 10001
	UpdatedBy *int64 `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s OfflineAnnouncementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineAnnouncementResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineAnnouncementResponseBody) GetAnnouncementId() *int64 {
	return s.AnnouncementId
}

func (s *OfflineAnnouncementResponseBody) GetChanged() *bool {
	return s.Changed
}

func (s *OfflineAnnouncementResponseBody) GetCode() *string {
	return s.Code
}

func (s *OfflineAnnouncementResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *OfflineAnnouncementResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OfflineAnnouncementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineAnnouncementResponseBody) GetStatus() *string {
	return s.Status
}

func (s *OfflineAnnouncementResponseBody) GetUpdatedBy() *int64 {
	return s.UpdatedBy
}

func (s *OfflineAnnouncementResponseBody) SetAnnouncementId(v int64) *OfflineAnnouncementResponseBody {
	s.AnnouncementId = &v
	return s
}

func (s *OfflineAnnouncementResponseBody) SetChanged(v bool) *OfflineAnnouncementResponseBody {
	s.Changed = &v
	return s
}

func (s *OfflineAnnouncementResponseBody) SetCode(v string) *OfflineAnnouncementResponseBody {
	s.Code = &v
	return s
}

func (s *OfflineAnnouncementResponseBody) SetGmtModified(v string) *OfflineAnnouncementResponseBody {
	s.GmtModified = &v
	return s
}

func (s *OfflineAnnouncementResponseBody) SetMessage(v string) *OfflineAnnouncementResponseBody {
	s.Message = &v
	return s
}

func (s *OfflineAnnouncementResponseBody) SetRequestId(v string) *OfflineAnnouncementResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineAnnouncementResponseBody) SetStatus(v string) *OfflineAnnouncementResponseBody {
	s.Status = &v
	return s
}

func (s *OfflineAnnouncementResponseBody) SetUpdatedBy(v int64) *OfflineAnnouncementResponseBody {
	s.UpdatedBy = &v
	return s
}

func (s *OfflineAnnouncementResponseBody) Validate() error {
	return dara.Validate(s)
}
