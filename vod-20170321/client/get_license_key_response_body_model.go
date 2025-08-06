// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicenseKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLicenseKeyResponseBody
	GetCode() *string
	SetData(v *GetLicenseKeyResponseBodyData) *GetLicenseKeyResponseBody
	GetData() *GetLicenseKeyResponseBodyData
	SetHttpStatusCode(v int32) *GetLicenseKeyResponseBody
	GetHttpStatusCode() *int32
	SetLogExt(v interface{}) *GetLicenseKeyResponseBody
	GetLogExt() interface{}
	SetMessage(v string) *GetLicenseKeyResponseBody
	GetMessage() *string
	SetPageNo(v int64) *GetLicenseKeyResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *GetLicenseKeyResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *GetLicenseKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLicenseKeyResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *GetLicenseKeyResponseBody
	GetTotal() *int64
}

type GetLicenseKeyResponseBody struct {
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetLicenseKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	LogExt         interface{}                    `json:"LogExt,omitempty" xml:"LogExt,omitempty"`
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNo         *int64                         `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int64                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Total          *int64                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetLicenseKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetLicenseKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLicenseKeyResponseBody) GetData() *GetLicenseKeyResponseBodyData {
	return s.Data
}

func (s *GetLicenseKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetLicenseKeyResponseBody) GetLogExt() interface{} {
	return s.LogExt
}

func (s *GetLicenseKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLicenseKeyResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetLicenseKeyResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetLicenseKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLicenseKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLicenseKeyResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetLicenseKeyResponseBody) SetCode(v string) *GetLicenseKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GetLicenseKeyResponseBody) SetData(v *GetLicenseKeyResponseBodyData) *GetLicenseKeyResponseBody {
	s.Data = v
	return s
}

func (s *GetLicenseKeyResponseBody) SetHttpStatusCode(v int32) *GetLicenseKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLicenseKeyResponseBody) SetLogExt(v interface{}) *GetLicenseKeyResponseBody {
	s.LogExt = v
	return s
}

func (s *GetLicenseKeyResponseBody) SetMessage(v string) *GetLicenseKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GetLicenseKeyResponseBody) SetPageNo(v int64) *GetLicenseKeyResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetLicenseKeyResponseBody) SetPageSize(v int64) *GetLicenseKeyResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetLicenseKeyResponseBody) SetRequestId(v string) *GetLicenseKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLicenseKeyResponseBody) SetSuccess(v bool) *GetLicenseKeyResponseBody {
	s.Success = &v
	return s
}

func (s *GetLicenseKeyResponseBody) SetTotal(v int64) *GetLicenseKeyResponseBody {
	s.Total = &v
	return s
}

func (s *GetLicenseKeyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLicenseKeyResponseBodyData struct {
	List  []interface{} `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Total *int64        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetLicenseKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLicenseKeyResponseBodyData) GetList() []interface{} {
	return s.List
}

func (s *GetLicenseKeyResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetLicenseKeyResponseBodyData) SetList(v []interface{}) *GetLicenseKeyResponseBodyData {
	s.List = v
	return s
}

func (s *GetLicenseKeyResponseBodyData) SetTotal(v int64) *GetLicenseKeyResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetLicenseKeyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
