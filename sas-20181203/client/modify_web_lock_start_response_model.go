// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockStartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebLockStartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebLockStartResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebLockStartResponseBody) *ModifyWebLockStartResponse
	GetBody() *ModifyWebLockStartResponseBody
}

type ModifyWebLockStartResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebLockStartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebLockStartResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockStartResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockStartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebLockStartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebLockStartResponse) GetBody() *ModifyWebLockStartResponseBody {
	return s.Body
}

func (s *ModifyWebLockStartResponse) SetHeaders(v map[string]*string) *ModifyWebLockStartResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockStartResponse) SetStatusCode(v int32) *ModifyWebLockStartResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebLockStartResponse) SetBody(v *ModifyWebLockStartResponseBody) *ModifyWebLockStartResponse {
	s.Body = v
	return s
}

func (s *ModifyWebLockStartResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
