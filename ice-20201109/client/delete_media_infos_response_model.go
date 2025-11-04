// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaInfosResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaInfosResponseBody) *DeleteMediaInfosResponse
	GetBody() *DeleteMediaInfosResponseBody
}

type DeleteMediaInfosResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaInfosResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaInfosResponse) GetBody() *DeleteMediaInfosResponseBody {
	return s.Body
}

func (s *DeleteMediaInfosResponse) SetHeaders(v map[string]*string) *DeleteMediaInfosResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaInfosResponse) SetStatusCode(v int32) *DeleteMediaInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaInfosResponse) SetBody(v *DeleteMediaInfosResponseBody) *DeleteMediaInfosResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
