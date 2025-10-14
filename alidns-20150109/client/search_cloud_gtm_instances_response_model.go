// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchCloudGtmInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchCloudGtmInstancesResponse
	GetStatusCode() *int32
	SetBody(v *SearchCloudGtmInstancesResponseBody) *SearchCloudGtmInstancesResponse
	GetBody() *SearchCloudGtmInstancesResponseBody
}

type SearchCloudGtmInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchCloudGtmInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchCloudGtmInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstancesResponse) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchCloudGtmInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchCloudGtmInstancesResponse) GetBody() *SearchCloudGtmInstancesResponseBody {
	return s.Body
}

func (s *SearchCloudGtmInstancesResponse) SetHeaders(v map[string]*string) *SearchCloudGtmInstancesResponse {
	s.Headers = v
	return s
}

func (s *SearchCloudGtmInstancesResponse) SetStatusCode(v int32) *SearchCloudGtmInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchCloudGtmInstancesResponse) SetBody(v *SearchCloudGtmInstancesResponseBody) *SearchCloudGtmInstancesResponse {
	s.Body = v
	return s
}

func (s *SearchCloudGtmInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
