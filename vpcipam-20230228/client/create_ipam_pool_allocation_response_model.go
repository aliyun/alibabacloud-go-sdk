// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamPoolAllocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpamPoolAllocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpamPoolAllocationResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpamPoolAllocationResponseBody) *CreateIpamPoolAllocationResponse
	GetBody() *CreateIpamPoolAllocationResponseBody
}

type CreateIpamPoolAllocationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpamPoolAllocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpamPoolAllocationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamPoolAllocationResponse) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolAllocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpamPoolAllocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpamPoolAllocationResponse) GetBody() *CreateIpamPoolAllocationResponseBody {
	return s.Body
}

func (s *CreateIpamPoolAllocationResponse) SetHeaders(v map[string]*string) *CreateIpamPoolAllocationResponse {
	s.Headers = v
	return s
}

func (s *CreateIpamPoolAllocationResponse) SetStatusCode(v int32) *CreateIpamPoolAllocationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpamPoolAllocationResponse) SetBody(v *CreateIpamPoolAllocationResponseBody) *CreateIpamPoolAllocationResponse {
	s.Body = v
	return s
}

func (s *CreateIpamPoolAllocationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
