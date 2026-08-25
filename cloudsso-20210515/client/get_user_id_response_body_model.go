// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserIdResponseBody
	GetRequestId() *string
	SetUserId(v string) *GetUserIdResponseBody
	GetUserId() *string
}

type GetUserIdResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A3A41736-A050-50B6-ABC5-590F376A0044
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the CloudSSO user.
	//
	// example:
	//
	// u-d8d1iox****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserIdResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetUserIdResponseBody) SetRequestId(v string) *GetUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserIdResponseBody) SetUserId(v string) *GetUserIdResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserIdResponseBody) Validate() error {
	return dara.Validate(s)
}
