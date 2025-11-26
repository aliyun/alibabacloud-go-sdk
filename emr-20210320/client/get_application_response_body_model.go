// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody
	GetApplication() *GetApplicationResponseBodyApplication
	SetRequestId(v string) *GetApplicationResponseBody
	GetRequestId() *string
}

type GetApplicationResponseBody struct {
	// The application.
	Application *GetApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) GetApplication() *GetApplicationResponseBodyApplication {
	return s.Application
}

func (s *GetApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationResponseBody) SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody {
	s.Application = v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponseBody) Validate() error {
	if s.Application != nil {
		if err := s.Application.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyApplication struct {
	// The list of actions supported by the application.
	Actions []*GetApplicationResponseBodyApplicationActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The application name.
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The status of the application.
	//
	// example:
	//
	// RUNNING
	ApplicationState *string `json:"ApplicationState,omitempty" xml:"ApplicationState,omitempty"`
	// The application version.
	//
	// example:
	//
	// 2.8.1
	ApplicationVersion *string `json:"ApplicationVersion,omitempty" xml:"ApplicationVersion,omitempty"`
	// The community version.
	//
	// example:
	//
	// 2.8.1
	CommunityVersion *string `json:"CommunityVersion,omitempty" xml:"CommunityVersion,omitempty"`
}

func (s GetApplicationResponseBodyApplication) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplication) GetActions() []*GetApplicationResponseBodyApplicationActions {
	return s.Actions
}

func (s *GetApplicationResponseBodyApplication) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *GetApplicationResponseBodyApplication) GetApplicationState() *string {
	return s.ApplicationState
}

func (s *GetApplicationResponseBodyApplication) GetApplicationVersion() *string {
	return s.ApplicationVersion
}

func (s *GetApplicationResponseBodyApplication) GetCommunityVersion() *string {
	return s.CommunityVersion
}

func (s *GetApplicationResponseBodyApplication) SetActions(v []*GetApplicationResponseBodyApplicationActions) *GetApplicationResponseBodyApplication {
	s.Actions = v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationName(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationName = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationState(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationState = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationVersion(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationVersion = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCommunityVersion(v string) *GetApplicationResponseBodyApplication {
	s.CommunityVersion = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApplicationResponseBodyApplicationActions struct {
	// The action name.
	//
	// example:
	//
	// decommission
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// The list of action parameters.
	ActionParams []*GetApplicationResponseBodyApplicationActionsActionParams `json:"ActionParams,omitempty" xml:"ActionParams,omitempty" type:"Repeated"`
	// The action command.
	//
	// example:
	//
	// decommission
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The component name.
	//
	// example:
	//
	// DataNode
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The description of the action.
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution scope. Valid values:
	//
	// 	- APPLICATION: the application scope.
	//
	// 	- COMPONENT: the component scope.
	//
	// 	- COMPONENT_INSTANCE: the component instance scope.
	//
	// example:
	//
	// COMPONENT_INSTANCE
	RunActionScope *string `json:"RunActionScope,omitempty" xml:"RunActionScope,omitempty"`
}

func (s GetApplicationResponseBodyApplicationActions) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationActions) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationActions) GetActionName() *string {
	return s.ActionName
}

func (s *GetApplicationResponseBodyApplicationActions) GetActionParams() []*GetApplicationResponseBodyApplicationActionsActionParams {
	return s.ActionParams
}

func (s *GetApplicationResponseBodyApplicationActions) GetCommand() *string {
	return s.Command
}

func (s *GetApplicationResponseBodyApplicationActions) GetComponentName() *string {
	return s.ComponentName
}

func (s *GetApplicationResponseBodyApplicationActions) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyApplicationActions) GetRunActionScope() *string {
	return s.RunActionScope
}

func (s *GetApplicationResponseBodyApplicationActions) SetActionName(v string) *GetApplicationResponseBodyApplicationActions {
	s.ActionName = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActions) SetActionParams(v []*GetApplicationResponseBodyApplicationActionsActionParams) *GetApplicationResponseBodyApplicationActions {
	s.ActionParams = v
	return s
}

func (s *GetApplicationResponseBodyApplicationActions) SetCommand(v string) *GetApplicationResponseBodyApplicationActions {
	s.Command = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActions) SetComponentName(v string) *GetApplicationResponseBodyApplicationActions {
	s.ComponentName = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActions) SetDescription(v string) *GetApplicationResponseBodyApplicationActions {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActions) SetRunActionScope(v string) *GetApplicationResponseBodyApplicationActions {
	s.RunActionScope = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActions) Validate() error {
	if s.ActionParams != nil {
		for _, item := range s.ActionParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApplicationResponseBodyApplicationActionsActionParams struct {
	// Description
	//
	// example:
	//
	// start
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// timeout
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The parameter value attribute.
	ValueAttribute *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute `json:"ValueAttribute,omitempty" xml:"ValueAttribute,omitempty" type:"Struct"`
}

func (s GetApplicationResponseBodyApplicationActionsActionParams) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationActionsActionParams) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationActionsActionParams) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyApplicationActionsActionParams) GetKey() *string {
	return s.Key
}

func (s *GetApplicationResponseBodyApplicationActionsActionParams) GetValueAttribute() *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute {
	return s.ValueAttribute
}

func (s *GetApplicationResponseBodyApplicationActionsActionParams) SetDescription(v string) *GetApplicationResponseBodyApplicationActionsActionParams {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActionsActionParams) SetKey(v string) *GetApplicationResponseBodyApplicationActionsActionParams {
	s.Key = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActionsActionParams) SetValueAttribute(v *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) *GetApplicationResponseBodyApplicationActionsActionParams {
	s.ValueAttribute = v
	return s
}

func (s *GetApplicationResponseBodyApplicationActionsActionParams) Validate() error {
	if s.ValueAttribute != nil {
		if err := s.ValueAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute struct {
	// Value description.
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The size of the value increment.
	//
	// example:
	//
	// 1
	ValueIncrementStep *string `json:"ValueIncrementStep,omitempty" xml:"ValueIncrementStep,omitempty"`
	// The maximum range of values.
	//
	// example:
	//
	// 100
	ValueMaximum *string `json:"ValueMaximum,omitempty" xml:"ValueMaximum,omitempty"`
	// The minimum range of values.
	//
	// example:
	//
	// 1
	ValueMinimum *string `json:"ValueMinimum,omitempty" xml:"ValueMinimum,omitempty"`
	// The value type of the column.
	//
	// example:
	//
	// int
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	// The unit of the value.
	//
	// example:
	//
	// number
	ValueUnit *string `json:"ValueUnit,omitempty" xml:"ValueUnit,omitempty"`
}

func (s GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) GetValueIncrementStep() *string {
	return s.ValueIncrementStep
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) GetValueMaximum() *string {
	return s.ValueMaximum
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) GetValueMinimum() *string {
	return s.ValueMinimum
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) GetValueType() *string {
	return s.ValueType
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) GetValueUnit() *string {
	return s.ValueUnit
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) SetDescription(v string) *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) SetValueIncrementStep(v string) *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute {
	s.ValueIncrementStep = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) SetValueMaximum(v string) *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute {
	s.ValueMaximum = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) SetValueMinimum(v string) *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute {
	s.ValueMinimum = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) SetValueType(v string) *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute {
	s.ValueType = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) SetValueUnit(v string) *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute {
	s.ValueUnit = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationActionsActionParamsValueAttribute) Validate() error {
	return dara.Validate(s)
}
