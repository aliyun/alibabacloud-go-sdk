// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunGenerateQuestionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunGenerateQuestionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunGenerateQuestionsResponse
	GetStatusCode() *int32
	SetBody(v *RunGenerateQuestionsResponseBody) *RunGenerateQuestionsResponse
	GetBody() *RunGenerateQuestionsResponseBody
}

type RunGenerateQuestionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunGenerateQuestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunGenerateQuestionsResponse) String() string {
	return dara.Prettify(s)
}

func (s RunGenerateQuestionsResponse) GoString() string {
	return s.String()
}

func (s *RunGenerateQuestionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunGenerateQuestionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunGenerateQuestionsResponse) GetBody() *RunGenerateQuestionsResponseBody {
	return s.Body
}

func (s *RunGenerateQuestionsResponse) SetHeaders(v map[string]*string) *RunGenerateQuestionsResponse {
	s.Headers = v
	return s
}

func (s *RunGenerateQuestionsResponse) SetStatusCode(v int32) *RunGenerateQuestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *RunGenerateQuestionsResponse) SetBody(v *RunGenerateQuestionsResponseBody) *RunGenerateQuestionsResponse {
	s.Body = v
	return s
}

func (s *RunGenerateQuestionsResponse) Validate() error {
	return dara.Validate(s)
}
