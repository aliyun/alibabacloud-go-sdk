// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCriteriaStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCriteriaStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCriteriaStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ListCriteriaStrategyResponseBody) *ListCriteriaStrategyResponse
	GetBody() *ListCriteriaStrategyResponseBody
}

type ListCriteriaStrategyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCriteriaStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCriteriaStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCriteriaStrategyResponse) GoString() string {
	return s.String()
}

func (s *ListCriteriaStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCriteriaStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCriteriaStrategyResponse) GetBody() *ListCriteriaStrategyResponseBody {
	return s.Body
}

func (s *ListCriteriaStrategyResponse) SetHeaders(v map[string]*string) *ListCriteriaStrategyResponse {
	s.Headers = v
	return s
}

func (s *ListCriteriaStrategyResponse) SetStatusCode(v int32) *ListCriteriaStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCriteriaStrategyResponse) SetBody(v *ListCriteriaStrategyResponseBody) *ListCriteriaStrategyResponse {
	s.Body = v
	return s
}

func (s *ListCriteriaStrategyResponse) Validate() error {
	return dara.Validate(s)
}
