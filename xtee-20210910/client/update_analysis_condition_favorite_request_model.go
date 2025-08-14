// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAnalysisConditionFavoriteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateAnalysisConditionFavoriteRequest
	GetLang() *string
	SetCondition(v string) *UpdateAnalysisConditionFavoriteRequest
	GetCondition() *string
	SetEventBeginTime(v int64) *UpdateAnalysisConditionFavoriteRequest
	GetEventBeginTime() *int64
	SetEventCode(v string) *UpdateAnalysisConditionFavoriteRequest
	GetEventCode() *string
	SetEventEndTime(v int64) *UpdateAnalysisConditionFavoriteRequest
	GetEventEndTime() *int64
	SetFieldName(v string) *UpdateAnalysisConditionFavoriteRequest
	GetFieldName() *string
	SetFieldValue(v string) *UpdateAnalysisConditionFavoriteRequest
	GetFieldValue() *string
	SetId(v int64) *UpdateAnalysisConditionFavoriteRequest
	GetId() *int64
	SetName(v string) *UpdateAnalysisConditionFavoriteRequest
	GetName() *string
	SetRegId(v string) *UpdateAnalysisConditionFavoriteRequest
	GetRegId() *string
	SetType(v string) *UpdateAnalysisConditionFavoriteRequest
	GetType() *string
}

type UpdateAnalysisConditionFavoriteRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Condition value.
	//
	// example:
	//
	// {\\"relationship\\":\\"and\\",\\"list\\":[{\\"deepCount\\":1,\\"left\\":{\\"hasRightVariable\\":true,\\"fieldType\\":\\"INT\\",\\"functionName\\":\\"\\",\\"leftVariableType\\":\\"NATIVE\\",\\"name\\":\\"DEtest222\\",\\"operatorCode\\":\\"equals\\"},\\"right\\":{\\"rightVariableType\\":\\"constant\\",\\"name\\":\\"9007199254\\",\\"functionName\\":\\"\\"},\\"operatorCode\\":\\"equals\\"}]}
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// Start time, accurate to milliseconds (ms).
	//
	// example:
	//
	// 1752076800000
	EventBeginTime *int64 `json:"eventBeginTime,omitempty" xml:"eventBeginTime,omitempty"`
	// Event code
	//
	// example:
	//
	// de_ajnoqe2016
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// example:
	//
	// 1753891199000
	EventEndTime *int64 `json:"eventEndTime,omitempty" xml:"eventEndTime,omitempty"`
	// Field name
	//
	// example:
	//
	// age
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// Field value
	//
	// example:
	//
	// 20
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	// Primary key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Condition name
	//
	// example:
	//
	// 查询条件1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Type, BASIC: Basic query, ADVANCE: Advanced query, BATCH: Batch query
	//
	// example:
	//
	// BASIC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAnalysisConditionFavoriteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAnalysisConditionFavoriteRequest) GoString() string {
	return s.String()
}

func (s *UpdateAnalysisConditionFavoriteRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateAnalysisConditionFavoriteRequest) GetCondition() *string {
	return s.Condition
}

func (s *UpdateAnalysisConditionFavoriteRequest) GetEventBeginTime() *int64 {
	return s.EventBeginTime
}

func (s *UpdateAnalysisConditionFavoriteRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *UpdateAnalysisConditionFavoriteRequest) GetEventEndTime() *int64 {
	return s.EventEndTime
}

func (s *UpdateAnalysisConditionFavoriteRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *UpdateAnalysisConditionFavoriteRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *UpdateAnalysisConditionFavoriteRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateAnalysisConditionFavoriteRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAnalysisConditionFavoriteRequest) GetRegId() *string {
	return s.RegId
}

func (s *UpdateAnalysisConditionFavoriteRequest) GetType() *string {
	return s.Type
}

func (s *UpdateAnalysisConditionFavoriteRequest) SetLang(v string) *UpdateAnalysisConditionFavoriteRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteRequest) SetCondition(v string) *UpdateAnalysisConditionFavoriteRequest {
	s.Condition = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteRequest) SetEventBeginTime(v int64) *UpdateAnalysisConditionFavoriteRequest {
	s.EventBeginTime = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteRequest) SetEventCode(v string) *UpdateAnalysisConditionFavoriteRequest {
	s.EventCode = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteRequest) SetEventEndTime(v int64) *UpdateAnalysisConditionFavoriteRequest {
	s.EventEndTime = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteRequest) SetFieldName(v string) *UpdateAnalysisConditionFavoriteRequest {
	s.FieldName = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteRequest) SetFieldValue(v string) *UpdateAnalysisConditionFavoriteRequest {
	s.FieldValue = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteRequest) SetId(v int64) *UpdateAnalysisConditionFavoriteRequest {
	s.Id = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteRequest) SetName(v string) *UpdateAnalysisConditionFavoriteRequest {
	s.Name = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteRequest) SetRegId(v string) *UpdateAnalysisConditionFavoriteRequest {
	s.RegId = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteRequest) SetType(v string) *UpdateAnalysisConditionFavoriteRequest {
	s.Type = &v
	return s
}

func (s *UpdateAnalysisConditionFavoriteRequest) Validate() error {
	return dara.Validate(s)
}
