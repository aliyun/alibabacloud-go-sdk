// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpNonePoolListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DedicatedIpNonePoolListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DedicatedIpNonePoolListResponse
	GetStatusCode() *int32
	SetBody(v *DedicatedIpNonePoolListResponseBody) *DedicatedIpNonePoolListResponse
	GetBody() *DedicatedIpNonePoolListResponseBody
}

type DedicatedIpNonePoolListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DedicatedIpNonePoolListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DedicatedIpNonePoolListResponse) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpNonePoolListResponse) GoString() string {
	return s.String()
}

func (s *DedicatedIpNonePoolListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DedicatedIpNonePoolListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DedicatedIpNonePoolListResponse) GetBody() *DedicatedIpNonePoolListResponseBody {
	return s.Body
}

func (s *DedicatedIpNonePoolListResponse) SetHeaders(v map[string]*string) *DedicatedIpNonePoolListResponse {
	s.Headers = v
	return s
}

func (s *DedicatedIpNonePoolListResponse) SetStatusCode(v int32) *DedicatedIpNonePoolListResponse {
	s.StatusCode = &v
	return s
}

func (s *DedicatedIpNonePoolListResponse) SetBody(v *DedicatedIpNonePoolListResponseBody) *DedicatedIpNonePoolListResponse {
	s.Body = v
	return s
}

func (s *DedicatedIpNonePoolListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
