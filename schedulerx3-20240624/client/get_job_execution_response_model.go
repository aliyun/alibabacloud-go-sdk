// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobExecutionResponse
	GetStatusCode() *int32
	SetBody(v *GetJobExecutionResponseBody) *GetJobExecutionResponse
	GetBody() *GetJobExecutionResponseBody
}

type GetJobExecutionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionResponse) GoString() string {
	return s.String()
}

func (s *GetJobExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobExecutionResponse) GetBody() *GetJobExecutionResponseBody {
	return s.Body
}

func (s *GetJobExecutionResponse) SetHeaders(v map[string]*string) *GetJobExecutionResponse {
	s.Headers = v
	return s
}

func (s *GetJobExecutionResponse) SetStatusCode(v int32) *GetJobExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobExecutionResponse) SetBody(v *GetJobExecutionResponseBody) *GetJobExecutionResponse {
	s.Body = v
	return s
}

func (s *GetJobExecutionResponse) Validate() error {
	return dara.Validate(s)
}
