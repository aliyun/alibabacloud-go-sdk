// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOncallScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateOncallScheduleResponseBody
	GetRequestId() *string
}

type UpdateOncallScheduleResponseBody struct {
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateOncallScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOncallScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOncallScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOncallScheduleResponseBody) SetRequestId(v string) *UpdateOncallScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOncallScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
