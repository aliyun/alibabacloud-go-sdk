// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveStreamBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveStreamBlockResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveStreamBlockResponseBody) *DeleteLiveStreamBlockResponse
	GetBody() *DeleteLiveStreamBlockResponseBody
}

type DeleteLiveStreamBlockResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveStreamBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveStreamBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamBlockResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveStreamBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveStreamBlockResponse) GetBody() *DeleteLiveStreamBlockResponseBody {
	return s.Body
}

func (s *DeleteLiveStreamBlockResponse) SetHeaders(v map[string]*string) *DeleteLiveStreamBlockResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveStreamBlockResponse) SetStatusCode(v int32) *DeleteLiveStreamBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveStreamBlockResponse) SetBody(v *DeleteLiveStreamBlockResponseBody) *DeleteLiveStreamBlockResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveStreamBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
