// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcIpamServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpcIpamServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpcIpamServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetVpcIpamServiceStatusResponseBody) *GetVpcIpamServiceStatusResponse
	GetBody() *GetVpcIpamServiceStatusResponseBody
}

type GetVpcIpamServiceStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpcIpamServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpcIpamServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpcIpamServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetVpcIpamServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpcIpamServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpcIpamServiceStatusResponse) GetBody() *GetVpcIpamServiceStatusResponseBody {
	return s.Body
}

func (s *GetVpcIpamServiceStatusResponse) SetHeaders(v map[string]*string) *GetVpcIpamServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetVpcIpamServiceStatusResponse) SetStatusCode(v int32) *GetVpcIpamServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcIpamServiceStatusResponse) SetBody(v *GetVpcIpamServiceStatusResponseBody) *GetVpcIpamServiceStatusResponse {
	s.Body = v
	return s
}

func (s *GetVpcIpamServiceStatusResponse) Validate() error {
	return dara.Validate(s)
}
