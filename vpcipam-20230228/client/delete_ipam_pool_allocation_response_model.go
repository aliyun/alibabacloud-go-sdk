// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamPoolAllocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpamPoolAllocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpamPoolAllocationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpamPoolAllocationResponseBody) *DeleteIpamPoolAllocationResponse
	GetBody() *DeleteIpamPoolAllocationResponseBody
}

type DeleteIpamPoolAllocationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpamPoolAllocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpamPoolAllocationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamPoolAllocationResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolAllocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpamPoolAllocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpamPoolAllocationResponse) GetBody() *DeleteIpamPoolAllocationResponseBody {
	return s.Body
}

func (s *DeleteIpamPoolAllocationResponse) SetHeaders(v map[string]*string) *DeleteIpamPoolAllocationResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpamPoolAllocationResponse) SetStatusCode(v int32) *DeleteIpamPoolAllocationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpamPoolAllocationResponse) SetBody(v *DeleteIpamPoolAllocationResponseBody) *DeleteIpamPoolAllocationResponse {
	s.Body = v
	return s
}

func (s *DeleteIpamPoolAllocationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
