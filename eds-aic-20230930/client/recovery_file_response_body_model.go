// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoveryFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *RecoveryFileResponseBody
	GetCount() *int64
	SetData(v []*RecoveryFileResponseBodyData) *RecoveryFileResponseBody
	GetData() []*RecoveryFileResponseBodyData
	SetRequestId(v string) *RecoveryFileResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RecoveryFileResponseBody
	GetTaskId() *string
}

type RecoveryFileResponseBody struct {
	// The number of restored instances.
	//
	// example:
	//
	// 97
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The backup file that is restored.
	//
	// example:
	//
	// 6AD56E39-430B-5401-AB4A-7B086454****
	Data []*RecoveryFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6AD56E39-430B-5401-AB4A-7B086454****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the batch task.
	//
	// example:
	//
	// t-5prhfo7wv1gjx****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecoveryFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoveryFileResponseBody) GoString() string {
	return s.String()
}

func (s *RecoveryFileResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *RecoveryFileResponseBody) GetData() []*RecoveryFileResponseBodyData {
	return s.Data
}

func (s *RecoveryFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoveryFileResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RecoveryFileResponseBody) SetCount(v int64) *RecoveryFileResponseBody {
	s.Count = &v
	return s
}

func (s *RecoveryFileResponseBody) SetData(v []*RecoveryFileResponseBodyData) *RecoveryFileResponseBody {
	s.Data = v
	return s
}

func (s *RecoveryFileResponseBody) SetRequestId(v string) *RecoveryFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoveryFileResponseBody) SetTaskId(v string) *RecoveryFileResponseBody {
	s.TaskId = &v
	return s
}

func (s *RecoveryFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecoveryFileResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-5prhfo7wv1gjx****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecoveryFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecoveryFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecoveryFileResponseBodyData) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *RecoveryFileResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *RecoveryFileResponseBodyData) SetAndroidInstanceId(v string) *RecoveryFileResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *RecoveryFileResponseBodyData) SetTaskId(v string) *RecoveryFileResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RecoveryFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
