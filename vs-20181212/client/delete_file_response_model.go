// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFileResponseBody) *DeleteFileResponse
	GetBody() *DeleteFileResponseBody
}

type DeleteFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFileResponse) GetBody() *DeleteFileResponseBody {
	return s.Body
}

func (s *DeleteFileResponse) SetHeaders(v map[string]*string) *DeleteFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileResponse) SetStatusCode(v int32) *DeleteFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileResponse) SetBody(v *DeleteFileResponseBody) *DeleteFileResponse {
	s.Body = v
	return s
}

func (s *DeleteFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
