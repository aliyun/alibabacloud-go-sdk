// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAsyncInvokeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAsyncInvokeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAsyncInvokeConfigResponse
	GetStatusCode() *int32
}

type DeleteAsyncInvokeConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAsyncInvokeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAsyncInvokeConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAsyncInvokeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAsyncInvokeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAsyncInvokeConfigResponse) SetHeaders(v map[string]*string) *DeleteAsyncInvokeConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAsyncInvokeConfigResponse) SetStatusCode(v int32) *DeleteAsyncInvokeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAsyncInvokeConfigResponse) Validate() error {
	return dara.Validate(s)
}
