// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBroadcastStickerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBroadcastStickerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBroadcastStickerResponse
	GetStatusCode() *int32
	SetBody(v *CreateBroadcastStickerResponseBody) *CreateBroadcastStickerResponse
	GetBody() *CreateBroadcastStickerResponseBody
}

type CreateBroadcastStickerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBroadcastStickerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBroadcastStickerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBroadcastStickerResponse) GoString() string {
	return s.String()
}

func (s *CreateBroadcastStickerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBroadcastStickerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBroadcastStickerResponse) GetBody() *CreateBroadcastStickerResponseBody {
	return s.Body
}

func (s *CreateBroadcastStickerResponse) SetHeaders(v map[string]*string) *CreateBroadcastStickerResponse {
	s.Headers = v
	return s
}

func (s *CreateBroadcastStickerResponse) SetStatusCode(v int32) *CreateBroadcastStickerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBroadcastStickerResponse) SetBody(v *CreateBroadcastStickerResponseBody) *CreateBroadcastStickerResponse {
	s.Body = v
	return s
}

func (s *CreateBroadcastStickerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
