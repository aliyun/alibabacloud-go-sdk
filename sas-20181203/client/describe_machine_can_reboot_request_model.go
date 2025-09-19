// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMachineCanRebootRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *DescribeMachineCanRebootRequest
	GetType() *string
	SetUuid(v string) *DescribeMachineCanRebootRequest
	GetUuid() *string
}

type DescribeMachineCanRebootRequest struct {
	// The type of the vulnerabilities. Valid values:
	//
	// 	- cve: Linux software vulnerabilities
	//
	// 	- sys: Windows system vulnerabilities
	//
	// example:
	//
	// sys
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 5b268326-273e-44fc-a0e3-9482435c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeMachineCanRebootRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMachineCanRebootRequest) GoString() string {
	return s.String()
}

func (s *DescribeMachineCanRebootRequest) GetType() *string {
	return s.Type
}

func (s *DescribeMachineCanRebootRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeMachineCanRebootRequest) SetType(v string) *DescribeMachineCanRebootRequest {
	s.Type = &v
	return s
}

func (s *DescribeMachineCanRebootRequest) SetUuid(v string) *DescribeMachineCanRebootRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeMachineCanRebootRequest) Validate() error {
	return dara.Validate(s)
}
