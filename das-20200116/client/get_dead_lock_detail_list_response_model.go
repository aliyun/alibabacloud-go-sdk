// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeadLockDetailListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeadLockDetailListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeadLockDetailListResponse
	GetStatusCode() *int32
	SetBody(v *GetDeadLockDetailListResponseBody) *GetDeadLockDetailListResponse
	GetBody() *GetDeadLockDetailListResponseBody
}

type GetDeadLockDetailListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeadLockDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeadLockDetailListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeadLockDetailListResponse) GoString() string {
	return s.String()
}

func (s *GetDeadLockDetailListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeadLockDetailListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeadLockDetailListResponse) GetBody() *GetDeadLockDetailListResponseBody {
	return s.Body
}

func (s *GetDeadLockDetailListResponse) SetHeaders(v map[string]*string) *GetDeadLockDetailListResponse {
	s.Headers = v
	return s
}

func (s *GetDeadLockDetailListResponse) SetStatusCode(v int32) *GetDeadLockDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeadLockDetailListResponse) SetBody(v *GetDeadLockDetailListResponseBody) *GetDeadLockDetailListResponse {
	s.Body = v
	return s
}

func (s *GetDeadLockDetailListResponse) Validate() error {
	return dara.Validate(s)
}
