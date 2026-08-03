// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInspectionScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInspectionScheduleResponseBody
	GetRequestId() *string
}

type ModifyInspectionScheduleResponseBody struct {
	// example:
	//
	// AD7E16AA-6B23-43BF-979C-07D957FB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInspectionScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInspectionScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInspectionScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInspectionScheduleResponseBody) SetRequestId(v string) *ModifyInspectionScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInspectionScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
