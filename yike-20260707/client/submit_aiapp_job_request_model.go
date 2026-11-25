// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIAppJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SubmitAIAppJobRequest
	GetAppId() *string
	SetClientToken(v string) *SubmitAIAppJobRequest
	GetClientToken() *string
	SetInputs(v string) *SubmitAIAppJobRequest
	GetInputs() *string
	SetUserData(v string) *SubmitAIAppJobRequest
	GetUserData() *string
}

type SubmitAIAppJobRequest struct {
	// The template ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Ensures the idempotence of the request. Generate a parameter value from your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length.
	//
	// example:
	//
	// ******20b48fb04***
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The task inputs.
	//
	// example:
	//
	// {
	//
	//   "InputPrompt": "xxxx"
	//
	// }
	Inputs *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// The custom fields.
	//
	// example:
	//
	// {\\"Key\\":\\"Value\\"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIAppJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIAppJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIAppJobRequest) GetAppId() *string {
	return s.AppId
}

func (s *SubmitAIAppJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitAIAppJobRequest) GetInputs() *string {
	return s.Inputs
}

func (s *SubmitAIAppJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIAppJobRequest) SetAppId(v string) *SubmitAIAppJobRequest {
	s.AppId = &v
	return s
}

func (s *SubmitAIAppJobRequest) SetClientToken(v string) *SubmitAIAppJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitAIAppJobRequest) SetInputs(v string) *SubmitAIAppJobRequest {
	s.Inputs = &v
	return s
}

func (s *SubmitAIAppJobRequest) SetUserData(v string) *SubmitAIAppJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIAppJobRequest) Validate() error {
	return dara.Validate(s)
}
