// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOriginProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOriginProtectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOriginProtectionResponseBody) *UpdateOriginProtectionResponse
	GetBody() *UpdateOriginProtectionResponseBody
}

type UpdateOriginProtectionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOriginProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOriginProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginProtectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateOriginProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOriginProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOriginProtectionResponse) GetBody() *UpdateOriginProtectionResponseBody {
	return s.Body
}

func (s *UpdateOriginProtectionResponse) SetHeaders(v map[string]*string) *UpdateOriginProtectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateOriginProtectionResponse) SetStatusCode(v int32) *UpdateOriginProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOriginProtectionResponse) SetBody(v *UpdateOriginProtectionResponseBody) *UpdateOriginProtectionResponse {
	s.Body = v
	return s
}

func (s *UpdateOriginProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
