// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstanceConfigResponseBodyData) *DescribeDBInstanceConfigResponseBody
	GetData() *DescribeDBInstanceConfigResponseBodyData
	SetRequestId(v string) *DescribeDBInstanceConfigResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceConfigResponseBody struct {
	Data *DescribeDBInstanceConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 66428721-xxxx-xxxx-xxxx-3BD1B67837E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigResponseBody) GetData() *DescribeDBInstanceConfigResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceConfigResponseBody) SetData(v *DescribeDBInstanceConfigResponseBodyData) *DescribeDBInstanceConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceConfigResponseBody) SetRequestId(v string) *DescribeDBInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceConfigResponseBodyData struct {
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string                                           `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Params       []*DescribeDBInstanceConfigResponseBodyDataParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceConfigResponseBodyData) GetParams() []*DescribeDBInstanceConfigResponseBodyDataParams {
	return s.Params
}

func (s *DescribeDBInstanceConfigResponseBodyData) SetDBInstanceId(v string) *DescribeDBInstanceConfigResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyData) SetParams(v []*DescribeDBInstanceConfigResponseBodyDataParams) *DescribeDBInstanceConfigResponseBodyData {
	s.Params = v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyData) Validate() error {
	if s.Params != nil {
		for _, item := range s.Params {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceConfigResponseBodyDataParams struct {
	// example:
	//
	// Maximum number of concurrently executed queries. Zero means unlimited.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// 1
	IsDynamic *int32 `json:"IsDynamic,omitempty" xml:"IsDynamic,omitempty"`
	// example:
	//
	// 1
	IsUserModifiable *int32 `json:"IsUserModifiable,omitempty" xml:"IsUserModifiable,omitempty"`
	// example:
	//
	// max_concurrent_queries
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// [0-100]
	Optional *string `json:"Optional,omitempty" xml:"Optional,omitempty"`
	// example:
	//
	// 0
	ParamRelyRule *string `json:"ParamRelyRule,omitempty" xml:"ParamRelyRule,omitempty"`
	// example:
	//
	// 100
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceConfigResponseBodyDataParams) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigResponseBodyDataParams) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) GetComment() *string {
	return s.Comment
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) GetIsDynamic() *int32 {
	return s.IsDynamic
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) GetIsUserModifiable() *int32 {
	return s.IsUserModifiable
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) GetName() *string {
	return s.Name
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) GetOptional() *string {
	return s.Optional
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) GetParamRelyRule() *string {
	return s.ParamRelyRule
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) SetComment(v string) *DescribeDBInstanceConfigResponseBodyDataParams {
	s.Comment = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) SetDefaultValue(v string) *DescribeDBInstanceConfigResponseBodyDataParams {
	s.DefaultValue = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) SetIsDynamic(v int32) *DescribeDBInstanceConfigResponseBodyDataParams {
	s.IsDynamic = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) SetIsUserModifiable(v int32) *DescribeDBInstanceConfigResponseBodyDataParams {
	s.IsUserModifiable = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) SetName(v string) *DescribeDBInstanceConfigResponseBodyDataParams {
	s.Name = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) SetOptional(v string) *DescribeDBInstanceConfigResponseBodyDataParams {
	s.Optional = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) SetParamRelyRule(v string) *DescribeDBInstanceConfigResponseBodyDataParams {
	s.ParamRelyRule = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) SetValue(v string) *DescribeDBInstanceConfigResponseBodyDataParams {
	s.Value = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyDataParams) Validate() error {
	return dara.Validate(s)
}
