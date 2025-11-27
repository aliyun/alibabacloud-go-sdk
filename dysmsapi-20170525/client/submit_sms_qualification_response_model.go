// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmsQualificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSmsQualificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSmsQualificationResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSmsQualificationResponseBody) *SubmitSmsQualificationResponse
	GetBody() *SubmitSmsQualificationResponseBody
}

type SubmitSmsQualificationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSmsQualificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSmsQualificationResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmsQualificationResponse) GoString() string {
	return s.String()
}

func (s *SubmitSmsQualificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSmsQualificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSmsQualificationResponse) GetBody() *SubmitSmsQualificationResponseBody {
	return s.Body
}

func (s *SubmitSmsQualificationResponse) SetHeaders(v map[string]*string) *SubmitSmsQualificationResponse {
	s.Headers = v
	return s
}

func (s *SubmitSmsQualificationResponse) SetStatusCode(v int32) *SubmitSmsQualificationResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSmsQualificationResponse) SetBody(v *SubmitSmsQualificationResponseBody) *SubmitSmsQualificationResponse {
	s.Body = v
	return s
}

func (s *SubmitSmsQualificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
