// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockRefreshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebLockRefreshResponseBody
	GetRequestId() *string
}

type ModifyWebLockRefreshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CF8C834-8028-5E01-96E2-0F065EA99F6D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockRefreshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockRefreshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebLockRefreshResponseBody) SetRequestId(v string) *ModifyWebLockRefreshResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebLockRefreshResponseBody) Validate() error {
	return dara.Validate(s)
}
