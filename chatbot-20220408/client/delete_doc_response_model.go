// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDocResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDocResponseBody) *DeleteDocResponse
	GetBody() *DeleteDocResponseBody
}

type DeleteDocResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDocResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocResponse) GoString() string {
	return s.String()
}

func (s *DeleteDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDocResponse) GetBody() *DeleteDocResponseBody {
	return s.Body
}

func (s *DeleteDocResponse) SetHeaders(v map[string]*string) *DeleteDocResponse {
	s.Headers = v
	return s
}

func (s *DeleteDocResponse) SetStatusCode(v int32) *DeleteDocResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDocResponse) SetBody(v *DeleteDocResponseBody) *DeleteDocResponse {
	s.Body = v
	return s
}

func (s *DeleteDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
