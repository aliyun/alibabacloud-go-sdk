// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeAmount(v int32) *UpdateLogstashRequest
	GetNodeAmount() *int32
	SetNodeSpec(v *UpdateLogstashRequestNodeSpec) *UpdateLogstashRequest
	GetNodeSpec() *UpdateLogstashRequestNodeSpec
	SetClientToken(v string) *UpdateLogstashRequest
	GetClientToken() *string
}

type UpdateLogstashRequest struct {
	// example:
	//
	// 3
	NodeAmount *int32                         `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	NodeSpec   *UpdateLogstashRequestNodeSpec `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateLogstashRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogstashRequest) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *UpdateLogstashRequest) GetNodeSpec() *UpdateLogstashRequestNodeSpec {
	return s.NodeSpec
}

func (s *UpdateLogstashRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateLogstashRequest) SetNodeAmount(v int32) *UpdateLogstashRequest {
	s.NodeAmount = &v
	return s
}

func (s *UpdateLogstashRequest) SetNodeSpec(v *UpdateLogstashRequestNodeSpec) *UpdateLogstashRequest {
	s.NodeSpec = v
	return s
}

func (s *UpdateLogstashRequest) SetClientToken(v string) *UpdateLogstashRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLogstashRequest) Validate() error {
	if s.NodeSpec != nil {
		if err := s.NodeSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLogstashRequestNodeSpec struct {
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// example:
	//
	// cloud_efficiency
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// elasticsearch.sn1ne.large
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s UpdateLogstashRequestNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashRequestNodeSpec) GoString() string {
	return s.String()
}

func (s *UpdateLogstashRequestNodeSpec) GetDisk() *int32 {
	return s.Disk
}

func (s *UpdateLogstashRequestNodeSpec) GetDiskType() *string {
	return s.DiskType
}

func (s *UpdateLogstashRequestNodeSpec) GetSpec() *string {
	return s.Spec
}

func (s *UpdateLogstashRequestNodeSpec) SetDisk(v int32) *UpdateLogstashRequestNodeSpec {
	s.Disk = &v
	return s
}

func (s *UpdateLogstashRequestNodeSpec) SetDiskType(v string) *UpdateLogstashRequestNodeSpec {
	s.DiskType = &v
	return s
}

func (s *UpdateLogstashRequestNodeSpec) SetSpec(v string) *UpdateLogstashRequestNodeSpec {
	s.Spec = &v
	return s
}

func (s *UpdateLogstashRequestNodeSpec) Validate() error {
	return dara.Validate(s)
}
