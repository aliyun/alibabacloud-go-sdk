// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnnouncementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnnouncementId(v int64) *CreateAnnouncementResponseBody
	GetAnnouncementId() *int64
	SetCode(v string) *CreateAnnouncementResponseBody
	GetCode() *string
	SetCreatedBy(v int64) *CreateAnnouncementResponseBody
	GetCreatedBy() *int64
	SetMessage(v string) *CreateAnnouncementResponseBody
	GetMessage() *string
	SetPublishedAt(v string) *CreateAnnouncementResponseBody
	GetPublishedAt() *string
	SetRequestId(v string) *CreateAnnouncementResponseBody
	GetRequestId() *string
	SetSourceType(v string) *CreateAnnouncementResponseBody
	GetSourceType() *string
	SetStatus(v string) *CreateAnnouncementResponseBody
	GetStatus() *string
}

type CreateAnnouncementResponseBody struct {
	// The business ID of the notice.
	//
	// example:
	//
	// 1001
	AnnouncementId *int64 `json:"announcementId,omitempty" xml:"announcementId,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The creator.
	//
	// example:
	//
	// 10001
	CreatedBy *int64 `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The response message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The publish time in ISO 8601 format. This field is empty for drafts.
	//
	// example:
	//
	// 2026-08-20T14:00:00+08:00
	PublishedAt *string `json:"publishedAt,omitempty" xml:"publishedAt,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The source type of the dictionary file. Valid values: OSS: Object Storage Service (OSS). ORIGIN: retains the previously uploaded dictionary.
	//
	// example:
	//
	// PLATFORM
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The refund status. You must query this field to confirm the refund status during processing. Valid values:
	//
	// - SUCCESS: All refunds succeeded.
	//
	// - FAIL: Failed.
	//
	// - WAIT_PAY: Waiting for refund.
	//
	// - EXPIRE: Expired.
	//
	// - PAYING: Refund in progress.
	//
	// - TERMINATE: Refund terminated.
	//
	// example:
	//
	// PUBLISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateAnnouncementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAnnouncementResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAnnouncementResponseBody) GetAnnouncementId() *int64 {
	return s.AnnouncementId
}

func (s *CreateAnnouncementResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAnnouncementResponseBody) GetCreatedBy() *int64 {
	return s.CreatedBy
}

func (s *CreateAnnouncementResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAnnouncementResponseBody) GetPublishedAt() *string {
	return s.PublishedAt
}

func (s *CreateAnnouncementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAnnouncementResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateAnnouncementResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateAnnouncementResponseBody) SetAnnouncementId(v int64) *CreateAnnouncementResponseBody {
	s.AnnouncementId = &v
	return s
}

func (s *CreateAnnouncementResponseBody) SetCode(v string) *CreateAnnouncementResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAnnouncementResponseBody) SetCreatedBy(v int64) *CreateAnnouncementResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *CreateAnnouncementResponseBody) SetMessage(v string) *CreateAnnouncementResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAnnouncementResponseBody) SetPublishedAt(v string) *CreateAnnouncementResponseBody {
	s.PublishedAt = &v
	return s
}

func (s *CreateAnnouncementResponseBody) SetRequestId(v string) *CreateAnnouncementResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAnnouncementResponseBody) SetSourceType(v string) *CreateAnnouncementResponseBody {
	s.SourceType = &v
	return s
}

func (s *CreateAnnouncementResponseBody) SetStatus(v string) *CreateAnnouncementResponseBody {
	s.Status = &v
	return s
}

func (s *CreateAnnouncementResponseBody) Validate() error {
	return dara.Validate(s)
}
