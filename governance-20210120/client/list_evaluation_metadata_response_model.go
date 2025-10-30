// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEvaluationMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEvaluationMetadataResponse
	GetStatusCode() *int32
	SetBody(v *ListEvaluationMetadataResponseBody) *ListEvaluationMetadataResponse
	GetBody() *ListEvaluationMetadataResponseBody
}

type ListEvaluationMetadataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvaluationMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvaluationMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetadataResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEvaluationMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEvaluationMetadataResponse) GetBody() *ListEvaluationMetadataResponseBody {
	return s.Body
}

func (s *ListEvaluationMetadataResponse) SetHeaders(v map[string]*string) *ListEvaluationMetadataResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluationMetadataResponse) SetStatusCode(v int32) *ListEvaluationMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluationMetadataResponse) SetBody(v *ListEvaluationMetadataResponseBody) *ListEvaluationMetadataResponse {
	s.Body = v
	return s
}

func (s *ListEvaluationMetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
