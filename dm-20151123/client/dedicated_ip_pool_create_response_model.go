// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DedicatedIpPoolCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DedicatedIpPoolCreateResponse
	GetStatusCode() *int32
	SetBody(v *DedicatedIpPoolCreateResponseBody) *DedicatedIpPoolCreateResponse
	GetBody() *DedicatedIpPoolCreateResponseBody
}

type DedicatedIpPoolCreateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DedicatedIpPoolCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DedicatedIpPoolCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolCreateResponse) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DedicatedIpPoolCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DedicatedIpPoolCreateResponse) GetBody() *DedicatedIpPoolCreateResponseBody {
	return s.Body
}

func (s *DedicatedIpPoolCreateResponse) SetHeaders(v map[string]*string) *DedicatedIpPoolCreateResponse {
	s.Headers = v
	return s
}

func (s *DedicatedIpPoolCreateResponse) SetStatusCode(v int32) *DedicatedIpPoolCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *DedicatedIpPoolCreateResponse) SetBody(v *DedicatedIpPoolCreateResponseBody) *DedicatedIpPoolCreateResponse {
	s.Body = v
	return s
}

func (s *DedicatedIpPoolCreateResponse) Validate() error {
	return dara.Validate(s)
}
