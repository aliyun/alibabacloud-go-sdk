// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressPoolLbStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmAddressPoolLbStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmAddressPoolLbStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmAddressPoolLbStrategyResponseBody) *UpdateCloudGtmAddressPoolLbStrategyResponse
	GetBody() *UpdateCloudGtmAddressPoolLbStrategyResponseBody
}

type UpdateCloudGtmAddressPoolLbStrategyResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmAddressPoolLbStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmAddressPoolLbStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressPoolLbStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponse) GetBody() *UpdateCloudGtmAddressPoolLbStrategyResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmAddressPoolLbStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponse) SetStatusCode(v int32) *UpdateCloudGtmAddressPoolLbStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponse) SetBody(v *UpdateCloudGtmAddressPoolLbStrategyResponseBody) *UpdateCloudGtmAddressPoolLbStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
