// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpamPoolAllocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIpamPoolAllocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIpamPoolAllocationResponse
	GetStatusCode() *int32
	SetBody(v *GetIpamPoolAllocationResponseBody) *GetIpamPoolAllocationResponse
	GetBody() *GetIpamPoolAllocationResponseBody
}

type GetIpamPoolAllocationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIpamPoolAllocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIpamPoolAllocationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIpamPoolAllocationResponse) GoString() string {
	return s.String()
}

func (s *GetIpamPoolAllocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIpamPoolAllocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIpamPoolAllocationResponse) GetBody() *GetIpamPoolAllocationResponseBody {
	return s.Body
}

func (s *GetIpamPoolAllocationResponse) SetHeaders(v map[string]*string) *GetIpamPoolAllocationResponse {
	s.Headers = v
	return s
}

func (s *GetIpamPoolAllocationResponse) SetStatusCode(v int32) *GetIpamPoolAllocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIpamPoolAllocationResponse) SetBody(v *GetIpamPoolAllocationResponseBody) *GetIpamPoolAllocationResponse {
	s.Body = v
	return s
}

func (s *GetIpamPoolAllocationResponse) Validate() error {
	return dara.Validate(s)
}
