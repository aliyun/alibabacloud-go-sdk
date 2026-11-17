// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkItemAuditResponseBody
	GetCode() *string
	SetItems(v []*GetWorkItemAuditResponseBodyItems) *GetWorkItemAuditResponseBody
	GetItems() []*GetWorkItemAuditResponseBodyItems
	SetMessage(v string) *GetWorkItemAuditResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkItemAuditResponseBody
	GetRequestId() *string
}

type GetWorkItemAuditResponseBody struct {
	// The common business status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The complete work item audit records sorted in descending order of event time.
	//
	// example:
	//
	// [{"eventId":"evt_example","eventType":"WORK_ITEM_UPDATED","occurredAt":"2026-11-05T08:00:00Z","status":"SUCCESS"}]
	Items []*GetWorkItemAuditResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The common error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetWorkItemAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemAuditResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkItemAuditResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkItemAuditResponseBody) GetItems() []*GetWorkItemAuditResponseBodyItems {
	return s.Items
}

func (s *GetWorkItemAuditResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkItemAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkItemAuditResponseBody) SetCode(v string) *GetWorkItemAuditResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkItemAuditResponseBody) SetItems(v []*GetWorkItemAuditResponseBodyItems) *GetWorkItemAuditResponseBody {
	s.Items = v
	return s
}

func (s *GetWorkItemAuditResponseBody) SetMessage(v string) *GetWorkItemAuditResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkItemAuditResponseBody) SetRequestId(v string) *GetWorkItemAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkItemAuditResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkItemAuditResponseBodyItems struct {
	// The list of field change details. Each item includes the changed field and its values before and after the change.
	//
	// example:
	//
	// [{"field":"title","before":"Old title","after":"New title"}]
	Changes interface{} `json:"changes,omitempty" xml:"changes,omitempty"`
	// The event description.
	//
	// example:
	//
	// The work item title is changed to Quarterly Sales Opportunity Review
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The unique identifier of the audit event.
	//
	// example:
	//
	// evt_20261105_0001
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// The audit event type, such as WORK_ITEM_UPDATED or TODO_COMPLETED.
	//
	// example:
	//
	// WORK_ITEM_UPDATED
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// The time when the event occurred, in ISO 8601 format.
	//
	// example:
	//
	// 2026-11-05T08:00:00Z
	OccurredAt *string `json:"occurredAt,omitempty" xml:"occurredAt,omitempty"`
	// The display name of the operator.
	//
	// example:
	//
	// Zhang San
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The user ID of the operator.
	//
	// example:
	//
	// 10001
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// The list of questions to be confirmed related to the event.
	//
	// example:
	//
	// [{"question":"Whether to continue execution"}]
	Questions []interface{} `json:"questions,omitempty" xml:"questions,omitempty" type:"Repeated"`
	// The event result payload as a JSON object. If no result is available, this field is empty.
	//
	// example:
	//
	// {"result":"ok"}
	ResultPayload interface{} `json:"resultPayload,omitempty" xml:"resultPayload,omitempty"`
	// The event status.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The event title.
	//
	// example:
	//
	// Investigate sales opportunities
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The ID of the associated to-do item. If no to-do item is associated, an empty string is returned.
	//
	// example:
	//
	// todo_20261105_0007
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s GetWorkItemAuditResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemAuditResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetWorkItemAuditResponseBodyItems) GetChanges() interface{} {
	return s.Changes
}

func (s *GetWorkItemAuditResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *GetWorkItemAuditResponseBodyItems) GetEventId() *string {
	return s.EventId
}

func (s *GetWorkItemAuditResponseBodyItems) GetEventType() *string {
	return s.EventType
}

func (s *GetWorkItemAuditResponseBodyItems) GetOccurredAt() *string {
	return s.OccurredAt
}

func (s *GetWorkItemAuditResponseBodyItems) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetWorkItemAuditResponseBodyItems) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *GetWorkItemAuditResponseBodyItems) GetQuestions() []interface{} {
	return s.Questions
}

func (s *GetWorkItemAuditResponseBodyItems) GetResultPayload() interface{} {
	return s.ResultPayload
}

func (s *GetWorkItemAuditResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *GetWorkItemAuditResponseBodyItems) GetTitle() *string {
	return s.Title
}

func (s *GetWorkItemAuditResponseBodyItems) GetTodoId() *string {
	return s.TodoId
}

func (s *GetWorkItemAuditResponseBodyItems) SetChanges(v interface{}) *GetWorkItemAuditResponseBodyItems {
	s.Changes = v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetDescription(v string) *GetWorkItemAuditResponseBodyItems {
	s.Description = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetEventId(v string) *GetWorkItemAuditResponseBodyItems {
	s.EventId = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetEventType(v string) *GetWorkItemAuditResponseBodyItems {
	s.EventType = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetOccurredAt(v string) *GetWorkItemAuditResponseBodyItems {
	s.OccurredAt = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetOperatorName(v string) *GetWorkItemAuditResponseBodyItems {
	s.OperatorName = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetOperatorUserId(v string) *GetWorkItemAuditResponseBodyItems {
	s.OperatorUserId = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetQuestions(v []interface{}) *GetWorkItemAuditResponseBodyItems {
	s.Questions = v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetResultPayload(v interface{}) *GetWorkItemAuditResponseBodyItems {
	s.ResultPayload = v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetStatus(v string) *GetWorkItemAuditResponseBodyItems {
	s.Status = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetTitle(v string) *GetWorkItemAuditResponseBodyItems {
	s.Title = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetTodoId(v string) *GetWorkItemAuditResponseBodyItems {
	s.TodoId = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
