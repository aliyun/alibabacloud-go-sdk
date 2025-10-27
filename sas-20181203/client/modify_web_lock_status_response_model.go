// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebLockStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebLockStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebLockStatusResponseBody) *ModifyWebLockStatusResponse
	GetBody() *ModifyWebLockStatusResponseBody
}

type ModifyWebLockStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebLockStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebLockStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebLockStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebLockStatusResponse) GetBody() *ModifyWebLockStatusResponseBody {
	return s.Body
}

func (s *ModifyWebLockStatusResponse) SetHeaders(v map[string]*string) *ModifyWebLockStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockStatusResponse) SetStatusCode(v int32) *ModifyWebLockStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebLockStatusResponse) SetBody(v *ModifyWebLockStatusResponseBody) *ModifyWebLockStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyWebLockStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
