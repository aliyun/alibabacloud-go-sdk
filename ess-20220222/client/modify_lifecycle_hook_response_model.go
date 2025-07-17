// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLifecycleHookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLifecycleHookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLifecycleHookResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLifecycleHookResponseBody) *ModifyLifecycleHookResponse
	GetBody() *ModifyLifecycleHookResponseBody
}

type ModifyLifecycleHookResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLifecycleHookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLifecycleHookResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLifecycleHookResponse) GoString() string {
	return s.String()
}

func (s *ModifyLifecycleHookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLifecycleHookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLifecycleHookResponse) GetBody() *ModifyLifecycleHookResponseBody {
	return s.Body
}

func (s *ModifyLifecycleHookResponse) SetHeaders(v map[string]*string) *ModifyLifecycleHookResponse {
	s.Headers = v
	return s
}

func (s *ModifyLifecycleHookResponse) SetStatusCode(v int32) *ModifyLifecycleHookResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLifecycleHookResponse) SetBody(v *ModifyLifecycleHookResponseBody) *ModifyLifecycleHookResponse {
	s.Body = v
	return s
}

func (s *ModifyLifecycleHookResponse) Validate() error {
	return dara.Validate(s)
}
