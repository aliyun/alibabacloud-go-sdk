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
	// The ID of the user group.
	//
	// example:
	//
	// ug-12341234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// AA8D67CB-345D-5CDA-986E-FFAC7D0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
