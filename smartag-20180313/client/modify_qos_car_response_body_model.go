// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosCarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyQosCarResponseBody
	GetRequestId() *string
}

type ModifyQosCarResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 551CD836-9E46-4F2C-A167-B4363180A647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyQosCarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosCarResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyQosCarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyQosCarResponseBody) SetRequestId(v string) *ModifyQosCarResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyQosCarResponseBody) Validate() error {
	return dara.Validate(s)
}
