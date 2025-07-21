// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpChangeWarmupTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DedicatedIpChangeWarmupTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DedicatedIpChangeWarmupTypeResponse
	GetStatusCode() *int32
	SetBody(v *DedicatedIpChangeWarmupTypeResponseBody) *DedicatedIpChangeWarmupTypeResponse
	GetBody() *DedicatedIpChangeWarmupTypeResponseBody
}

type DedicatedIpChangeWarmupTypeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DedicatedIpChangeWarmupTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DedicatedIpChangeWarmupTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpChangeWarmupTypeResponse) GoString() string {
	return s.String()
}

func (s *DedicatedIpChangeWarmupTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DedicatedIpChangeWarmupTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DedicatedIpChangeWarmupTypeResponse) GetBody() *DedicatedIpChangeWarmupTypeResponseBody {
	return s.Body
}

func (s *DedicatedIpChangeWarmupTypeResponse) SetHeaders(v map[string]*string) *DedicatedIpChangeWarmupTypeResponse {
	s.Headers = v
	return s
}

func (s *DedicatedIpChangeWarmupTypeResponse) SetStatusCode(v int32) *DedicatedIpChangeWarmupTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DedicatedIpChangeWarmupTypeResponse) SetBody(v *DedicatedIpChangeWarmupTypeResponseBody) *DedicatedIpChangeWarmupTypeResponse {
	s.Body = v
	return s
}

func (s *DedicatedIpChangeWarmupTypeResponse) Validate() error {
	return dara.Validate(s)
}
