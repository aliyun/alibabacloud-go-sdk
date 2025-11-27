// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendBatchCardSmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendBatchCardSmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendBatchCardSmsResponse
	GetStatusCode() *int32
	SetBody(v *SendBatchCardSmsResponseBody) *SendBatchCardSmsResponse
	GetBody() *SendBatchCardSmsResponseBody
}

type SendBatchCardSmsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendBatchCardSmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendBatchCardSmsResponse) String() string {
	return dara.Prettify(s)
}

func (s SendBatchCardSmsResponse) GoString() string {
	return s.String()
}

func (s *SendBatchCardSmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendBatchCardSmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendBatchCardSmsResponse) GetBody() *SendBatchCardSmsResponseBody {
	return s.Body
}

func (s *SendBatchCardSmsResponse) SetHeaders(v map[string]*string) *SendBatchCardSmsResponse {
	s.Headers = v
	return s
}

func (s *SendBatchCardSmsResponse) SetStatusCode(v int32) *SendBatchCardSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendBatchCardSmsResponse) SetBody(v *SendBatchCardSmsResponseBody) *SendBatchCardSmsResponse {
	s.Body = v
	return s
}

func (s *SendBatchCardSmsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
