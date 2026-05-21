// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTempFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTempFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTempFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTempFileResponseBody) *DeleteTempFileResponse
	GetBody() *DeleteTempFileResponseBody
}

type DeleteTempFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTempFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTempFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTempFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteTempFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTempFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTempFileResponse) GetBody() *DeleteTempFileResponseBody {
	return s.Body
}

func (s *DeleteTempFileResponse) SetHeaders(v map[string]*string) *DeleteTempFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteTempFileResponse) SetStatusCode(v int32) *DeleteTempFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTempFileResponse) SetBody(v *DeleteTempFileResponseBody) *DeleteTempFileResponse {
	s.Body = v
	return s
}

func (s *DeleteTempFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
