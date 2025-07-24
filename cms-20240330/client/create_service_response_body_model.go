// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPid(v string) *CreateServiceResponseBody
	GetPid() *string
	SetRequestId(v string) *CreateServiceResponseBody
	GetRequestId() *string
	SetServiceId(v string) *CreateServiceResponseBody
	GetServiceId() *string
}

type CreateServiceResponseBody struct {
	// example:
	//
	// cwzxvuc6uo@d60088ad4797d26
	Pid *string `json:"pid,omitempty" xml:"pid,omitempty"`
	// example:
	//
	// 3A2FA9E9-9CF1-5CB1-A808-52828F14310D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// cwzxvuc6uo@4bc6b15ad81f166174ffb
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) GetPid() *string {
	return s.Pid
}

func (s *CreateServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceResponseBody) SetPid(v string) *CreateServiceResponseBody {
	s.Pid = &v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceId(v string) *CreateServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
