// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLifecycleHookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLifecycleHookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLifecycleHookResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLifecycleHookResponseBody) *DeleteLifecycleHookResponse
	GetBody() *DeleteLifecycleHookResponseBody
}

type DeleteLifecycleHookResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLifecycleHookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLifecycleHookResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLifecycleHookResponse) GoString() string {
	return s.String()
}

func (s *DeleteLifecycleHookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLifecycleHookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLifecycleHookResponse) GetBody() *DeleteLifecycleHookResponseBody {
	return s.Body
}

func (s *DeleteLifecycleHookResponse) SetHeaders(v map[string]*string) *DeleteLifecycleHookResponse {
	s.Headers = v
	return s
}

func (s *DeleteLifecycleHookResponse) SetStatusCode(v int32) *DeleteLifecycleHookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLifecycleHookResponse) SetBody(v *DeleteLifecycleHookResponseBody) *DeleteLifecycleHookResponse {
	s.Body = v
	return s
}

func (s *DeleteLifecycleHookResponse) Validate() error {
	return dara.Validate(s)
}
