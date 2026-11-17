// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableOutputItemShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableOutputItemShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableOutputItemShareResponse
	GetStatusCode() *int32
	SetBody(v *DisableOutputItemShareResponseBody) *DisableOutputItemShareResponse
	GetBody() *DisableOutputItemShareResponseBody
}

type DisableOutputItemShareResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableOutputItemShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableOutputItemShareResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableOutputItemShareResponse) GoString() string {
	return s.String()
}

func (s *DisableOutputItemShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableOutputItemShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableOutputItemShareResponse) GetBody() *DisableOutputItemShareResponseBody {
	return s.Body
}

func (s *DisableOutputItemShareResponse) SetHeaders(v map[string]*string) *DisableOutputItemShareResponse {
	s.Headers = v
	return s
}

func (s *DisableOutputItemShareResponse) SetStatusCode(v int32) *DisableOutputItemShareResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableOutputItemShareResponse) SetBody(v *DisableOutputItemShareResponseBody) *DisableOutputItemShareResponse {
	s.Body = v
	return s
}

func (s *DisableOutputItemShareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
