// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDtsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetDtsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetDtsJobResponse
	GetStatusCode() *int32
	SetBody(v *ResetDtsJobResponseBody) *ResetDtsJobResponse
	GetBody() *ResetDtsJobResponseBody
}

type ResetDtsJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetDtsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetDtsJobResponse) GoString() string {
	return s.String()
}

func (s *ResetDtsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetDtsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetDtsJobResponse) GetBody() *ResetDtsJobResponseBody {
	return s.Body
}

func (s *ResetDtsJobResponse) SetHeaders(v map[string]*string) *ResetDtsJobResponse {
	s.Headers = v
	return s
}

func (s *ResetDtsJobResponse) SetStatusCode(v int32) *ResetDtsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDtsJobResponse) SetBody(v *ResetDtsJobResponseBody) *ResetDtsJobResponse {
	s.Body = v
	return s
}

func (s *ResetDtsJobResponse) Validate() error {
	return dara.Validate(s)
}
