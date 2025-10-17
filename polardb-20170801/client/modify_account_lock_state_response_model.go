// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountLockStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccountLockStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccountLockStateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccountLockStateResponseBody) *ModifyAccountLockStateResponse
	GetBody() *ModifyAccountLockStateResponseBody
}

type ModifyAccountLockStateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountLockStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountLockStateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountLockStateResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountLockStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccountLockStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccountLockStateResponse) GetBody() *ModifyAccountLockStateResponseBody {
	return s.Body
}

func (s *ModifyAccountLockStateResponse) SetHeaders(v map[string]*string) *ModifyAccountLockStateResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountLockStateResponse) SetStatusCode(v int32) *ModifyAccountLockStateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountLockStateResponse) SetBody(v *ModifyAccountLockStateResponseBody) *ModifyAccountLockStateResponse {
	s.Body = v
	return s
}

func (s *ModifyAccountLockStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
