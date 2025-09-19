// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudVendorAccountAKResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudVendorAccountAKResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudVendorAccountAKResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudVendorAccountAKResponseBody) *DeleteCloudVendorAccountAKResponse
	GetBody() *DeleteCloudVendorAccountAKResponseBody
}

type DeleteCloudVendorAccountAKResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudVendorAccountAKResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudVendorAccountAKResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudVendorAccountAKResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudVendorAccountAKResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudVendorAccountAKResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudVendorAccountAKResponse) GetBody() *DeleteCloudVendorAccountAKResponseBody {
	return s.Body
}

func (s *DeleteCloudVendorAccountAKResponse) SetHeaders(v map[string]*string) *DeleteCloudVendorAccountAKResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudVendorAccountAKResponse) SetStatusCode(v int32) *DeleteCloudVendorAccountAKResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudVendorAccountAKResponse) SetBody(v *DeleteCloudVendorAccountAKResponseBody) *DeleteCloudVendorAccountAKResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudVendorAccountAKResponse) Validate() error {
	return dara.Validate(s)
}
