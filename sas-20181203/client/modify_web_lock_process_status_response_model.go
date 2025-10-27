// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockProcessStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebLockProcessStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebLockProcessStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebLockProcessStatusResponseBody) *ModifyWebLockProcessStatusResponse
	GetBody() *ModifyWebLockProcessStatusResponseBody
}

type ModifyWebLockProcessStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebLockProcessStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebLockProcessStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockProcessStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockProcessStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebLockProcessStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebLockProcessStatusResponse) GetBody() *ModifyWebLockProcessStatusResponseBody {
	return s.Body
}

func (s *ModifyWebLockProcessStatusResponse) SetHeaders(v map[string]*string) *ModifyWebLockProcessStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockProcessStatusResponse) SetStatusCode(v int32) *ModifyWebLockProcessStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebLockProcessStatusResponse) SetBody(v *ModifyWebLockProcessStatusResponseBody) *ModifyWebLockProcessStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyWebLockProcessStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
