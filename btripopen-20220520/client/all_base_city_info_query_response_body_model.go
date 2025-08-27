// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllBaseCityInfoQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AllBaseCityInfoQueryResponseBody
	GetCode() *string
	SetMessage(v string) *AllBaseCityInfoQueryResponseBody
	GetMessage() *string
	SetModule(v *AllBaseCityInfoQueryResponseBodyModule) *AllBaseCityInfoQueryResponseBody
	GetModule() *AllBaseCityInfoQueryResponseBodyModule
	SetRequestId(v string) *AllBaseCityInfoQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AllBaseCityInfoQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *AllBaseCityInfoQueryResponseBody
	GetTraceId() *string
}

type AllBaseCityInfoQueryResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// demo
	Message *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Module  *AllBaseCityInfoQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s AllBaseCityInfoQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllBaseCityInfoQueryResponseBody) GoString() string {
	return s.String()
}

func (s *AllBaseCityInfoQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *AllBaseCityInfoQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AllBaseCityInfoQueryResponseBody) GetModule() *AllBaseCityInfoQueryResponseBodyModule {
	return s.Module
}

func (s *AllBaseCityInfoQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllBaseCityInfoQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AllBaseCityInfoQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AllBaseCityInfoQueryResponseBody) SetCode(v string) *AllBaseCityInfoQueryResponseBody {
	s.Code = &v
	return s
}

func (s *AllBaseCityInfoQueryResponseBody) SetMessage(v string) *AllBaseCityInfoQueryResponseBody {
	s.Message = &v
	return s
}

func (s *AllBaseCityInfoQueryResponseBody) SetModule(v *AllBaseCityInfoQueryResponseBodyModule) *AllBaseCityInfoQueryResponseBody {
	s.Module = v
	return s
}

func (s *AllBaseCityInfoQueryResponseBody) SetRequestId(v string) *AllBaseCityInfoQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllBaseCityInfoQueryResponseBody) SetSuccess(v bool) *AllBaseCityInfoQueryResponseBody {
	s.Success = &v
	return s
}

func (s *AllBaseCityInfoQueryResponseBody) SetTraceId(v string) *AllBaseCityInfoQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *AllBaseCityInfoQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type AllBaseCityInfoQueryResponseBodyModule struct {
	AllCityBaseInfoList []*AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList `json:"all_city_base_info_list,omitempty" xml:"all_city_base_info_list,omitempty" type:"Repeated"`
}

func (s AllBaseCityInfoQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s AllBaseCityInfoQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *AllBaseCityInfoQueryResponseBodyModule) GetAllCityBaseInfoList() []*AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList {
	return s.AllCityBaseInfoList
}

func (s *AllBaseCityInfoQueryResponseBodyModule) SetAllCityBaseInfoList(v []*AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) *AllBaseCityInfoQueryResponseBodyModule {
	s.AllCityBaseInfoList = v
	return s
}

func (s *AllBaseCityInfoQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList struct {
	// example:
	//
	// 330122
	Adcode *string `json:"adcode,omitempty" xml:"adcode,omitempty"`
	// example:
	//
	// 0571
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 3
	CityLevel *string `json:"city_level,omitempty" xml:"city_level,omitempty"`
	// example:
	//
	// 桐庐
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 中国，浙江省，杭州市，桐庐
	CnNameTree *string `json:"cn_name_tree,omitempty" xml:"cn_name_tree,omitempty"`
	// example:
	//
	// 1
	Id            *int64    `json:"id,omitempty" xml:"id,omitempty"`
	OtherNameList []*string `json:"other_name_list,omitempty" xml:"other_name_list,omitempty" type:"Repeated"`
}

func (s AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) String() string {
	return dara.Prettify(s)
}

func (s AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) GoString() string {
	return s.String()
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) GetAdcode() *string {
	return s.Adcode
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) GetCityCode() *string {
	return s.CityCode
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) GetCityLevel() *string {
	return s.CityLevel
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) GetCityName() *string {
	return s.CityName
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) GetCnNameTree() *string {
	return s.CnNameTree
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) GetId() *int64 {
	return s.Id
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) GetOtherNameList() []*string {
	return s.OtherNameList
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) SetAdcode(v string) *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList {
	s.Adcode = &v
	return s
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) SetCityCode(v string) *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList {
	s.CityCode = &v
	return s
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) SetCityLevel(v string) *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList {
	s.CityLevel = &v
	return s
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) SetCityName(v string) *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList {
	s.CityName = &v
	return s
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) SetCnNameTree(v string) *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList {
	s.CnNameTree = &v
	return s
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) SetId(v int64) *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList {
	s.Id = &v
	return s
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) SetOtherNameList(v []*string) *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList {
	s.OtherNameList = v
	return s
}

func (s *AllBaseCityInfoQueryResponseBodyModuleAllCityBaseInfoList) Validate() error {
	return dara.Validate(s)
}
