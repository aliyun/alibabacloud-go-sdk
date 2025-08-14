// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshUploadMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *RefreshUploadMediaRequest
	GetMediaId() *string
}

type RefreshUploadMediaRequest struct {
	// The ID of the media asset.
	//
	// example:
	//
	// 411bed50018971edb60b0764a0ec6***
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s RefreshUploadMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshUploadMediaRequest) GoString() string {
	return s.String()
}

func (s *RefreshUploadMediaRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *RefreshUploadMediaRequest) SetMediaId(v string) *RefreshUploadMediaRequest {
	s.MediaId = &v
	return s
}

func (s *RefreshUploadMediaRequest) Validate() error {
	return dara.Validate(s)
}
