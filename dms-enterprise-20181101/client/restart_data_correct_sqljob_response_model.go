// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDataCorrectSQLJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartDataCorrectSQLJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartDataCorrectSQLJobResponse
	GetStatusCode() *int32
	SetBody(v *RestartDataCorrectSQLJobResponseBody) *RestartDataCorrectSQLJobResponse
	GetBody() *RestartDataCorrectSQLJobResponseBody
}

type RestartDataCorrectSQLJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDataCorrectSQLJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartDataCorrectSQLJobResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartDataCorrectSQLJobResponse) GoString() string {
	return s.String()
}

func (s *RestartDataCorrectSQLJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartDataCorrectSQLJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartDataCorrectSQLJobResponse) GetBody() *RestartDataCorrectSQLJobResponseBody {
	return s.Body
}

func (s *RestartDataCorrectSQLJobResponse) SetHeaders(v map[string]*string) *RestartDataCorrectSQLJobResponse {
	s.Headers = v
	return s
}

func (s *RestartDataCorrectSQLJobResponse) SetStatusCode(v int32) *RestartDataCorrectSQLJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDataCorrectSQLJobResponse) SetBody(v *RestartDataCorrectSQLJobResponseBody) *RestartDataCorrectSQLJobResponse {
	s.Body = v
	return s
}

func (s *RestartDataCorrectSQLJobResponse) Validate() error {
	return dara.Validate(s)
}
