// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteLiveDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteLiveDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteLiveDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteLiveDomainConfigsResponseBody) *BatchDeleteLiveDomainConfigsResponse
	GetBody() *BatchDeleteLiveDomainConfigsResponseBody
}

type BatchDeleteLiveDomainConfigsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteLiveDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteLiveDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteLiveDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteLiveDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteLiveDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteLiveDomainConfigsResponse) GetBody() *BatchDeleteLiveDomainConfigsResponseBody {
	return s.Body
}

func (s *BatchDeleteLiveDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchDeleteLiveDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteLiveDomainConfigsResponse) SetStatusCode(v int32) *BatchDeleteLiveDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteLiveDomainConfigsResponse) SetBody(v *BatchDeleteLiveDomainConfigsResponseBody) *BatchDeleteLiveDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteLiveDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
