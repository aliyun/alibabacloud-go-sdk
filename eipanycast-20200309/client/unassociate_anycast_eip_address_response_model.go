// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateAnycastEipAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassociateAnycastEipAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassociateAnycastEipAddressResponse
	GetStatusCode() *int32
	SetBody(v *UnassociateAnycastEipAddressResponseBody) *UnassociateAnycastEipAddressResponse
	GetBody() *UnassociateAnycastEipAddressResponseBody
}

type UnassociateAnycastEipAddressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassociateAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassociateAnycastEipAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassociateAnycastEipAddressResponse) GoString() string {
	return s.String()
}

func (s *UnassociateAnycastEipAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassociateAnycastEipAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassociateAnycastEipAddressResponse) GetBody() *UnassociateAnycastEipAddressResponseBody {
	return s.Body
}

func (s *UnassociateAnycastEipAddressResponse) SetHeaders(v map[string]*string) *UnassociateAnycastEipAddressResponse {
	s.Headers = v
	return s
}

func (s *UnassociateAnycastEipAddressResponse) SetStatusCode(v int32) *UnassociateAnycastEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassociateAnycastEipAddressResponse) SetBody(v *UnassociateAnycastEipAddressResponseBody) *UnassociateAnycastEipAddressResponse {
	s.Body = v
	return s
}

func (s *UnassociateAnycastEipAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
