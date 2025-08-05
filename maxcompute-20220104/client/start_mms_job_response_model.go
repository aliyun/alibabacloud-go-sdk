// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMmsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartMmsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartMmsJobResponse
	GetStatusCode() *int32
	SetBody(v *StartMmsJobResponseBody) *StartMmsJobResponse
	GetBody() *StartMmsJobResponseBody
}

type StartMmsJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartMmsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartMmsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartMmsJobResponse) GoString() string {
	return s.String()
}

func (s *StartMmsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartMmsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartMmsJobResponse) GetBody() *StartMmsJobResponseBody {
	return s.Body
}

func (s *StartMmsJobResponse) SetHeaders(v map[string]*string) *StartMmsJobResponse {
	s.Headers = v
	return s
}

func (s *StartMmsJobResponse) SetStatusCode(v int32) *StartMmsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartMmsJobResponse) SetBody(v *StartMmsJobResponseBody) *StartMmsJobResponse {
	s.Body = v
	return s
}

func (s *StartMmsJobResponse) Validate() error {
	return dara.Validate(s)
}
