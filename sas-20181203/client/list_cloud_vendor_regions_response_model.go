// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudVendorRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudVendorRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudVendorRegionsResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudVendorRegionsResponseBody) *ListCloudVendorRegionsResponse
	GetBody() *ListCloudVendorRegionsResponseBody
}

type ListCloudVendorRegionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudVendorRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudVendorRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudVendorRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListCloudVendorRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudVendorRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudVendorRegionsResponse) GetBody() *ListCloudVendorRegionsResponseBody {
	return s.Body
}

func (s *ListCloudVendorRegionsResponse) SetHeaders(v map[string]*string) *ListCloudVendorRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListCloudVendorRegionsResponse) SetStatusCode(v int32) *ListCloudVendorRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudVendorRegionsResponse) SetBody(v *ListCloudVendorRegionsResponseBody) *ListCloudVendorRegionsResponse {
	s.Body = v
	return s
}

func (s *ListCloudVendorRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
