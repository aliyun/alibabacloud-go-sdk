// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressionVariableFunctionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeExpressionVariableFunctionListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeExpressionVariableFunctionListResponseBodyResultObject) *DescribeExpressionVariableFunctionListResponseBody
	GetResultObject() []*DescribeExpressionVariableFunctionListResponseBodyResultObject
}

type DescribeExpressionVariableFunctionListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 6E8817D5-5354-5953-84B1-D98379F036DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeExpressionVariableFunctionListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeExpressionVariableFunctionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariableFunctionListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariableFunctionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressionVariableFunctionListResponseBody) GetResultObject() []*DescribeExpressionVariableFunctionListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeExpressionVariableFunctionListResponseBody) SetRequestId(v string) *DescribeExpressionVariableFunctionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponseBody) SetResultObject(v []*DescribeExpressionVariableFunctionListResponseBodyResultObject) *DescribeExpressionVariableFunctionListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponseBody) Validate() error {
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

type DescribeExpressionVariableFunctionListResponseBodyResultObject struct {
	// Description information.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Function name
	//
	// example:
	//
	// concat
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Maximum number of parameters
	//
	// example:
	//
	// 4
	MaxParamSize *int64 `json:"maxParamSize,omitempty" xml:"maxParamSize,omitempty"`
	// Minimum number of parameters
	//
	// example:
	//
	// 2
	MinParamSize *int64 `json:"minParamSize,omitempty" xml:"minParamSize,omitempty"`
	// Parameter types
	//
	// example:
	//
	// *STRING
	ParamTypes *string `json:"paramTypes,omitempty" xml:"paramTypes,omitempty"`
	// Whether it is directly invoked
	//
	// example:
	//
	// true
	Redirect *bool `json:"redirect,omitempty" xml:"redirect,omitempty"`
	// Method return types
	//
	// example:
	//
	// STRING
	ReturnTypes *string `json:"returnTypes,omitempty" xml:"returnTypes,omitempty"`
	// Function value
	//
	// example:
	//
	// concat
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeExpressionVariableFunctionListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariableFunctionListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) GetKey() *string {
	return s.Key
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) GetMaxParamSize() *int64 {
	return s.MaxParamSize
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) GetMinParamSize() *int64 {
	return s.MinParamSize
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) GetParamTypes() *string {
	return s.ParamTypes
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) GetRedirect() *bool {
	return s.Redirect
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) GetReturnTypes() *string {
	return s.ReturnTypes
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) GetValue() *string {
	return s.Value
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) SetDescription(v string) *DescribeExpressionVariableFunctionListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) SetKey(v string) *DescribeExpressionVariableFunctionListResponseBodyResultObject {
	s.Key = &v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) SetMaxParamSize(v int64) *DescribeExpressionVariableFunctionListResponseBodyResultObject {
	s.MaxParamSize = &v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) SetMinParamSize(v int64) *DescribeExpressionVariableFunctionListResponseBodyResultObject {
	s.MinParamSize = &v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) SetParamTypes(v string) *DescribeExpressionVariableFunctionListResponseBodyResultObject {
	s.ParamTypes = &v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) SetRedirect(v bool) *DescribeExpressionVariableFunctionListResponseBodyResultObject {
	s.Redirect = &v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) SetReturnTypes(v string) *DescribeExpressionVariableFunctionListResponseBodyResultObject {
	s.ReturnTypes = &v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) SetValue(v string) *DescribeExpressionVariableFunctionListResponseBodyResultObject {
	s.Value = &v
	return s
}

func (s *DescribeExpressionVariableFunctionListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
