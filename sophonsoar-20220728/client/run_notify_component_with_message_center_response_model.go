// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNotifyComponentWithMessageCenterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunNotifyComponentWithMessageCenterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunNotifyComponentWithMessageCenterResponse
	GetStatusCode() *int32
	SetBody(v *RunNotifyComponentWithMessageCenterResponseBody) *RunNotifyComponentWithMessageCenterResponse
	GetBody() *RunNotifyComponentWithMessageCenterResponseBody
}

type RunNotifyComponentWithMessageCenterResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunNotifyComponentWithMessageCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunNotifyComponentWithMessageCenterResponse) String() string {
	return dara.Prettify(s)
}

func (s RunNotifyComponentWithMessageCenterResponse) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithMessageCenterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunNotifyComponentWithMessageCenterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunNotifyComponentWithMessageCenterResponse) GetBody() *RunNotifyComponentWithMessageCenterResponseBody {
	return s.Body
}

func (s *RunNotifyComponentWithMessageCenterResponse) SetHeaders(v map[string]*string) *RunNotifyComponentWithMessageCenterResponse {
	s.Headers = v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponse) SetStatusCode(v int32) *RunNotifyComponentWithMessageCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponse) SetBody(v *RunNotifyComponentWithMessageCenterResponseBody) *RunNotifyComponentWithMessageCenterResponse {
	s.Body = v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
