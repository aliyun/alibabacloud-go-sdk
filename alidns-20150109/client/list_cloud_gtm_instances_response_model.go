// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudGtmInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudGtmInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudGtmInstancesResponseBody) *ListCloudGtmInstancesResponse
	GetBody() *ListCloudGtmInstancesResponseBody
}

type ListCloudGtmInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudGtmInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudGtmInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudGtmInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudGtmInstancesResponse) GetBody() *ListCloudGtmInstancesResponseBody {
	return s.Body
}

func (s *ListCloudGtmInstancesResponse) SetHeaders(v map[string]*string) *ListCloudGtmInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudGtmInstancesResponse) SetStatusCode(v int32) *ListCloudGtmInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudGtmInstancesResponse) SetBody(v *ListCloudGtmInstancesResponseBody) *ListCloudGtmInstancesResponse {
	s.Body = v
	return s
}

func (s *ListCloudGtmInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
