// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAutoRenewAttributeResponseBody
	GetRequestId() *string
}

type ModifyAutoRenewAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4CE6DF97-AEA4-484F-906F-C407EE******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAutoRenewAttributeResponseBody) SetRequestId(v string) *ModifyAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAutoRenewAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
