// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTuringTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitTuringTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitTuringTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitTuringTaskResponseBody) *SubmitTuringTaskResponse
	GetBody() *SubmitTuringTaskResponseBody
}

type SubmitTuringTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTuringTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTuringTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitTuringTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitTuringTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitTuringTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitTuringTaskResponse) GetBody() *SubmitTuringTaskResponseBody {
	return s.Body
}

func (s *SubmitTuringTaskResponse) SetHeaders(v map[string]*string) *SubmitTuringTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitTuringTaskResponse) SetStatusCode(v int32) *SubmitTuringTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTuringTaskResponse) SetBody(v *SubmitTuringTaskResponseBody) *SubmitTuringTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitTuringTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
