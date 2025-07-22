// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeadLockHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeadLockHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeadLockHistoryResponse
	GetStatusCode() *int32
	SetBody(v *GetDeadLockHistoryResponseBody) *GetDeadLockHistoryResponse
	GetBody() *GetDeadLockHistoryResponseBody
}

type GetDeadLockHistoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeadLockHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeadLockHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeadLockHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetDeadLockHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeadLockHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeadLockHistoryResponse) GetBody() *GetDeadLockHistoryResponseBody {
	return s.Body
}

func (s *GetDeadLockHistoryResponse) SetHeaders(v map[string]*string) *GetDeadLockHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetDeadLockHistoryResponse) SetStatusCode(v int32) *GetDeadLockHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeadLockHistoryResponse) SetBody(v *GetDeadLockHistoryResponseBody) *GetDeadLockHistoryResponse {
	s.Body = v
	return s
}

func (s *GetDeadLockHistoryResponse) Validate() error {
	return dara.Validate(s)
}
