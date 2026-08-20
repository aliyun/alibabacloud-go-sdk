// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccountParameterResponseBody
	GetRequestId() *string
}

type ModifyAccountParameterResponseBody struct {
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D190D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountParameterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountParameterResponseBody) SetRequestId(v string) *ModifyAccountParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountParameterResponseBody) Validate() error {
	return dara.Validate(s)
}
