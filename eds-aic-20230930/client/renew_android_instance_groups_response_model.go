// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAndroidInstanceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewAndroidInstanceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewAndroidInstanceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *RenewAndroidInstanceGroupsResponseBody) *RenewAndroidInstanceGroupsResponse
	GetBody() *RenewAndroidInstanceGroupsResponseBody
}

type RenewAndroidInstanceGroupsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewAndroidInstanceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewAndroidInstanceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewAndroidInstanceGroupsResponse) GoString() string {
	return s.String()
}

func (s *RenewAndroidInstanceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewAndroidInstanceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewAndroidInstanceGroupsResponse) GetBody() *RenewAndroidInstanceGroupsResponseBody {
	return s.Body
}

func (s *RenewAndroidInstanceGroupsResponse) SetHeaders(v map[string]*string) *RenewAndroidInstanceGroupsResponse {
	s.Headers = v
	return s
}

func (s *RenewAndroidInstanceGroupsResponse) SetStatusCode(v int32) *RenewAndroidInstanceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewAndroidInstanceGroupsResponse) SetBody(v *RenewAndroidInstanceGroupsResponseBody) *RenewAndroidInstanceGroupsResponse {
	s.Body = v
	return s
}

func (s *RenewAndroidInstanceGroupsResponse) Validate() error {
	return dara.Validate(s)
}
