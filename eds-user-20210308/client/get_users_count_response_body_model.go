// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUsersCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupCount(v int32) *GetUsersCountResponseBody
	GetGroupCount() *int32
	SetMaxUserNumber(v int32) *GetUsersCountResponseBody
	GetMaxUserNumber() *int32
	SetOrgCount(v int32) *GetUsersCountResponseBody
	GetOrgCount() *int32
	SetRequestId(v string) *GetUsersCountResponseBody
	GetRequestId() *string
	SetUserCount(v int32) *GetUsersCountResponseBody
	GetUserCount() *int32
}

type GetUsersCountResponseBody struct {
	// example:
	//
	// 0
	GroupCount *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// example:
	//
	// 10000
	MaxUserNumber *int32 `json:"MaxUserNumber,omitempty" xml:"MaxUserNumber,omitempty"`
	// example:
	//
	// 278
	OrgCount *int32 `json:"OrgCount,omitempty" xml:"OrgCount,omitempty"`
	// example:
	//
	// 9677D40F-0040-5956-A0EB-11B8B88****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s GetUsersCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUsersCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetUsersCountResponseBody) GetGroupCount() *int32 {
	return s.GroupCount
}

func (s *GetUsersCountResponseBody) GetMaxUserNumber() *int32 {
	return s.MaxUserNumber
}

func (s *GetUsersCountResponseBody) GetOrgCount() *int32 {
	return s.OrgCount
}

func (s *GetUsersCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUsersCountResponseBody) GetUserCount() *int32 {
	return s.UserCount
}

func (s *GetUsersCountResponseBody) SetGroupCount(v int32) *GetUsersCountResponseBody {
	s.GroupCount = &v
	return s
}

func (s *GetUsersCountResponseBody) SetMaxUserNumber(v int32) *GetUsersCountResponseBody {
	s.MaxUserNumber = &v
	return s
}

func (s *GetUsersCountResponseBody) SetOrgCount(v int32) *GetUsersCountResponseBody {
	s.OrgCount = &v
	return s
}

func (s *GetUsersCountResponseBody) SetRequestId(v string) *GetUsersCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUsersCountResponseBody) SetUserCount(v int32) *GetUsersCountResponseBody {
	s.UserCount = &v
	return s
}

func (s *GetUsersCountResponseBody) Validate() error {
	return dara.Validate(s)
}
