// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudClusterAllUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCloudClusterAllUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCloudClusterAllUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetCloudClusterAllUrlResponseBody) *GetCloudClusterAllUrlResponse
	GetBody() *GetCloudClusterAllUrlResponseBody
}

type GetCloudClusterAllUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudClusterAllUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudClusterAllUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCloudClusterAllUrlResponse) GoString() string {
	return s.String()
}

func (s *GetCloudClusterAllUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCloudClusterAllUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCloudClusterAllUrlResponse) GetBody() *GetCloudClusterAllUrlResponseBody {
	return s.Body
}

func (s *GetCloudClusterAllUrlResponse) SetHeaders(v map[string]*string) *GetCloudClusterAllUrlResponse {
	s.Headers = v
	return s
}

func (s *GetCloudClusterAllUrlResponse) SetStatusCode(v int32) *GetCloudClusterAllUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudClusterAllUrlResponse) SetBody(v *GetCloudClusterAllUrlResponseBody) *GetCloudClusterAllUrlResponse {
	s.Body = v
	return s
}

func (s *GetCloudClusterAllUrlResponse) Validate() error {
	return dara.Validate(s)
}
