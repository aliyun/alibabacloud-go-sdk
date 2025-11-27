// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceTDEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceTDEResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceTDEResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 777C4593-8053-427B-99E2-105593277CAB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceTDEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceTDEResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceTDEResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceTDEResponseBody) SetRequestId(v string) *ModifyDBInstanceTDEResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceTDEResponseBody) Validate() error {
	return dara.Validate(s)
}
