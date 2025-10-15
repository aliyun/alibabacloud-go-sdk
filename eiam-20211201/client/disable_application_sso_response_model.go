// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationSsoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableApplicationSsoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableApplicationSsoResponse
	GetStatusCode() *int32
	SetBody(v *DisableApplicationSsoResponseBody) *DisableApplicationSsoResponse
	GetBody() *DisableApplicationSsoResponseBody
}

type DisableApplicationSsoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationSsoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationSsoResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationSsoResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationSsoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableApplicationSsoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableApplicationSsoResponse) GetBody() *DisableApplicationSsoResponseBody {
	return s.Body
}

func (s *DisableApplicationSsoResponse) SetHeaders(v map[string]*string) *DisableApplicationSsoResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationSsoResponse) SetStatusCode(v int32) *DisableApplicationSsoResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationSsoResponse) SetBody(v *DisableApplicationSsoResponseBody) *DisableApplicationSsoResponse {
	s.Body = v
	return s
}

func (s *DisableApplicationSsoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
