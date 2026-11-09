// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *CreateGroupIdRequest
	GetGroupId() *string
	SetInstanceId(v string) *CreateGroupIdRequest
	GetInstanceId() *string
	SetTag(v []*CreateGroupIdRequestTag) *CreateGroupIdRequest
	GetTag() []*CreateGroupIdRequestTag
}

type CreateGroupIdRequest struct {
	// The group ID to be created. Valid values:
	//
	// - The value must start with `GID_` or `GID-` and can contain only letters, digits, hyphens (-), and underscores (_). The value must be 7 to 64 characters in length.
	//
	// - A group ID cannot be modified after it is created. For more information, see [Terms](https://help.aliyun.com/document_detail/42420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance where the group ID resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-0pp1ldu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The array of tag key-value pairs.
	Tag []*CreateGroupIdRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupIdRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupIdRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateGroupIdRequest) GetTag() []*CreateGroupIdRequestTag {
	return s.Tag
}

func (s *CreateGroupIdRequest) SetGroupId(v string) *CreateGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupIdRequest) SetInstanceId(v string) *CreateGroupIdRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateGroupIdRequest) SetTag(v []*CreateGroupIdRequestTag) *CreateGroupIdRequest {
	s.Tag = v
	return s
}

func (s *CreateGroupIdRequest) Validate() error {
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

type CreateGroupIdRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateGroupIdRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupIdRequestTag) GoString() string {
	return s.String()
}

func (s *CreateGroupIdRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateGroupIdRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateGroupIdRequestTag) SetKey(v string) *CreateGroupIdRequestTag {
	s.Key = &v
	return s
}

func (s *CreateGroupIdRequestTag) SetValue(v string) *CreateGroupIdRequestTag {
	s.Value = &v
	return s
}

func (s *CreateGroupIdRequestTag) Validate() error {
	return dara.Validate(s)
}
