// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceTDEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceTDEResponseBody
	GetRequestId() *string
}

type ModifyInstanceTDEResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceTDEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceTDEResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTDEResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceTDEResponseBody) SetRequestId(v string) *ModifyInstanceTDEResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceTDEResponseBody) Validate() error {
	return dara.Validate(s)
}
