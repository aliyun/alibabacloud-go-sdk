// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnAssociateEnsEipAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnAssociateEnsEipAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnAssociateEnsEipAddressResponse
	GetStatusCode() *int32
	SetBody(v *UnAssociateEnsEipAddressResponseBody) *UnAssociateEnsEipAddressResponse
	GetBody() *UnAssociateEnsEipAddressResponseBody
}

type UnAssociateEnsEipAddressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnAssociateEnsEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnAssociateEnsEipAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s UnAssociateEnsEipAddressResponse) GoString() string {
	return s.String()
}

func (s *UnAssociateEnsEipAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnAssociateEnsEipAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnAssociateEnsEipAddressResponse) GetBody() *UnAssociateEnsEipAddressResponseBody {
	return s.Body
}

func (s *UnAssociateEnsEipAddressResponse) SetHeaders(v map[string]*string) *UnAssociateEnsEipAddressResponse {
	s.Headers = v
	return s
}

func (s *UnAssociateEnsEipAddressResponse) SetStatusCode(v int32) *UnAssociateEnsEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *UnAssociateEnsEipAddressResponse) SetBody(v *UnAssociateEnsEipAddressResponseBody) *UnAssociateEnsEipAddressResponse {
	s.Body = v
	return s
}

func (s *UnAssociateEnsEipAddressResponse) Validate() error {
	return dara.Validate(s)
}
