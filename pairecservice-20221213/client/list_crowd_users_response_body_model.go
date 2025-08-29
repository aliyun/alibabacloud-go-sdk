// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrowdUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListCrowdUsersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCrowdUsersResponseBody
	GetTotalCount() *int64
	SetUsers(v []*string) *ListCrowdUsersResponseBody
	GetUsers() []*string
}

type ListCrowdUsersResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F0AB6527-093F-5C44-B3BD-42C8C210C619
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int64    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users      []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListCrowdUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCrowdUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrowdUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrowdUsersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCrowdUsersResponseBody) GetUsers() []*string {
	return s.Users
}

func (s *ListCrowdUsersResponseBody) SetRequestId(v string) *ListCrowdUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCrowdUsersResponseBody) SetTotalCount(v int64) *ListCrowdUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCrowdUsersResponseBody) SetUsers(v []*string) *ListCrowdUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListCrowdUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
