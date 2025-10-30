// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWuyingServerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWuyingServerAttributeResponseBody
	GetRequestId() *string
}

type ModifyWuyingServerAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWuyingServerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWuyingServerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWuyingServerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWuyingServerAttributeResponseBody) SetRequestId(v string) *ModifyWuyingServerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWuyingServerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
