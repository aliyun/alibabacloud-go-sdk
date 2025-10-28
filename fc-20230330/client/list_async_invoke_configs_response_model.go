// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncInvokeConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAsyncInvokeConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAsyncInvokeConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListAsyncInvokeConfigOutput) *ListAsyncInvokeConfigsResponse
	GetBody() *ListAsyncInvokeConfigOutput
}

type ListAsyncInvokeConfigsResponse struct {
	Headers    map[string]*string           `json:"headers" xml:"headers"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAsyncInvokeConfigOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAsyncInvokeConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncInvokeConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAsyncInvokeConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAsyncInvokeConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAsyncInvokeConfigsResponse) GetBody() *ListAsyncInvokeConfigOutput {
	return s.Body
}

func (s *ListAsyncInvokeConfigsResponse) SetHeaders(v map[string]*string) *ListAsyncInvokeConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAsyncInvokeConfigsResponse) SetStatusCode(v int32) *ListAsyncInvokeConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsyncInvokeConfigsResponse) SetBody(v *ListAsyncInvokeConfigOutput) *ListAsyncInvokeConfigsResponse {
	s.Body = v
	return s
}

func (s *ListAsyncInvokeConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
