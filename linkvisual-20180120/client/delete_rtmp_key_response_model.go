// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRtmpKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRtmpKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRtmpKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRtmpKeyResponseBody) *DeleteRtmpKeyResponse
	GetBody() *DeleteRtmpKeyResponseBody
}

type DeleteRtmpKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRtmpKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRtmpKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRtmpKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteRtmpKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRtmpKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRtmpKeyResponse) GetBody() *DeleteRtmpKeyResponseBody {
	return s.Body
}

func (s *DeleteRtmpKeyResponse) SetHeaders(v map[string]*string) *DeleteRtmpKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteRtmpKeyResponse) SetStatusCode(v int32) *DeleteRtmpKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRtmpKeyResponse) SetBody(v *DeleteRtmpKeyResponseBody) *DeleteRtmpKeyResponse {
	s.Body = v
	return s
}

func (s *DeleteRtmpKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
