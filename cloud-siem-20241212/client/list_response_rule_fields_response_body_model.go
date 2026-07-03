// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResponseRuleFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListResponseRuleFields(v []*ListResponseRuleFieldsResponseBodyListResponseRuleFields) *ListResponseRuleFieldsResponseBody
	GetListResponseRuleFields() []*ListResponseRuleFieldsResponseBodyListResponseRuleFields
	SetMaxResults(v int32) *ListResponseRuleFieldsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListResponseRuleFieldsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResponseRuleFieldsResponseBody
	GetRequestId() *string
}

type ListResponseRuleFieldsResponseBody struct {
	// The list of response rule fields.
	ListResponseRuleFields []*ListResponseRuleFieldsResponseBodyListResponseRuleFields `json:"ListResponseRuleFields,omitempty" xml:"ListResponseRuleFields,omitempty" type:"Repeated"`
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. Valid values: leave this parameter empty for the first query or if no more results exist. If a next query exists, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResponseRuleFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResponseRuleFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResponseRuleFieldsResponseBody) GetListResponseRuleFields() []*ListResponseRuleFieldsResponseBodyListResponseRuleFields {
	return s.ListResponseRuleFields
}

func (s *ListResponseRuleFieldsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResponseRuleFieldsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResponseRuleFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResponseRuleFieldsResponseBody) SetListResponseRuleFields(v []*ListResponseRuleFieldsResponseBodyListResponseRuleFields) *ListResponseRuleFieldsResponseBody {
	s.ListResponseRuleFields = v
	return s
}

func (s *ListResponseRuleFieldsResponseBody) SetMaxResults(v int32) *ListResponseRuleFieldsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBody) SetNextToken(v string) *ListResponseRuleFieldsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBody) SetRequestId(v string) *ListResponseRuleFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBody) Validate() error {
	if s.ListResponseRuleFields != nil {
		for _, item := range s.ListResponseRuleFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResponseRuleFieldsResponseBodyListResponseRuleFields struct {
	// The data type of the automated response rule condition field.
	//
	// example:
	//
	// ip
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The whitelisted field.
	//
	// example:
	//
	// appname
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The name of the rule field.
	//
	// example:
	//
	// OriginIP,ConsoleLog,CPUTime,Duration,ErrorCode,ErrorMessage,ResponseSize,ResponseStatus,RoutineName,ClientRequestID,LogTimestamp,FetchStatus,SubRequestID
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The normalization object type to which the field belongs.
	//
	// example:
	//
	// alert
	FieldNormalization *string `json:"FieldNormalization,omitempty" xml:"FieldNormalization,omitempty"`
	// The list of optional enumeration values for the field. This parameter is not returned if no enumeration values are available.
	RightValue []*ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue `json:"RightValue,omitempty" xml:"RightValue,omitempty" type:"Repeated"`
	// The English descriptions of the operators.
	SupportOperators []*ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators `json:"SupportOperators,omitempty" xml:"SupportOperators,omitempty" type:"Repeated"`
}

func (s ListResponseRuleFieldsResponseBodyListResponseRuleFields) String() string {
	return dara.Prettify(s)
}

func (s ListResponseRuleFieldsResponseBodyListResponseRuleFields) GoString() string {
	return s.String()
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) GetDataType() *string {
	return s.DataType
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) GetField() *string {
	return s.Field
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) GetFieldName() *string {
	return s.FieldName
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) GetFieldNormalization() *string {
	return s.FieldNormalization
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) GetRightValue() []*ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue {
	return s.RightValue
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) GetSupportOperators() []*ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators {
	return s.SupportOperators
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) SetDataType(v string) *ListResponseRuleFieldsResponseBodyListResponseRuleFields {
	s.DataType = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) SetField(v string) *ListResponseRuleFieldsResponseBodyListResponseRuleFields {
	s.Field = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) SetFieldName(v string) *ListResponseRuleFieldsResponseBodyListResponseRuleFields {
	s.FieldName = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) SetFieldNormalization(v string) *ListResponseRuleFieldsResponseBodyListResponseRuleFields {
	s.FieldNormalization = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) SetRightValue(v []*ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue) *ListResponseRuleFieldsResponseBodyListResponseRuleFields {
	s.RightValue = v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) SetSupportOperators(v []*ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) *ListResponseRuleFieldsResponseBodyListResponseRuleFields {
	s.SupportOperators = v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFields) Validate() error {
	if s.RightValue != nil {
		for _, item := range s.RightValue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SupportOperators != nil {
		for _, item := range s.SupportOperators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue struct {
	// The right-side value.
	//
	// example:
	//
	// dev_selectdb_cluster
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The display name of the enumeration value.
	//
	// example:
	//
	// high
	ValueName *string `json:"ValueName,omitempty" xml:"ValueName,omitempty"`
}

func (s ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue) String() string {
	return dara.Prettify(s)
}

func (s ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue) GoString() string {
	return s.String()
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue) GetValue() *string {
	return s.Value
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue) GetValueName() *string {
	return s.ValueName
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue) SetValue(v string) *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue {
	s.Value = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue) SetValueName(v string) *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue {
	s.ValueName = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsRightValue) Validate() error {
	return dara.Validate(s)
}

type ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators struct {
	// Indicates whether a right-side value is required. Valid values:
	//
	// - true: Required.
	//
	// - false: Not required.
	//
	// example:
	//
	// true
	HasRightValue *string `json:"HasRightValue,omitempty" xml:"HasRightValue,omitempty"`
	// The position of the operator in the operator list.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The aggregation method for the dispatch rule condition. Valid values:
	//
	// - `=`: equal to
	//
	// - `<>`: not equal to
	//
	// - `in`: contains
	//
	// - `not in`: does not contain
	//
	// - `REGEXP`: matches the regular expression
	//
	// - `NOT REGEXP`: does not match the regular expression
	//
	// example:
	//
	// BETWEEN
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The display name of the operator.
	//
	// example:
	//
	// autotest-operator
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The data types supported by the current operator, separated by commas.
	//
	// example:
	//
	// ip
	SupportDataType *string `json:"SupportDataType,omitempty" xml:"SupportDataType,omitempty"`
}

func (s ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) String() string {
	return dara.Prettify(s)
}

func (s ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) GoString() string {
	return s.String()
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) GetHasRightValue() *string {
	return s.HasRightValue
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) GetIndex() *int32 {
	return s.Index
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) GetOperator() *string {
	return s.Operator
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) GetOperatorName() *string {
	return s.OperatorName
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) GetSupportDataType() *string {
	return s.SupportDataType
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) SetHasRightValue(v string) *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators {
	s.HasRightValue = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) SetIndex(v int32) *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators {
	s.Index = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) SetOperator(v string) *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators {
	s.Operator = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) SetOperatorName(v string) *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators {
	s.OperatorName = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) SetSupportDataType(v string) *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators {
	s.SupportDataType = &v
	return s
}

func (s *ListResponseRuleFieldsResponseBodyListResponseRuleFieldsSupportOperators) Validate() error {
	return dara.Validate(s)
}
