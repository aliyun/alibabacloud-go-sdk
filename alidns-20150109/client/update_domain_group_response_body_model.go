// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *UpdateDomainGroupResponseBody
	GetGroupId() *string
	SetGroupName(v string) *UpdateDomainGroupResponseBody
	GetGroupName() *string
	SetRequestId(v string) *UpdateDomainGroupResponseBody
	GetRequestId() *string
}

type UpdateDomainGroupResponseBody struct {
	// The ID of the domain name group.
	//
	// example:
	//
	// 2223
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The new name of the domain name group.
	//
	// example:
	//
	// NewName
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDomainGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateDomainGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateDomainGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDomainGroupResponseBody) SetGroupId(v string) *UpdateDomainGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *UpdateDomainGroupResponseBody) SetGroupName(v string) *UpdateDomainGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *UpdateDomainGroupResponseBody) SetRequestId(v string) *UpdateDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
