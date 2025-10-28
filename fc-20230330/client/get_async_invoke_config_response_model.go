// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncInvokeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsyncInvokeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsyncInvokeConfigResponse
	GetStatusCode() *int32
	SetBody(v *AsyncConfig) *GetAsyncInvokeConfigResponse
	GetBody() *AsyncConfig
}

type GetAsyncInvokeConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncConfig       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncInvokeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncInvokeConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncInvokeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsyncInvokeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsyncInvokeConfigResponse) GetBody() *AsyncConfig {
	return s.Body
}

func (s *GetAsyncInvokeConfigResponse) SetHeaders(v map[string]*string) *GetAsyncInvokeConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncInvokeConfigResponse) SetStatusCode(v int32) *GetAsyncInvokeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncInvokeConfigResponse) SetBody(v *AsyncConfig) *GetAsyncInvokeConfigResponse {
	s.Body = v
	return s
}

func (s *GetAsyncInvokeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
