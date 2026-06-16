// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppNotificationSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppNotificationSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppNotificationSceneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppNotificationSceneResponseBody) *DeleteAppNotificationSceneResponse
	GetBody() *DeleteAppNotificationSceneResponseBody
}

type DeleteAppNotificationSceneResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppNotificationSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppNotificationSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppNotificationSceneResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppNotificationSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppNotificationSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppNotificationSceneResponse) GetBody() *DeleteAppNotificationSceneResponseBody {
	return s.Body
}

func (s *DeleteAppNotificationSceneResponse) SetHeaders(v map[string]*string) *DeleteAppNotificationSceneResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppNotificationSceneResponse) SetStatusCode(v int32) *DeleteAppNotificationSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppNotificationSceneResponse) SetBody(v *DeleteAppNotificationSceneResponseBody) *DeleteAppNotificationSceneResponse {
	s.Body = v
	return s
}

func (s *DeleteAppNotificationSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
