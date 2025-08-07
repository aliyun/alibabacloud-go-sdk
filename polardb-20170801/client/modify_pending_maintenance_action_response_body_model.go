// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPendingMaintenanceActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *ModifyPendingMaintenanceActionResponseBody
	GetIds() *string
	SetRequestId(v string) *ModifyPendingMaintenanceActionResponseBody
	GetRequestId() *string
}

type ModifyPendingMaintenanceActionResponseBody struct {
	// The ID of the task.
	//
	// example:
	//
	// 111111
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 93061E17-B56A-4324-BC95-D0FFD2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPendingMaintenanceActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPendingMaintenanceActionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPendingMaintenanceActionResponseBody) GetIds() *string {
	return s.Ids
}

func (s *ModifyPendingMaintenanceActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPendingMaintenanceActionResponseBody) SetIds(v string) *ModifyPendingMaintenanceActionResponseBody {
	s.Ids = &v
	return s
}

func (s *ModifyPendingMaintenanceActionResponseBody) SetRequestId(v string) *ModifyPendingMaintenanceActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPendingMaintenanceActionResponseBody) Validate() error {
	return dara.Validate(s)
}
