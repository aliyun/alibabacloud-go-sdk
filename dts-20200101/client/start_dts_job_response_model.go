// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDtsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDtsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDtsJobResponse
	GetStatusCode() *int32
	SetBody(v *StartDtsJobResponseBody) *StartDtsJobResponse
	GetBody() *StartDtsJobResponseBody
}

type StartDtsJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDtsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDtsJobResponse) GoString() string {
	return s.String()
}

func (s *StartDtsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDtsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDtsJobResponse) GetBody() *StartDtsJobResponseBody {
	return s.Body
}

func (s *StartDtsJobResponse) SetHeaders(v map[string]*string) *StartDtsJobResponse {
	s.Headers = v
	return s
}

func (s *StartDtsJobResponse) SetStatusCode(v int32) *StartDtsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDtsJobResponse) SetBody(v *StartDtsJobResponseBody) *StartDtsJobResponse {
	s.Body = v
	return s
}

func (s *StartDtsJobResponse) Validate() error {
	return dara.Validate(s)
}
