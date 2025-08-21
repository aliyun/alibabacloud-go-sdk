// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyPrepayInstanceSpecRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ModifyPrepayInstanceSpecRequest
	GetInstanceType() *string
}

type ModifyPrepayInstanceSpecRequest struct {
	// The ID of the instance that you want to upgrade or downgrade.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The updated configuration.
	//
	// example:
	//
	// ens.sn1.tiny
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ModifyPrepayInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPrepayInstanceSpecRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyPrepayInstanceSpecRequest) SetInstanceId(v string) *ModifyPrepayInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetInstanceType(v string) *ModifyPrepayInstanceSpecRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) Validate() error {
	return dara.Validate(s)
}
