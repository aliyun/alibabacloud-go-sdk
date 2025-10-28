// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAsyncInvokeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutAsyncInvokeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutAsyncInvokeConfigResponse
	GetStatusCode() *int32
	SetBody(v *AsyncConfig) *PutAsyncInvokeConfigResponse
	GetBody() *AsyncConfig
}

type PutAsyncInvokeConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncConfig       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutAsyncInvokeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PutAsyncInvokeConfigResponse) GoString() string {
	return s.String()
}

func (s *PutAsyncInvokeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutAsyncInvokeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutAsyncInvokeConfigResponse) GetBody() *AsyncConfig {
	return s.Body
}

func (s *PutAsyncInvokeConfigResponse) SetHeaders(v map[string]*string) *PutAsyncInvokeConfigResponse {
	s.Headers = v
	return s
}

func (s *PutAsyncInvokeConfigResponse) SetStatusCode(v int32) *PutAsyncInvokeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PutAsyncInvokeConfigResponse) SetBody(v *AsyncConfig) *PutAsyncInvokeConfigResponse {
	s.Body = v
	return s
}

func (s *PutAsyncInvokeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
