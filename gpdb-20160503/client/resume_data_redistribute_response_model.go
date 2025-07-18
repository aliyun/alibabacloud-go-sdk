// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeDataRedistributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeDataRedistributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeDataRedistributeResponse
	GetStatusCode() *int32
	SetBody(v *ResumeDataRedistributeResponseBody) *ResumeDataRedistributeResponse
	GetBody() *ResumeDataRedistributeResponseBody
}

type ResumeDataRedistributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeDataRedistributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeDataRedistributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeDataRedistributeResponse) GoString() string {
	return s.String()
}

func (s *ResumeDataRedistributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeDataRedistributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeDataRedistributeResponse) GetBody() *ResumeDataRedistributeResponseBody {
	return s.Body
}

func (s *ResumeDataRedistributeResponse) SetHeaders(v map[string]*string) *ResumeDataRedistributeResponse {
	s.Headers = v
	return s
}

func (s *ResumeDataRedistributeResponse) SetStatusCode(v int32) *ResumeDataRedistributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeDataRedistributeResponse) SetBody(v *ResumeDataRedistributeResponseBody) *ResumeDataRedistributeResponse {
	s.Body = v
	return s
}

func (s *ResumeDataRedistributeResponse) Validate() error {
	return dara.Validate(s)
}
