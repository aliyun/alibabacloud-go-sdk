// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKgEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKgEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKgEntityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKgEntityResponseBody) *DeleteKgEntityResponse
	GetBody() *DeleteKgEntityResponseBody
}

type DeleteKgEntityResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKgEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKgEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKgEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteKgEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKgEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKgEntityResponse) GetBody() *DeleteKgEntityResponseBody {
	return s.Body
}

func (s *DeleteKgEntityResponse) SetHeaders(v map[string]*string) *DeleteKgEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteKgEntityResponse) SetStatusCode(v int32) *DeleteKgEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKgEntityResponse) SetBody(v *DeleteKgEntityResponseBody) *DeleteKgEntityResponse {
	s.Body = v
	return s
}

func (s *DeleteKgEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
