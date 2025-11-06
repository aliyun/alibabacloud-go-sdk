// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupDomainAutoRenewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SetupDomainAutoRenewRequest
	GetInstanceId() *string
	SetOperation(v string) *SetupDomainAutoRenewRequest
	GetOperation() *string
}

type SetupDomainAutoRenewRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// S2019270W570xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SET
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s SetupDomainAutoRenewRequest) String() string {
	return dara.Prettify(s)
}

func (s SetupDomainAutoRenewRequest) GoString() string {
	return s.String()
}

func (s *SetupDomainAutoRenewRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetupDomainAutoRenewRequest) GetOperation() *string {
	return s.Operation
}

func (s *SetupDomainAutoRenewRequest) SetInstanceId(v string) *SetupDomainAutoRenewRequest {
	s.InstanceId = &v
	return s
}

func (s *SetupDomainAutoRenewRequest) SetOperation(v string) *SetupDomainAutoRenewRequest {
	s.Operation = &v
	return s
}

func (s *SetupDomainAutoRenewRequest) Validate() error {
	return dara.Validate(s)
}
