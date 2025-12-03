// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantResponse
	GetStatusCode() *int32
	SetBody(v *GrantResponseBody) *GrantResponse
	GetBody() *GrantResponseBody
}

type GrantResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantResponse) GoString() string {
	return s.String()
}

func (s *GrantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantResponse) GetBody() *GrantResponseBody {
	return s.Body
}

func (s *GrantResponse) SetHeaders(v map[string]*string) *GrantResponse {
	s.Headers = v
	return s
}

func (s *GrantResponse) SetStatusCode(v int32) *GrantResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantResponse) SetBody(v *GrantResponseBody) *GrantResponse {
	s.Body = v
	return s
}

func (s *GrantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
