// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAegisForLingjunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuids(v []*string) *InstallAegisForLingjunRequest
	GetUuids() []*string
}

type InstallAegisForLingjunRequest struct {
	// List of unique UUIDs for Lingjun bare metal.
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s InstallAegisForLingjunRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallAegisForLingjunRequest) GoString() string {
	return s.String()
}

func (s *InstallAegisForLingjunRequest) GetUuids() []*string {
	return s.Uuids
}

func (s *InstallAegisForLingjunRequest) SetUuids(v []*string) *InstallAegisForLingjunRequest {
	s.Uuids = v
	return s
}

func (s *InstallAegisForLingjunRequest) Validate() error {
	return dara.Validate(s)
}
