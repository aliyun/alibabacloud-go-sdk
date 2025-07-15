// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetLiveDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSetLiveDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSetLiveDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *BatchSetLiveDomainConfigsResponseBody) *BatchSetLiveDomainConfigsResponse
	GetBody() *BatchSetLiveDomainConfigsResponseBody
}

type BatchSetLiveDomainConfigsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSetLiveDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSetLiveDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSetLiveDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchSetLiveDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSetLiveDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSetLiveDomainConfigsResponse) GetBody() *BatchSetLiveDomainConfigsResponseBody {
	return s.Body
}

func (s *BatchSetLiveDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchSetLiveDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchSetLiveDomainConfigsResponse) SetStatusCode(v int32) *BatchSetLiveDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSetLiveDomainConfigsResponse) SetBody(v *BatchSetLiveDomainConfigsResponseBody) *BatchSetLiveDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *BatchSetLiveDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
