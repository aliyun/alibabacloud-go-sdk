// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoTagJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIVideoTagJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIVideoTagJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIVideoTagJobResponseBody) *SubmitAIVideoTagJobResponse
	GetBody() *SubmitAIVideoTagJobResponseBody
}

type SubmitAIVideoTagJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIVideoTagJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIVideoTagJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoTagJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoTagJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIVideoTagJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIVideoTagJobResponse) GetBody() *SubmitAIVideoTagJobResponseBody {
	return s.Body
}

func (s *SubmitAIVideoTagJobResponse) SetHeaders(v map[string]*string) *SubmitAIVideoTagJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIVideoTagJobResponse) SetStatusCode(v int32) *SubmitAIVideoTagJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIVideoTagJobResponse) SetBody(v *SubmitAIVideoTagJobResponseBody) *SubmitAIVideoTagJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIVideoTagJobResponse) Validate() error {
	return dara.Validate(s)
}
