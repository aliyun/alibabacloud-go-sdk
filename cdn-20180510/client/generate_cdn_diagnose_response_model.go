// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCdnDiagnoseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateCdnDiagnoseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateCdnDiagnoseResponse
	GetStatusCode() *int32
	SetBody(v *GenerateCdnDiagnoseResponseBody) *GenerateCdnDiagnoseResponse
	GetBody() *GenerateCdnDiagnoseResponseBody
}

type GenerateCdnDiagnoseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCdnDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCdnDiagnoseResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateCdnDiagnoseResponse) GoString() string {
	return s.String()
}

func (s *GenerateCdnDiagnoseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateCdnDiagnoseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateCdnDiagnoseResponse) GetBody() *GenerateCdnDiagnoseResponseBody {
	return s.Body
}

func (s *GenerateCdnDiagnoseResponse) SetHeaders(v map[string]*string) *GenerateCdnDiagnoseResponse {
	s.Headers = v
	return s
}

func (s *GenerateCdnDiagnoseResponse) SetStatusCode(v int32) *GenerateCdnDiagnoseResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCdnDiagnoseResponse) SetBody(v *GenerateCdnDiagnoseResponseBody) *GenerateCdnDiagnoseResponse {
	s.Body = v
	return s
}

func (s *GenerateCdnDiagnoseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
