// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppNotificationSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppNotificationSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppNotificationSceneResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppNotificationSceneResponseBody) *CreateAppNotificationSceneResponse
	GetBody() *CreateAppNotificationSceneResponseBody
}

type CreateAppNotificationSceneResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppNotificationSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppNotificationSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppNotificationSceneResponse) GoString() string {
	return s.String()
}

func (s *CreateAppNotificationSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppNotificationSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppNotificationSceneResponse) GetBody() *CreateAppNotificationSceneResponseBody {
	return s.Body
}

func (s *CreateAppNotificationSceneResponse) SetHeaders(v map[string]*string) *CreateAppNotificationSceneResponse {
	s.Headers = v
	return s
}

func (s *CreateAppNotificationSceneResponse) SetStatusCode(v int32) *CreateAppNotificationSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppNotificationSceneResponse) SetBody(v *CreateAppNotificationSceneResponseBody) *CreateAppNotificationSceneResponse {
	s.Body = v
	return s
}

func (s *CreateAppNotificationSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
