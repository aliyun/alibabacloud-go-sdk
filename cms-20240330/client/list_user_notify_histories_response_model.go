// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserNotifyHistoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserNotifyHistoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserNotifyHistoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserNotifyHistoriesResponseBody) *ListUserNotifyHistoriesResponse
	GetBody() *ListUserNotifyHistoriesResponseBody
}

type ListUserNotifyHistoriesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserNotifyHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserNotifyHistoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserNotifyHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListUserNotifyHistoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserNotifyHistoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserNotifyHistoriesResponse) GetBody() *ListUserNotifyHistoriesResponseBody {
	return s.Body
}

func (s *ListUserNotifyHistoriesResponse) SetHeaders(v map[string]*string) *ListUserNotifyHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListUserNotifyHistoriesResponse) SetStatusCode(v int32) *ListUserNotifyHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserNotifyHistoriesResponse) SetBody(v *ListUserNotifyHistoriesResponseBody) *ListUserNotifyHistoriesResponse {
	s.Body = v
	return s
}

func (s *ListUserNotifyHistoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
