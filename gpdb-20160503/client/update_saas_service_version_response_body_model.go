// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSaasServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateSaasServiceVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSaasServiceVersionResponseBody
	GetRequestId() *string
	SetServiceId(v string) *UpdateSaasServiceVersionResponseBody
	GetServiceId() *string
}

type UpdateSaasServiceVersionResponseBody struct {
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
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// drama-123456
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s UpdateSaasServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSaasServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSaasServiceVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSaasServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSaasServiceVersionResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateSaasServiceVersionResponseBody) SetMessage(v string) *UpdateSaasServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSaasServiceVersionResponseBody) SetRequestId(v string) *UpdateSaasServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSaasServiceVersionResponseBody) SetServiceId(v string) *UpdateSaasServiceVersionResponseBody {
	s.ServiceId = &v
	return s
}

func (s *UpdateSaasServiceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
