// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAxnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAxnResponse
	GetStatusCode() *int32
	SetBody(v *BindAxnResponseBody) *BindAxnResponse
	GetBody() *BindAxnResponseBody
}

type BindAxnResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxnResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAxnResponse) GoString() string {
	return s.String()
}

func (s *BindAxnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAxnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAxnResponse) GetBody() *BindAxnResponseBody {
	return s.Body
}

func (s *BindAxnResponse) SetHeaders(v map[string]*string) *BindAxnResponse {
	s.Headers = v
	return s
}

func (s *BindAxnResponse) SetStatusCode(v int32) *BindAxnResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxnResponse) SetBody(v *BindAxnResponseBody) *BindAxnResponse {
	s.Body = v
	return s
}

func (s *BindAxnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
