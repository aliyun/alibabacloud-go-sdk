// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateInstanceWithPrivateDNSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateInstanceWithPrivateDNSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateInstanceWithPrivateDNSResponse
	GetStatusCode() *int32
	SetBody(v *DissociateInstanceWithPrivateDNSResponseBody) *DissociateInstanceWithPrivateDNSResponse
	GetBody() *DissociateInstanceWithPrivateDNSResponseBody
}

type DissociateInstanceWithPrivateDNSResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateInstanceWithPrivateDNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateInstanceWithPrivateDNSResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateInstanceWithPrivateDNSResponse) GoString() string {
	return s.String()
}

func (s *DissociateInstanceWithPrivateDNSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateInstanceWithPrivateDNSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateInstanceWithPrivateDNSResponse) GetBody() *DissociateInstanceWithPrivateDNSResponseBody {
	return s.Body
}

func (s *DissociateInstanceWithPrivateDNSResponse) SetHeaders(v map[string]*string) *DissociateInstanceWithPrivateDNSResponse {
	s.Headers = v
	return s
}

func (s *DissociateInstanceWithPrivateDNSResponse) SetStatusCode(v int32) *DissociateInstanceWithPrivateDNSResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateInstanceWithPrivateDNSResponse) SetBody(v *DissociateInstanceWithPrivateDNSResponseBody) *DissociateInstanceWithPrivateDNSResponse {
	s.Body = v
	return s
}

func (s *DissociateInstanceWithPrivateDNSResponse) Validate() error {
	return dara.Validate(s)
}
