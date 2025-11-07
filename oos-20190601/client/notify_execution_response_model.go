// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *NotifyExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *NotifyExecutionResponse
	GetStatusCode() *int32
	SetBody(v *NotifyExecutionResponseBody) *NotifyExecutionResponse
	GetBody() *NotifyExecutionResponseBody
}

type NotifyExecutionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NotifyExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NotifyExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s NotifyExecutionResponse) GoString() string {
	return s.String()
}

func (s *NotifyExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *NotifyExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *NotifyExecutionResponse) GetBody() *NotifyExecutionResponseBody {
	return s.Body
}

func (s *NotifyExecutionResponse) SetHeaders(v map[string]*string) *NotifyExecutionResponse {
	s.Headers = v
	return s
}

func (s *NotifyExecutionResponse) SetStatusCode(v int32) *NotifyExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *NotifyExecutionResponse) SetBody(v *NotifyExecutionResponseBody) *NotifyExecutionResponse {
	s.Body = v
	return s
}

func (s *NotifyExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
