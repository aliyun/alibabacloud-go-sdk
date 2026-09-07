// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdUsersCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdUserCount(v int32) *GetAdUsersCountResponseBody
	GetAdUserCount() *int32
	SetRequestId(v string) *GetAdUsersCountResponseBody
	GetRequestId() *string
}

type GetAdUsersCountResponseBody struct {
	// example:
	//
	// 1000
	AdUserCount *int32 `json:"AdUserCount,omitempty" xml:"AdUserCount,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAdUsersCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAdUsersCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdUsersCountResponseBody) GetAdUserCount() *int32 {
	return s.AdUserCount
}

func (s *GetAdUsersCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAdUsersCountResponseBody) SetAdUserCount(v int32) *GetAdUsersCountResponseBody {
	s.AdUserCount = &v
	return s
}

func (s *GetAdUsersCountResponseBody) SetRequestId(v string) *GetAdUsersCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdUsersCountResponseBody) Validate() error {
	return dara.Validate(s)
}
