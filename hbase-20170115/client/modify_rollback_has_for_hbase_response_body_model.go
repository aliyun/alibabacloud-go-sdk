// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRollbackHasForHbaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRollbackHasForHbaseResponseBody
	GetRequestId() *string
}

type ModifyRollbackHasForHbaseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRollbackHasForHbaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRollbackHasForHbaseResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRollbackHasForHbaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRollbackHasForHbaseResponseBody) SetRequestId(v string) *ModifyRollbackHasForHbaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRollbackHasForHbaseResponseBody) Validate() error {
	return dara.Validate(s)
}
