// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseSaasServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PauseSaasServiceResponseBody
	GetRequestId() *string
	SetServiceId(v string) *PauseSaasServiceResponseBody
	GetServiceId() *string
}

type PauseSaasServiceResponseBody struct {
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// drama-123456
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s PauseSaasServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseSaasServiceResponseBody) GoString() string {
	return s.String()
}

func (s *PauseSaasServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseSaasServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *PauseSaasServiceResponseBody) SetRequestId(v string) *PauseSaasServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseSaasServiceResponseBody) SetServiceId(v string) *PauseSaasServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *PauseSaasServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
