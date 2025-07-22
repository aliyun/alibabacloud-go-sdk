// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInstanceDasConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DisableInstanceDasConfigRequest
	GetEngine() *string
	SetInstanceId(v string) *DisableInstanceDasConfigRequest
	GetInstanceId() *string
	SetScaleType(v string) *DisableInstanceDasConfigRequest
	GetScaleType() *string
}

type DisableInstanceDasConfigRequest struct {
	// The database engine. Set the value to Redis.
	//
	// This parameter is required.
	//
	// example:
	//
	// Redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1nti25tc7bq5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of auto scaling. Valid values:
	//
	// 	- **specScale**: The specifications of a database instance are automatically scaled up or down.
	//
	// 	- **shardScale**: The number of shards for a database instance is automatically increased or decreased.
	//
	// 	- **bandwidthScale**: The bandwidth of a database instance is automatically increased or decreased.
	//
	// This parameter is required.
	//
	// example:
	//
	// bandwidthScale
	ScaleType *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
}

func (s DisableInstanceDasConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableInstanceDasConfigRequest) GoString() string {
	return s.String()
}

func (s *DisableInstanceDasConfigRequest) GetEngine() *string {
	return s.Engine
}

func (s *DisableInstanceDasConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableInstanceDasConfigRequest) GetScaleType() *string {
	return s.ScaleType
}

func (s *DisableInstanceDasConfigRequest) SetEngine(v string) *DisableInstanceDasConfigRequest {
	s.Engine = &v
	return s
}

func (s *DisableInstanceDasConfigRequest) SetInstanceId(v string) *DisableInstanceDasConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableInstanceDasConfigRequest) SetScaleType(v string) *DisableInstanceDasConfigRequest {
	s.ScaleType = &v
	return s
}

func (s *DisableInstanceDasConfigRequest) Validate() error {
	return dara.Validate(s)
}
