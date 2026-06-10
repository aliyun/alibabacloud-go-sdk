// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAgentForClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeAgentForClusterResponseBody
	GetRequestId() *string
	SetCode(v string) *UpgradeAgentForClusterResponseBody
	GetCode() *string
	SetData(v *UpgradeAgentForClusterResponseBodyData) *UpgradeAgentForClusterResponseBody
	GetData() *UpgradeAgentForClusterResponseBodyData
	SetMessage(v string) *UpgradeAgentForClusterResponseBody
	GetMessage() *string
}

type UpgradeAgentForClusterResponseBody struct {
	// Request ID, which can be used for end-to-end Diagnosis
	//
	// example:
	//
	// B149FD9C-ED5C-5765-B3AD-05AA4A4D64D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Status code
	//
	// - `code == Success` indicates successful authorization;
	//
	// - Other status codes indicate authorization failure. When authorization fails, view the `message` field to obtain detailed error message;
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned Data.
	Data *UpgradeAgentForClusterResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error message
	//
	// - If `code == Success`, this field is empty;
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpgradeAgentForClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentForClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeAgentForClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeAgentForClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpgradeAgentForClusterResponseBody) GetData() *UpgradeAgentForClusterResponseBodyData {
	return s.Data
}

func (s *UpgradeAgentForClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeAgentForClusterResponseBody) SetRequestId(v string) *UpgradeAgentForClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeAgentForClusterResponseBody) SetCode(v string) *UpgradeAgentForClusterResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeAgentForClusterResponseBody) SetData(v *UpgradeAgentForClusterResponseBodyData) *UpgradeAgentForClusterResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeAgentForClusterResponseBody) SetMessage(v string) *UpgradeAgentForClusterResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeAgentForClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeAgentForClusterResponseBodyData struct {
	// Job ID.
	//
	// You can use this job ID to invoke the GetAgentTask API to view the execution status of the job.
	//
	// example:
	//
	// 7523e9e0ddc74d99a5236f4f4d5056e6
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UpgradeAgentForClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentForClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeAgentForClusterResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeAgentForClusterResponseBodyData) SetTaskId(v string) *UpgradeAgentForClusterResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpgradeAgentForClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
