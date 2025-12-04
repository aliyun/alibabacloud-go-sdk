// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGeneralConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGeneralConfigsResponseBody
	GetCode() *string
	SetData(v []*ListGeneralConfigsResponseBodyData) *ListGeneralConfigsResponseBody
	GetData() []*ListGeneralConfigsResponseBodyData
	SetHttpStatusCode(v int32) *ListGeneralConfigsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGeneralConfigsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListGeneralConfigsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGeneralConfigsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListGeneralConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGeneralConfigsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListGeneralConfigsResponseBody
	GetTotalCount() *int32
}

type ListGeneralConfigsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListGeneralConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGeneralConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGeneralConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGeneralConfigsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGeneralConfigsResponseBody) GetData() []*ListGeneralConfigsResponseBodyData {
	return s.Data
}

func (s *ListGeneralConfigsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGeneralConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGeneralConfigsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGeneralConfigsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGeneralConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGeneralConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGeneralConfigsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListGeneralConfigsResponseBody) SetCode(v string) *ListGeneralConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGeneralConfigsResponseBody) SetData(v []*ListGeneralConfigsResponseBodyData) *ListGeneralConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListGeneralConfigsResponseBody) SetHttpStatusCode(v int32) *ListGeneralConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGeneralConfigsResponseBody) SetMessage(v string) *ListGeneralConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGeneralConfigsResponseBody) SetPageNumber(v int32) *ListGeneralConfigsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListGeneralConfigsResponseBody) SetPageSize(v int32) *ListGeneralConfigsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListGeneralConfigsResponseBody) SetRequestId(v string) *ListGeneralConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGeneralConfigsResponseBody) SetSuccess(v bool) *ListGeneralConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGeneralConfigsResponseBody) SetTotalCount(v int32) *ListGeneralConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGeneralConfigsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGeneralConfigsResponseBodyData struct {
	// example:
	//
	// xxx
	ConfigDesc *string `json:"ConfigDesc,omitempty" xml:"ConfigDesc,omitempty"`
	// example:
	//
	// xxx
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// example:
	//
	// xxx
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// example:
	//
	// CustomSemanticSearch
	ConfigValueType *string `json:"ConfigValueType,omitempty" xml:"ConfigValueType,omitempty"`
}

func (s ListGeneralConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGeneralConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGeneralConfigsResponseBodyData) GetConfigDesc() *string {
	return s.ConfigDesc
}

func (s *ListGeneralConfigsResponseBodyData) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *ListGeneralConfigsResponseBodyData) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ListGeneralConfigsResponseBodyData) GetConfigValueType() *string {
	return s.ConfigValueType
}

func (s *ListGeneralConfigsResponseBodyData) SetConfigDesc(v string) *ListGeneralConfigsResponseBodyData {
	s.ConfigDesc = &v
	return s
}

func (s *ListGeneralConfigsResponseBodyData) SetConfigKey(v string) *ListGeneralConfigsResponseBodyData {
	s.ConfigKey = &v
	return s
}

func (s *ListGeneralConfigsResponseBodyData) SetConfigValue(v string) *ListGeneralConfigsResponseBodyData {
	s.ConfigValue = &v
	return s
}

func (s *ListGeneralConfigsResponseBodyData) SetConfigValueType(v string) *ListGeneralConfigsResponseBodyData {
	s.ConfigValueType = &v
	return s
}

func (s *ListGeneralConfigsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
