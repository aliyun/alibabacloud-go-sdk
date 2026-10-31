// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRectifyFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateRectifyFlowRequestCreateCommand) *CreateRectifyFlowRequest
	GetCreateCommand() *CreateRectifyFlowRequestCreateCommand
	SetOpTenantId(v int64) *CreateRectifyFlowRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateRectifyFlowRequest
	GetOpUserId() *string
}

type CreateRectifyFlowRequest struct {
	// The business request parameters.
	//
	// This parameter is required.
	CreateCommand *CreateRectifyFlowRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreateRectifyFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRectifyFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateRectifyFlowRequest) GetCreateCommand() *CreateRectifyFlowRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateRectifyFlowRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateRectifyFlowRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateRectifyFlowRequest) SetCreateCommand(v *CreateRectifyFlowRequestCreateCommand) *CreateRectifyFlowRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateRectifyFlowRequest) SetOpTenantId(v int64) *CreateRectifyFlowRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateRectifyFlowRequest) SetOpUserId(v string) *CreateRectifyFlowRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateRectifyFlowRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRectifyFlowRequestCreateCommand struct {
	// Specifies whether to associate with an existing process. Default value: false.
	BindToRectify *bool `json:"BindToRectify,omitempty" xml:"BindToRectify,omitempty"`
	// This parameter is required when you associate with an existing process.
	//
	// example:
	//
	// 117970862340
	BindToRectifyId *int64 `json:"BindToRectifyId,omitempty" xml:"BindToRectifyId,omitempty"`
	// The operation description. This parameter is optional. Maximum length: 128 characters.
	//
	// example:
	//
	// Initiate quality issue rectification for the customer table
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The governance object type. Valid values: TABLE, DATASOURCE_TABLE, DATASOURCE, INDEX, REALTIME_LOGICAL_TABLE, and QD_FEATURE.
	//
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	GovernItemType *string `json:"GovernItemType,omitempty" xml:"GovernItemType,omitempty"`
	// The list of issue IDs. Issues that cannot be rectified are skipped. If all issues are skipped, the original business error is returned.
	//
	// This parameter is required.
	GovernObjectIds []*int64 `json:"GovernObjectIds,omitempty" xml:"GovernObjectIds,omitempty" type:"Repeated"`
	// The rectification name. This parameter is required when you create a process and can contain up to 128 characters. It is ignored when you associate with an existing process.
	//
	// example:
	//
	// Quality issue rectification for the customer table
	RectifyName *string `json:"RectifyName,omitempty" xml:"RectifyName,omitempty"`
	// The rectification priority. Valid values: HIGH, MIDDLE, and LOW. Default value: MIDDLE. This parameter is ignored and not verified when you associate with an existing process.
	//
	// example:
	//
	// MIDDLE
	RectifyPriority *string `json:"RectifyPriority,omitempty" xml:"RectifyPriority,omitempty"`
	// The rectification scope. Valid values: QUALITY_RULE_SCOPE, QUALITY_RULE_TASK_SCOPE, QUALITY_GOVERN_OBJECT_SCOPE, and QUALITY_WATCH_SCOPE. The default value is the same as the value on the page. A rule scope may include other issues.
	//
	// example:
	//
	// QUALITY_GOVERN_OBJECT_SCOPE
	RectifyScope *string `json:"RectifyScope,omitempty" xml:"RectifyScope,omitempty"`
	// The ID of the user who performs the rectification. Default value: the current user. This parameter is ignored when you associate with an existing process.
	//
	// example:
	//
	// user001
	RectifyUser *string `json:"RectifyUser,omitempty" xml:"RectifyUser,omitempty"`
	// The ID of the user who performs the verification. Default value: the current user. This parameter is ignored when you associate with an existing process.
	//
	// example:
	//
	// user002
	VerifyUser *string `json:"VerifyUser,omitempty" xml:"VerifyUser,omitempty"`
}

func (s CreateRectifyFlowRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateRectifyFlowRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateRectifyFlowRequestCreateCommand) GetBindToRectify() *bool {
	return s.BindToRectify
}

func (s *CreateRectifyFlowRequestCreateCommand) GetBindToRectifyId() *int64 {
	return s.BindToRectifyId
}

func (s *CreateRectifyFlowRequestCreateCommand) GetComment() *string {
	return s.Comment
}

func (s *CreateRectifyFlowRequestCreateCommand) GetGovernItemType() *string {
	return s.GovernItemType
}

func (s *CreateRectifyFlowRequestCreateCommand) GetGovernObjectIds() []*int64 {
	return s.GovernObjectIds
}

func (s *CreateRectifyFlowRequestCreateCommand) GetRectifyName() *string {
	return s.RectifyName
}

func (s *CreateRectifyFlowRequestCreateCommand) GetRectifyPriority() *string {
	return s.RectifyPriority
}

func (s *CreateRectifyFlowRequestCreateCommand) GetRectifyScope() *string {
	return s.RectifyScope
}

func (s *CreateRectifyFlowRequestCreateCommand) GetRectifyUser() *string {
	return s.RectifyUser
}

func (s *CreateRectifyFlowRequestCreateCommand) GetVerifyUser() *string {
	return s.VerifyUser
}

func (s *CreateRectifyFlowRequestCreateCommand) SetBindToRectify(v bool) *CreateRectifyFlowRequestCreateCommand {
	s.BindToRectify = &v
	return s
}

func (s *CreateRectifyFlowRequestCreateCommand) SetBindToRectifyId(v int64) *CreateRectifyFlowRequestCreateCommand {
	s.BindToRectifyId = &v
	return s
}

func (s *CreateRectifyFlowRequestCreateCommand) SetComment(v string) *CreateRectifyFlowRequestCreateCommand {
	s.Comment = &v
	return s
}

func (s *CreateRectifyFlowRequestCreateCommand) SetGovernItemType(v string) *CreateRectifyFlowRequestCreateCommand {
	s.GovernItemType = &v
	return s
}

func (s *CreateRectifyFlowRequestCreateCommand) SetGovernObjectIds(v []*int64) *CreateRectifyFlowRequestCreateCommand {
	s.GovernObjectIds = v
	return s
}

func (s *CreateRectifyFlowRequestCreateCommand) SetRectifyName(v string) *CreateRectifyFlowRequestCreateCommand {
	s.RectifyName = &v
	return s
}

func (s *CreateRectifyFlowRequestCreateCommand) SetRectifyPriority(v string) *CreateRectifyFlowRequestCreateCommand {
	s.RectifyPriority = &v
	return s
}

func (s *CreateRectifyFlowRequestCreateCommand) SetRectifyScope(v string) *CreateRectifyFlowRequestCreateCommand {
	s.RectifyScope = &v
	return s
}

func (s *CreateRectifyFlowRequestCreateCommand) SetRectifyUser(v string) *CreateRectifyFlowRequestCreateCommand {
	s.RectifyUser = &v
	return s
}

func (s *CreateRectifyFlowRequestCreateCommand) SetVerifyUser(v string) *CreateRectifyFlowRequestCreateCommand {
	s.VerifyUser = &v
	return s
}

func (s *CreateRectifyFlowRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
