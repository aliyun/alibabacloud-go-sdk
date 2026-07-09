// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluatorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEvaluatorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEvaluatorsResponse
	GetStatusCode() *int32
	SetBody(v *ListEvaluatorsResponseBody) *ListEvaluatorsResponse
	GetBody() *ListEvaluatorsResponseBody
}

type ListEvaluatorsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvaluatorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvaluatorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluatorsResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluatorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEvaluatorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEvaluatorsResponse) GetBody() *ListEvaluatorsResponseBody {
	return s.Body
}

func (s *ListEvaluatorsResponse) SetHeaders(v map[string]*string) *ListEvaluatorsResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluatorsResponse) SetStatusCode(v int32) *ListEvaluatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluatorsResponse) SetBody(v *ListEvaluatorsResponseBody) *ListEvaluatorsResponse {
	s.Body = v
	return s
}

func (s *ListEvaluatorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
