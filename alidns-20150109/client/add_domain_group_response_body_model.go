// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *AddDomainGroupResponseBody
	GetGroupId() *string
	SetGroupName(v string) *AddDomainGroupResponseBody
	GetGroupName() *string
	SetRequestId(v string) *AddDomainGroupResponseBody
	GetRequestId() *string
}

type AddDomainGroupResponseBody struct {
	// The ID of the domain name group.
	//
	// example:
	//
	// 2223
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the domain name group.
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

func (s AddDomainGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *AddDomainGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *AddDomainGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDomainGroupResponseBody) SetGroupId(v string) *AddDomainGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *AddDomainGroupResponseBody) SetGroupName(v string) *AddDomainGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *AddDomainGroupResponseBody) SetRequestId(v string) *AddDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDomainGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
