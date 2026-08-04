// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSaasServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanUpgrade(v bool) *CheckSaasServiceVersionResponseBody
	GetCanUpgrade() *bool
	SetMessage(v string) *CheckSaasServiceVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckSaasServiceVersionResponseBody
	GetRequestId() *string
	SetServiceId(v string) *CheckSaasServiceVersionResponseBody
	GetServiceId() *string
}

type CheckSaasServiceVersionResponseBody struct {
	// Indicates whether the service can be upgraded.
	//
	// example:
	//
	// true
	CanUpgrade *bool `json:"CanUpgrade,omitempty" xml:"CanUpgrade,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// agdb-xxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CheckSaasServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSaasServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSaasServiceVersionResponseBody) GetCanUpgrade() *bool {
	return s.CanUpgrade
}

func (s *CheckSaasServiceVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckSaasServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckSaasServiceVersionResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *CheckSaasServiceVersionResponseBody) SetCanUpgrade(v bool) *CheckSaasServiceVersionResponseBody {
	s.CanUpgrade = &v
	return s
}

func (s *CheckSaasServiceVersionResponseBody) SetMessage(v string) *CheckSaasServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CheckSaasServiceVersionResponseBody) SetRequestId(v string) *CheckSaasServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSaasServiceVersionResponseBody) SetServiceId(v string) *CheckSaasServiceVersionResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CheckSaasServiceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
