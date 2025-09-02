// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityResultsByEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQualityResultsByEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQualityResultsByEntityResponse
	GetStatusCode() *int32
	SetBody(v *ListQualityResultsByEntityResponseBody) *ListQualityResultsByEntityResponse
	GetBody() *ListQualityResultsByEntityResponseBody
}

type ListQualityResultsByEntityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQualityResultsByEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQualityResultsByEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQualityResultsByEntityResponse) GoString() string {
	return s.String()
}

func (s *ListQualityResultsByEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQualityResultsByEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQualityResultsByEntityResponse) GetBody() *ListQualityResultsByEntityResponseBody {
	return s.Body
}

func (s *ListQualityResultsByEntityResponse) SetHeaders(v map[string]*string) *ListQualityResultsByEntityResponse {
	s.Headers = v
	return s
}

func (s *ListQualityResultsByEntityResponse) SetStatusCode(v int32) *ListQualityResultsByEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQualityResultsByEntityResponse) SetBody(v *ListQualityResultsByEntityResponseBody) *ListQualityResultsByEntityResponse {
	s.Body = v
	return s
}

func (s *ListQualityResultsByEntityResponse) Validate() error {
	return dara.Validate(s)
}
