// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkPackageEnabledResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNetworkPackageEnabledResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNetworkPackageEnabledResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNetworkPackageEnabledResponseBody) *ModifyNetworkPackageEnabledResponse
	GetBody() *ModifyNetworkPackageEnabledResponseBody
}

type ModifyNetworkPackageEnabledResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNetworkPackageEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNetworkPackageEnabledResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkPackageEnabledResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageEnabledResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNetworkPackageEnabledResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNetworkPackageEnabledResponse) GetBody() *ModifyNetworkPackageEnabledResponseBody {
	return s.Body
}

func (s *ModifyNetworkPackageEnabledResponse) SetHeaders(v map[string]*string) *ModifyNetworkPackageEnabledResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkPackageEnabledResponse) SetStatusCode(v int32) *ModifyNetworkPackageEnabledResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNetworkPackageEnabledResponse) SetBody(v *ModifyNetworkPackageEnabledResponseBody) *ModifyNetworkPackageEnabledResponse {
	s.Body = v
	return s
}

func (s *ModifyNetworkPackageEnabledResponse) Validate() error {
	return dara.Validate(s)
}
