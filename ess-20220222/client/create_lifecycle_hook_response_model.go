// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLifecycleHookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLifecycleHookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLifecycleHookResponse
	GetStatusCode() *int32
	SetBody(v *CreateLifecycleHookResponseBody) *CreateLifecycleHookResponse
	GetBody() *CreateLifecycleHookResponseBody
}

type CreateLifecycleHookResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLifecycleHookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLifecycleHookResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLifecycleHookResponse) GoString() string {
	return s.String()
}

func (s *CreateLifecycleHookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLifecycleHookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLifecycleHookResponse) GetBody() *CreateLifecycleHookResponseBody {
	return s.Body
}

func (s *CreateLifecycleHookResponse) SetHeaders(v map[string]*string) *CreateLifecycleHookResponse {
	s.Headers = v
	return s
}

func (s *CreateLifecycleHookResponse) SetStatusCode(v int32) *CreateLifecycleHookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLifecycleHookResponse) SetBody(v *CreateLifecycleHookResponseBody) *CreateLifecycleHookResponse {
	s.Body = v
	return s
}

func (s *CreateLifecycleHookResponse) Validate() error {
	return dara.Validate(s)
}
