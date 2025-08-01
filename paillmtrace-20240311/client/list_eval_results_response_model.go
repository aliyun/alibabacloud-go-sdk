// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvalResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEvalResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEvalResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListEvalResultsResponseBody) *ListEvalResultsResponse
	GetBody() *ListEvalResultsResponseBody
}

type ListEvalResultsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvalResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvalResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEvalResultsResponse) GoString() string {
	return s.String()
}

func (s *ListEvalResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEvalResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEvalResultsResponse) GetBody() *ListEvalResultsResponseBody {
	return s.Body
}

func (s *ListEvalResultsResponse) SetHeaders(v map[string]*string) *ListEvalResultsResponse {
	s.Headers = v
	return s
}

func (s *ListEvalResultsResponse) SetStatusCode(v int32) *ListEvalResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvalResultsResponse) SetBody(v *ListEvalResultsResponseBody) *ListEvalResultsResponse {
	s.Body = v
	return s
}

func (s *ListEvalResultsResponse) Validate() error {
	return dara.Validate(s)
}
