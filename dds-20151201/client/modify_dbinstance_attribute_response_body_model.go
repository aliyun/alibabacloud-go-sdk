// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceAttributeResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceAttributeResponseBody struct {
	// example:
	//
	// 6826DEE3-B374-5DB7-909D-C8978827C534
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceAttributeResponseBody) SetRequestId(v string) *ModifyDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
