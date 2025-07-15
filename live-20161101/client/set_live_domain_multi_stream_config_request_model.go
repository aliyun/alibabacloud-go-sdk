// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainMultiStreamConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *SetLiveDomainMultiStreamConfigRequest
	GetDomain() *string
	SetOwnerId(v int64) *SetLiveDomainMultiStreamConfigRequest
	GetOwnerId() *int64
	SetSwitch(v string) *SetLiveDomainMultiStreamConfigRequest
	GetSwitch() *string
}

type SetLiveDomainMultiStreamConfigRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable the dual-stream disaster recovery feature. Valid values:
	//
	// 	- **on**: enables the feature.
	//
	// 	- **off**: disables the feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	Switch *string `json:"Switch,omitempty" xml:"Switch,omitempty"`
}

func (s SetLiveDomainMultiStreamConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainMultiStreamConfigRequest) GoString() string {
	return s.String()
}

func (s *SetLiveDomainMultiStreamConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *SetLiveDomainMultiStreamConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLiveDomainMultiStreamConfigRequest) GetSwitch() *string {
	return s.Switch
}

func (s *SetLiveDomainMultiStreamConfigRequest) SetDomain(v string) *SetLiveDomainMultiStreamConfigRequest {
	s.Domain = &v
	return s
}

func (s *SetLiveDomainMultiStreamConfigRequest) SetOwnerId(v int64) *SetLiveDomainMultiStreamConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLiveDomainMultiStreamConfigRequest) SetSwitch(v string) *SetLiveDomainMultiStreamConfigRequest {
	s.Switch = &v
	return s
}

func (s *SetLiveDomainMultiStreamConfigRequest) Validate() error {
	return dara.Validate(s)
}
