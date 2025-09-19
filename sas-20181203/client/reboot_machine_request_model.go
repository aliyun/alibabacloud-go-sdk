// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuid(v string) *RebootMachineRequest
	GetUuid() *string
}

type RebootMachineRequest struct {
	// The UUID of the server that you want to restart.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7151f27e-1d51-4e98-a540-8936a****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RebootMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootMachineRequest) GoString() string {
	return s.String()
}

func (s *RebootMachineRequest) GetUuid() *string {
	return s.Uuid
}

func (s *RebootMachineRequest) SetUuid(v string) *RebootMachineRequest {
	s.Uuid = &v
	return s
}

func (s *RebootMachineRequest) Validate() error {
	return dara.Validate(s)
}
