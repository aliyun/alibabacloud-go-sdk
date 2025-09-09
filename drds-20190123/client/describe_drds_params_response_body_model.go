// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*DescribeDrdsParamsResponseBodyList) *DescribeDrdsParamsResponseBody
	GetList() []*DescribeDrdsParamsResponseBodyList
	SetRequestId(v string) *DescribeDrdsParamsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsParamsResponseBody
	GetSuccess() *bool
}

type DescribeDrdsParamsResponseBody struct {
	// Indicates information about parameters.
	List []*DescribeDrdsParamsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Indicates the ID of the request.
	//
	// example:
	//
	// 2F7F8080-9132-4279-85D0-B7E5C4305162
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsParamsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsResponseBody) GetList() []*DescribeDrdsParamsResponseBodyList {
	return s.List
}

func (s *DescribeDrdsParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsParamsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsParamsResponseBody) SetList(v []*DescribeDrdsParamsResponseBodyList) *DescribeDrdsParamsResponseBody {
	s.List = v
	return s
}

func (s *DescribeDrdsParamsResponseBody) SetRequestId(v string) *DescribeDrdsParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsParamsResponseBody) SetSuccess(v bool) *DescribeDrdsParamsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsParamsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsParamsResponseBodyList struct {
	// Indicates the name of the database.
	//
	// example:
	//
	// drds_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates whether a restart is required.
	//
	// example:
	//
	// true
	NeedRestart *bool `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	// Indicates the default value of a parameter.
	//
	// example:
	//
	// 1000
	ParamDefaultValue *string `json:"ParamDefaultValue,omitempty" xml:"ParamDefaultValue,omitempty"`
	// Indicates the description of the parameter.
	ParamDesc *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	// Indicates the name of the parameter.
	//
	// example:
	//
	// SLOW_SQL_TIME
	ParamEnglishName *string `json:"ParamEnglishName,omitempty" xml:"ParamEnglishName,omitempty"`
	// Indicates the parameter level.
	//
	// example:
	//
	// INSTANCE
	ParamLevel *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	// Indicates the name of the parameter.
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// Indicates the value range of the parameter.
	//
	// example:
	//
	// [1000-900000]
	ParamRanges *string `json:"ParamRanges,omitempty" xml:"ParamRanges,omitempty"`
	// Indicates the type of the parameter.
	//
	// example:
	//
	// CONFIG
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Indicates the value of the parameter.
	//
	// example:
	//
	// 1000
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
	// Indicates the name of the variable.
	//
	// example:
	//
	// slowSqlTime
	ParamVariableName *string `json:"ParamVariableName,omitempty" xml:"ParamVariableName,omitempty"`
}

func (s DescribeDrdsParamsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsParamsResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeDrdsParamsResponseBodyList) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsParamsResponseBodyList) GetNeedRestart() *bool {
	return s.NeedRestart
}

func (s *DescribeDrdsParamsResponseBodyList) GetParamDefaultValue() *string {
	return s.ParamDefaultValue
}

func (s *DescribeDrdsParamsResponseBodyList) GetParamDesc() *string {
	return s.ParamDesc
}

func (s *DescribeDrdsParamsResponseBodyList) GetParamEnglishName() *string {
	return s.ParamEnglishName
}

func (s *DescribeDrdsParamsResponseBodyList) GetParamLevel() *string {
	return s.ParamLevel
}

func (s *DescribeDrdsParamsResponseBodyList) GetParamName() *string {
	return s.ParamName
}

func (s *DescribeDrdsParamsResponseBodyList) GetParamRanges() *string {
	return s.ParamRanges
}

func (s *DescribeDrdsParamsResponseBodyList) GetParamType() *string {
	return s.ParamType
}

func (s *DescribeDrdsParamsResponseBodyList) GetParamValue() *string {
	return s.ParamValue
}

func (s *DescribeDrdsParamsResponseBodyList) GetParamVariableName() *string {
	return s.ParamVariableName
}

func (s *DescribeDrdsParamsResponseBodyList) SetDbName(v string) *DescribeDrdsParamsResponseBodyList {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetNeedRestart(v bool) *DescribeDrdsParamsResponseBodyList {
	s.NeedRestart = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamDefaultValue(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamDefaultValue = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamDesc(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamDesc = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamEnglishName(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamEnglishName = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamLevel(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamLevel = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamName(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamName = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamRanges(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamRanges = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamType(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamType = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamValue(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamValue = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) SetParamVariableName(v string) *DescribeDrdsParamsResponseBodyList {
	s.ParamVariableName = &v
	return s
}

func (s *DescribeDrdsParamsResponseBodyList) Validate() error {
	return dara.Validate(s)
}
