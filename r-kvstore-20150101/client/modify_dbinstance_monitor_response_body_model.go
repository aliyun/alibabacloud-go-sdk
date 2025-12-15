// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceMonitorResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceMonitorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2BE6E619-A657-42E3-AD2D-18F8428A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceMonitorResponseBody) SetRequestId(v string) *ModifyDBInstanceMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
