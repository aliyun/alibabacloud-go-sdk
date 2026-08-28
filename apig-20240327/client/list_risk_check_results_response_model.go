// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRiskCheckResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRiskCheckResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRiskCheckResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListRiskCheckResultsResponseBody) *ListRiskCheckResultsResponse
	GetBody() *ListRiskCheckResultsResponseBody
}

type ListRiskCheckResultsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRiskCheckResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRiskCheckResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRiskCheckResultsResponse) GoString() string {
	return s.String()
}

func (s *ListRiskCheckResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRiskCheckResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRiskCheckResultsResponse) GetBody() *ListRiskCheckResultsResponseBody {
	return s.Body
}

func (s *ListRiskCheckResultsResponse) SetHeaders(v map[string]*string) *ListRiskCheckResultsResponse {
	s.Headers = v
	return s
}

func (s *ListRiskCheckResultsResponse) SetStatusCode(v int32) *ListRiskCheckResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRiskCheckResultsResponse) SetBody(v *ListRiskCheckResultsResponseBody) *ListRiskCheckResultsResponse {
	s.Body = v
	return s
}

func (s *ListRiskCheckResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
