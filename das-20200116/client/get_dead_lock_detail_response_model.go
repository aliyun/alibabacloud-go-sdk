// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeadLockDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeadLockDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeadLockDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDeadLockDetailResponseBody) *GetDeadLockDetailResponse
	GetBody() *GetDeadLockDetailResponseBody
}

type GetDeadLockDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeadLockDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeadLockDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeadLockDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDeadLockDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeadLockDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeadLockDetailResponse) GetBody() *GetDeadLockDetailResponseBody {
	return s.Body
}

func (s *GetDeadLockDetailResponse) SetHeaders(v map[string]*string) *GetDeadLockDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDeadLockDetailResponse) SetStatusCode(v int32) *GetDeadLockDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeadLockDetailResponse) SetBody(v *GetDeadLockDetailResponseBody) *GetDeadLockDetailResponse {
	s.Body = v
	return s
}

func (s *GetDeadLockDetailResponse) Validate() error {
	return dara.Validate(s)
}
