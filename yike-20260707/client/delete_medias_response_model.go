// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediasResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediasResponseBody) *DeleteMediasResponse
	GetBody() *DeleteMediasResponseBody
}

type DeleteMediasResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediasResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediasResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediasResponse) GetBody() *DeleteMediasResponseBody {
	return s.Body
}

func (s *DeleteMediasResponse) SetHeaders(v map[string]*string) *DeleteMediasResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediasResponse) SetStatusCode(v int32) *DeleteMediasResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediasResponse) SetBody(v *DeleteMediasResponseBody) *DeleteMediasResponse {
	s.Body = v
	return s
}

func (s *DeleteMediasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
