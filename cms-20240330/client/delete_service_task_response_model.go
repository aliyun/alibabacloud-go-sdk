// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceTaskResponseBody) *DeleteServiceTaskResponse
	GetBody() *DeleteServiceTaskResponseBody
}

type DeleteServiceTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceTaskResponse) GetBody() *DeleteServiceTaskResponseBody {
	return s.Body
}

func (s *DeleteServiceTaskResponse) SetHeaders(v map[string]*string) *DeleteServiceTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceTaskResponse) SetStatusCode(v int32) *DeleteServiceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceTaskResponse) SetBody(v *DeleteServiceTaskResponseBody) *DeleteServiceTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
