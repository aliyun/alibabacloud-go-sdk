// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInspectionScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateInspectionScheduleResponseBody
	GetRequestId() *string
}

type CreateInspectionScheduleResponseBody struct {
	// example:
	//
	// 561AFBF1-BE20-44DB-9BD1-6988B53E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInspectionScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInspectionScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInspectionScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInspectionScheduleResponseBody) SetRequestId(v string) *CreateInspectionScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInspectionScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
