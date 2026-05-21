// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOperationEventScheduleTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateOperationEventScheduleTimeResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateOperationEventScheduleTimeResponseBody
	GetRequestId() *string
}

type UpdateOperationEventScheduleTimeResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 11F64C4C-EC50-5472-BC5D-7FD9F51499F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOperationEventScheduleTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOperationEventScheduleTimeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOperationEventScheduleTimeResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateOperationEventScheduleTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOperationEventScheduleTimeResponseBody) SetData(v bool) *UpdateOperationEventScheduleTimeResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateOperationEventScheduleTimeResponseBody) SetRequestId(v string) *UpdateOperationEventScheduleTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOperationEventScheduleTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
