// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErAttachmentId(v string) *GetErAttachmentRequest
	GetErAttachmentId() *string
	SetErId(v string) *GetErAttachmentRequest
	GetErId() *string
	SetRegionId(v string) *GetErAttachmentRequest
	GetRegionId() *string
}

type GetErAttachmentRequest struct {
	// The ID of the Lingjun HUB network connection instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-attachment-i1ioibyf
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// Lingjun HUB ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetErAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetErAttachmentRequest) GoString() string {
	return s.String()
}

func (s *GetErAttachmentRequest) GetErAttachmentId() *string {
	return s.ErAttachmentId
}

func (s *GetErAttachmentRequest) GetErId() *string {
	return s.ErId
}

func (s *GetErAttachmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetErAttachmentRequest) SetErAttachmentId(v string) *GetErAttachmentRequest {
	s.ErAttachmentId = &v
	return s
}

func (s *GetErAttachmentRequest) SetErId(v string) *GetErAttachmentRequest {
	s.ErId = &v
	return s
}

func (s *GetErAttachmentRequest) SetRegionId(v string) *GetErAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *GetErAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
