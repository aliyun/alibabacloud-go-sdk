// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudVendorAccountAKResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCloudVendorAccountAKResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCloudVendorAccountAKResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCloudVendorAccountAKResponseBody) *ModifyCloudVendorAccountAKResponse
	GetBody() *ModifyCloudVendorAccountAKResponseBody
}

type ModifyCloudVendorAccountAKResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudVendorAccountAKResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudVendorAccountAKResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudVendorAccountAKResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudVendorAccountAKResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCloudVendorAccountAKResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCloudVendorAccountAKResponse) GetBody() *ModifyCloudVendorAccountAKResponseBody {
	return s.Body
}

func (s *ModifyCloudVendorAccountAKResponse) SetHeaders(v map[string]*string) *ModifyCloudVendorAccountAKResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudVendorAccountAKResponse) SetStatusCode(v int32) *ModifyCloudVendorAccountAKResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponse) SetBody(v *ModifyCloudVendorAccountAKResponseBody) *ModifyCloudVendorAccountAKResponse {
	s.Body = v
	return s
}

func (s *ModifyCloudVendorAccountAKResponse) Validate() error {
	return dara.Validate(s)
}
