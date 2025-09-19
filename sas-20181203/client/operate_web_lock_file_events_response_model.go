// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateWebLockFileEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateWebLockFileEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateWebLockFileEventsResponse
	GetStatusCode() *int32
	SetBody(v *OperateWebLockFileEventsResponseBody) *OperateWebLockFileEventsResponse
	GetBody() *OperateWebLockFileEventsResponseBody
}

type OperateWebLockFileEventsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateWebLockFileEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateWebLockFileEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateWebLockFileEventsResponse) GoString() string {
	return s.String()
}

func (s *OperateWebLockFileEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateWebLockFileEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateWebLockFileEventsResponse) GetBody() *OperateWebLockFileEventsResponseBody {
	return s.Body
}

func (s *OperateWebLockFileEventsResponse) SetHeaders(v map[string]*string) *OperateWebLockFileEventsResponse {
	s.Headers = v
	return s
}

func (s *OperateWebLockFileEventsResponse) SetStatusCode(v int32) *OperateWebLockFileEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateWebLockFileEventsResponse) SetBody(v *OperateWebLockFileEventsResponseBody) *OperateWebLockFileEventsResponse {
	s.Body = v
	return s
}

func (s *OperateWebLockFileEventsResponse) Validate() error {
	return dara.Validate(s)
}
