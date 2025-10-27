// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCloudVendorAccountAKResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCloudVendorAccountAKResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCloudVendorAccountAKResponse
	GetStatusCode() *int32
	SetBody(v *AddCloudVendorAccountAKResponseBody) *AddCloudVendorAccountAKResponse
	GetBody() *AddCloudVendorAccountAKResponseBody
}

type AddCloudVendorAccountAKResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCloudVendorAccountAKResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCloudVendorAccountAKResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCloudVendorAccountAKResponse) GoString() string {
	return s.String()
}

func (s *AddCloudVendorAccountAKResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCloudVendorAccountAKResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCloudVendorAccountAKResponse) GetBody() *AddCloudVendorAccountAKResponseBody {
	return s.Body
}

func (s *AddCloudVendorAccountAKResponse) SetHeaders(v map[string]*string) *AddCloudVendorAccountAKResponse {
	s.Headers = v
	return s
}

func (s *AddCloudVendorAccountAKResponse) SetStatusCode(v int32) *AddCloudVendorAccountAKResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCloudVendorAccountAKResponse) SetBody(v *AddCloudVendorAccountAKResponseBody) *AddCloudVendorAccountAKResponse {
	s.Body = v
	return s
}

func (s *AddCloudVendorAccountAKResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
