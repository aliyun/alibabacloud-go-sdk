// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamResourceCidrsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpamResourceCidrsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpamResourceCidrsResponse
	GetStatusCode() *int32
	SetBody(v *ListIpamResourceCidrsResponseBody) *ListIpamResourceCidrsResponse
	GetBody() *ListIpamResourceCidrsResponseBody
}

type ListIpamResourceCidrsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamResourceCidrsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamResourceCidrsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceCidrsResponse) GoString() string {
	return s.String()
}

func (s *ListIpamResourceCidrsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpamResourceCidrsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpamResourceCidrsResponse) GetBody() *ListIpamResourceCidrsResponseBody {
	return s.Body
}

func (s *ListIpamResourceCidrsResponse) SetHeaders(v map[string]*string) *ListIpamResourceCidrsResponse {
	s.Headers = v
	return s
}

func (s *ListIpamResourceCidrsResponse) SetStatusCode(v int32) *ListIpamResourceCidrsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamResourceCidrsResponse) SetBody(v *ListIpamResourceCidrsResponseBody) *ListIpamResourceCidrsResponse {
	s.Body = v
	return s
}

func (s *ListIpamResourceCidrsResponse) Validate() error {
	return dara.Validate(s)
}
