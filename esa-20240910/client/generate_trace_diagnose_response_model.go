// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTraceDiagnoseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateTraceDiagnoseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateTraceDiagnoseResponse
	GetStatusCode() *int32
	SetBody(v *GenerateTraceDiagnoseResponseBody) *GenerateTraceDiagnoseResponse
	GetBody() *GenerateTraceDiagnoseResponseBody
}

type GenerateTraceDiagnoseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTraceDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTraceDiagnoseResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateTraceDiagnoseResponse) GoString() string {
	return s.String()
}

func (s *GenerateTraceDiagnoseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateTraceDiagnoseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateTraceDiagnoseResponse) GetBody() *GenerateTraceDiagnoseResponseBody {
	return s.Body
}

func (s *GenerateTraceDiagnoseResponse) SetHeaders(v map[string]*string) *GenerateTraceDiagnoseResponse {
	s.Headers = v
	return s
}

func (s *GenerateTraceDiagnoseResponse) SetStatusCode(v int32) *GenerateTraceDiagnoseResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTraceDiagnoseResponse) SetBody(v *GenerateTraceDiagnoseResponseBody) *GenerateTraceDiagnoseResponse {
	s.Body = v
	return s
}

func (s *GenerateTraceDiagnoseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
