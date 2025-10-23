// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DedicatedIpPoolListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DedicatedIpPoolListResponse
	GetStatusCode() *int32
	SetBody(v *DedicatedIpPoolListResponseBody) *DedicatedIpPoolListResponse
	GetBody() *DedicatedIpPoolListResponseBody
}

type DedicatedIpPoolListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DedicatedIpPoolListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DedicatedIpPoolListResponse) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolListResponse) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DedicatedIpPoolListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DedicatedIpPoolListResponse) GetBody() *DedicatedIpPoolListResponseBody {
	return s.Body
}

func (s *DedicatedIpPoolListResponse) SetHeaders(v map[string]*string) *DedicatedIpPoolListResponse {
	s.Headers = v
	return s
}

func (s *DedicatedIpPoolListResponse) SetStatusCode(v int32) *DedicatedIpPoolListResponse {
	s.StatusCode = &v
	return s
}

func (s *DedicatedIpPoolListResponse) SetBody(v *DedicatedIpPoolListResponseBody) *DedicatedIpPoolListResponse {
	s.Body = v
	return s
}

func (s *DedicatedIpPoolListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
