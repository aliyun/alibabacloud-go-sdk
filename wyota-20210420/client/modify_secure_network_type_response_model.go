// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecureNetworkTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecureNetworkTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecureNetworkTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecureNetworkTypeResponseBody) *ModifySecureNetworkTypeResponse
	GetBody() *ModifySecureNetworkTypeResponseBody
}

type ModifySecureNetworkTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecureNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecureNetworkTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecureNetworkTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifySecureNetworkTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecureNetworkTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecureNetworkTypeResponse) GetBody() *ModifySecureNetworkTypeResponseBody {
	return s.Body
}

func (s *ModifySecureNetworkTypeResponse) SetHeaders(v map[string]*string) *ModifySecureNetworkTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifySecureNetworkTypeResponse) SetStatusCode(v int32) *ModifySecureNetworkTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecureNetworkTypeResponse) SetBody(v *ModifySecureNetworkTypeResponseBody) *ModifySecureNetworkTypeResponse {
	s.Body = v
	return s
}

func (s *ModifySecureNetworkTypeResponse) Validate() error {
	return dara.Validate(s)
}
