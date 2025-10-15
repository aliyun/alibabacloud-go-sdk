// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableApplicationResponse
	GetStatusCode() *int32
	SetBody(v *DisableApplicationResponseBody) *DisableApplicationResponse
	GetBody() *DisableApplicationResponseBody
}

type DisableApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableApplicationResponse) GetBody() *DisableApplicationResponseBody {
	return s.Body
}

func (s *DisableApplicationResponse) SetHeaders(v map[string]*string) *DisableApplicationResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationResponse) SetStatusCode(v int32) *DisableApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationResponse) SetBody(v *DisableApplicationResponseBody) *DisableApplicationResponse {
	s.Body = v
	return s
}

func (s *DisableApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
