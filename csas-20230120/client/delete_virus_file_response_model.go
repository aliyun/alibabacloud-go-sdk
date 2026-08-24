// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirusFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVirusFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVirusFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVirusFileResponseBody) *DeleteVirusFileResponse
	GetBody() *DeleteVirusFileResponseBody
}

type DeleteVirusFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirusFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVirusFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirusFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirusFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVirusFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVirusFileResponse) GetBody() *DeleteVirusFileResponseBody {
	return s.Body
}

func (s *DeleteVirusFileResponse) SetHeaders(v map[string]*string) *DeleteVirusFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirusFileResponse) SetStatusCode(v int32) *DeleteVirusFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirusFileResponse) SetBody(v *DeleteVirusFileResponseBody) *DeleteVirusFileResponse {
	s.Body = v
	return s
}

func (s *DeleteVirusFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
