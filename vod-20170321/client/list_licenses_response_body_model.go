// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLicensesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLicensesResponseBody
	GetCode() *string
	SetData(v *ListLicensesResponseBodyData) *ListLicensesResponseBody
	GetData() *ListLicensesResponseBodyData
	SetHttpStatusCode(v int32) *ListLicensesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListLicensesResponseBody
	GetMessage() *string
	SetPageNo(v int64) *ListLicensesResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListLicensesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListLicensesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLicensesResponseBody
	GetSuccess() *bool
}

type ListLicensesResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListLicensesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNo         *int64                        `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int64                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLicensesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLicensesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLicensesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLicensesResponseBody) GetData() *ListLicensesResponseBodyData {
	return s.Data
}

func (s *ListLicensesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListLicensesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLicensesResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLicensesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLicensesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLicensesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLicensesResponseBody) SetCode(v string) *ListLicensesResponseBody {
	s.Code = &v
	return s
}

func (s *ListLicensesResponseBody) SetData(v *ListLicensesResponseBodyData) *ListLicensesResponseBody {
	s.Data = v
	return s
}

func (s *ListLicensesResponseBody) SetHttpStatusCode(v int32) *ListLicensesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLicensesResponseBody) SetMessage(v string) *ListLicensesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLicensesResponseBody) SetPageNo(v int64) *ListLicensesResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListLicensesResponseBody) SetPageSize(v int64) *ListLicensesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLicensesResponseBody) SetRequestId(v string) *ListLicensesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLicensesResponseBody) SetSuccess(v bool) *ListLicensesResponseBody {
	s.Success = &v
	return s
}

func (s *ListLicensesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLicensesResponseBodyData struct {
	List  []interface{} `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Total *int64        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListLicensesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLicensesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLicensesResponseBodyData) GetList() []interface{} {
	return s.List
}

func (s *ListLicensesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListLicensesResponseBodyData) SetList(v []interface{}) *ListLicensesResponseBodyData {
	s.List = v
	return s
}

func (s *ListLicensesResponseBodyData) SetTotal(v int64) *ListLicensesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListLicensesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
