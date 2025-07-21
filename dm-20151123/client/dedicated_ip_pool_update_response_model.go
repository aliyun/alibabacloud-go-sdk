// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DedicatedIpPoolUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DedicatedIpPoolUpdateResponse
	GetStatusCode() *int32
	SetBody(v *DedicatedIpPoolUpdateResponseBody) *DedicatedIpPoolUpdateResponse
	GetBody() *DedicatedIpPoolUpdateResponseBody
}

type DedicatedIpPoolUpdateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DedicatedIpPoolUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DedicatedIpPoolUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolUpdateResponse) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DedicatedIpPoolUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DedicatedIpPoolUpdateResponse) GetBody() *DedicatedIpPoolUpdateResponseBody {
	return s.Body
}

func (s *DedicatedIpPoolUpdateResponse) SetHeaders(v map[string]*string) *DedicatedIpPoolUpdateResponse {
	s.Headers = v
	return s
}

func (s *DedicatedIpPoolUpdateResponse) SetStatusCode(v int32) *DedicatedIpPoolUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *DedicatedIpPoolUpdateResponse) SetBody(v *DedicatedIpPoolUpdateResponseBody) *DedicatedIpPoolUpdateResponse {
	s.Body = v
	return s
}

func (s *DedicatedIpPoolUpdateResponse) Validate() error {
	return dara.Validate(s)
}
