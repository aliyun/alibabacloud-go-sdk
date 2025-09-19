// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageRiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageRiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageRiskResponse
	GetStatusCode() *int32
	SetBody(v *ListImageRiskResponseBody) *ListImageRiskResponse
	GetBody() *ListImageRiskResponseBody
}

type ListImageRiskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageRiskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageRiskResponse) GoString() string {
	return s.String()
}

func (s *ListImageRiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageRiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageRiskResponse) GetBody() *ListImageRiskResponseBody {
	return s.Body
}

func (s *ListImageRiskResponse) SetHeaders(v map[string]*string) *ListImageRiskResponse {
	s.Headers = v
	return s
}

func (s *ListImageRiskResponse) SetStatusCode(v int32) *ListImageRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageRiskResponse) SetBody(v *ListImageRiskResponseBody) *ListImageRiskResponse {
	s.Body = v
	return s
}

func (s *ListImageRiskResponse) Validate() error {
	return dara.Validate(s)
}
