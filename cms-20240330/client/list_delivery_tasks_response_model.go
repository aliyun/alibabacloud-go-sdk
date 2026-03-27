// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeliveryTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeliveryTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListDeliveryTasksResponseBody) *ListDeliveryTasksResponse
	GetBody() *ListDeliveryTasksResponseBody
}

type ListDeliveryTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeliveryTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeliveryTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDeliveryTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeliveryTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeliveryTasksResponse) GetBody() *ListDeliveryTasksResponseBody {
	return s.Body
}

func (s *ListDeliveryTasksResponse) SetHeaders(v map[string]*string) *ListDeliveryTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDeliveryTasksResponse) SetStatusCode(v int32) *ListDeliveryTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeliveryTasksResponse) SetBody(v *ListDeliveryTasksResponseBody) *ListDeliveryTasksResponse {
	s.Body = v
	return s
}

func (s *ListDeliveryTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
