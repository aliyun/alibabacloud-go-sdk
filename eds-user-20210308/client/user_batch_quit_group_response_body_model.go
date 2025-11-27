// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserBatchQuitGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UserBatchQuitGroupResponseBody
	GetRequestId() *string
}

type UserBatchQuitGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 890JNJNF-SADASSDFS-SDFSDF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UserBatchQuitGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UserBatchQuitGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UserBatchQuitGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UserBatchQuitGroupResponseBody) SetRequestId(v string) *UserBatchQuitGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UserBatchQuitGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
