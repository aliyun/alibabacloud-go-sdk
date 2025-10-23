// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DedicatedIpPoolDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DedicatedIpPoolDeleteResponse
	GetStatusCode() *int32
	SetBody(v *DedicatedIpPoolDeleteResponseBody) *DedicatedIpPoolDeleteResponse
	GetBody() *DedicatedIpPoolDeleteResponseBody
}

type DedicatedIpPoolDeleteResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DedicatedIpPoolDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DedicatedIpPoolDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolDeleteResponse) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DedicatedIpPoolDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DedicatedIpPoolDeleteResponse) GetBody() *DedicatedIpPoolDeleteResponseBody {
	return s.Body
}

func (s *DedicatedIpPoolDeleteResponse) SetHeaders(v map[string]*string) *DedicatedIpPoolDeleteResponse {
	s.Headers = v
	return s
}

func (s *DedicatedIpPoolDeleteResponse) SetStatusCode(v int32) *DedicatedIpPoolDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *DedicatedIpPoolDeleteResponse) SetBody(v *DedicatedIpPoolDeleteResponseBody) *DedicatedIpPoolDeleteResponse {
	s.Body = v
	return s
}

func (s *DedicatedIpPoolDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
