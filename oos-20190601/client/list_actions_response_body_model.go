// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*ListActionsResponseBodyActions) *ListActionsResponseBody
	GetActions() []*ListActionsResponseBodyActions
	SetMaxResults(v int32) *ListActionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListActionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListActionsResponseBody
	GetRequestId() *string
}

type ListActionsResponseBody struct {
	// The details of the actions.
	Actions []*ListActionsResponseBodyActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F9154C02-F847-4563-BB6A-6DD01A4F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListActionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListActionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListActionsResponseBody) GetActions() []*ListActionsResponseBodyActions {
	return s.Actions
}

func (s *ListActionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListActionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListActionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListActionsResponseBody) SetActions(v []*ListActionsResponseBodyActions) *ListActionsResponseBody {
	s.Actions = v
	return s
}

func (s *ListActionsResponseBody) SetMaxResults(v int32) *ListActionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListActionsResponseBody) SetNextToken(v string) *ListActionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListActionsResponseBody) SetRequestId(v string) *ListActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListActionsResponseBody) Validate() error {
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

type ListActionsResponseBodyActions struct {
	// The type of the action.
	//
	// 1.  Atomic actions
	//
	//     	- Atomic.API
	//
	//     	- Atomic.Trigger
	//
	//     	- Atomic.Control
	//
	//     	- Atomic.Embedded
	//
	// 2.  Cloud product actions
	//
	//     	- Product.ECS
	//
	//     	- Product.RDS
	//
	//     	- Product.VPC
	//
	//     	- Product.FC
	//
	//     	- ...
	//
	// example:
	//
	// ACS::Template
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The time when the action was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the action.
	//
	// example:
	//
	// ReplaceSystemDisk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the action.
	//
	// example:
	//
	// ACS::ECS::ReplaceSystemDisk
	OOSActionName *string `json:"OOSActionName,omitempty" xml:"OOSActionName,omitempty"`
	// The number of times that the action is used.
	//
	// example:
	//
	// 5
	Popularity *int32 `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	// The parameters of the action.
	//
	// example:
	//
	// { "ImageId": { "Description": "The mirror ID you will use when resetting the system", "Type": "String" }, "InstanceId": { "Description": "the instance id that you will handle .", "Type": "String" } }
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The version of the template that corresponds to the action.
	//
	// >  For atomic actions, this parameter is not returned.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s ListActionsResponseBodyActions) String() string {
	return dara.Prettify(s)
}

func (s ListActionsResponseBodyActions) GoString() string {
	return s.String()
}

func (s *ListActionsResponseBodyActions) GetActionType() *string {
	return s.ActionType
}

func (s *ListActionsResponseBodyActions) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *ListActionsResponseBodyActions) GetDescription() *string {
	return s.Description
}

func (s *ListActionsResponseBodyActions) GetOOSActionName() *string {
	return s.OOSActionName
}

func (s *ListActionsResponseBodyActions) GetPopularity() *int32 {
	return s.Popularity
}

func (s *ListActionsResponseBodyActions) GetProperties() *string {
	return s.Properties
}

func (s *ListActionsResponseBodyActions) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListActionsResponseBodyActions) SetActionType(v string) *ListActionsResponseBodyActions {
	s.ActionType = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetCreatedDate(v string) *ListActionsResponseBodyActions {
	s.CreatedDate = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetDescription(v string) *ListActionsResponseBodyActions {
	s.Description = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetOOSActionName(v string) *ListActionsResponseBodyActions {
	s.OOSActionName = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetPopularity(v int32) *ListActionsResponseBodyActions {
	s.Popularity = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetProperties(v string) *ListActionsResponseBodyActions {
	s.Properties = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetTemplateVersion(v string) *ListActionsResponseBodyActions {
	s.TemplateVersion = &v
	return s
}

func (s *ListActionsResponseBodyActions) Validate() error {
	return dara.Validate(s)
}
