// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllRootVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAllRootVariableResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeAllRootVariableResponseBodyResultObject) *DescribeAllRootVariableResponseBody
	GetResultObject() []*DescribeAllRootVariableResponseBodyResultObject
}

type DescribeAllRootVariableResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeAllRootVariableResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeAllRootVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllRootVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllRootVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAllRootVariableResponseBody) GetResultObject() []*DescribeAllRootVariableResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAllRootVariableResponseBody) SetRequestId(v string) *DescribeAllRootVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllRootVariableResponseBody) SetResultObject(v []*DescribeAllRootVariableResponseBodyResultObject) *DescribeAllRootVariableResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAllRootVariableResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAllRootVariableResponseBodyResultObject struct {
	// Variable code
	//
	// example:
	//
	// age
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Variable description.
	//
	// example:
	//
	// 年龄
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Display type and group label
	//
	// example:
	//
	// NATIVE
	DisplayType *string `json:"displayType,omitempty" xml:"displayType,omitempty"`
	// Favorite flag
	//
	// example:
	//
	// true
	FavoriteFlag *bool `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	// Field ranking
	//
	// example:
	//
	// 1
	FieldRank *int64 `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	// Field type.
	//
	// example:
	//
	// STRING
	FieldType *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	// Variable ID.
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Input field type.
	//
	// example:
	//
	// STRING
	InputFieldType *string `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	// Variable input.
	//
	// example:
	//
	// age
	Inputs *string `json:"inputs,omitempty" xml:"inputs,omitempty"`
	// Variable name.
	//
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Maximum cross-sectional area of the checkbox.
	OutputThreshold *DescribeAllRootVariableResponseBodyResultObjectOutputThreshold `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	// Data source
	//
	// example:
	//
	// SAF
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// Title.
	//
	// example:
	//
	// 年龄
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Variable type.
	//
	// example:
	//
	// NATIVE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAllRootVariableResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllRootVariableResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetCode() *string {
	return s.Code
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetFieldRank() *int64 {
	return s.FieldRank
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetOutputThreshold() *DescribeAllRootVariableResponseBodyResultObjectOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeAllRootVariableResponseBodyResultObject) GetType() *string {
	return s.Type
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetCode(v string) *DescribeAllRootVariableResponseBodyResultObject {
	s.Code = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetDescription(v string) *DescribeAllRootVariableResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetDisplayType(v string) *DescribeAllRootVariableResponseBodyResultObject {
	s.DisplayType = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetFavoriteFlag(v bool) *DescribeAllRootVariableResponseBodyResultObject {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetFieldRank(v int64) *DescribeAllRootVariableResponseBodyResultObject {
	s.FieldRank = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetFieldType(v string) *DescribeAllRootVariableResponseBodyResultObject {
	s.FieldType = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetId(v int64) *DescribeAllRootVariableResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetInputFieldType(v string) *DescribeAllRootVariableResponseBodyResultObject {
	s.InputFieldType = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetInputs(v string) *DescribeAllRootVariableResponseBodyResultObject {
	s.Inputs = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetName(v string) *DescribeAllRootVariableResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetOutputThreshold(v *DescribeAllRootVariableResponseBodyResultObjectOutputThreshold) *DescribeAllRootVariableResponseBodyResultObject {
	s.OutputThreshold = v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetSourceType(v string) *DescribeAllRootVariableResponseBodyResultObject {
	s.SourceType = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetTitle(v string) *DescribeAllRootVariableResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) SetType(v string) *DescribeAllRootVariableResponseBodyResultObject {
	s.Type = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObject) Validate() error {
	if s.OutputThreshold != nil {
		if err := s.OutputThreshold.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAllRootVariableResponseBodyResultObjectOutputThreshold struct {
	// Maximum value
	//
	// example:
	//
	// 1000
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	// Minimum value.
	//
	// example:
	//
	// 10
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeAllRootVariableResponseBodyResultObjectOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllRootVariableResponseBodyResultObjectOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeAllRootVariableResponseBodyResultObjectOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeAllRootVariableResponseBodyResultObjectOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeAllRootVariableResponseBodyResultObjectOutputThreshold) SetMaxValue(v float64) *DescribeAllRootVariableResponseBodyResultObjectOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObjectOutputThreshold) SetMinValue(v float64) *DescribeAllRootVariableResponseBodyResultObjectOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeAllRootVariableResponseBodyResultObjectOutputThreshold) Validate() error {
	return dara.Validate(s)
}
