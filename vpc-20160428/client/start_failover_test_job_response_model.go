// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartFailoverTestJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartFailoverTestJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartFailoverTestJobResponse
	GetStatusCode() *int32
	SetBody(v *StartFailoverTestJobResponseBody) *StartFailoverTestJobResponse
	GetBody() *StartFailoverTestJobResponseBody
}

type StartFailoverTestJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartFailoverTestJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartFailoverTestJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartFailoverTestJobResponse) GoString() string {
	return s.String()
}

func (s *StartFailoverTestJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartFailoverTestJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartFailoverTestJobResponse) GetBody() *StartFailoverTestJobResponseBody {
	return s.Body
}

func (s *StartFailoverTestJobResponse) SetHeaders(v map[string]*string) *StartFailoverTestJobResponse {
	s.Headers = v
	return s
}

func (s *StartFailoverTestJobResponse) SetStatusCode(v int32) *StartFailoverTestJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartFailoverTestJobResponse) SetBody(v *StartFailoverTestJobResponseBody) *StartFailoverTestJobResponse {
	s.Body = v
	return s
}

func (s *StartFailoverTestJobResponse) Validate() error {
	return dara.Validate(s)
}
