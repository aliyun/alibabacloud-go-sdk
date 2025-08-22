// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *CreateGroupResponseBody
	GetGroupId() *string
	SetRequestId(v string) *CreateGroupResponseBody
	GetRequestId() *string
}

type CreateGroupResponseBody struct {
	// example:
	//
	// g-4267dcfbf1b6d128c87adf0e95f
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// E2D0E863-1651-5E58-823F-B451C8C24615
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupResponseBody) SetGroupId(v string) *CreateGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
