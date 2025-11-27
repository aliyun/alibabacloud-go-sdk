// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeSignatureQualificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeSignatureQualificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeSignatureQualificationResponse
	GetStatusCode() *int32
	SetBody(v *ChangeSignatureQualificationResponseBody) *ChangeSignatureQualificationResponse
	GetBody() *ChangeSignatureQualificationResponseBody
}

type ChangeSignatureQualificationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeSignatureQualificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeSignatureQualificationResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeSignatureQualificationResponse) GoString() string {
	return s.String()
}

func (s *ChangeSignatureQualificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeSignatureQualificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeSignatureQualificationResponse) GetBody() *ChangeSignatureQualificationResponseBody {
	return s.Body
}

func (s *ChangeSignatureQualificationResponse) SetHeaders(v map[string]*string) *ChangeSignatureQualificationResponse {
	s.Headers = v
	return s
}

func (s *ChangeSignatureQualificationResponse) SetStatusCode(v int32) *ChangeSignatureQualificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeSignatureQualificationResponse) SetBody(v *ChangeSignatureQualificationResponseBody) *ChangeSignatureQualificationResponse {
	s.Body = v
	return s
}

func (s *ChangeSignatureQualificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
