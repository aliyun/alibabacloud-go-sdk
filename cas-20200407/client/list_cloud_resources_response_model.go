// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudResourcesResponseBody) *ListCloudResourcesResponse
	GetBody() *ListCloudResourcesResponseBody
}

type ListCloudResourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudResourcesResponse) GetBody() *ListCloudResourcesResponseBody {
	return s.Body
}

func (s *ListCloudResourcesResponse) SetHeaders(v map[string]*string) *ListCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudResourcesResponse) SetStatusCode(v int32) *ListCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudResourcesResponse) SetBody(v *ListCloudResourcesResponseBody) *ListCloudResourcesResponse {
	s.Body = v
	return s
}

func (s *ListCloudResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
