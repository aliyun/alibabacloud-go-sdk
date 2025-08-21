// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigDomainSecurityProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigDomainSecurityProfileResponseBody
	GetRequestId() *string
}

type ConfigDomainSecurityProfileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9728769F-9466-534E-BE12-CAB29A675828
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigDomainSecurityProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigDomainSecurityProfileResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigDomainSecurityProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigDomainSecurityProfileResponseBody) SetRequestId(v string) *ConfigDomainSecurityProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigDomainSecurityProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
