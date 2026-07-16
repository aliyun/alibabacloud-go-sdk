// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWaitingRoomPreviewPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPageHtml(v string) *GetWaitingRoomPreviewPageRequest
	GetCustomPageHtml() *string
}

type GetWaitingRoomPreviewPageRequest struct {
	// The custom waiting room page content. This parameter is required when the waiting room type is custom. The content must be URL-encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello%20world!
	CustomPageHtml *string `json:"CustomPageHtml,omitempty" xml:"CustomPageHtml,omitempty"`
}

func (s GetWaitingRoomPreviewPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWaitingRoomPreviewPageRequest) GoString() string {
	return s.String()
}

func (s *GetWaitingRoomPreviewPageRequest) GetCustomPageHtml() *string {
	return s.CustomPageHtml
}

func (s *GetWaitingRoomPreviewPageRequest) SetCustomPageHtml(v string) *GetWaitingRoomPreviewPageRequest {
	s.CustomPageHtml = &v
	return s
}

func (s *GetWaitingRoomPreviewPageRequest) Validate() error {
	return dara.Validate(s)
}
