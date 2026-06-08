// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppGroupDeleteProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppGroupDeleteProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppGroupDeleteProtectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppGroupDeleteProtectionResponseBody) *UpdateAppGroupDeleteProtectionResponse
	GetBody() *UpdateAppGroupDeleteProtectionResponseBody
}

type UpdateAppGroupDeleteProtectionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppGroupDeleteProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppGroupDeleteProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppGroupDeleteProtectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppGroupDeleteProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppGroupDeleteProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppGroupDeleteProtectionResponse) GetBody() *UpdateAppGroupDeleteProtectionResponseBody {
	return s.Body
}

func (s *UpdateAppGroupDeleteProtectionResponse) SetHeaders(v map[string]*string) *UpdateAppGroupDeleteProtectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppGroupDeleteProtectionResponse) SetStatusCode(v int32) *UpdateAppGroupDeleteProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppGroupDeleteProtectionResponse) SetBody(v *UpdateAppGroupDeleteProtectionResponseBody) *UpdateAppGroupDeleteProtectionResponse {
	s.Body = v
	return s
}

func (s *UpdateAppGroupDeleteProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
