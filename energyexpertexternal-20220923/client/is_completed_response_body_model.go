// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsCompletedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *IsCompletedResponseBodyData) *IsCompletedResponseBody
	GetData() *IsCompletedResponseBodyData
	SetRequestId(v string) *IsCompletedResponseBody
	GetRequestId() *string
}

type IsCompletedResponseBody struct {
	// The response parameters.
	Data *IsCompletedResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s IsCompletedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IsCompletedResponseBody) GoString() string {
	return s.String()
}

func (s *IsCompletedResponseBody) GetData() *IsCompletedResponseBodyData {
	return s.Data
}

func (s *IsCompletedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IsCompletedResponseBody) SetData(v *IsCompletedResponseBodyData) *IsCompletedResponseBody {
	s.Data = v
	return s
}

func (s *IsCompletedResponseBody) SetRequestId(v string) *IsCompletedResponseBody {
	s.RequestId = &v
	return s
}

func (s *IsCompletedResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IsCompletedResponseBodyData struct {
	// Modified time in milliseconds, e.g. 1711438780000.
	//
	// example:
	//
	// 1711438780000
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// The unique key of this generation task.
	//
	// example:
	//
	// 550c2b7b-f2e0-4176-ab0a-53ea4b355721
	TaskKey *string `json:"taskKey,omitempty" xml:"taskKey,omitempty"`
	// Unused temporarily.
	//
	// example:
	//
	// null
	TaskShortResult *string `json:"taskShortResult,omitempty" xml:"taskShortResult,omitempty"`
	// The status of the report generation task. The possible values are `running`, `success`, and `error`, which mean generating, generating succeeded, and generating failed, respectively. If you encounter a result generation failure, check the model, correct the model, and then generate the result again.
	//
	// example:
	//
	// running
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s IsCompletedResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s IsCompletedResponseBodyData) GoString() string {
	return s.String()
}

func (s *IsCompletedResponseBodyData) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *IsCompletedResponseBodyData) GetTaskKey() *string {
	return s.TaskKey
}

func (s *IsCompletedResponseBodyData) GetTaskShortResult() *string {
	return s.TaskShortResult
}

func (s *IsCompletedResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *IsCompletedResponseBodyData) SetModifiedTime(v int64) *IsCompletedResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *IsCompletedResponseBodyData) SetTaskKey(v string) *IsCompletedResponseBodyData {
	s.TaskKey = &v
	return s
}

func (s *IsCompletedResponseBodyData) SetTaskShortResult(v string) *IsCompletedResponseBodyData {
	s.TaskShortResult = &v
	return s
}

func (s *IsCompletedResponseBodyData) SetTaskStatus(v string) *IsCompletedResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *IsCompletedResponseBodyData) Validate() error {
	return dara.Validate(s)
}
