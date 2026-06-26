// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeSaasServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumeSaasServiceResponseBody
	GetRequestId() *string
	SetServiceId(v string) *ResumeSaasServiceResponseBody
	GetServiceId() *string
}

type ResumeSaasServiceResponseBody struct {
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// drama-123456
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ResumeSaasServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeSaasServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeSaasServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeSaasServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *ResumeSaasServiceResponseBody) SetRequestId(v string) *ResumeSaasServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeSaasServiceResponseBody) SetServiceId(v string) *ResumeSaasServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *ResumeSaasServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
