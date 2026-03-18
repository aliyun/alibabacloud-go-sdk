// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackExecutionResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStackExecutionResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStackExecutionResultResponse
	GetStatusCode() *int32
	SetBody(v *GetStackExecutionResultResponseBody) *GetStackExecutionResultResponse
	GetBody() *GetStackExecutionResultResponseBody
}

type GetStackExecutionResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStackExecutionResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStackExecutionResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStackExecutionResultResponse) GoString() string {
	return s.String()
}

func (s *GetStackExecutionResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStackExecutionResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStackExecutionResultResponse) GetBody() *GetStackExecutionResultResponseBody {
	return s.Body
}

func (s *GetStackExecutionResultResponse) SetHeaders(v map[string]*string) *GetStackExecutionResultResponse {
	s.Headers = v
	return s
}

func (s *GetStackExecutionResultResponse) SetStatusCode(v int32) *GetStackExecutionResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackExecutionResultResponse) SetBody(v *GetStackExecutionResultResponseBody) *GetStackExecutionResultResponse {
	s.Body = v
	return s
}

func (s *GetStackExecutionResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
