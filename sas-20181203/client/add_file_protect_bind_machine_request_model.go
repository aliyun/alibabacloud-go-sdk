// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFileProtectBindMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertUuids(v []*string) *AddFileProtectBindMachineRequest
	GetAlertUuids() []*string
	SetBlockUuids(v []*string) *AddFileProtectBindMachineRequest
	GetBlockUuids() []*string
	SetNoneUuids(v []*string) *AddFileProtectBindMachineRequest
	GetNoneUuids() []*string
}

type AddFileProtectBindMachineRequest struct {
	AlertUuids []*string `json:"AlertUuids,omitempty" xml:"AlertUuids,omitempty" type:"Repeated"`
	BlockUuids []*string `json:"BlockUuids,omitempty" xml:"BlockUuids,omitempty" type:"Repeated"`
	NoneUuids  []*string `json:"NoneUuids,omitempty" xml:"NoneUuids,omitempty" type:"Repeated"`
}

func (s AddFileProtectBindMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFileProtectBindMachineRequest) GoString() string {
	return s.String()
}

func (s *AddFileProtectBindMachineRequest) GetAlertUuids() []*string {
	return s.AlertUuids
}

func (s *AddFileProtectBindMachineRequest) GetBlockUuids() []*string {
	return s.BlockUuids
}

func (s *AddFileProtectBindMachineRequest) GetNoneUuids() []*string {
	return s.NoneUuids
}

func (s *AddFileProtectBindMachineRequest) SetAlertUuids(v []*string) *AddFileProtectBindMachineRequest {
	s.AlertUuids = v
	return s
}

func (s *AddFileProtectBindMachineRequest) SetBlockUuids(v []*string) *AddFileProtectBindMachineRequest {
	s.BlockUuids = v
	return s
}

func (s *AddFileProtectBindMachineRequest) SetNoneUuids(v []*string) *AddFileProtectBindMachineRequest {
	s.NoneUuids = v
	return s
}

func (s *AddFileProtectBindMachineRequest) Validate() error {
	return dara.Validate(s)
}
