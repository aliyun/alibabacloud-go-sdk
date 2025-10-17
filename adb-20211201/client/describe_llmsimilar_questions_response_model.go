// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLLMSimilarQuestionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLLMSimilarQuestionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLLMSimilarQuestionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLLMSimilarQuestionsResponseBody) *DescribeLLMSimilarQuestionsResponse
	GetBody() *DescribeLLMSimilarQuestionsResponseBody
}

type DescribeLLMSimilarQuestionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLLMSimilarQuestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLLMSimilarQuestionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLLMSimilarQuestionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLLMSimilarQuestionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLLMSimilarQuestionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLLMSimilarQuestionsResponse) GetBody() *DescribeLLMSimilarQuestionsResponseBody {
	return s.Body
}

func (s *DescribeLLMSimilarQuestionsResponse) SetHeaders(v map[string]*string) *DescribeLLMSimilarQuestionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponse) SetStatusCode(v int32) *DescribeLLMSimilarQuestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponse) SetBody(v *DescribeLLMSimilarQuestionsResponseBody) *DescribeLLMSimilarQuestionsResponse {
	s.Body = v
	return s
}

func (s *DescribeLLMSimilarQuestionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
