// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLLMAnswerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLLMAnswerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLLMAnswerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLLMAnswerResponseBody) *DescribeLLMAnswerResponse
	GetBody() *DescribeLLMAnswerResponseBody
}

type DescribeLLMAnswerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLLMAnswerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLLMAnswerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLLMAnswerResponse) GoString() string {
	return s.String()
}

func (s *DescribeLLMAnswerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLLMAnswerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLLMAnswerResponse) GetBody() *DescribeLLMAnswerResponseBody {
	return s.Body
}

func (s *DescribeLLMAnswerResponse) SetHeaders(v map[string]*string) *DescribeLLMAnswerResponse {
	s.Headers = v
	return s
}

func (s *DescribeLLMAnswerResponse) SetStatusCode(v int32) *DescribeLLMAnswerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLLMAnswerResponse) SetBody(v *DescribeLLMAnswerResponseBody) *DescribeLLMAnswerResponse {
	s.Body = v
	return s
}

func (s *DescribeLLMAnswerResponse) Validate() error {
	return dara.Validate(s)
}
