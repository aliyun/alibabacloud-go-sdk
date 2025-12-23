// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamPoolAllocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIpamPoolAllocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIpamPoolAllocationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIpamPoolAllocationResponseBody) *UpdateIpamPoolAllocationResponse
	GetBody() *UpdateIpamPoolAllocationResponseBody
}

type UpdateIpamPoolAllocationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpamPoolAllocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpamPoolAllocationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamPoolAllocationResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpamPoolAllocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIpamPoolAllocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIpamPoolAllocationResponse) GetBody() *UpdateIpamPoolAllocationResponseBody {
	return s.Body
}

func (s *UpdateIpamPoolAllocationResponse) SetHeaders(v map[string]*string) *UpdateIpamPoolAllocationResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpamPoolAllocationResponse) SetStatusCode(v int32) *UpdateIpamPoolAllocationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpamPoolAllocationResponse) SetBody(v *UpdateIpamPoolAllocationResponseBody) *UpdateIpamPoolAllocationResponse {
	s.Body = v
	return s
}

func (s *UpdateIpamPoolAllocationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
