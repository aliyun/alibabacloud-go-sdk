// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSdkListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSdkListResponseBody
	GetCode() *string
	SetData(v *GetSdkListResponseBodyData) *GetSdkListResponseBody
	GetData() *GetSdkListResponseBodyData
	SetHttpStatusCode(v int32) *GetSdkListResponseBody
	GetHttpStatusCode() *int32
	SetLogExt(v interface{}) *GetSdkListResponseBody
	GetLogExt() interface{}
	SetMessage(v string) *GetSdkListResponseBody
	GetMessage() *string
	SetPageNo(v int64) *GetSdkListResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *GetSdkListResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *GetSdkListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSdkListResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *GetSdkListResponseBody
	GetTotal() *int64
}

type GetSdkListResponseBody struct {
	Code           *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetSdkListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	LogExt         interface{}                 `json:"LogExt,omitempty" xml:"LogExt,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNo         *int64                      `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int64                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Total          *int64                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSdkListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSdkListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSdkListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSdkListResponseBody) GetData() *GetSdkListResponseBodyData {
	return s.Data
}

func (s *GetSdkListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSdkListResponseBody) GetLogExt() interface{} {
	return s.LogExt
}

func (s *GetSdkListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSdkListResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetSdkListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetSdkListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSdkListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSdkListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetSdkListResponseBody) SetCode(v string) *GetSdkListResponseBody {
	s.Code = &v
	return s
}

func (s *GetSdkListResponseBody) SetData(v *GetSdkListResponseBodyData) *GetSdkListResponseBody {
	s.Data = v
	return s
}

func (s *GetSdkListResponseBody) SetHttpStatusCode(v int32) *GetSdkListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSdkListResponseBody) SetLogExt(v interface{}) *GetSdkListResponseBody {
	s.LogExt = v
	return s
}

func (s *GetSdkListResponseBody) SetMessage(v string) *GetSdkListResponseBody {
	s.Message = &v
	return s
}

func (s *GetSdkListResponseBody) SetPageNo(v int64) *GetSdkListResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetSdkListResponseBody) SetPageSize(v int64) *GetSdkListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSdkListResponseBody) SetRequestId(v string) *GetSdkListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSdkListResponseBody) SetSuccess(v bool) *GetSdkListResponseBody {
	s.Success = &v
	return s
}

func (s *GetSdkListResponseBody) SetTotal(v int64) *GetSdkListResponseBody {
	s.Total = &v
	return s
}

func (s *GetSdkListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSdkListResponseBodyData struct {
	Businesses map[string][]*DataBusinessesValue `json:"Businesses,omitempty" xml:"Businesses,omitempty"`
}

func (s GetSdkListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSdkListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSdkListResponseBodyData) GetBusinesses() map[string][]*DataBusinessesValue {
	return s.Businesses
}

func (s *GetSdkListResponseBodyData) SetBusinesses(v map[string][]*DataBusinessesValue) *GetSdkListResponseBodyData {
	s.Businesses = v
	return s
}

func (s *GetSdkListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
