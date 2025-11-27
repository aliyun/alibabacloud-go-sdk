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
	// The ID of the request.
	//
	// example:
	//
	// 52B9805C-432C-4ED1-83FD-2F916B6D2733
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
