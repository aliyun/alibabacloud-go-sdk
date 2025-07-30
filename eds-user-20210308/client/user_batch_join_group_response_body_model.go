// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserBatchJoinGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UserBatchJoinGroupResponseBody
	GetRequestId() *string
}

type UserBatchJoinGroupResponseBody struct {
	// example:
	//
	// 7A2C3803-C975-5871-A232-80A91009****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UserBatchJoinGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UserBatchJoinGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UserBatchJoinGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UserBatchJoinGroupResponseBody) SetRequestId(v string) *UserBatchJoinGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UserBatchJoinGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
