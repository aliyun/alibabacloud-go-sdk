// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListGroupIdRequest
	GetInstanceId() *string
	SetTag(v []*ListGroupIdRequestTag) *ListGroupIdRequest
	GetTag() []*ListGroupIdRequestTag
}

type ListGroupIdRequest struct {
	// The ID of the instance for which you want to list group IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-0pp1ldu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of tags.
	Tag []*ListGroupIdRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupIdRequest) GoString() string {
	return s.String()
}

func (s *ListGroupIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupIdRequest) GetTag() []*ListGroupIdRequestTag {
	return s.Tag
}

func (s *ListGroupIdRequest) SetInstanceId(v string) *ListGroupIdRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGroupIdRequest) SetTag(v []*ListGroupIdRequestTag) *ListGroupIdRequest {
	s.Tag = v
	return s
}

func (s *ListGroupIdRequest) Validate() error {
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

type ListGroupIdRequestTag struct {
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

func (s ListGroupIdRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListGroupIdRequestTag) GoString() string {
	return s.String()
}

func (s *ListGroupIdRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListGroupIdRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListGroupIdRequestTag) SetKey(v string) *ListGroupIdRequestTag {
	s.Key = &v
	return s
}

func (s *ListGroupIdRequestTag) SetValue(v string) *ListGroupIdRequestTag {
	s.Value = &v
	return s
}

func (s *ListGroupIdRequestTag) Validate() error {
	return dara.Validate(s)
}
