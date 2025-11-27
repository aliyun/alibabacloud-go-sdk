// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEipAddressWithRCInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateEipAddressWithRCInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateEipAddressWithRCInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AssociateEipAddressWithRCInstanceResponseBody) *AssociateEipAddressWithRCInstanceResponse
	GetBody() *AssociateEipAddressWithRCInstanceResponseBody
}

type AssociateEipAddressWithRCInstanceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateEipAddressWithRCInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateEipAddressWithRCInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateEipAddressWithRCInstanceResponse) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressWithRCInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateEipAddressWithRCInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateEipAddressWithRCInstanceResponse) GetBody() *AssociateEipAddressWithRCInstanceResponseBody {
	return s.Body
}

func (s *AssociateEipAddressWithRCInstanceResponse) SetHeaders(v map[string]*string) *AssociateEipAddressWithRCInstanceResponse {
	s.Headers = v
	return s
}

func (s *AssociateEipAddressWithRCInstanceResponse) SetStatusCode(v int32) *AssociateEipAddressWithRCInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateEipAddressWithRCInstanceResponse) SetBody(v *AssociateEipAddressWithRCInstanceResponseBody) *AssociateEipAddressWithRCInstanceResponse {
	s.Body = v
	return s
}

func (s *AssociateEipAddressWithRCInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
