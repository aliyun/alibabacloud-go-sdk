// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkitemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkitemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkitemResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkitemResponseBody) *DeleteWorkitemResponse
	GetBody() *DeleteWorkitemResponseBody
}

type DeleteWorkitemResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkitemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkitemResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkitemResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkitemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkitemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkitemResponse) GetBody() *DeleteWorkitemResponseBody {
	return s.Body
}

func (s *DeleteWorkitemResponse) SetHeaders(v map[string]*string) *DeleteWorkitemResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkitemResponse) SetStatusCode(v int32) *DeleteWorkitemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkitemResponse) SetBody(v *DeleteWorkitemResponseBody) *DeleteWorkitemResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkitemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
