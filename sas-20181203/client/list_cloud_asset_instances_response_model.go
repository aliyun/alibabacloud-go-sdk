// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAssetInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudAssetInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudAssetInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudAssetInstancesResponseBody) *ListCloudAssetInstancesResponse
	GetBody() *ListCloudAssetInstancesResponseBody
}

type ListCloudAssetInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudAssetInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudAssetInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudAssetInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudAssetInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudAssetInstancesResponse) GetBody() *ListCloudAssetInstancesResponseBody {
	return s.Body
}

func (s *ListCloudAssetInstancesResponse) SetHeaders(v map[string]*string) *ListCloudAssetInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudAssetInstancesResponse) SetStatusCode(v int32) *ListCloudAssetInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudAssetInstancesResponse) SetBody(v *ListCloudAssetInstancesResponseBody) *ListCloudAssetInstancesResponse {
	s.Body = v
	return s
}

func (s *ListCloudAssetInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
