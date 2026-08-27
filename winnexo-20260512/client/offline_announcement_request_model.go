// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineAnnouncementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnouncementId(v int64) *OfflineAnnouncementRequest
	GetAnnouncementId() *int64
	SetTenantId(v string) *OfflineAnnouncementRequest
	GetTenantId() *string
}

type OfflineAnnouncementRequest struct {
	// The business ID of the announcement.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	AnnouncementId *int64 `json:"announcementId,omitempty" xml:"announcementId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this explicitly with --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s OfflineAnnouncementRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineAnnouncementRequest) GoString() string {
	return s.String()
}

func (s *OfflineAnnouncementRequest) GetAnnouncementId() *int64 {
	return s.AnnouncementId
}

func (s *OfflineAnnouncementRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *OfflineAnnouncementRequest) SetAnnouncementId(v int64) *OfflineAnnouncementRequest {
	s.AnnouncementId = &v
	return s
}

func (s *OfflineAnnouncementRequest) SetTenantId(v string) *OfflineAnnouncementRequest {
	s.TenantId = &v
	return s
}

func (s *OfflineAnnouncementRequest) Validate() error {
	return dara.Validate(s)
}
