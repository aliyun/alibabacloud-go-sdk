// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryHistoryJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int32) *CreateDeliveryHistoryJobResponseBody
	GetJobId() *int32
	SetRequestId(v string) *CreateDeliveryHistoryJobResponseBody
	GetRequestId() *string
}

type CreateDeliveryHistoryJobResponseBody struct {
	// The ID of the historical event delivery task.
	//
	// example:
	//
	// 16602
	JobId *int32 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9D356A34-D5A9-41CD-9915-837B7F9D8722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeliveryHistoryJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryHistoryJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeliveryHistoryJobResponseBody) GetJobId() *int32 {
	return s.JobId
}

func (s *CreateDeliveryHistoryJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeliveryHistoryJobResponseBody) SetJobId(v int32) *CreateDeliveryHistoryJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateDeliveryHistoryJobResponseBody) SetRequestId(v string) *CreateDeliveryHistoryJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeliveryHistoryJobResponseBody) Validate() error {
	return dara.Validate(s)
}
