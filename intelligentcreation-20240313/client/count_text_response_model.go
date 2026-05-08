// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CountTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CountTextResponse
	GetStatusCode() *int32
	SetBody(v *CountTextResponseBody) *CountTextResponse
	GetBody() *CountTextResponseBody
}

type CountTextResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountTextResponse) String() string {
	return dara.Prettify(s)
}

func (s CountTextResponse) GoString() string {
	return s.String()
}

func (s *CountTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CountTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CountTextResponse) GetBody() *CountTextResponseBody {
	return s.Body
}

func (s *CountTextResponse) SetHeaders(v map[string]*string) *CountTextResponse {
	s.Headers = v
	return s
}

func (s *CountTextResponse) SetStatusCode(v int32) *CountTextResponse {
	s.StatusCode = &v
	return s
}

func (s *CountTextResponse) SetBody(v *CountTextResponseBody) *CountTextResponse {
	s.Body = v
	return s
}

func (s *CountTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
