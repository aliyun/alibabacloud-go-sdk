// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamPoolCidrsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpamPoolCidrsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpamPoolCidrsResponse
	GetStatusCode() *int32
	SetBody(v *ListIpamPoolCidrsResponseBody) *ListIpamPoolCidrsResponse
	GetBody() *ListIpamPoolCidrsResponseBody
}

type ListIpamPoolCidrsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamPoolCidrsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamPoolCidrsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolCidrsResponse) GoString() string {
	return s.String()
}

func (s *ListIpamPoolCidrsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpamPoolCidrsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpamPoolCidrsResponse) GetBody() *ListIpamPoolCidrsResponseBody {
	return s.Body
}

func (s *ListIpamPoolCidrsResponse) SetHeaders(v map[string]*string) *ListIpamPoolCidrsResponse {
	s.Headers = v
	return s
}

func (s *ListIpamPoolCidrsResponse) SetStatusCode(v int32) *ListIpamPoolCidrsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamPoolCidrsResponse) SetBody(v *ListIpamPoolCidrsResponseBody) *ListIpamPoolCidrsResponse {
	s.Body = v
	return s
}

func (s *ListIpamPoolCidrsResponse) Validate() error {
	return dara.Validate(s)
}
