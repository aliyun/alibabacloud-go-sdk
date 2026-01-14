// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *UpdateDomainStateResponseBody
	GetDomain() *string
	SetRequestId(v string) *UpdateDomainStateResponseBody
	GetRequestId() *string
	SetState(v string) *UpdateDomainStateResponseBody
	GetState() *string
}

type UpdateDomainStateResponseBody struct {
	// The accelerated domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ICP filing status of the accelerated domain name. Valid values:
	//
	// 	- **illegal:*	- The domain name is illegal.
	//
	// 	- **inactive:*	- The domain name has not completed ICP filing.
	//
	// 	- **active:*	- The domain name has a valid ICP number.
	//
	// 	- **unknown:*	- The ICP filing status is unknown.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s UpdateDomainStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainStateResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *UpdateDomainStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDomainStateResponseBody) GetState() *string {
	return s.State
}

func (s *UpdateDomainStateResponseBody) SetDomain(v string) *UpdateDomainStateResponseBody {
	s.Domain = &v
	return s
}

func (s *UpdateDomainStateResponseBody) SetRequestId(v string) *UpdateDomainStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainStateResponseBody) SetState(v string) *UpdateDomainStateResponseBody {
	s.State = &v
	return s
}

func (s *UpdateDomainStateResponseBody) Validate() error {
	return dara.Validate(s)
}
