// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudauthstSceneListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudauthstSceneListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudauthstSceneListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudauthstSceneListResponseBody) *DescribeCloudauthstSceneListResponse
	GetBody() *DescribeCloudauthstSceneListResponseBody
}

type DescribeCloudauthstSceneListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudauthstSceneListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudauthstSceneListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudauthstSceneListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudauthstSceneListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudauthstSceneListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudauthstSceneListResponse) GetBody() *DescribeCloudauthstSceneListResponseBody {
	return s.Body
}

func (s *DescribeCloudauthstSceneListResponse) SetHeaders(v map[string]*string) *DescribeCloudauthstSceneListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudauthstSceneListResponse) SetStatusCode(v int32) *DescribeCloudauthstSceneListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudauthstSceneListResponse) SetBody(v *DescribeCloudauthstSceneListResponseBody) *DescribeCloudauthstSceneListResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudauthstSceneListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
