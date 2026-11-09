// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteGroupIdRequest
	GetGroupId() *string
	SetInstanceId(v string) *DeleteGroupIdRequest
	GetInstanceId() *string
}

type DeleteGroupIdRequest struct {
	// The group ID that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the group ID that you want to delete belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-0pp1ldu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupIdRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteGroupIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteGroupIdRequest) SetGroupId(v string) *DeleteGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteGroupIdRequest) SetInstanceId(v string) *DeleteGroupIdRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteGroupIdRequest) Validate() error {
	return dara.Validate(s)
}
