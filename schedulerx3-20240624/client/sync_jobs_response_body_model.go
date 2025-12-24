// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SyncJobsResponseBody
	GetCode() *int32
	SetMessage(v string) *SyncJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncJobsResponseBody
	GetSuccess() *bool
}

type SyncJobsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Not found: appName not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6BCE89B3-E882-511D-9A75-D452A56EC4B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncJobsResponseBody) GoString() string {
	return s.String()
}

func (s *SyncJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SyncJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncJobsResponseBody) SetCode(v int32) *SyncJobsResponseBody {
	s.Code = &v
	return s
}

func (s *SyncJobsResponseBody) SetMessage(v string) *SyncJobsResponseBody {
	s.Message = &v
	return s
}

func (s *SyncJobsResponseBody) SetRequestId(v string) *SyncJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncJobsResponseBody) SetSuccess(v bool) *SyncJobsResponseBody {
	s.Success = &v
	return s
}

func (s *SyncJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
