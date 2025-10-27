// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockUnbindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebLockUnbindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebLockUnbindResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebLockUnbindResponseBody) *ModifyWebLockUnbindResponse
	GetBody() *ModifyWebLockUnbindResponseBody
}

type ModifyWebLockUnbindResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebLockUnbindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebLockUnbindResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockUnbindResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockUnbindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebLockUnbindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebLockUnbindResponse) GetBody() *ModifyWebLockUnbindResponseBody {
	return s.Body
}

func (s *ModifyWebLockUnbindResponse) SetHeaders(v map[string]*string) *ModifyWebLockUnbindResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockUnbindResponse) SetStatusCode(v int32) *ModifyWebLockUnbindResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebLockUnbindResponse) SetBody(v *ModifyWebLockUnbindResponseBody) *ModifyWebLockUnbindResponse {
	s.Body = v
	return s
}

func (s *ModifyWebLockUnbindResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
