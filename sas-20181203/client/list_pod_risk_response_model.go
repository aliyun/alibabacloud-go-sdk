// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPodRiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPodRiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPodRiskResponse
	GetStatusCode() *int32
	SetBody(v *ListPodRiskResponseBody) *ListPodRiskResponse
	GetBody() *ListPodRiskResponseBody
}

type ListPodRiskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPodRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPodRiskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPodRiskResponse) GoString() string {
	return s.String()
}

func (s *ListPodRiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPodRiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPodRiskResponse) GetBody() *ListPodRiskResponseBody {
	return s.Body
}

func (s *ListPodRiskResponse) SetHeaders(v map[string]*string) *ListPodRiskResponse {
	s.Headers = v
	return s
}

func (s *ListPodRiskResponse) SetStatusCode(v int32) *ListPodRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPodRiskResponse) SetBody(v *ListPodRiskResponseBody) *ListPodRiskResponse {
	s.Body = v
	return s
}

func (s *ListPodRiskResponse) Validate() error {
	return dara.Validate(s)
}
