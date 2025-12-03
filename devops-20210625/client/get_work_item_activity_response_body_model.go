// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemActivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivities(v []*GetWorkItemActivityResponseBodyActivities) *GetWorkItemActivityResponseBody
	GetActivities() []*GetWorkItemActivityResponseBodyActivities
	SetErrorCode(v string) *GetWorkItemActivityResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetWorkItemActivityResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GetWorkItemActivityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkItemActivityResponseBody
	GetSuccess() *bool
}

type GetWorkItemActivityResponseBody struct {
	Activities []*GetWorkItemActivityResponseBodyActivities `json:"activities,omitempty" xml:"activities,omitempty" type:"Repeated"`
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetWorkItemActivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemActivityResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkItemActivityResponseBody) GetActivities() []*GetWorkItemActivityResponseBodyActivities {
	return s.Activities
}

func (s *GetWorkItemActivityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWorkItemActivityResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetWorkItemActivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkItemActivityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkItemActivityResponseBody) SetActivities(v []*GetWorkItemActivityResponseBodyActivities) *GetWorkItemActivityResponseBody {
	s.Activities = v
	return s
}

func (s *GetWorkItemActivityResponseBody) SetErrorCode(v string) *GetWorkItemActivityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkItemActivityResponseBody) SetErrorMsg(v string) *GetWorkItemActivityResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetWorkItemActivityResponseBody) SetRequestId(v string) *GetWorkItemActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkItemActivityResponseBody) SetSuccess(v bool) *GetWorkItemActivityResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkItemActivityResponseBody) Validate() error {
	if s.Activities != nil {
		for _, item := range s.Activities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkItemActivityResponseBodyActivities struct {
	// example:
	//
	// update
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// example:
	//
	// 3201131
	EventId *int64 `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// 1640867079624
	EventTime *int64 `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	// example:
	//
	// workitem.updated
	EventType *string                                              `json:"eventType,omitempty" xml:"eventType,omitempty"`
	NewValue  []*GetWorkItemActivityResponseBodyActivitiesNewValue `json:"newValue,omitempty" xml:"newValue,omitempty" type:"Repeated"`
	OldValue  []*GetWorkItemActivityResponseBodyActivitiesOldValue `json:"oldValue,omitempty" xml:"oldValue,omitempty" type:"Repeated"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// 3201132
	ParentEventId *int64 `json:"parentEventId,omitempty" xml:"parentEventId,omitempty"`
	// example:
	//
	// public
	Property *GetWorkItemActivityResponseBodyActivitiesProperty `json:"property,omitempty" xml:"property,omitempty" type:"Struct"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty" xml:"resourceIdentifier,omitempty"`
}

func (s GetWorkItemActivityResponseBodyActivities) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemActivityResponseBodyActivities) GoString() string {
	return s.String()
}

func (s *GetWorkItemActivityResponseBodyActivities) GetActionType() *string {
	return s.ActionType
}

func (s *GetWorkItemActivityResponseBodyActivities) GetEventId() *int64 {
	return s.EventId
}

func (s *GetWorkItemActivityResponseBodyActivities) GetEventTime() *int64 {
	return s.EventTime
}

func (s *GetWorkItemActivityResponseBodyActivities) GetEventType() *string {
	return s.EventType
}

func (s *GetWorkItemActivityResponseBodyActivities) GetNewValue() []*GetWorkItemActivityResponseBodyActivitiesNewValue {
	return s.NewValue
}

func (s *GetWorkItemActivityResponseBodyActivities) GetOldValue() []*GetWorkItemActivityResponseBodyActivitiesOldValue {
	return s.OldValue
}

func (s *GetWorkItemActivityResponseBodyActivities) GetOperator() *string {
	return s.Operator
}

func (s *GetWorkItemActivityResponseBodyActivities) GetParentEventId() *int64 {
	return s.ParentEventId
}

func (s *GetWorkItemActivityResponseBodyActivities) GetProperty() *GetWorkItemActivityResponseBodyActivitiesProperty {
	return s.Property
}

func (s *GetWorkItemActivityResponseBodyActivities) GetResourceIdentifier() *string {
	return s.ResourceIdentifier
}

func (s *GetWorkItemActivityResponseBodyActivities) SetActionType(v string) *GetWorkItemActivityResponseBodyActivities {
	s.ActionType = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetEventId(v int64) *GetWorkItemActivityResponseBodyActivities {
	s.EventId = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetEventTime(v int64) *GetWorkItemActivityResponseBodyActivities {
	s.EventTime = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetEventType(v string) *GetWorkItemActivityResponseBodyActivities {
	s.EventType = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetNewValue(v []*GetWorkItemActivityResponseBodyActivitiesNewValue) *GetWorkItemActivityResponseBodyActivities {
	s.NewValue = v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetOldValue(v []*GetWorkItemActivityResponseBodyActivitiesOldValue) *GetWorkItemActivityResponseBodyActivities {
	s.OldValue = v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetOperator(v string) *GetWorkItemActivityResponseBodyActivities {
	s.Operator = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetParentEventId(v int64) *GetWorkItemActivityResponseBodyActivities {
	s.ParentEventId = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetProperty(v *GetWorkItemActivityResponseBodyActivitiesProperty) *GetWorkItemActivityResponseBodyActivities {
	s.Property = v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) SetResourceIdentifier(v string) *GetWorkItemActivityResponseBodyActivities {
	s.ResourceIdentifier = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivities) Validate() error {
	if s.NewValue != nil {
		for _, item := range s.NewValue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OldValue != nil {
		for _, item := range s.OldValue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Property != nil {
		if err := s.Property.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkItemActivityResponseBodyActivitiesNewValue struct {
	// example:
	//
	// Sprint-221124
	DisplayValue *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	// example:
	//
	// bed1cca179badeb501a72d1194
	PlainValue *string `json:"plainValue,omitempty" xml:"plainValue,omitempty"`
	// example:
	//
	// Sprint
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetWorkItemActivityResponseBodyActivitiesNewValue) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemActivityResponseBodyActivitiesNewValue) GoString() string {
	return s.String()
}

func (s *GetWorkItemActivityResponseBodyActivitiesNewValue) GetDisplayValue() *string {
	return s.DisplayValue
}

func (s *GetWorkItemActivityResponseBodyActivitiesNewValue) GetPlainValue() *string {
	return s.PlainValue
}

func (s *GetWorkItemActivityResponseBodyActivitiesNewValue) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetWorkItemActivityResponseBodyActivitiesNewValue) SetDisplayValue(v string) *GetWorkItemActivityResponseBodyActivitiesNewValue {
	s.DisplayValue = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesNewValue) SetPlainValue(v string) *GetWorkItemActivityResponseBodyActivitiesNewValue {
	s.PlainValue = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesNewValue) SetResourceType(v string) *GetWorkItemActivityResponseBodyActivitiesNewValue {
	s.ResourceType = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesNewValue) Validate() error {
	return dara.Validate(s)
}

type GetWorkItemActivityResponseBodyActivitiesOldValue struct {
	// example:
	//
	// Sprint-221124
	DisplayValue *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	// example:
	//
	// bed1cca179badeb501a72d1194
	PlainValue *string `json:"plainValue,omitempty" xml:"plainValue,omitempty"`
	// example:
	//
	// Sprint
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetWorkItemActivityResponseBodyActivitiesOldValue) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemActivityResponseBodyActivitiesOldValue) GoString() string {
	return s.String()
}

func (s *GetWorkItemActivityResponseBodyActivitiesOldValue) GetDisplayValue() *string {
	return s.DisplayValue
}

func (s *GetWorkItemActivityResponseBodyActivitiesOldValue) GetPlainValue() *string {
	return s.PlainValue
}

func (s *GetWorkItemActivityResponseBodyActivitiesOldValue) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetWorkItemActivityResponseBodyActivitiesOldValue) SetDisplayValue(v string) *GetWorkItemActivityResponseBodyActivitiesOldValue {
	s.DisplayValue = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesOldValue) SetPlainValue(v string) *GetWorkItemActivityResponseBodyActivitiesOldValue {
	s.PlainValue = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesOldValue) SetResourceType(v string) *GetWorkItemActivityResponseBodyActivitiesOldValue {
	s.ResourceType = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesOldValue) Validate() error {
	return dara.Validate(s)
}

type GetWorkItemActivityResponseBodyActivitiesProperty struct {
	// example:
	//
	// 标题
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// subject
	PropertyIdentifier *string `json:"propertyIdentifier,omitempty" xml:"propertyIdentifier,omitempty"`
	// example:
	//
	// subject
	PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty"`
	// example:
	//
	// null
	PropertyType *string `json:"propertyType,omitempty" xml:"propertyType,omitempty"`
}

func (s GetWorkItemActivityResponseBodyActivitiesProperty) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemActivityResponseBodyActivitiesProperty) GoString() string {
	return s.String()
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) GetPropertyIdentifier() *string {
	return s.PropertyIdentifier
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) GetPropertyName() *string {
	return s.PropertyName
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) GetPropertyType() *string {
	return s.PropertyType
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) SetDisplayName(v string) *GetWorkItemActivityResponseBodyActivitiesProperty {
	s.DisplayName = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) SetPropertyIdentifier(v string) *GetWorkItemActivityResponseBodyActivitiesProperty {
	s.PropertyIdentifier = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) SetPropertyName(v string) *GetWorkItemActivityResponseBodyActivitiesProperty {
	s.PropertyName = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) SetPropertyType(v string) *GetWorkItemActivityResponseBodyActivitiesProperty {
	s.PropertyType = &v
	return s
}

func (s *GetWorkItemActivityResponseBodyActivitiesProperty) Validate() error {
	return dara.Validate(s)
}
