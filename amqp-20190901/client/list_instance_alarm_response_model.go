// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceAlarmResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceAlarmResponseBody) *ListInstanceAlarmResponse
	GetBody() *ListInstanceAlarmResponseBody
}

type ListInstanceAlarmResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAlarmResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceAlarmResponse) GetBody() *ListInstanceAlarmResponseBody {
	return s.Body
}

func (s *ListInstanceAlarmResponse) SetHeaders(v map[string]*string) *ListInstanceAlarmResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceAlarmResponse) SetStatusCode(v int32) *ListInstanceAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceAlarmResponse) SetBody(v *ListInstanceAlarmResponseBody) *ListInstanceAlarmResponse {
	s.Body = v
	return s
}

func (s *ListInstanceAlarmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
