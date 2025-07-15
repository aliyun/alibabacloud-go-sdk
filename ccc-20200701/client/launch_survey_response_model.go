// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchSurveyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LaunchSurveyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LaunchSurveyResponse
	GetStatusCode() *int32
	SetBody(v *LaunchSurveyResponseBody) *LaunchSurveyResponse
	GetBody() *LaunchSurveyResponseBody
}

type LaunchSurveyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LaunchSurveyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LaunchSurveyResponse) String() string {
	return dara.Prettify(s)
}

func (s LaunchSurveyResponse) GoString() string {
	return s.String()
}

func (s *LaunchSurveyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LaunchSurveyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LaunchSurveyResponse) GetBody() *LaunchSurveyResponseBody {
	return s.Body
}

func (s *LaunchSurveyResponse) SetHeaders(v map[string]*string) *LaunchSurveyResponse {
	s.Headers = v
	return s
}

func (s *LaunchSurveyResponse) SetStatusCode(v int32) *LaunchSurveyResponse {
	s.StatusCode = &v
	return s
}

func (s *LaunchSurveyResponse) SetBody(v *LaunchSurveyResponseBody) *LaunchSurveyResponse {
	s.Body = v
	return s
}

func (s *LaunchSurveyResponse) Validate() error {
	return dara.Validate(s)
}
