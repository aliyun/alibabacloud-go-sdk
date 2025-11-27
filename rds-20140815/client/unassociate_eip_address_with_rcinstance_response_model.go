// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateEipAddressWithRCInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassociateEipAddressWithRCInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassociateEipAddressWithRCInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UnassociateEipAddressWithRCInstanceResponseBody) *UnassociateEipAddressWithRCInstanceResponse
	GetBody() *UnassociateEipAddressWithRCInstanceResponseBody
}

type UnassociateEipAddressWithRCInstanceResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassociateEipAddressWithRCInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassociateEipAddressWithRCInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassociateEipAddressWithRCInstanceResponse) GoString() string {
	return s.String()
}

func (s *UnassociateEipAddressWithRCInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassociateEipAddressWithRCInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassociateEipAddressWithRCInstanceResponse) GetBody() *UnassociateEipAddressWithRCInstanceResponseBody {
	return s.Body
}

func (s *UnassociateEipAddressWithRCInstanceResponse) SetHeaders(v map[string]*string) *UnassociateEipAddressWithRCInstanceResponse {
	s.Headers = v
	return s
}

func (s *UnassociateEipAddressWithRCInstanceResponse) SetStatusCode(v int32) *UnassociateEipAddressWithRCInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassociateEipAddressWithRCInstanceResponse) SetBody(v *UnassociateEipAddressWithRCInstanceResponseBody) *UnassociateEipAddressWithRCInstanceResponse {
	s.Body = v
	return s
}

func (s *UnassociateEipAddressWithRCInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
