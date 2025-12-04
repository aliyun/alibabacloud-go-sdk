// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErAttachmentId(v string) *UpdateErAttachmentRequest
	GetErAttachmentId() *string
	SetErAttachmentName(v string) *UpdateErAttachmentRequest
	GetErAttachmentName() *string
	SetErId(v string) *UpdateErAttachmentRequest
	GetErId() *string
	SetRegionId(v string) *UpdateErAttachmentRequest
	GetRegionId() *string
}

type UpdateErAttachmentRequest struct {
	// The connection ID of the Lingjun HUB network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-attachment-i1ioibyf
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// Lingjun HUB Network Instance Connection Name
	//
	// example:
	//
	// er-attachment-wulanchabu-main
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	// Lingjun HUB ID
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

func (s UpdateErAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateErAttachmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateErAttachmentRequest) GetErAttachmentId() *string {
	return s.ErAttachmentId
}

func (s *UpdateErAttachmentRequest) GetErAttachmentName() *string {
	return s.ErAttachmentName
}

func (s *UpdateErAttachmentRequest) GetErId() *string {
	return s.ErId
}

func (s *UpdateErAttachmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateErAttachmentRequest) SetErAttachmentId(v string) *UpdateErAttachmentRequest {
	s.ErAttachmentId = &v
	return s
}

func (s *UpdateErAttachmentRequest) SetErAttachmentName(v string) *UpdateErAttachmentRequest {
	s.ErAttachmentName = &v
	return s
}

func (s *UpdateErAttachmentRequest) SetErId(v string) *UpdateErAttachmentRequest {
	s.ErId = &v
	return s
}

func (s *UpdateErAttachmentRequest) SetRegionId(v string) *UpdateErAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateErAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
