// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceSafetyLockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceSafetyLockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceSafetyLockResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceSafetyLockResponseBody) *UpdateServiceSafetyLockResponse
	GetBody() *UpdateServiceSafetyLockResponseBody
}

type UpdateServiceSafetyLockResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceSafetyLockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceSafetyLockResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceSafetyLockResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceSafetyLockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceSafetyLockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceSafetyLockResponse) GetBody() *UpdateServiceSafetyLockResponseBody {
	return s.Body
}

func (s *UpdateServiceSafetyLockResponse) SetHeaders(v map[string]*string) *UpdateServiceSafetyLockResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceSafetyLockResponse) SetStatusCode(v int32) *UpdateServiceSafetyLockResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceSafetyLockResponse) SetBody(v *UpdateServiceSafetyLockResponseBody) *UpdateServiceSafetyLockResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceSafetyLockResponse) Validate() error {
	return dara.Validate(s)
}
