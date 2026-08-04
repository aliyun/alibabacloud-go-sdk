// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContacterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContacterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContacterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContacterResponseBody) *DeleteContacterResponse
	GetBody() *DeleteContacterResponseBody
}

type DeleteContacterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContacterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContacterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContacterResponse) GoString() string {
	return s.String()
}

func (s *DeleteContacterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContacterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContacterResponse) GetBody() *DeleteContacterResponseBody {
	return s.Body
}

func (s *DeleteContacterResponse) SetHeaders(v map[string]*string) *DeleteContacterResponse {
	s.Headers = v
	return s
}

func (s *DeleteContacterResponse) SetStatusCode(v int32) *DeleteContacterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContacterResponse) SetBody(v *DeleteContacterResponseBody) *DeleteContacterResponse {
	s.Body = v
	return s
}

func (s *DeleteContacterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
