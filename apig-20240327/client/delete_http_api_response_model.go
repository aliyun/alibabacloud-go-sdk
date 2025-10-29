// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHttpApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHttpApiResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHttpApiResponseBody) *DeleteHttpApiResponse
	GetBody() *DeleteHttpApiResponseBody
}

type DeleteHttpApiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHttpApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHttpApiResponse) GetBody() *DeleteHttpApiResponseBody {
	return s.Body
}

func (s *DeleteHttpApiResponse) SetHeaders(v map[string]*string) *DeleteHttpApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpApiResponse) SetStatusCode(v int32) *DeleteHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpApiResponse) SetBody(v *DeleteHttpApiResponseBody) *DeleteHttpApiResponse {
	s.Body = v
	return s
}

func (s *DeleteHttpApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
