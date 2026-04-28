// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelApiResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelApiResponseBody) *DeleteModelApiResponse
	GetBody() *DeleteModelApiResponseBody
}

type DeleteModelApiResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelApiResponse) GetBody() *DeleteModelApiResponseBody {
	return s.Body
}

func (s *DeleteModelApiResponse) SetHeaders(v map[string]*string) *DeleteModelApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelApiResponse) SetStatusCode(v int32) *DeleteModelApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelApiResponse) SetBody(v *DeleteModelApiResponseBody) *DeleteModelApiResponse {
	s.Body = v
	return s
}

func (s *DeleteModelApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
