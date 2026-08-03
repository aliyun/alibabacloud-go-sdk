// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInspectionScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteInspectionScheduleResponseBody
	GetRequestId() *string
}

type DeleteInspectionScheduleResponseBody struct {
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInspectionScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInspectionScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInspectionScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInspectionScheduleResponseBody) SetRequestId(v string) *DeleteInspectionScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInspectionScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
