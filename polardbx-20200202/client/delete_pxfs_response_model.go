// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePxfsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePxfsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePxfsResponse
	GetStatusCode() *int32
	SetBody(v *DeletePxfsResponseBody) *DeletePxfsResponse
	GetBody() *DeletePxfsResponseBody
}

type DeletePxfsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePxfsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePxfsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfsResponse) GoString() string {
	return s.String()
}

func (s *DeletePxfsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePxfsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePxfsResponse) GetBody() *DeletePxfsResponseBody {
	return s.Body
}

func (s *DeletePxfsResponse) SetHeaders(v map[string]*string) *DeletePxfsResponse {
	s.Headers = v
	return s
}

func (s *DeletePxfsResponse) SetStatusCode(v int32) *DeletePxfsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePxfsResponse) SetBody(v *DeletePxfsResponseBody) *DeletePxfsResponse {
	s.Body = v
	return s
}

func (s *DeletePxfsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
