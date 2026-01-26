// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseAlarmResponse
	GetStatusCode() *int32
	SetBody(v *CloseAlarmResponseBody) *CloseAlarmResponse
	GetBody() *CloseAlarmResponseBody
}

type CloseAlarmResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseAlarmResponse) GoString() string {
	return s.String()
}

func (s *CloseAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseAlarmResponse) GetBody() *CloseAlarmResponseBody {
	return s.Body
}

func (s *CloseAlarmResponse) SetHeaders(v map[string]*string) *CloseAlarmResponse {
	s.Headers = v
	return s
}

func (s *CloseAlarmResponse) SetStatusCode(v int32) *CloseAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseAlarmResponse) SetBody(v *CloseAlarmResponseBody) *CloseAlarmResponse {
	s.Body = v
	return s
}

func (s *CloseAlarmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
