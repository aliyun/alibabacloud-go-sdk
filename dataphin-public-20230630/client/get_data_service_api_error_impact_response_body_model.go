// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiErrorImpactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataServiceApiErrorImpactResponseBody
	GetCode() *string
	SetData(v *GetDataServiceApiErrorImpactResponseBodyData) *GetDataServiceApiErrorImpactResponseBody
	GetData() *GetDataServiceApiErrorImpactResponseBodyData
	SetHttpStatusCode(v int32) *GetDataServiceApiErrorImpactResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceApiErrorImpactResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataServiceApiErrorImpactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceApiErrorImpactResponseBody
	GetSuccess() *bool
}

type GetDataServiceApiErrorImpactResponseBody struct {
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDataServiceApiErrorImpactResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataServiceApiErrorImpactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiErrorImpactResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiErrorImpactResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceApiErrorImpactResponseBody) GetData() *GetDataServiceApiErrorImpactResponseBodyData {
	return s.Data
}

func (s *GetDataServiceApiErrorImpactResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceApiErrorImpactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceApiErrorImpactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceApiErrorImpactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceApiErrorImpactResponseBody) SetCode(v string) *GetDataServiceApiErrorImpactResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBody) SetData(v *GetDataServiceApiErrorImpactResponseBodyData) *GetDataServiceApiErrorImpactResponseBody {
	s.Data = v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBody) SetHttpStatusCode(v int32) *GetDataServiceApiErrorImpactResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBody) SetMessage(v string) *GetDataServiceApiErrorImpactResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBody) SetRequestId(v string) *GetDataServiceApiErrorImpactResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBody) SetSuccess(v bool) *GetDataServiceApiErrorImpactResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceApiErrorImpactResponseBodyData struct {
	ErrorApiList []*GetDataServiceApiErrorImpactResponseBodyDataErrorApiList `json:"ErrorApiList,omitempty" xml:"ErrorApiList,omitempty" type:"Repeated"`
	ErrorAppList []*GetDataServiceApiErrorImpactResponseBodyDataErrorAppList `json:"ErrorAppList,omitempty" xml:"ErrorAppList,omitempty" type:"Repeated"`
}

func (s GetDataServiceApiErrorImpactResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiErrorImpactResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiErrorImpactResponseBodyData) GetErrorApiList() []*GetDataServiceApiErrorImpactResponseBodyDataErrorApiList {
	return s.ErrorApiList
}

func (s *GetDataServiceApiErrorImpactResponseBodyData) GetErrorAppList() []*GetDataServiceApiErrorImpactResponseBodyDataErrorAppList {
	return s.ErrorAppList
}

func (s *GetDataServiceApiErrorImpactResponseBodyData) SetErrorApiList(v []*GetDataServiceApiErrorImpactResponseBodyDataErrorApiList) *GetDataServiceApiErrorImpactResponseBodyData {
	s.ErrorApiList = v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBodyData) SetErrorAppList(v []*GetDataServiceApiErrorImpactResponseBodyDataErrorAppList) *GetDataServiceApiErrorImpactResponseBodyData {
	s.ErrorAppList = v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceApiErrorImpactResponseBodyDataErrorApiList struct {
	// example:
	//
	// test
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// example:
	//
	// 2012
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 100
	ErrorCount *int64 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
}

func (s GetDataServiceApiErrorImpactResponseBodyDataErrorApiList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiErrorImpactResponseBodyDataErrorApiList) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorApiList) GetApiName() *string {
	return s.ApiName
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorApiList) GetAppId() *int64 {
	return s.AppId
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorApiList) GetErrorCount() *int64 {
	return s.ErrorCount
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorApiList) SetApiName(v string) *GetDataServiceApiErrorImpactResponseBodyDataErrorApiList {
	s.ApiName = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorApiList) SetAppId(v int64) *GetDataServiceApiErrorImpactResponseBodyDataErrorApiList {
	s.AppId = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorApiList) SetErrorCount(v int64) *GetDataServiceApiErrorImpactResponseBodyDataErrorApiList {
	s.ErrorCount = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorApiList) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceApiErrorImpactResponseBodyDataErrorAppList struct {
	// appId
	//
	// example:
	//
	// 1021
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// appKey
	//
	// example:
	//
	// 100211
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// app1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 10
	ErrorCount *int64 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
}

func (s GetDataServiceApiErrorImpactResponseBodyDataErrorAppList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiErrorImpactResponseBodyDataErrorAppList) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList) GetAppId() *int32 {
	return s.AppId
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList) GetAppName() *string {
	return s.AppName
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList) GetErrorCount() *int64 {
	return s.ErrorCount
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList) SetAppId(v int32) *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList {
	s.AppId = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList) SetAppKey(v int64) *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList {
	s.AppKey = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList) SetAppName(v string) *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList {
	s.AppName = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList) SetErrorCount(v int64) *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList {
	s.ErrorCount = &v
	return s
}

func (s *GetDataServiceApiErrorImpactResponseBodyDataErrorAppList) Validate() error {
	return dara.Validate(s)
}
