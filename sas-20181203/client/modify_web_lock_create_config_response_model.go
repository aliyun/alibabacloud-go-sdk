// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockCreateConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebLockCreateConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebLockCreateConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebLockCreateConfigResponseBody) *ModifyWebLockCreateConfigResponse
	GetBody() *ModifyWebLockCreateConfigResponseBody
}

type ModifyWebLockCreateConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebLockCreateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebLockCreateConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockCreateConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockCreateConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebLockCreateConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebLockCreateConfigResponse) GetBody() *ModifyWebLockCreateConfigResponseBody {
	return s.Body
}

func (s *ModifyWebLockCreateConfigResponse) SetHeaders(v map[string]*string) *ModifyWebLockCreateConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockCreateConfigResponse) SetStatusCode(v int32) *ModifyWebLockCreateConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebLockCreateConfigResponse) SetBody(v *ModifyWebLockCreateConfigResponseBody) *ModifyWebLockCreateConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyWebLockCreateConfigResponse) Validate() error {
	return dara.Validate(s)
}
