// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInstanceDasConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableInstanceDasConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableInstanceDasConfigResponse
	GetStatusCode() *int32
	SetBody(v *DisableInstanceDasConfigResponseBody) *DisableInstanceDasConfigResponse
	GetBody() *DisableInstanceDasConfigResponseBody
}

type DisableInstanceDasConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableInstanceDasConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableInstanceDasConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableInstanceDasConfigResponse) GoString() string {
	return s.String()
}

func (s *DisableInstanceDasConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableInstanceDasConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableInstanceDasConfigResponse) GetBody() *DisableInstanceDasConfigResponseBody {
	return s.Body
}

func (s *DisableInstanceDasConfigResponse) SetHeaders(v map[string]*string) *DisableInstanceDasConfigResponse {
	s.Headers = v
	return s
}

func (s *DisableInstanceDasConfigResponse) SetStatusCode(v int32) *DisableInstanceDasConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableInstanceDasConfigResponse) SetBody(v *DisableInstanceDasConfigResponseBody) *DisableInstanceDasConfigResponse {
	s.Body = v
	return s
}

func (s *DisableInstanceDasConfigResponse) Validate() error {
	return dara.Validate(s)
}
