// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDomainGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ChangeDomainGroupResponseBody
	GetGroupId() *string
	SetGroupName(v string) *ChangeDomainGroupResponseBody
	GetGroupName() *string
	SetRequestId(v string) *ChangeDomainGroupResponseBody
	GetRequestId() *string
}

type ChangeDomainGroupResponseBody struct {
	// The ID of the target domain name group.
	//
	// example:
	//
	// 2223
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the destination domain name group.
	//
	// example:
	//
	// MyGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeDomainGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDomainGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *ChangeDomainGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *ChangeDomainGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeDomainGroupResponseBody) SetGroupId(v string) *ChangeDomainGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *ChangeDomainGroupResponseBody) SetGroupName(v string) *ChangeDomainGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *ChangeDomainGroupResponseBody) SetRequestId(v string) *ChangeDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDomainGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
