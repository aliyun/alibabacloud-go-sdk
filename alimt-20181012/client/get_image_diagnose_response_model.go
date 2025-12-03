// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDiagnoseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageDiagnoseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageDiagnoseResponse
	GetStatusCode() *int32
	SetBody(v *GetImageDiagnoseResponseBody) *GetImageDiagnoseResponse
	GetBody() *GetImageDiagnoseResponseBody
}

type GetImageDiagnoseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageDiagnoseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageDiagnoseResponse) GoString() string {
	return s.String()
}

func (s *GetImageDiagnoseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageDiagnoseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageDiagnoseResponse) GetBody() *GetImageDiagnoseResponseBody {
	return s.Body
}

func (s *GetImageDiagnoseResponse) SetHeaders(v map[string]*string) *GetImageDiagnoseResponse {
	s.Headers = v
	return s
}

func (s *GetImageDiagnoseResponse) SetStatusCode(v int32) *GetImageDiagnoseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageDiagnoseResponse) SetBody(v *GetImageDiagnoseResponseBody) *GetImageDiagnoseResponse {
	s.Body = v
	return s
}

func (s *GetImageDiagnoseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
