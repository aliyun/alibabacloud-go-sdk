// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyActivityAdmissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyActivityAdmissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyActivityAdmissionResponse
	GetStatusCode() *int32
	SetBody(v *VerifyActivityAdmissionResponseBody) *VerifyActivityAdmissionResponse
	GetBody() *VerifyActivityAdmissionResponseBody
}

type VerifyActivityAdmissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyActivityAdmissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyActivityAdmissionResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyActivityAdmissionResponse) GoString() string {
	return s.String()
}

func (s *VerifyActivityAdmissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyActivityAdmissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyActivityAdmissionResponse) GetBody() *VerifyActivityAdmissionResponseBody {
	return s.Body
}

func (s *VerifyActivityAdmissionResponse) SetHeaders(v map[string]*string) *VerifyActivityAdmissionResponse {
	s.Headers = v
	return s
}

func (s *VerifyActivityAdmissionResponse) SetStatusCode(v int32) *VerifyActivityAdmissionResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyActivityAdmissionResponse) SetBody(v *VerifyActivityAdmissionResponseBody) *VerifyActivityAdmissionResponse {
	s.Body = v
	return s
}

func (s *VerifyActivityAdmissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
