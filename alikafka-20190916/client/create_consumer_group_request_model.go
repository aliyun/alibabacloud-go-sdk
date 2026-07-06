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
	// Group name.
	//
	// - Reserved instance: Supports uppercase and lowercase letters, numbers, underscores (_), hyphens (-), and periods (.), limited to 3-64 characters.
	//
	// - Serverless instance: Can only contain letters, numbers, and special characters "@._\\*$#^!&-", limited to 1-249 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-0pp1l9z8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Tag list.
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateConsumerGroupRequestTag struct {
	// The tag key of the resource.
	//
	// - N ranges from 1 to 20.
	//
	// - Cannot be empty.
	//
	// - Supports up to 128 characters, cannot start with aliyun or acs:, and cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// - N ranges from 1 to 20.
	//
	// - Can be empty.
	//
	// - Supports up to 128 characters, cannot start with aliyun or acs:, and cannot contain `http://` or `https://`.
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
