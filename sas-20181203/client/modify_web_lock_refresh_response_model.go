// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockRefreshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebLockRefreshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebLockRefreshResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebLockRefreshResponseBody) *ModifyWebLockRefreshResponse
	GetBody() *ModifyWebLockRefreshResponseBody
}

type ModifyWebLockRefreshResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebLockRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebLockRefreshResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockRefreshResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockRefreshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebLockRefreshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebLockRefreshResponse) GetBody() *ModifyWebLockRefreshResponseBody {
	return s.Body
}

func (s *ModifyWebLockRefreshResponse) SetHeaders(v map[string]*string) *ModifyWebLockRefreshResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockRefreshResponse) SetStatusCode(v int32) *ModifyWebLockRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebLockRefreshResponse) SetBody(v *ModifyWebLockRefreshResponseBody) *ModifyWebLockRefreshResponse {
	s.Body = v
	return s
}

func (s *ModifyWebLockRefreshResponse) Validate() error {
	return dara.Validate(s)
}
