// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsGroupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *OnsGroupListRequest
	GetGroupId() *string
	SetGroupType(v string) *OnsGroupListRequest
	GetGroupType() *string
	SetInstanceId(v string) *OnsGroupListRequest
	GetInstanceId() *string
	SetTag(v []*OnsGroupListRequestTag) *OnsGroupListRequest
	GetTag() []*OnsGroupListRequestTag
}

type OnsGroupListRequest struct {
	// This parameter is required only when you query specific consumer groups by using the fuzzy search method. If this parameter is not configured, the system queries all consumer groups that can be accessed by the current account.
	//
	// If you set this parameter to GID_ABC, the information about the consumer groups whose IDs contain GID_ABC is returned. For example, the information about the GID_test_GID_ABC_123 and GID_ABC_356 consumer groups is returned.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The protocol over which the queried consumer group publishes and subscribes to messages. All clients in a consumer group communicate with the ApsaraMQ for RocketMQ broker over the same protocol. You must create different consumer groups for TCP clients and HTTP clients. Valid values:
	//
	// 	- **tcp**: specifies that the consumer group publishes or subscribes to messages over TCP. This value is the default value.
	//
	// 	- **http**: specifies that the consumer group publishes or subscribes to messages over HTTP.
	//
	// example:
	//
	// tcp
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The ID of the instance to which the consumer group you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of tags that are attached to the consumer group. A maximum of 20 tags can be included in the list.
	Tag []*OnsGroupListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsGroupListRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupListRequest) GoString() string {
	return s.String()
}

func (s *OnsGroupListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsGroupListRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *OnsGroupListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsGroupListRequest) GetTag() []*OnsGroupListRequestTag {
	return s.Tag
}

func (s *OnsGroupListRequest) SetGroupId(v string) *OnsGroupListRequest {
	s.GroupId = &v
	return s
}

func (s *OnsGroupListRequest) SetGroupType(v string) *OnsGroupListRequest {
	s.GroupType = &v
	return s
}

func (s *OnsGroupListRequest) SetInstanceId(v string) *OnsGroupListRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsGroupListRequest) SetTag(v []*OnsGroupListRequestTag) *OnsGroupListRequest {
	s.Tag = v
	return s
}

func (s *OnsGroupListRequest) Validate() error {
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

type OnsGroupListRequestTag struct {
	// The key of the tag that is attached to the consumer group. This parameter is not required. If you configure this parameter, you must configure the **Key*	- parameter.***	- If you configure both the Key and Value parameters, the consumer groups are filtered based on the specified tag. If you do not configure these parameters, all consumer groups are queried.
	//
	// 	- The value of this parameter cannot be an empty string.
	//
	// 	- The tag value must be 1 to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// CartService
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is attached to the group. This parameter is not required. If you configure this parameter, you must configure the **Key*	- parameter.***	- If you configure both the Key and Value parameters, the consumer groups are filtered based on the specified tag. If you do not configure these parameters, all consumer groups are queried.
	//
	// 	- The value of this parameter can be an empty string.
	//
	// 	- The tag key must be 1 to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ServiceA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsGroupListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s OnsGroupListRequestTag) GoString() string {
	return s.String()
}

func (s *OnsGroupListRequestTag) GetKey() *string {
	return s.Key
}

func (s *OnsGroupListRequestTag) GetValue() *string {
	return s.Value
}

func (s *OnsGroupListRequestTag) SetKey(v string) *OnsGroupListRequestTag {
	s.Key = &v
	return s
}

func (s *OnsGroupListRequestTag) SetValue(v string) *OnsGroupListRequestTag {
	s.Value = &v
	return s
}

func (s *OnsGroupListRequestTag) Validate() error {
	return dara.Validate(s)
}
