// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultimodalAsyncModerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MultimodalAsyncModerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MultimodalAsyncModerationResponse
	GetStatusCode() *int32
	SetBody(v *MultimodalAsyncModerationResponseBody) *MultimodalAsyncModerationResponse
	GetBody() *MultimodalAsyncModerationResponseBody
}

type MultimodalAsyncModerationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MultimodalAsyncModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultimodalAsyncModerationResponse) String() string {
	return dara.Prettify(s)
}

func (s MultimodalAsyncModerationResponse) GoString() string {
	return s.String()
}

func (s *MultimodalAsyncModerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MultimodalAsyncModerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MultimodalAsyncModerationResponse) GetBody() *MultimodalAsyncModerationResponseBody {
	return s.Body
}

func (s *MultimodalAsyncModerationResponse) SetHeaders(v map[string]*string) *MultimodalAsyncModerationResponse {
	s.Headers = v
	return s
}

func (s *MultimodalAsyncModerationResponse) SetStatusCode(v int32) *MultimodalAsyncModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *MultimodalAsyncModerationResponse) SetBody(v *MultimodalAsyncModerationResponseBody) *MultimodalAsyncModerationResponse {
	s.Body = v
	return s
}

func (s *MultimodalAsyncModerationResponse) Validate() error {
	return dara.Validate(s)
}
