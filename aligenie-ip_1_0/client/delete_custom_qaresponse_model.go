// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomQAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomQAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomQAResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomQAResponseBody) *DeleteCustomQAResponse
	GetBody() *DeleteCustomQAResponseBody
}

type DeleteCustomQAResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomQAResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomQAResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomQAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomQAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomQAResponse) GetBody() *DeleteCustomQAResponseBody {
	return s.Body
}

func (s *DeleteCustomQAResponse) SetHeaders(v map[string]*string) *DeleteCustomQAResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomQAResponse) SetStatusCode(v int32) *DeleteCustomQAResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomQAResponse) SetBody(v *DeleteCustomQAResponseBody) *DeleteCustomQAResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomQAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
