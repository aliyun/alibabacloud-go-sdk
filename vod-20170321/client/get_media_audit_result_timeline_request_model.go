// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAuditResultTimelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetMediaAuditResultTimelineRequest
	GetMediaId() *string
}

type GetMediaAuditResultTimelineRequest struct {
	// The ID of the video.
	//
	// This parameter is required.
	//
	// example:
	//
	// 93ab850b4f6f*****54b6e91d24d81d4
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetMediaAuditResultTimelineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaAuditResultTimelineRequest) SetMediaId(v string) *GetMediaAuditResultTimelineRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaAuditResultTimelineRequest) Validate() error {
	return dara.Validate(s)
}
