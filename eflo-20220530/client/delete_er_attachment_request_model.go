// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteErAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErAttachmentId(v string) *DeleteErAttachmentRequest
	GetErAttachmentId() *string
	SetErId(v string) *DeleteErAttachmentRequest
	GetErId() *string
	SetRegionId(v string) *DeleteErAttachmentRequest
	GetRegionId() *string
}

type DeleteErAttachmentRequest struct {
	// The ID of the network connection instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// er-attachment-5n3nsmvl
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// Lingjun HUB Id
	//
	// This parameter is required.
	//
	// example:
	//
	// er-opy1wrfv
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

func (s DeleteErAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteErAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteErAttachmentRequest) GetErAttachmentId() *string {
	return s.ErAttachmentId
}

func (s *DeleteErAttachmentRequest) GetErId() *string {
	return s.ErId
}

func (s *DeleteErAttachmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteErAttachmentRequest) SetErAttachmentId(v string) *DeleteErAttachmentRequest {
	s.ErAttachmentId = &v
	return s
}

func (s *DeleteErAttachmentRequest) SetErId(v string) *DeleteErAttachmentRequest {
	s.ErId = &v
	return s
}

func (s *DeleteErAttachmentRequest) SetRegionId(v string) *DeleteErAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteErAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
