// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceVectorSupportStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceVectorSupportStatusResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceVectorSupportStatusResponseBody struct {
	// example:
	//
	// 16C62438-491B-5C02-9B49-BA924A1372A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceVectorSupportStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceVectorSupportStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceVectorSupportStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceVectorSupportStatusResponseBody) SetRequestId(v string) *ModifyDBInstanceVectorSupportStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceVectorSupportStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
