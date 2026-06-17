// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSaasServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSaasServiceResponseBody
	GetRequestId() *string
	SetServiceId(v string) *DeleteSaasServiceResponseBody
	GetServiceId() *string
}

type DeleteSaasServiceResponseBody struct {
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
	// agdb-xxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteSaasServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSaasServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSaasServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSaasServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *DeleteSaasServiceResponseBody) SetRequestId(v string) *DeleteSaasServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSaasServiceResponseBody) SetServiceId(v string) *DeleteSaasServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DeleteSaasServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
