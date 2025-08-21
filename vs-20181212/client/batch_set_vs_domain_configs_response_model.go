// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetVsDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSetVsDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSetVsDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *BatchSetVsDomainConfigsResponseBody) *BatchSetVsDomainConfigsResponse
	GetBody() *BatchSetVsDomainConfigsResponseBody
}

type BatchSetVsDomainConfigsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSetVsDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSetVsDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSetVsDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchSetVsDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSetVsDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSetVsDomainConfigsResponse) GetBody() *BatchSetVsDomainConfigsResponseBody {
	return s.Body
}

func (s *BatchSetVsDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchSetVsDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchSetVsDomainConfigsResponse) SetStatusCode(v int32) *BatchSetVsDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSetVsDomainConfigsResponse) SetBody(v *BatchSetVsDomainConfigsResponseBody) *BatchSetVsDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *BatchSetVsDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
