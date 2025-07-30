// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIntranetAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIntranetAttributeResponseBody
	GetRequestId() *string
}

type ModifyIntranetAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 97AC8948-D7E4-457E-BE03-850CF04E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIntranetAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIntranetAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIntranetAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIntranetAttributeResponseBody) SetRequestId(v string) *ModifyIntranetAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIntranetAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
