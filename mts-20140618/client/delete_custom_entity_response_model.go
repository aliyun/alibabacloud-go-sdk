// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomEntityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomEntityResponseBody) *DeleteCustomEntityResponse
	GetBody() *DeleteCustomEntityResponseBody
}

type DeleteCustomEntityResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomEntityResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomEntityResponse) GetBody() *DeleteCustomEntityResponseBody {
	return s.Body
}

func (s *DeleteCustomEntityResponse) SetHeaders(v map[string]*string) *DeleteCustomEntityResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomEntityResponse) SetStatusCode(v int32) *DeleteCustomEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomEntityResponse) SetBody(v *DeleteCustomEntityResponseBody) *DeleteCustomEntityResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
