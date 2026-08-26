// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableTrustedOriginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableTrustedOriginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableTrustedOriginResponse
	GetStatusCode() *int32
	SetBody(v *DisableTrustedOriginResponseBody) *DisableTrustedOriginResponse
	GetBody() *DisableTrustedOriginResponseBody
}

type DisableTrustedOriginResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableTrustedOriginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableTrustedOriginResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableTrustedOriginResponse) GoString() string {
	return s.String()
}

func (s *DisableTrustedOriginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableTrustedOriginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableTrustedOriginResponse) GetBody() *DisableTrustedOriginResponseBody {
	return s.Body
}

func (s *DisableTrustedOriginResponse) SetHeaders(v map[string]*string) *DisableTrustedOriginResponse {
	s.Headers = v
	return s
}

func (s *DisableTrustedOriginResponse) SetStatusCode(v int32) *DisableTrustedOriginResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableTrustedOriginResponse) SetBody(v *DisableTrustedOriginResponseBody) *DisableTrustedOriginResponse {
	s.Body = v
	return s
}

func (s *DisableTrustedOriginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
