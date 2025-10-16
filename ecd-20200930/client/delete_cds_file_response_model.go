// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCdsFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCdsFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCdsFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCdsFileResponseBody) *DeleteCdsFileResponse
	GetBody() *DeleteCdsFileResponseBody
}

type DeleteCdsFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCdsFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCdsFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCdsFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteCdsFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCdsFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCdsFileResponse) GetBody() *DeleteCdsFileResponseBody {
	return s.Body
}

func (s *DeleteCdsFileResponse) SetHeaders(v map[string]*string) *DeleteCdsFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteCdsFileResponse) SetStatusCode(v int32) *DeleteCdsFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCdsFileResponse) SetBody(v *DeleteCdsFileResponseBody) *DeleteCdsFileResponse {
	s.Body = v
	return s
}

func (s *DeleteCdsFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
