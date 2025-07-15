// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlreadyDelete(v bool) *CreateLiveMessageGroupResponseBody
	GetAlreadyDelete() *bool
	SetAlreadyExists(v bool) *CreateLiveMessageGroupResponseBody
	GetAlreadyExists() *bool
	SetGroupId(v string) *CreateLiveMessageGroupResponseBody
	GetGroupId() *string
	SetRequestId(v string) *CreateLiveMessageGroupResponseBody
	GetRequestId() *string
}

type CreateLiveMessageGroupResponseBody struct {
	// Indicates whether the group is deleted. If the group existed and is deleted, the group ID is unavailable. We recommend that you create a new group.
	//
	// example:
	//
	// true
	AlreadyDelete *bool `json:"AlreadyDelete,omitempty" xml:"AlreadyDelete,omitempty"`
	// Indicates whether the group already exists.
	//
	// example:
	//
	// true
	AlreadyExists *bool `json:"AlreadyExists,omitempty" xml:"AlreadyExists,omitempty"`
	// The ID of the group created.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A8C7B033-B339-1A58-B0E0-7B9197BA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLiveMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveMessageGroupResponseBody) GetAlreadyDelete() *bool {
	return s.AlreadyDelete
}

func (s *CreateLiveMessageGroupResponseBody) GetAlreadyExists() *bool {
	return s.AlreadyExists
}

func (s *CreateLiveMessageGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateLiveMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLiveMessageGroupResponseBody) SetAlreadyDelete(v bool) *CreateLiveMessageGroupResponseBody {
	s.AlreadyDelete = &v
	return s
}

func (s *CreateLiveMessageGroupResponseBody) SetAlreadyExists(v bool) *CreateLiveMessageGroupResponseBody {
	s.AlreadyExists = &v
	return s
}

func (s *CreateLiveMessageGroupResponseBody) SetGroupId(v string) *CreateLiveMessageGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateLiveMessageGroupResponseBody) SetRequestId(v string) *CreateLiveMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveMessageGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
