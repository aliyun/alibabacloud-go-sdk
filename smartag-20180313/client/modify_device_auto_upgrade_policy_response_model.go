// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceAutoUpgradePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDeviceAutoUpgradePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDeviceAutoUpgradePolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDeviceAutoUpgradePolicyResponseBody) *ModifyDeviceAutoUpgradePolicyResponse
	GetBody() *ModifyDeviceAutoUpgradePolicyResponseBody
}

type ModifyDeviceAutoUpgradePolicyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDeviceAutoUpgradePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDeviceAutoUpgradePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceAutoUpgradePolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceAutoUpgradePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDeviceAutoUpgradePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDeviceAutoUpgradePolicyResponse) GetBody() *ModifyDeviceAutoUpgradePolicyResponseBody {
	return s.Body
}

func (s *ModifyDeviceAutoUpgradePolicyResponse) SetHeaders(v map[string]*string) *ModifyDeviceAutoUpgradePolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyResponse) SetStatusCode(v int32) *ModifyDeviceAutoUpgradePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyResponse) SetBody(v *ModifyDeviceAutoUpgradePolicyResponseBody) *ModifyDeviceAutoUpgradePolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
