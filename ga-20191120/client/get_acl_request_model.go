// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *GetAclRequest
	GetAclId() *string
	SetRegionId(v string) *GetAclRequest
	GetRegionId() *string
}

type GetAclRequest struct {
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-hp34s2h0xx1ht4nwo****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The operation that you want to perform. Set the value to **GetAcl**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAclRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAclRequest) GoString() string {
	return s.String()
}

func (s *GetAclRequest) GetAclId() *string {
	return s.AclId
}

func (s *GetAclRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAclRequest) SetAclId(v string) *GetAclRequest {
	s.AclId = &v
	return s
}

func (s *GetAclRequest) SetRegionId(v string) *GetAclRequest {
	s.RegionId = &v
	return s
}

func (s *GetAclRequest) Validate() error {
	return dara.Validate(s)
}
