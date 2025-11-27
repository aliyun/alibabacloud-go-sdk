// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteVsDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteVsDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteVsDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteVsDomainConfigsResponseBody) *BatchDeleteVsDomainConfigsResponse
	GetBody() *BatchDeleteVsDomainConfigsResponseBody
}

type BatchDeleteVsDomainConfigsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteVsDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteVsDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteVsDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteVsDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteVsDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteVsDomainConfigsResponse) GetBody() *BatchDeleteVsDomainConfigsResponseBody {
	return s.Body
}

func (s *BatchDeleteVsDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchDeleteVsDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteVsDomainConfigsResponse) SetStatusCode(v int32) *BatchDeleteVsDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteVsDomainConfigsResponse) SetBody(v *BatchDeleteVsDomainConfigsResponseBody) *BatchDeleteVsDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteVsDomainConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
