// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicensesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLicensesResponseBody
	GetCode() *string
	SetData(v *GetLicensesResponseBodyData) *GetLicensesResponseBody
	GetData() *GetLicensesResponseBodyData
	SetHttpStatusCode(v int32) *GetLicensesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetLicensesResponseBody
	GetMessage() *string
	SetPageNo(v int64) *GetLicensesResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *GetLicensesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *GetLicensesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLicensesResponseBody
	GetSuccess() *bool
}

type GetLicensesResponseBody struct {
	Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetLicensesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNo         *int64                       `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int64                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLicensesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLicensesResponseBody) GoString() string {
	return s.String()
}

func (s *GetLicensesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLicensesResponseBody) GetData() *GetLicensesResponseBodyData {
	return s.Data
}

func (s *GetLicensesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetLicensesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLicensesResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetLicensesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetLicensesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLicensesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLicensesResponseBody) SetCode(v string) *GetLicensesResponseBody {
	s.Code = &v
	return s
}

func (s *GetLicensesResponseBody) SetData(v *GetLicensesResponseBodyData) *GetLicensesResponseBody {
	s.Data = v
	return s
}

func (s *GetLicensesResponseBody) SetHttpStatusCode(v int32) *GetLicensesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLicensesResponseBody) SetMessage(v string) *GetLicensesResponseBody {
	s.Message = &v
	return s
}

func (s *GetLicensesResponseBody) SetPageNo(v int64) *GetLicensesResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetLicensesResponseBody) SetPageSize(v int64) *GetLicensesResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetLicensesResponseBody) SetRequestId(v string) *GetLicensesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLicensesResponseBody) SetSuccess(v bool) *GetLicensesResponseBody {
	s.Success = &v
	return s
}

func (s *GetLicensesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLicensesResponseBodyData struct {
	List  []interface{} `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Total *int64        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetLicensesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLicensesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLicensesResponseBodyData) GetList() []interface{} {
	return s.List
}

func (s *GetLicensesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetLicensesResponseBodyData) SetList(v []interface{}) *GetLicensesResponseBodyData {
	s.List = v
	return s
}

func (s *GetLicensesResponseBodyData) SetTotal(v int64) *GetLicensesResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetLicensesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
