// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectTrafficQosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyExpressConnectTrafficQosResponseBody
	GetRequestId() *string
}

type ModifyExpressConnectTrafficQosResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B457
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyExpressConnectTrafficQosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectTrafficQosResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectTrafficQosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyExpressConnectTrafficQosResponseBody) SetRequestId(v string) *ModifyExpressConnectTrafficQosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosResponseBody) Validate() error {
	return dara.Validate(s)
}
