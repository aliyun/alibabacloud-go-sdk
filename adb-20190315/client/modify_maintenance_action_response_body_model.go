// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaintenanceActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *ModifyMaintenanceActionResponseBody
	GetIds() *string
	SetRequestId(v string) *ModifyMaintenanceActionResponseBody
	GetRequestId() *string
}

type ModifyMaintenanceActionResponseBody struct {
	// The O\\&M event ID.
	//
	// example:
	//
	// 11111
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7856CBE7-5BD0-4EE1-AC62-749392******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMaintenanceActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaintenanceActionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMaintenanceActionResponseBody) GetIds() *string {
	return s.Ids
}

func (s *ModifyMaintenanceActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMaintenanceActionResponseBody) SetIds(v string) *ModifyMaintenanceActionResponseBody {
	s.Ids = &v
	return s
}

func (s *ModifyMaintenanceActionResponseBody) SetRequestId(v string) *ModifyMaintenanceActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMaintenanceActionResponseBody) Validate() error {
	return dara.Validate(s)
}
