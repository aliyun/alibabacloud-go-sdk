// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateInstanceWithPrivateDNSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateInstanceWithPrivateDNSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateInstanceWithPrivateDNSResponse
	GetStatusCode() *int32
	SetBody(v *AssociateInstanceWithPrivateDNSResponseBody) *AssociateInstanceWithPrivateDNSResponse
	GetBody() *AssociateInstanceWithPrivateDNSResponseBody
}

type AssociateInstanceWithPrivateDNSResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateInstanceWithPrivateDNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateInstanceWithPrivateDNSResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateInstanceWithPrivateDNSResponse) GoString() string {
	return s.String()
}

func (s *AssociateInstanceWithPrivateDNSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateInstanceWithPrivateDNSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateInstanceWithPrivateDNSResponse) GetBody() *AssociateInstanceWithPrivateDNSResponseBody {
	return s.Body
}

func (s *AssociateInstanceWithPrivateDNSResponse) SetHeaders(v map[string]*string) *AssociateInstanceWithPrivateDNSResponse {
	s.Headers = v
	return s
}

func (s *AssociateInstanceWithPrivateDNSResponse) SetStatusCode(v int32) *AssociateInstanceWithPrivateDNSResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateInstanceWithPrivateDNSResponse) SetBody(v *AssociateInstanceWithPrivateDNSResponseBody) *AssociateInstanceWithPrivateDNSResponse {
	s.Body = v
	return s
}

func (s *AssociateInstanceWithPrivateDNSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
