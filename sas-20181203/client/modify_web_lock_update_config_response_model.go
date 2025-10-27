// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockUpdateConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebLockUpdateConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebLockUpdateConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebLockUpdateConfigResponseBody) *ModifyWebLockUpdateConfigResponse
	GetBody() *ModifyWebLockUpdateConfigResponseBody
}

type ModifyWebLockUpdateConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebLockUpdateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebLockUpdateConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockUpdateConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockUpdateConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebLockUpdateConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebLockUpdateConfigResponse) GetBody() *ModifyWebLockUpdateConfigResponseBody {
	return s.Body
}

func (s *ModifyWebLockUpdateConfigResponse) SetHeaders(v map[string]*string) *ModifyWebLockUpdateConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockUpdateConfigResponse) SetStatusCode(v int32) *ModifyWebLockUpdateConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebLockUpdateConfigResponse) SetBody(v *ModifyWebLockUpdateConfigResponseBody) *ModifyWebLockUpdateConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyWebLockUpdateConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
