// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAuditResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetMediaAuditResultRequest
	GetMediaId() *string
}

type GetMediaAuditResultRequest struct {
	// The ID of the video or image.
	//
	// This parameter is required.
	//
	// example:
	//
	// 93ab850b4f6f*****54b6e91d24d81d4
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetMediaAuditResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditResultRequest) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaAuditResultRequest) SetMediaId(v string) *GetMediaAuditResultRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaAuditResultRequest) Validate() error {
	return dara.Validate(s)
}
