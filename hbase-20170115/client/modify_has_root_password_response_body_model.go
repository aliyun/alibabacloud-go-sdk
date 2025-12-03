// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHasRootPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHasRootPasswordResponseBody
	GetRequestId() *string
}

type ModifyHasRootPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHasRootPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHasRootPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHasRootPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHasRootPasswordResponseBody) SetRequestId(v string) *ModifyHasRootPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHasRootPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
