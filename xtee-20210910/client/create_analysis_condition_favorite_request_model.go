// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnalysisConditionFavoriteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateAnalysisConditionFavoriteRequest
	GetLang() *string
	SetCondition(v string) *CreateAnalysisConditionFavoriteRequest
	GetCondition() *string
	SetEventBeginTime(v int64) *CreateAnalysisConditionFavoriteRequest
	GetEventBeginTime() *int64
	SetEventCodes(v string) *CreateAnalysisConditionFavoriteRequest
	GetEventCodes() *string
	SetEventEndTime(v int64) *CreateAnalysisConditionFavoriteRequest
	GetEventEndTime() *int64
	SetFieldName(v string) *CreateAnalysisConditionFavoriteRequest
	GetFieldName() *string
	SetFieldValue(v string) *CreateAnalysisConditionFavoriteRequest
	GetFieldValue() *string
	SetName(v string) *CreateAnalysisConditionFavoriteRequest
	GetName() *string
	SetRegId(v string) *CreateAnalysisConditionFavoriteRequest
	GetRegId() *string
	SetType(v string) *CreateAnalysisConditionFavoriteRequest
	GetType() *string
}

type CreateAnalysisConditionFavoriteRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// {"relationship":"and","list":[{"deepCount":1,"left":{"hasRightVariable":true,"fieldType":"INT","functionName":"","leftVariableType":"NATIVE","name":"DEtest222","operatorCode":"equals"},"right":{"rightVariableType":"constant","name":"11","functionName":""},"operatorCode":"equals"}]}
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// Start time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1751299200000
	EventBeginTime *int64 `json:"eventBeginTime,omitempty" xml:"eventBeginTime,omitempty"`
	// Event codes, separated by commas
	//
	// This parameter is required.
	//
	// example:
	//
	// ["de_ahqhsw7665","de_agbzfi5134"]
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1753372799000
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
	// 10
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	// Condition favorite name
	//
	// This parameter is required.
	//
	// example:
	//
	// 条件一
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Query type
	//
	// This parameter is required.
	//
	// example:
	//
	// BASIC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAnalysisConditionFavoriteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAnalysisConditionFavoriteRequest) GoString() string {
	return s.String()
}

func (s *CreateAnalysisConditionFavoriteRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateAnalysisConditionFavoriteRequest) GetCondition() *string {
	return s.Condition
}

func (s *CreateAnalysisConditionFavoriteRequest) GetEventBeginTime() *int64 {
	return s.EventBeginTime
}

func (s *CreateAnalysisConditionFavoriteRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *CreateAnalysisConditionFavoriteRequest) GetEventEndTime() *int64 {
	return s.EventEndTime
}

func (s *CreateAnalysisConditionFavoriteRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *CreateAnalysisConditionFavoriteRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *CreateAnalysisConditionFavoriteRequest) GetName() *string {
	return s.Name
}

func (s *CreateAnalysisConditionFavoriteRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateAnalysisConditionFavoriteRequest) GetType() *string {
	return s.Type
}

func (s *CreateAnalysisConditionFavoriteRequest) SetLang(v string) *CreateAnalysisConditionFavoriteRequest {
	s.Lang = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteRequest) SetCondition(v string) *CreateAnalysisConditionFavoriteRequest {
	s.Condition = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteRequest) SetEventBeginTime(v int64) *CreateAnalysisConditionFavoriteRequest {
	s.EventBeginTime = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteRequest) SetEventCodes(v string) *CreateAnalysisConditionFavoriteRequest {
	s.EventCodes = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteRequest) SetEventEndTime(v int64) *CreateAnalysisConditionFavoriteRequest {
	s.EventEndTime = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteRequest) SetFieldName(v string) *CreateAnalysisConditionFavoriteRequest {
	s.FieldName = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteRequest) SetFieldValue(v string) *CreateAnalysisConditionFavoriteRequest {
	s.FieldValue = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteRequest) SetName(v string) *CreateAnalysisConditionFavoriteRequest {
	s.Name = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteRequest) SetRegId(v string) *CreateAnalysisConditionFavoriteRequest {
	s.RegId = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteRequest) SetType(v string) *CreateAnalysisConditionFavoriteRequest {
	s.Type = &v
	return s
}

func (s *CreateAnalysisConditionFavoriteRequest) Validate() error {
	return dara.Validate(s)
}
