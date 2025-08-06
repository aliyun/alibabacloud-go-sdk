// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSolutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSolutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSolutionResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSolutionResponseBody) *SubmitSolutionResponse
	GetBody() *SubmitSolutionResponseBody
}

type SubmitSolutionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSolutionResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSolutionResponse) GoString() string {
	return s.String()
}

func (s *SubmitSolutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSolutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSolutionResponse) GetBody() *SubmitSolutionResponseBody {
	return s.Body
}

func (s *SubmitSolutionResponse) SetHeaders(v map[string]*string) *SubmitSolutionResponse {
	s.Headers = v
	return s
}

func (s *SubmitSolutionResponse) SetStatusCode(v int32) *SubmitSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSolutionResponse) SetBody(v *SubmitSolutionResponseBody) *SubmitSolutionResponse {
	s.Body = v
	return s
}

func (s *SubmitSolutionResponse) Validate() error {
	return dara.Validate(s)
}
