// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalysisConditionFavoriteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAnalysisConditionFavoriteListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeAnalysisConditionFavoriteListResponseBodyResultObject) *DescribeAnalysisConditionFavoriteListResponseBody
	GetResultObject() []*DescribeAnalysisConditionFavoriteListResponseBodyResultObject
}

type DescribeAnalysisConditionFavoriteListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject []*DescribeAnalysisConditionFavoriteListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeAnalysisConditionFavoriteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisConditionFavoriteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisConditionFavoriteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAnalysisConditionFavoriteListResponseBody) GetResultObject() []*DescribeAnalysisConditionFavoriteListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAnalysisConditionFavoriteListResponseBody) SetRequestId(v string) *DescribeAnalysisConditionFavoriteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponseBody) SetResultObject(v []*DescribeAnalysisConditionFavoriteListResponseBodyResultObject) *DescribeAnalysisConditionFavoriteListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAnalysisConditionFavoriteListResponseBodyResultObject struct {
	// Condition value.
	//
	// example:
	//
	// {"currentId":1,"deepCount":0,"list":[{"currentId":2,"deepCount":1,"left":{"code":"accountId","fieldType":"STRING","functionName":"","hasRightVariable":false,"name":"accountId"},"operatorCode":"isNotEmptyWrapper","parentId":1,"right":{"code":"\\"A\\nB\\nC\\"","functionName":"","name":"","rightVariableType":"constant"}}],"parentId":0,"relationship":"and"}
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// Event start timestamp.
	//
	// example:
	//
	// 1752076800000
	EventBeginTime *int64 `json:"eventBeginTime,omitempty" xml:"eventBeginTime,omitempty"`
	// Event codes.
	//
	// example:
	//
	// de_ahespg8137
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Event end time.
	//
	// example:
	//
	// 1753891199000
	EventEndTime *int64 `json:"eventEndTime,omitempty" xml:"eventEndTime,omitempty"`
	// Field name.
	//
	// example:
	//
	// age
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// Field value.
	//
	// example:
	//
	// 20
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Condition name
	//
	// example:
	//
	// 查询条件1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Type, BASIC: Basic query, ADVANCE: Advanced query, BATCH: Batch query
	//
	// example:
	//
	// BASIC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAnalysisConditionFavoriteListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisConditionFavoriteListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) GetCondition() *string {
	return s.Condition
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) GetEventBeginTime() *int64 {
	return s.EventBeginTime
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) GetEventEndTime() *int64 {
	return s.EventEndTime
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) GetType() *string {
	return s.Type
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) SetCondition(v string) *DescribeAnalysisConditionFavoriteListResponseBodyResultObject {
	s.Condition = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) SetEventBeginTime(v int64) *DescribeAnalysisConditionFavoriteListResponseBodyResultObject {
	s.EventBeginTime = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) SetEventCodes(v string) *DescribeAnalysisConditionFavoriteListResponseBodyResultObject {
	s.EventCodes = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) SetEventEndTime(v int64) *DescribeAnalysisConditionFavoriteListResponseBodyResultObject {
	s.EventEndTime = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) SetFieldName(v string) *DescribeAnalysisConditionFavoriteListResponseBodyResultObject {
	s.FieldName = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) SetFieldValue(v string) *DescribeAnalysisConditionFavoriteListResponseBodyResultObject {
	s.FieldValue = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) SetId(v int64) *DescribeAnalysisConditionFavoriteListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) SetName(v string) *DescribeAnalysisConditionFavoriteListResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) SetType(v string) *DescribeAnalysisConditionFavoriteListResponseBodyResultObject {
	s.Type = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
