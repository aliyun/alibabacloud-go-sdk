// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateRecord interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *TemplateRecord
	GetBizId() *string
	SetCopyStatus(v string) *TemplateRecord
	GetCopyStatus() *string
	SetGmtCreate(v string) *TemplateRecord
	GetGmtCreate() *string
	SetGmtModified(v string) *TemplateRecord
	GetGmtModified() *string
	SetId(v int64) *TemplateRecord
	GetId() *int64
	SetTemplateId(v string) *TemplateRecord
	GetTemplateId() *string
	SetUserId(v string) *TemplateRecord
	GetUserId() *string
}

type TemplateRecord struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CopyStatus *string `json:"CopyStatus,omitempty" xml:"CopyStatus,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s TemplateRecord) String() string {
	return dara.Prettify(s)
}

func (s TemplateRecord) GoString() string {
	return s.String()
}

func (s *TemplateRecord) GetBizId() *string {
	return s.BizId
}

func (s *TemplateRecord) GetCopyStatus() *string {
	return s.CopyStatus
}

func (s *TemplateRecord) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TemplateRecord) GetGmtModified() *string {
	return s.GmtModified
}

func (s *TemplateRecord) GetId() *int64 {
	return s.Id
}

func (s *TemplateRecord) GetTemplateId() *string {
	return s.TemplateId
}

func (s *TemplateRecord) GetUserId() *string {
	return s.UserId
}

func (s *TemplateRecord) SetBizId(v string) *TemplateRecord {
	s.BizId = &v
	return s
}

func (s *TemplateRecord) SetCopyStatus(v string) *TemplateRecord {
	s.CopyStatus = &v
	return s
}

func (s *TemplateRecord) SetGmtCreate(v string) *TemplateRecord {
	s.GmtCreate = &v
	return s
}

func (s *TemplateRecord) SetGmtModified(v string) *TemplateRecord {
	s.GmtModified = &v
	return s
}

func (s *TemplateRecord) SetId(v int64) *TemplateRecord {
	s.Id = &v
	return s
}

func (s *TemplateRecord) SetTemplateId(v string) *TemplateRecord {
	s.TemplateId = &v
	return s
}

func (s *TemplateRecord) SetUserId(v string) *TemplateRecord {
	s.UserId = &v
	return s
}

func (s *TemplateRecord) Validate() error {
	return dara.Validate(s)
}
