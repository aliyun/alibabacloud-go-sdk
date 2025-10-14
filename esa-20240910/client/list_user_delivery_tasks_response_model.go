// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDeliveryTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserDeliveryTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserDeliveryTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListUserDeliveryTasksResponseBody) *ListUserDeliveryTasksResponse
	GetBody() *ListUserDeliveryTasksResponseBody
}

type ListUserDeliveryTasksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserDeliveryTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserDeliveryTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserDeliveryTasksResponse) GoString() string {
	return s.String()
}

func (s *ListUserDeliveryTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserDeliveryTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserDeliveryTasksResponse) GetBody() *ListUserDeliveryTasksResponseBody {
	return s.Body
}

func (s *ListUserDeliveryTasksResponse) SetHeaders(v map[string]*string) *ListUserDeliveryTasksResponse {
	s.Headers = v
	return s
}

func (s *ListUserDeliveryTasksResponse) SetStatusCode(v int32) *ListUserDeliveryTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserDeliveryTasksResponse) SetBody(v *ListUserDeliveryTasksResponseBody) *ListUserDeliveryTasksResponse {
	s.Body = v
	return s
}

func (s *ListUserDeliveryTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
