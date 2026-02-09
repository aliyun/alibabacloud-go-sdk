// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTuringTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckTuringTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckTuringTaskResponse
	GetStatusCode() *int32
	SetBody(v *CheckTuringTaskResponseBody) *CheckTuringTaskResponse
	GetBody() *CheckTuringTaskResponseBody
}

type CheckTuringTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckTuringTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckTuringTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckTuringTaskResponse) GoString() string {
	return s.String()
}

func (s *CheckTuringTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckTuringTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckTuringTaskResponse) GetBody() *CheckTuringTaskResponseBody {
	return s.Body
}

func (s *CheckTuringTaskResponse) SetHeaders(v map[string]*string) *CheckTuringTaskResponse {
	s.Headers = v
	return s
}

func (s *CheckTuringTaskResponse) SetStatusCode(v int32) *CheckTuringTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckTuringTaskResponse) SetBody(v *CheckTuringTaskResponseBody) *CheckTuringTaskResponse {
	s.Body = v
	return s
}

func (s *CheckTuringTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
