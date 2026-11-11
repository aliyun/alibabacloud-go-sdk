// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccountResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudAccountResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudAccountResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudAccountResourcesResponseBody) *ListCloudAccountResourcesResponse
	GetBody() *ListCloudAccountResourcesResponseBody
}

type ListCloudAccountResourcesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudAccountResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudAccountResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudAccountResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudAccountResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudAccountResourcesResponse) GetBody() *ListCloudAccountResourcesResponseBody {
	return s.Body
}

func (s *ListCloudAccountResourcesResponse) SetHeaders(v map[string]*string) *ListCloudAccountResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudAccountResourcesResponse) SetStatusCode(v int32) *ListCloudAccountResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudAccountResourcesResponse) SetBody(v *ListCloudAccountResourcesResponseBody) *ListCloudAccountResourcesResponse {
	s.Body = v
	return s
}

func (s *ListCloudAccountResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
