// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPreprocessJobsConsoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitPreprocessJobsConsoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitPreprocessJobsConsoleResponse
	GetStatusCode() *int32
	SetBody(v *SubmitPreprocessJobsConsoleResponseBody) *SubmitPreprocessJobsConsoleResponse
	GetBody() *SubmitPreprocessJobsConsoleResponseBody
}

type SubmitPreprocessJobsConsoleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitPreprocessJobsConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitPreprocessJobsConsoleResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitPreprocessJobsConsoleResponse) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsConsoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitPreprocessJobsConsoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitPreprocessJobsConsoleResponse) GetBody() *SubmitPreprocessJobsConsoleResponseBody {
	return s.Body
}

func (s *SubmitPreprocessJobsConsoleResponse) SetHeaders(v map[string]*string) *SubmitPreprocessJobsConsoleResponse {
	s.Headers = v
	return s
}

func (s *SubmitPreprocessJobsConsoleResponse) SetStatusCode(v int32) *SubmitPreprocessJobsConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitPreprocessJobsConsoleResponse) SetBody(v *SubmitPreprocessJobsConsoleResponseBody) *SubmitPreprocessJobsConsoleResponse {
	s.Body = v
	return s
}

func (s *SubmitPreprocessJobsConsoleResponse) Validate() error {
	return dara.Validate(s)
}
