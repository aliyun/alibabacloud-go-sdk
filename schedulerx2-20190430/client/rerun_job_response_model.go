// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RerunJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RerunJobResponse
	GetStatusCode() *int32
	SetBody(v *RerunJobResponseBody) *RerunJobResponse
	GetBody() *RerunJobResponseBody
}

type RerunJobResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RerunJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RerunJobResponse) String() string {
	return dara.Prettify(s)
}

func (s RerunJobResponse) GoString() string {
	return s.String()
}

func (s *RerunJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RerunJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RerunJobResponse) GetBody() *RerunJobResponseBody {
	return s.Body
}

func (s *RerunJobResponse) SetHeaders(v map[string]*string) *RerunJobResponse {
	s.Headers = v
	return s
}

func (s *RerunJobResponse) SetStatusCode(v int32) *RerunJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RerunJobResponse) SetBody(v *RerunJobResponseBody) *RerunJobResponse {
	s.Body = v
	return s
}

func (s *RerunJobResponse) Validate() error {
	return dara.Validate(s)
}
