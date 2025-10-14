// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSiteDeliveryTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSiteDeliveryTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSiteDeliveryTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListSiteDeliveryTasksResponseBody) *ListSiteDeliveryTasksResponse
	GetBody() *ListSiteDeliveryTasksResponseBody
}

type ListSiteDeliveryTasksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSiteDeliveryTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSiteDeliveryTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSiteDeliveryTasksResponse) GoString() string {
	return s.String()
}

func (s *ListSiteDeliveryTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSiteDeliveryTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSiteDeliveryTasksResponse) GetBody() *ListSiteDeliveryTasksResponseBody {
	return s.Body
}

func (s *ListSiteDeliveryTasksResponse) SetHeaders(v map[string]*string) *ListSiteDeliveryTasksResponse {
	s.Headers = v
	return s
}

func (s *ListSiteDeliveryTasksResponse) SetStatusCode(v int32) *ListSiteDeliveryTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSiteDeliveryTasksResponse) SetBody(v *ListSiteDeliveryTasksResponseBody) *ListSiteDeliveryTasksResponse {
	s.Body = v
	return s
}

func (s *ListSiteDeliveryTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
