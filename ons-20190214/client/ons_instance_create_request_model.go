// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *OnsInstanceCreateRequest
	GetInstanceName() *string
	SetRemark(v string) *OnsInstanceCreateRequest
	GetRemark() *string
}

type OnsInstanceCreateRequest struct {
	// The name of the instance. The name must meet the following rules:
	//
	// 	- The name of the instance must be unique in the region where the instance is deployed.
	//
	// 	- The name must be 3 to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// Test instance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// Description
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s OnsInstanceCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceCreateRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceCreateRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *OnsInstanceCreateRequest) GetRemark() *string {
	return s.Remark
}

func (s *OnsInstanceCreateRequest) SetInstanceName(v string) *OnsInstanceCreateRequest {
	s.InstanceName = &v
	return s
}

func (s *OnsInstanceCreateRequest) SetRemark(v string) *OnsInstanceCreateRequest {
	s.Remark = &v
	return s
}

func (s *OnsInstanceCreateRequest) Validate() error {
	return dara.Validate(s)
}
