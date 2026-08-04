// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySaasServiceDeletionProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifySaasServiceDeletionProtectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySaasServiceDeletionProtectionResponseBody
	GetRequestId() *string
	SetServiceId(v string) *ModifySaasServiceDeletionProtectionResponseBody
	GetServiceId() *string
}

type ModifySaasServiceDeletionProtectionResponseBody struct {
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7565770E-7C45-462D-BA4A-8A5396F2CAD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// agdb-xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ModifySaasServiceDeletionProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySaasServiceDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySaasServiceDeletionProtectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySaasServiceDeletionProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySaasServiceDeletionProtectionResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *ModifySaasServiceDeletionProtectionResponseBody) SetMessage(v string) *ModifySaasServiceDeletionProtectionResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySaasServiceDeletionProtectionResponseBody) SetRequestId(v string) *ModifySaasServiceDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySaasServiceDeletionProtectionResponseBody) SetServiceId(v string) *ModifySaasServiceDeletionProtectionResponseBody {
	s.ServiceId = &v
	return s
}

func (s *ModifySaasServiceDeletionProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
