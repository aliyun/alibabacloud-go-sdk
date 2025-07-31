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
	// 434D7127-6229-4355-BA50-7A3685A725DF
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
