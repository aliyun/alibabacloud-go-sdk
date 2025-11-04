// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDispatchTag(v string) *StartAIWorkflowRequest
	GetDispatchTag() *string
	SetInputs(v string) *StartAIWorkflowRequest
	GetInputs() *string
	SetUserData(v string) *StartAIWorkflowRequest
	GetUserData() *string
	SetWorkflowId(v string) *StartAIWorkflowRequest
	GetWorkflowId() *string
}

type StartAIWorkflowRequest struct {
	// The tag for the task.
	//
	// example:
	//
	// default
	DispatchTag *string `json:"DispatchTag,omitempty" xml:"DispatchTag,omitempty"`
	// A JSON string containing the specific input parameters, such as information about the media assets, standard live streams, or RTC streams.
	//
	// example:
	//
	// {
	//
	//     "live_url": {
	//
	//         "Url": "rtmp://test.com/test_app/test_stream?auth_key=test",
	//
	//         "MaxIdleTime": 20
	//
	//     },
	//
	//     "source_language_id": "es"
	//
	// }
	Inputs *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// A user-defined parameter for passing custom metadata.
	//
	// example:
	//
	// {
	//
	// "url":"https://test.com"
	//
	// }
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the workflow template.
	//
	// example:
	//
	// ****3f44-f1f6-477e-9364-c5e6c49e****
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s StartAIWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAIWorkflowRequest) GoString() string {
	return s.String()
}

func (s *StartAIWorkflowRequest) GetDispatchTag() *string {
	return s.DispatchTag
}

func (s *StartAIWorkflowRequest) GetInputs() *string {
	return s.Inputs
}

func (s *StartAIWorkflowRequest) GetUserData() *string {
	return s.UserData
}

func (s *StartAIWorkflowRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *StartAIWorkflowRequest) SetDispatchTag(v string) *StartAIWorkflowRequest {
	s.DispatchTag = &v
	return s
}

func (s *StartAIWorkflowRequest) SetInputs(v string) *StartAIWorkflowRequest {
	s.Inputs = &v
	return s
}

func (s *StartAIWorkflowRequest) SetUserData(v string) *StartAIWorkflowRequest {
	s.UserData = &v
	return s
}

func (s *StartAIWorkflowRequest) SetWorkflowId(v string) *StartAIWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *StartAIWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
