// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindResponse
	GetStatusCode() *int32
	SetBody(v *BindResponseBody) *BindResponse
	GetBody() *BindResponseBody
}

type BindResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindResponse) String() string {
	return dara.Prettify(s)
}

func (s BindResponse) GoString() string {
	return s.String()
}

func (s *BindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindResponse) GetBody() *BindResponseBody {
	return s.Body
}

func (s *BindResponse) SetHeaders(v map[string]*string) *BindResponse {
	s.Headers = v
	return s
}

func (s *BindResponse) SetStatusCode(v int32) *BindResponse {
	s.StatusCode = &v
	return s
}

func (s *BindResponse) SetBody(v *BindResponseBody) *BindResponse {
	s.Body = v
	return s
}

func (s *BindResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
