// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaResponseBody) *DeleteMediaResponse
	GetBody() *DeleteMediaResponseBody
}

type DeleteMediaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaResponse) GetBody() *DeleteMediaResponseBody {
	return s.Body
}

func (s *DeleteMediaResponse) SetHeaders(v map[string]*string) *DeleteMediaResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaResponse) SetStatusCode(v int32) *DeleteMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaResponse) SetBody(v *DeleteMediaResponseBody) *DeleteMediaResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
