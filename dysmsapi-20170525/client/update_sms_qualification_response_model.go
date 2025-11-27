// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsQualificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmsQualificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmsQualificationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmsQualificationResponseBody) *UpdateSmsQualificationResponse
	GetBody() *UpdateSmsQualificationResponseBody
}

type UpdateSmsQualificationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmsQualificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmsQualificationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsQualificationResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmsQualificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmsQualificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmsQualificationResponse) GetBody() *UpdateSmsQualificationResponseBody {
	return s.Body
}

func (s *UpdateSmsQualificationResponse) SetHeaders(v map[string]*string) *UpdateSmsQualificationResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmsQualificationResponse) SetStatusCode(v int32) *UpdateSmsQualificationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmsQualificationResponse) SetBody(v *UpdateSmsQualificationResponseBody) *UpdateSmsQualificationResponse {
	s.Body = v
	return s
}

func (s *UpdateSmsQualificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
