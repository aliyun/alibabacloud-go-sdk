// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHighlightTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *CreateHighlightTaskShrinkRequest
	GetCredentialConfigShrink() *string
	SetEditShrink(v string) *CreateHighlightTaskShrinkRequest
	GetEditShrink() *string
	SetHighlightShrink(v string) *CreateHighlightTaskShrinkRequest
	GetHighlightShrink() *string
	SetMode(v string) *CreateHighlightTaskShrinkRequest
	GetMode() *string
	SetNotificationShrink(v string) *CreateHighlightTaskShrinkRequest
	GetNotificationShrink() *string
	SetOutputShrink(v string) *CreateHighlightTaskShrinkRequest
	GetOutputShrink() *string
	SetProjectName(v string) *CreateHighlightTaskShrinkRequest
	GetProjectName() *string
	SetSourcesShrink(v string) *CreateHighlightTaskShrinkRequest
	GetSourcesShrink() *string
	SetTagsShrink(v string) *CreateHighlightTaskShrinkRequest
	GetTagsShrink() *string
	SetType(v string) *CreateHighlightTaskShrinkRequest
	GetType() *string
	SetUserData(v string) *CreateHighlightTaskShrinkRequest
	GetUserData() *string
}

type CreateHighlightTaskShrinkRequest struct {
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	EditShrink             *string `json:"Edit,omitempty" xml:"Edit,omitempty"`
	HighlightShrink        *string `json:"Highlight,omitempty" xml:"Highlight,omitempty"`
	// example:
	//
	// Average
	Mode               *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// This parameter is required.
	OutputShrink *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// example:
	//
	// {"test":"val1"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Retrieval
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// {"ID": "testuid","Name": "test-user","Avatar": "http://test.com/testuid"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateHighlightTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHighlightTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateHighlightTaskShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateHighlightTaskShrinkRequest) GetEditShrink() *string {
	return s.EditShrink
}

func (s *CreateHighlightTaskShrinkRequest) GetHighlightShrink() *string {
	return s.HighlightShrink
}

func (s *CreateHighlightTaskShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateHighlightTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateHighlightTaskShrinkRequest) GetOutputShrink() *string {
	return s.OutputShrink
}

func (s *CreateHighlightTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateHighlightTaskShrinkRequest) GetSourcesShrink() *string {
	return s.SourcesShrink
}

func (s *CreateHighlightTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateHighlightTaskShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateHighlightTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateHighlightTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateHighlightTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateHighlightTaskShrinkRequest) SetEditShrink(v string) *CreateHighlightTaskShrinkRequest {
	s.EditShrink = &v
	return s
}

func (s *CreateHighlightTaskShrinkRequest) SetHighlightShrink(v string) *CreateHighlightTaskShrinkRequest {
	s.HighlightShrink = &v
	return s
}

func (s *CreateHighlightTaskShrinkRequest) SetMode(v string) *CreateHighlightTaskShrinkRequest {
	s.Mode = &v
	return s
}

func (s *CreateHighlightTaskShrinkRequest) SetNotificationShrink(v string) *CreateHighlightTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateHighlightTaskShrinkRequest) SetOutputShrink(v string) *CreateHighlightTaskShrinkRequest {
	s.OutputShrink = &v
	return s
}

func (s *CreateHighlightTaskShrinkRequest) SetProjectName(v string) *CreateHighlightTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateHighlightTaskShrinkRequest) SetSourcesShrink(v string) *CreateHighlightTaskShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreateHighlightTaskShrinkRequest) SetTagsShrink(v string) *CreateHighlightTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateHighlightTaskShrinkRequest) SetType(v string) *CreateHighlightTaskShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateHighlightTaskShrinkRequest) SetUserData(v string) *CreateHighlightTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateHighlightTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
