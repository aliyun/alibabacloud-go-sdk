// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTitleDiagnoseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTitleDiagnoseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTitleDiagnoseResponse
	GetStatusCode() *int32
	SetBody(v *GetTitleDiagnoseResponseBody) *GetTitleDiagnoseResponse
	GetBody() *GetTitleDiagnoseResponseBody
}

type GetTitleDiagnoseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTitleDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTitleDiagnoseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTitleDiagnoseResponse) GoString() string {
	return s.String()
}

func (s *GetTitleDiagnoseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTitleDiagnoseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTitleDiagnoseResponse) GetBody() *GetTitleDiagnoseResponseBody {
	return s.Body
}

func (s *GetTitleDiagnoseResponse) SetHeaders(v map[string]*string) *GetTitleDiagnoseResponse {
	s.Headers = v
	return s
}

func (s *GetTitleDiagnoseResponse) SetStatusCode(v int32) *GetTitleDiagnoseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTitleDiagnoseResponse) SetBody(v *GetTitleDiagnoseResponseBody) *GetTitleDiagnoseResponse {
	s.Body = v
	return s
}

func (s *GetTitleDiagnoseResponse) Validate() error {
	return dara.Validate(s)
}
