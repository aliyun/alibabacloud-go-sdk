// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomLineResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomLineResponseBody) *DeleteCustomLineResponse
	GetBody() *DeleteCustomLineResponseBody
}

type DeleteCustomLineResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomLineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomLineResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomLineResponse) GetBody() *DeleteCustomLineResponseBody {
	return s.Body
}

func (s *DeleteCustomLineResponse) SetHeaders(v map[string]*string) *DeleteCustomLineResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomLineResponse) SetStatusCode(v int32) *DeleteCustomLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomLineResponse) SetBody(v *DeleteCustomLineResponseBody) *DeleteCustomLineResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
