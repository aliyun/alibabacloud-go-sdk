// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactBlockListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactBlockListId(v string) *DeleteContactBlockListRequest
	GetContactBlockListId() *string
	SetInstanceId(v string) *DeleteContactBlockListRequest
	GetInstanceId() *string
	SetOperator(v string) *DeleteContactBlockListRequest
	GetOperator() *string
}

type DeleteContactBlockListRequest struct {
	// Contact list ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c6320d3c-fa45-4011-b3b1-acdfabe3a8c6
	ContactBlockListId *string `json:"ContactBlockListId,omitempty" xml:"ContactBlockListId,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c6320d3c-fa45-4011-b3b1-acdfabe3a8c6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Operator information
	//
	// example:
	//
	// 160131
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s DeleteContactBlockListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactBlockListRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactBlockListRequest) GetContactBlockListId() *string {
	return s.ContactBlockListId
}

func (s *DeleteContactBlockListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteContactBlockListRequest) GetOperator() *string {
	return s.Operator
}

func (s *DeleteContactBlockListRequest) SetContactBlockListId(v string) *DeleteContactBlockListRequest {
	s.ContactBlockListId = &v
	return s
}

func (s *DeleteContactBlockListRequest) SetInstanceId(v string) *DeleteContactBlockListRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteContactBlockListRequest) SetOperator(v string) *DeleteContactBlockListRequest {
	s.Operator = &v
	return s
}

func (s *DeleteContactBlockListRequest) Validate() error {
	return dara.Validate(s)
}
