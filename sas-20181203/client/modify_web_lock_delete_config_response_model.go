// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockDeleteConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebLockDeleteConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebLockDeleteConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebLockDeleteConfigResponseBody) *ModifyWebLockDeleteConfigResponse
	GetBody() *ModifyWebLockDeleteConfigResponseBody
}

type ModifyWebLockDeleteConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebLockDeleteConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebLockDeleteConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockDeleteConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockDeleteConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebLockDeleteConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebLockDeleteConfigResponse) GetBody() *ModifyWebLockDeleteConfigResponseBody {
	return s.Body
}

func (s *ModifyWebLockDeleteConfigResponse) SetHeaders(v map[string]*string) *ModifyWebLockDeleteConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockDeleteConfigResponse) SetStatusCode(v int32) *ModifyWebLockDeleteConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebLockDeleteConfigResponse) SetBody(v *ModifyWebLockDeleteConfigResponseBody) *ModifyWebLockDeleteConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyWebLockDeleteConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
