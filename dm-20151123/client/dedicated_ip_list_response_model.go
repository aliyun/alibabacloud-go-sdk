// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DedicatedIpListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DedicatedIpListResponse
	GetStatusCode() *int32
	SetBody(v *DedicatedIpListResponseBody) *DedicatedIpListResponse
	GetBody() *DedicatedIpListResponseBody
}

type DedicatedIpListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DedicatedIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DedicatedIpListResponse) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpListResponse) GoString() string {
	return s.String()
}

func (s *DedicatedIpListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DedicatedIpListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DedicatedIpListResponse) GetBody() *DedicatedIpListResponseBody {
	return s.Body
}

func (s *DedicatedIpListResponse) SetHeaders(v map[string]*string) *DedicatedIpListResponse {
	s.Headers = v
	return s
}

func (s *DedicatedIpListResponse) SetStatusCode(v int32) *DedicatedIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *DedicatedIpListResponse) SetBody(v *DedicatedIpListResponseBody) *DedicatedIpListResponse {
	s.Body = v
	return s
}

func (s *DedicatedIpListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
