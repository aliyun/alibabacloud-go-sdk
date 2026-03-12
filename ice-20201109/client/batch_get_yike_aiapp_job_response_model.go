// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetYikeAIAppJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetYikeAIAppJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetYikeAIAppJobResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetYikeAIAppJobResponseBody) *BatchGetYikeAIAppJobResponse
	GetBody() *BatchGetYikeAIAppJobResponseBody
}

type BatchGetYikeAIAppJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetYikeAIAppJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetYikeAIAppJobResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAIAppJobResponse) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAIAppJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetYikeAIAppJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetYikeAIAppJobResponse) GetBody() *BatchGetYikeAIAppJobResponseBody {
	return s.Body
}

func (s *BatchGetYikeAIAppJobResponse) SetHeaders(v map[string]*string) *BatchGetYikeAIAppJobResponse {
	s.Headers = v
	return s
}

func (s *BatchGetYikeAIAppJobResponse) SetStatusCode(v int32) *BatchGetYikeAIAppJobResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetYikeAIAppJobResponse) SetBody(v *BatchGetYikeAIAppJobResponseBody) *BatchGetYikeAIAppJobResponse {
	s.Body = v
	return s
}

func (s *BatchGetYikeAIAppJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
