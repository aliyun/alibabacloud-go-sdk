// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagCloudResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TagCloudResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TagCloudResourcesResponse
	GetStatusCode() *int32
	SetBody(v *TagCloudResourcesResponseBody) *TagCloudResourcesResponse
	GetBody() *TagCloudResourcesResponseBody
}

type TagCloudResourcesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagCloudResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s TagCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagCloudResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TagCloudResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TagCloudResourcesResponse) GetBody() *TagCloudResourcesResponseBody {
	return s.Body
}

func (s *TagCloudResourcesResponse) SetHeaders(v map[string]*string) *TagCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagCloudResourcesResponse) SetStatusCode(v int32) *TagCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagCloudResourcesResponse) SetBody(v *TagCloudResourcesResponseBody) *TagCloudResourcesResponse {
	s.Body = v
	return s
}

func (s *TagCloudResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
