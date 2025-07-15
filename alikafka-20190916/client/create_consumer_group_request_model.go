// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *CreateConsumerGroupRequest
	GetConsumerId() *string
	SetInstanceId(v string) *CreateConsumerGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateConsumerGroupRequest
	GetRegionId() *string
	SetRemark(v string) *CreateConsumerGroupRequest
	GetRemark() *string
	SetTag(v []*CreateConsumerGroupRequestTag) *CreateConsumerGroupRequest
	GetTag() []*CreateConsumerGroupRequestTag
}

type CreateConsumerGroupRequest struct {
	// The name of the consumer group.
	//
	// 	- The value can contain only letters, digits, hyphens (-), and underscores (_), and the value must contain at least one letter or digit.
	//
	// 	- The value must be 3 to 128 characters in length. If the value that you specify contains more than 128 characters, the system automatically truncates the value to 128 characters.
	//
	// 	- After a consumer group is created, you cannot change the name of the consumer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-0pp1l9z8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the consumer group.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The tags.
	Tag []*CreateConsumerGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *CreateConsumerGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateConsumerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateConsumerGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateConsumerGroupRequest) GetTag() []*CreateConsumerGroupRequestTag {
	return s.Tag
}

func (s *CreateConsumerGroupRequest) SetConsumerId(v string) *CreateConsumerGroupRequest {
	s.ConsumerId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetInstanceId(v string) *CreateConsumerGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetRegionId(v string) *CreateConsumerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetRemark(v string) *CreateConsumerGroupRequest {
	s.Remark = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetTag(v []*CreateConsumerGroupRequestTag) *CreateConsumerGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateConsumerGroupRequest) Validate() error {
	return dara.Validate(s)
}

type CreateConsumerGroupRequestTag struct {
	// The tag key.
	//
	// 	- You must specify this parameter.
	//
	// 	- The tag key can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// 	- You can leave this parameter empty.
	//
	// 	- The tag value can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConsumerGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateConsumerGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateConsumerGroupRequestTag) SetKey(v string) *CreateConsumerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateConsumerGroupRequestTag) SetValue(v string) *CreateConsumerGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateConsumerGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
