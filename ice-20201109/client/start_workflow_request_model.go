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
	// Specifies whether to skip the input path verification for the workflow. This parameter takes effect only when the workflow input is an OSS file. We recommend that you do not skip the verification to avoid errors caused by incorrect paths. If this parameter is not specified, the default value is false. Valid values:
	//
	// - **true**: Skip the verification.
	//
	// - **false**: Do not skip the verification.
	//
	// example:
	//
	// false
	SkipInputVerification *bool `json:"SkipInputVerification,omitempty" xml:"SkipInputVerification,omitempty"`
	// The workflow input. Currently, media asset types and OSS files are supported.
	//
	// Type: the supported media object type. Valid values:
	//
	// - OSS: an OSS file.
	//
	// - Media: a media asset ID.
	//
	// Media: the media value. Valid values:
	//
	// - If Type is set to OSS, the value is a URL that supports the OSS protocol and HTTP protocol.
	//
	// - If Type is set to Media, the value is a media asset ID.
	//
	// example:
	//
	// {
	//
	//       "Type": "Media",
	//
	//       "Media": "******30706071edbfe290b488******"
	//
	// } or
	//
	// {
	//
	//       "Type": "OSS",
	//
	//       "Media": "oss://bucket.oss-ap-southeast-1.aliyuncs.com/A/B/C/test1.flv"
	//
	// }
	TaskInput *string `json:"TaskInput,omitempty" xml:"TaskInput,omitempty"`
	// The custom settings in JSON format. The maximum length is 512 bytes. [Custom callback URL configuration](https://help.aliyun.com/document_detail/451631.html) is supported.
	//
	// example:
	//
	// {"NotifyAddress":"https://xx.xx.xxx"} or {"NotifyAddress":"ice-callback-demo"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The workflow template ID. You can view the template ID in the [Intelligent Media Services console](https://ims.console.aliyun.com/settings/workflow/list) by navigating to Configuration Management > Workflow Template.
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
