// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomTextResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomTextResponseBody) *DeleteCustomTextResponse
	GetBody() *DeleteCustomTextResponseBody
}

type DeleteCustomTextResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomTextResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTextResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomTextResponse) GetBody() *DeleteCustomTextResponseBody {
	return s.Body
}

func (s *DeleteCustomTextResponse) SetHeaders(v map[string]*string) *DeleteCustomTextResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomTextResponse) SetStatusCode(v int32) *DeleteCustomTextResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomTextResponse) SetBody(v *DeleteCustomTextResponseBody) *DeleteCustomTextResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
