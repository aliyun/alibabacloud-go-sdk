// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteInstanceRequest
	GetInstanceId() *string
	SetTag(v []*DeleteInstanceRequestTag) *DeleteInstanceRequest
	GetTag() []*DeleteInstanceRequestTag
}

type DeleteInstanceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-sh-ae502ee79ef8
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The tag of objects that match the lifecycle rule. You can specify multiple tags.
	//
	// example:
	//
	// Key， Value
	Tag []*DeleteInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceRequest) GetTag() []*DeleteInstanceRequestTag {
	return s.Tag
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetTag(v []*DeleteInstanceRequestTag) *DeleteInstanceRequest {
	s.Tag = v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
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

type DeleteInstanceRequestTag struct {
	// The key of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeleteInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *DeleteInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *DeleteInstanceRequestTag) SetKey(v string) *DeleteInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *DeleteInstanceRequestTag) SetValue(v string) *DeleteInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *DeleteInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
