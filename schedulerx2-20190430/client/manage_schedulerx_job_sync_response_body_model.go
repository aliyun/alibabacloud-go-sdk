// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageSchedulerxJobSyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ManageSchedulerxJobSyncResponseBody
	GetCode() *int32
	SetMessage(v string) *ManageSchedulerxJobSyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *ManageSchedulerxJobSyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ManageSchedulerxJobSyncResponseBody
	GetSuccess() *bool
}

type ManageSchedulerxJobSyncResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// jobId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB58C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ManageSchedulerxJobSyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManageSchedulerxJobSyncResponseBody) GoString() string {
	return s.String()
}

func (s *ManageSchedulerxJobSyncResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ManageSchedulerxJobSyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ManageSchedulerxJobSyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManageSchedulerxJobSyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ManageSchedulerxJobSyncResponseBody) SetCode(v int32) *ManageSchedulerxJobSyncResponseBody {
	s.Code = &v
	return s
}

func (s *ManageSchedulerxJobSyncResponseBody) SetMessage(v string) *ManageSchedulerxJobSyncResponseBody {
	s.Message = &v
	return s
}

func (s *ManageSchedulerxJobSyncResponseBody) SetRequestId(v string) *ManageSchedulerxJobSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManageSchedulerxJobSyncResponseBody) SetSuccess(v bool) *ManageSchedulerxJobSyncResponseBody {
	s.Success = &v
	return s
}

func (s *ManageSchedulerxJobSyncResponseBody) Validate() error {
	return dara.Validate(s)
}
