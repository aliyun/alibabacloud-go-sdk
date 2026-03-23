// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReadWriteSplittingConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyReadWriteSplittingConnectionResponseBody
	GetRequestId() *string
}

type ModifyReadWriteSplittingConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyReadWriteSplittingConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyReadWriteSplittingConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyReadWriteSplittingConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyReadWriteSplittingConnectionResponseBody) SetRequestId(v string) *ModifyReadWriteSplittingConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
