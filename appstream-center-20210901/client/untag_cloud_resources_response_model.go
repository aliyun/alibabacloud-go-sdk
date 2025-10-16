// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagCloudResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UntagCloudResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UntagCloudResourcesResponse
	GetStatusCode() *int32
	SetBody(v *UntagCloudResourcesResponseBody) *UntagCloudResourcesResponse
	GetBody() *UntagCloudResourcesResponseBody
}

type UntagCloudResourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagCloudResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s UntagCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagCloudResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UntagCloudResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UntagCloudResourcesResponse) GetBody() *UntagCloudResourcesResponseBody {
	return s.Body
}

func (s *UntagCloudResourcesResponse) SetHeaders(v map[string]*string) *UntagCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagCloudResourcesResponse) SetStatusCode(v int32) *UntagCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagCloudResourcesResponse) SetBody(v *UntagCloudResourcesResponseBody) *UntagCloudResourcesResponse {
	s.Body = v
	return s
}

func (s *UntagCloudResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
