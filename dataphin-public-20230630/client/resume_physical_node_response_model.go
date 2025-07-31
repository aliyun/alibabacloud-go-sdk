// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumePhysicalNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumePhysicalNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumePhysicalNodeResponse
	GetStatusCode() *int32
	SetBody(v *ResumePhysicalNodeResponseBody) *ResumePhysicalNodeResponse
	GetBody() *ResumePhysicalNodeResponseBody
}

type ResumePhysicalNodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumePhysicalNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumePhysicalNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumePhysicalNodeResponse) GoString() string {
	return s.String()
}

func (s *ResumePhysicalNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumePhysicalNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumePhysicalNodeResponse) GetBody() *ResumePhysicalNodeResponseBody {
	return s.Body
}

func (s *ResumePhysicalNodeResponse) SetHeaders(v map[string]*string) *ResumePhysicalNodeResponse {
	s.Headers = v
	return s
}

func (s *ResumePhysicalNodeResponse) SetStatusCode(v int32) *ResumePhysicalNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumePhysicalNodeResponse) SetBody(v *ResumePhysicalNodeResponseBody) *ResumePhysicalNodeResponse {
	s.Body = v
	return s
}

func (s *ResumePhysicalNodeResponse) Validate() error {
	return dara.Validate(s)
}
