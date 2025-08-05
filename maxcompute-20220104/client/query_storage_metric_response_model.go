// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStorageMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryStorageMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryStorageMetricResponse
	GetStatusCode() *int32
	SetBody(v *QueryStorageMetricResponseBody) *QueryStorageMetricResponse
	GetBody() *QueryStorageMetricResponseBody
}

type QueryStorageMetricResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryStorageMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryStorageMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryStorageMetricResponse) GoString() string {
	return s.String()
}

func (s *QueryStorageMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryStorageMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryStorageMetricResponse) GetBody() *QueryStorageMetricResponseBody {
	return s.Body
}

func (s *QueryStorageMetricResponse) SetHeaders(v map[string]*string) *QueryStorageMetricResponse {
	s.Headers = v
	return s
}

func (s *QueryStorageMetricResponse) SetStatusCode(v int32) *QueryStorageMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStorageMetricResponse) SetBody(v *QueryStorageMetricResponseBody) *QueryStorageMetricResponse {
	s.Body = v
	return s
}

func (s *QueryStorageMetricResponse) Validate() error {
	return dara.Validate(s)
}
