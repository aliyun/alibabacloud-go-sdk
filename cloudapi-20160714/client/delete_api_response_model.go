// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApiResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApiResponseBody) *DeleteApiResponse
	GetBody() *DeleteApiResponseBody
}

type DeleteApiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApiResponse) GetBody() *DeleteApiResponseBody {
	return s.Body
}

func (s *DeleteApiResponse) SetHeaders(v map[string]*string) *DeleteApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiResponse) SetStatusCode(v int32) *DeleteApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiResponse) SetBody(v *DeleteApiResponseBody) *DeleteApiResponse {
	s.Body = v
	return s
}

func (s *DeleteApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
