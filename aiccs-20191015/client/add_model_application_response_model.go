// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddModelApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddModelApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddModelApplicationResponse
	GetStatusCode() *int32
	SetBody(v *AddModelApplicationResponseBody) *AddModelApplicationResponse
	GetBody() *AddModelApplicationResponseBody
}

type AddModelApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddModelApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddModelApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s AddModelApplicationResponse) GoString() string {
	return s.String()
}

func (s *AddModelApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddModelApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddModelApplicationResponse) GetBody() *AddModelApplicationResponseBody {
	return s.Body
}

func (s *AddModelApplicationResponse) SetHeaders(v map[string]*string) *AddModelApplicationResponse {
	s.Headers = v
	return s
}

func (s *AddModelApplicationResponse) SetStatusCode(v int32) *AddModelApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddModelApplicationResponse) SetBody(v *AddModelApplicationResponseBody) *AddModelApplicationResponse {
	s.Body = v
	return s
}

func (s *AddModelApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
