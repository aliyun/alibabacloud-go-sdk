// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBgpGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBgpGroupAttributeResponseBody
	GetRequestId() *string
}

type ModifyBgpGroupAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8C3C6D7C-A1CE-4FD8-BC57-DC493A55F76F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBgpGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBgpGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBgpGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBgpGroupAttributeResponseBody) SetRequestId(v string) *ModifyBgpGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBgpGroupAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
