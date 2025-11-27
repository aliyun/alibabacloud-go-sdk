// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLogisticsSmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendLogisticsSmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendLogisticsSmsResponse
	GetStatusCode() *int32
	SetBody(v *SendLogisticsSmsResponseBody) *SendLogisticsSmsResponse
	GetBody() *SendLogisticsSmsResponseBody
}

type SendLogisticsSmsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendLogisticsSmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendLogisticsSmsResponse) String() string {
	return dara.Prettify(s)
}

func (s SendLogisticsSmsResponse) GoString() string {
	return s.String()
}

func (s *SendLogisticsSmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendLogisticsSmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendLogisticsSmsResponse) GetBody() *SendLogisticsSmsResponseBody {
	return s.Body
}

func (s *SendLogisticsSmsResponse) SetHeaders(v map[string]*string) *SendLogisticsSmsResponse {
	s.Headers = v
	return s
}

func (s *SendLogisticsSmsResponse) SetStatusCode(v int32) *SendLogisticsSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendLogisticsSmsResponse) SetBody(v *SendLogisticsSmsResponseBody) *SendLogisticsSmsResponse {
	s.Body = v
	return s
}

func (s *SendLogisticsSmsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
