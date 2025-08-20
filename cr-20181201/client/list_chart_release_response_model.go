// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChartReleaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChartReleaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChartReleaseResponse
	GetStatusCode() *int32
	SetBody(v *ListChartReleaseResponseBody) *ListChartReleaseResponse
	GetBody() *ListChartReleaseResponseBody
}

type ListChartReleaseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChartReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChartReleaseResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChartReleaseResponse) GoString() string {
	return s.String()
}

func (s *ListChartReleaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChartReleaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChartReleaseResponse) GetBody() *ListChartReleaseResponseBody {
	return s.Body
}

func (s *ListChartReleaseResponse) SetHeaders(v map[string]*string) *ListChartReleaseResponse {
	s.Headers = v
	return s
}

func (s *ListChartReleaseResponse) SetStatusCode(v int32) *ListChartReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChartReleaseResponse) SetBody(v *ListChartReleaseResponseBody) *ListChartReleaseResponse {
	s.Body = v
	return s
}

func (s *ListChartReleaseResponse) Validate() error {
	return dara.Validate(s)
}
