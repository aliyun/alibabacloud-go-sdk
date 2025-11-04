// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlayInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePlayInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePlayInfoResponse
	GetStatusCode() *int32
	SetBody(v *DeletePlayInfoResponseBody) *DeletePlayInfoResponse
	GetBody() *DeletePlayInfoResponseBody
}

type DeletePlayInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePlayInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePlayInfoResponse) GoString() string {
	return s.String()
}

func (s *DeletePlayInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePlayInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePlayInfoResponse) GetBody() *DeletePlayInfoResponseBody {
	return s.Body
}

func (s *DeletePlayInfoResponse) SetHeaders(v map[string]*string) *DeletePlayInfoResponse {
	s.Headers = v
	return s
}

func (s *DeletePlayInfoResponse) SetStatusCode(v int32) *DeletePlayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePlayInfoResponse) SetBody(v *DeletePlayInfoResponseBody) *DeletePlayInfoResponse {
	s.Body = v
	return s
}

func (s *DeletePlayInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
