// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobExecutionThreadDumpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobExecutionThreadDumpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobExecutionThreadDumpResponse
	GetStatusCode() *int32
	SetBody(v *GetJobExecutionThreadDumpResponseBody) *GetJobExecutionThreadDumpResponse
	GetBody() *GetJobExecutionThreadDumpResponseBody
}

type GetJobExecutionThreadDumpResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobExecutionThreadDumpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobExecutionThreadDumpResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionThreadDumpResponse) GoString() string {
	return s.String()
}

func (s *GetJobExecutionThreadDumpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobExecutionThreadDumpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobExecutionThreadDumpResponse) GetBody() *GetJobExecutionThreadDumpResponseBody {
	return s.Body
}

func (s *GetJobExecutionThreadDumpResponse) SetHeaders(v map[string]*string) *GetJobExecutionThreadDumpResponse {
	s.Headers = v
	return s
}

func (s *GetJobExecutionThreadDumpResponse) SetStatusCode(v int32) *GetJobExecutionThreadDumpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobExecutionThreadDumpResponse) SetBody(v *GetJobExecutionThreadDumpResponseBody) *GetJobExecutionThreadDumpResponse {
	s.Body = v
	return s
}

func (s *GetJobExecutionThreadDumpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
