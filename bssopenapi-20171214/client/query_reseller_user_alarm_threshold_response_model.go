// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResellerUserAlarmThresholdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryResellerUserAlarmThresholdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryResellerUserAlarmThresholdResponse
	GetStatusCode() *int32
	SetBody(v *QueryResellerUserAlarmThresholdResponseBody) *QueryResellerUserAlarmThresholdResponse
	GetBody() *QueryResellerUserAlarmThresholdResponseBody
}

type QueryResellerUserAlarmThresholdResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryResellerUserAlarmThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryResellerUserAlarmThresholdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryResellerUserAlarmThresholdResponse) GoString() string {
	return s.String()
}

func (s *QueryResellerUserAlarmThresholdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryResellerUserAlarmThresholdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryResellerUserAlarmThresholdResponse) GetBody() *QueryResellerUserAlarmThresholdResponseBody {
	return s.Body
}

func (s *QueryResellerUserAlarmThresholdResponse) SetHeaders(v map[string]*string) *QueryResellerUserAlarmThresholdResponse {
	s.Headers = v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponse) SetStatusCode(v int32) *QueryResellerUserAlarmThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponse) SetBody(v *QueryResellerUserAlarmThresholdResponseBody) *QueryResellerUserAlarmThresholdResponse {
	s.Body = v
	return s
}

func (s *QueryResellerUserAlarmThresholdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
