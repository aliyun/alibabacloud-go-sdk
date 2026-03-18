// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateQuotaScheduleResponseBody
	GetData() *string
	SetRequestId(v string) *UpdateQuotaScheduleResponseBody
	GetRequestId() *string
}

type UpdateQuotaScheduleResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateQuotaScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaScheduleResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateQuotaScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQuotaScheduleResponseBody) SetData(v string) *UpdateQuotaScheduleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateQuotaScheduleResponseBody) SetRequestId(v string) *UpdateQuotaScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQuotaScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
