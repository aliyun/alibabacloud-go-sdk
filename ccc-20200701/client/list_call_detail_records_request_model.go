// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallDetailRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListCallDetailRecordsRequest
	GetAgentId() *string
	SetCalledNumber(v string) *ListCallDetailRecordsRequest
	GetCalledNumber() *string
	SetCallingNumber(v string) *ListCallDetailRecordsRequest
	GetCallingNumber() *string
	SetContactDisposition(v string) *ListCallDetailRecordsRequest
	GetContactDisposition() *string
	SetContactDispositionList(v string) *ListCallDetailRecordsRequest
	GetContactDispositionList() *string
	SetContactId(v string) *ListCallDetailRecordsRequest
	GetContactId() *string
	SetContactType(v string) *ListCallDetailRecordsRequest
	GetContactType() *string
	SetContactTypeList(v string) *ListCallDetailRecordsRequest
	GetContactTypeList() *string
	SetCriteria(v string) *ListCallDetailRecordsRequest
	GetCriteria() *string
	SetEarlyMediaStateList(v string) *ListCallDetailRecordsRequest
	GetEarlyMediaStateList() *string
	SetEndTime(v int64) *ListCallDetailRecordsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListCallDetailRecordsRequest
	GetInstanceId() *string
	SetOrderByField(v string) *ListCallDetailRecordsRequest
	GetOrderByField() *string
	SetPageNumber(v int32) *ListCallDetailRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCallDetailRecordsRequest
	GetPageSize() *int32
	SetSatisfactionDescriptionList(v string) *ListCallDetailRecordsRequest
	GetSatisfactionDescriptionList() *string
	SetSatisfactionList(v string) *ListCallDetailRecordsRequest
	GetSatisfactionList() *string
	SetSatisfactionSurveyChannel(v string) *ListCallDetailRecordsRequest
	GetSatisfactionSurveyChannel() *string
	SetSkillGroupId(v string) *ListCallDetailRecordsRequest
	GetSkillGroupId() *string
	SetSortOrder(v string) *ListCallDetailRecordsRequest
	GetSortOrder() *string
	SetStartTime(v int64) *ListCallDetailRecordsRequest
	GetStartTime() *int64
}

type ListCallDetailRecordsRequest struct {
	// Filter by agent ID.
	//
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Filter by called number.
	//
	// example:
	//
	// 1320523****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// Filter by calling number.
	//
	// example:
	//
	// 07353988****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// Filter by disposition type. Note: Disposition reasons such as voicemail, transfer to agent failed, queuing timeout, queuing overflow, and IVR abnormal are only displayed if the customer has configured a disposition reason node. If no such node is configured and there is no transfer-to-agent module in the IVR, the disposition reason defaults to "Abandoned in IVR".
	//
	// example:
	//
	// Success
	ContactDisposition *string `json:"ContactDisposition,omitempty" xml:"ContactDisposition,omitempty"`
	// Filter by hang-up reason list.
	//
	// example:
	//
	// ["Success","NoAnswer"]
	ContactDispositionList *string `json:"ContactDispositionList,omitempty" xml:"ContactDispositionList,omitempty"`
	// Query the record of a specific call by specifying a contactId. The contactId can be obtained from the softphone software development kit (SDK) during a call. If this parameter is provided, all other query parameters are automatically ignored.
	//
	// example:
	//
	// job-12515239414412****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// Filter by call type.
	//
	// example:
	//
	// Outbound
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// Filter by contact type list.
	//
	// example:
	//
	// ["Internal","Inbound"]
	ContactTypeList *string `json:"ContactTypeList,omitempty" xml:"ContactTypeList,omitempty"`
	// Perform a fuzzy query based on the calling or called number. The value must be a JSON string containing only one field, phoneNumber, which can be the full number or a partial segment of either the calling or called number.
	//
	// example:
	//
	// {"phoneNumber":"0735"}
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// Filter by a list of reasons for failed connection.
	//
	// example:
	//
	// ["NotConnected","NoAnswer"]
	EarlyMediaStateList *string `json:"EarlyMediaStateList,omitempty" xml:"EarlyMediaStateList,omitempty"`
	// End time of the historical data to retrieve. The default value is the current time, in UNIX timestamp format with millisecond precision.
	//
	// example:
	//
	// 1532707199000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Sorting field. Optional. Default value is startTime (call start time).
	//
	// example:
	//
	// startTime
	OrderByField *string `json:"OrderByField,omitempty" xml:"OrderByField,omitempty"`
	// Page number for paging, ranging from 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, ranging from 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filter by satisfaction description list. The description content is Custom by the Customer.
	//
	// example:
	//
	// ["满意","一般"]
	SatisfactionDescriptionList *string `json:"SatisfactionDescriptionList,omitempty" xml:"SatisfactionDescriptionList,omitempty"`
	// Filter by satisfaction List. Separate multiple satisfaction Results with commas.
	//
	// example:
	//
	// ["1","3"]
	SatisfactionList *string `json:"SatisfactionList,omitempty" xml:"SatisfactionList,omitempty"`
	// Filter by satisfaction survey channel.
	//
	// example:
	//
	// IVR
	SatisfactionSurveyChannel *string `json:"SatisfactionSurveyChannel,omitempty" xml:"SatisfactionSurveyChannel,omitempty"`
	// Filter by skill group ID.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// Sorting order. This parameter is optional and defaults to descending.
	//
	// example:
	//
	// DESC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// Start time of the historical data to retrieve. The default value is 00:00:00 of the current day, in UNIX timestamp format with millisecond precision.
	//
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListCallDetailRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListCallDetailRecordsRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *ListCallDetailRecordsRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *ListCallDetailRecordsRequest) GetContactDisposition() *string {
	return s.ContactDisposition
}

func (s *ListCallDetailRecordsRequest) GetContactDispositionList() *string {
	return s.ContactDispositionList
}

func (s *ListCallDetailRecordsRequest) GetContactId() *string {
	return s.ContactId
}

func (s *ListCallDetailRecordsRequest) GetContactType() *string {
	return s.ContactType
}

func (s *ListCallDetailRecordsRequest) GetContactTypeList() *string {
	return s.ContactTypeList
}

func (s *ListCallDetailRecordsRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *ListCallDetailRecordsRequest) GetEarlyMediaStateList() *string {
	return s.EarlyMediaStateList
}

func (s *ListCallDetailRecordsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListCallDetailRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCallDetailRecordsRequest) GetOrderByField() *string {
	return s.OrderByField
}

func (s *ListCallDetailRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallDetailRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallDetailRecordsRequest) GetSatisfactionDescriptionList() *string {
	return s.SatisfactionDescriptionList
}

func (s *ListCallDetailRecordsRequest) GetSatisfactionList() *string {
	return s.SatisfactionList
}

func (s *ListCallDetailRecordsRequest) GetSatisfactionSurveyChannel() *string {
	return s.SatisfactionSurveyChannel
}

func (s *ListCallDetailRecordsRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListCallDetailRecordsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListCallDetailRecordsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCallDetailRecordsRequest) SetAgentId(v string) *ListCallDetailRecordsRequest {
	s.AgentId = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetCalledNumber(v string) *ListCallDetailRecordsRequest {
	s.CalledNumber = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetCallingNumber(v string) *ListCallDetailRecordsRequest {
	s.CallingNumber = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetContactDisposition(v string) *ListCallDetailRecordsRequest {
	s.ContactDisposition = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetContactDispositionList(v string) *ListCallDetailRecordsRequest {
	s.ContactDispositionList = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetContactId(v string) *ListCallDetailRecordsRequest {
	s.ContactId = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetContactType(v string) *ListCallDetailRecordsRequest {
	s.ContactType = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetContactTypeList(v string) *ListCallDetailRecordsRequest {
	s.ContactTypeList = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetCriteria(v string) *ListCallDetailRecordsRequest {
	s.Criteria = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetEarlyMediaStateList(v string) *ListCallDetailRecordsRequest {
	s.EarlyMediaStateList = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetEndTime(v int64) *ListCallDetailRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetInstanceId(v string) *ListCallDetailRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetOrderByField(v string) *ListCallDetailRecordsRequest {
	s.OrderByField = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetPageNumber(v int32) *ListCallDetailRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetPageSize(v int32) *ListCallDetailRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetSatisfactionDescriptionList(v string) *ListCallDetailRecordsRequest {
	s.SatisfactionDescriptionList = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetSatisfactionList(v string) *ListCallDetailRecordsRequest {
	s.SatisfactionList = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetSatisfactionSurveyChannel(v string) *ListCallDetailRecordsRequest {
	s.SatisfactionSurveyChannel = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetSkillGroupId(v string) *ListCallDetailRecordsRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetSortOrder(v string) *ListCallDetailRecordsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetStartTime(v int64) *ListCallDetailRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *ListCallDetailRecordsRequest) Validate() error {
	return dara.Validate(s)
}
