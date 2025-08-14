// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSkipInputVerification(v bool) *StartWorkflowRequest
	GetSkipInputVerification() *bool
	SetTaskInput(v string) *StartWorkflowRequest
	GetTaskInput() *string
	SetUserData(v string) *StartWorkflowRequest
	GetUserData() *string
	SetWorkflowId(v string) *StartWorkflowRequest
	GetWorkflowId() *string
}

type StartWorkflowRequest struct {
	SkipInputVerification *bool `json:"SkipInputVerification,omitempty" xml:"SkipInputVerification,omitempty"`
	// The workflow input. Only media assets are supported.
	//
	// example:
	//
	// {
	//
	//       "Type": "Media",
	//
	//       "Media": "******30706071edbfe290b488******"
	//
	// }
	TaskInput *string `json:"TaskInput,omitempty" xml:"TaskInput,omitempty"`
	// The user-defined data in the JSON format, which cannot be up to 512 bytes in length. You can specify a custom callback URL. For more information, see [Configure a callback upon editing completion](https://help.aliyun.com/document_detail/451631.html).
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the workflow template. To view the template ID, log on to the [IMS console](https://ims.console.aliyun.com/settings/workflow/list) and choose Configurations > Workflow Template.
	//
	// example:
	//
	// ******f0e54971ecbffd472190******
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s StartWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s StartWorkflowRequest) GoString() string {
	return s.String()
}

func (s *StartWorkflowRequest) GetSkipInputVerification() *bool {
	return s.SkipInputVerification
}

func (s *StartWorkflowRequest) GetTaskInput() *string {
	return s.TaskInput
}

func (s *StartWorkflowRequest) GetUserData() *string {
	return s.UserData
}

func (s *StartWorkflowRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *StartWorkflowRequest) SetSkipInputVerification(v bool) *StartWorkflowRequest {
	s.SkipInputVerification = &v
	return s
}

func (s *StartWorkflowRequest) SetTaskInput(v string) *StartWorkflowRequest {
	s.TaskInput = &v
	return s
}

func (s *StartWorkflowRequest) SetUserData(v string) *StartWorkflowRequest {
	s.UserData = &v
	return s
}

func (s *StartWorkflowRequest) SetWorkflowId(v string) *StartWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *StartWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
