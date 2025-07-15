// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKafkaClientIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKafkaClientIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKafkaClientIpResponse
	GetStatusCode() *int32
	SetBody(v *GetKafkaClientIpResponseBody) *GetKafkaClientIpResponse
	GetBody() *GetKafkaClientIpResponseBody
}

type GetKafkaClientIpResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKafkaClientIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKafkaClientIpResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKafkaClientIpResponse) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKafkaClientIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKafkaClientIpResponse) GetBody() *GetKafkaClientIpResponseBody {
	return s.Body
}

func (s *GetKafkaClientIpResponse) SetHeaders(v map[string]*string) *GetKafkaClientIpResponse {
	s.Headers = v
	return s
}

func (s *GetKafkaClientIpResponse) SetStatusCode(v int32) *GetKafkaClientIpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKafkaClientIpResponse) SetBody(v *GetKafkaClientIpResponseBody) *GetKafkaClientIpResponse {
	s.Body = v
	return s
}

func (s *GetKafkaClientIpResponse) Validate() error {
	return dara.Validate(s)
}
