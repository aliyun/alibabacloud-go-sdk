// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSearchConditionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteSearchConditionRequest
	GetName() *string
	SetSourceIp(v string) *DeleteSearchConditionRequest
	GetSourceIp() *string
	SetType(v string) *DeleteSearchConditionRequest
	GetType() *string
}

type DeleteSearchConditionRequest struct {
	// The name of the frequently used search condition.
	//
	// >  You can call the [DescribeSearchCondition](~~DescribeSearchCondition~~) operation to query frequently used search conditions.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 19.12.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the asset. Default value: ecs. Valid values:
	//
	// 	- **ecs**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **cloud_product**: an Alibaba Cloud service.
	//
	// example:
	//
	// ecs
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteSearchConditionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSearchConditionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSearchConditionRequest) GetName() *string {
	return s.Name
}

func (s *DeleteSearchConditionRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteSearchConditionRequest) GetType() *string {
	return s.Type
}

func (s *DeleteSearchConditionRequest) SetName(v string) *DeleteSearchConditionRequest {
	s.Name = &v
	return s
}

func (s *DeleteSearchConditionRequest) SetSourceIp(v string) *DeleteSearchConditionRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteSearchConditionRequest) SetType(v string) *DeleteSearchConditionRequest {
	s.Type = &v
	return s
}

func (s *DeleteSearchConditionRequest) Validate() error {
	return dara.Validate(s)
}
