// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAllowedIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAllowedIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAllowedIpResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAllowedIpResponseBody) *UpdateAllowedIpResponse
	GetBody() *UpdateAllowedIpResponseBody
}

type UpdateAllowedIpResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAllowedIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAllowedIpResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAllowedIpResponse) GoString() string {
	return s.String()
}

func (s *UpdateAllowedIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAllowedIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAllowedIpResponse) GetBody() *UpdateAllowedIpResponseBody {
	return s.Body
}

func (s *UpdateAllowedIpResponse) SetHeaders(v map[string]*string) *UpdateAllowedIpResponse {
	s.Headers = v
	return s
}

func (s *UpdateAllowedIpResponse) SetStatusCode(v int32) *UpdateAllowedIpResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAllowedIpResponse) SetBody(v *UpdateAllowedIpResponseBody) *UpdateAllowedIpResponse {
	s.Body = v
	return s
}

func (s *UpdateAllowedIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
