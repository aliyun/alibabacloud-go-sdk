// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindSvcMapBindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FindSvcMapBindResponseBody
	GetCode() *string
	SetCurrent(v int32) *FindSvcMapBindResponseBody
	GetCurrent() *int32
	SetData(v map[string]interface{}) *FindSvcMapBindResponseBody
	GetData() map[string]interface{}
	SetErrorName(v string) *FindSvcMapBindResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *FindSvcMapBindResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *FindSvcMapBindResponseBody
	GetMessage() *string
	SetPages(v int32) *FindSvcMapBindResponseBody
	GetPages() *int32
	SetRequestId(v string) *FindSvcMapBindResponseBody
	GetRequestId() *string
	SetSize(v int32) *FindSvcMapBindResponseBody
	GetSize() *int32
	SetSuccess(v bool) *FindSvcMapBindResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *FindSvcMapBindResponseBody
	GetTotal() *int32
}

type FindSvcMapBindResponseBody struct {
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Current   *int32                 `json:"Current,omitempty" xml:"Current,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorName *string                `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Pages     *int32                 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32                 `json:"Size,omitempty" xml:"Size,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s FindSvcMapBindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindSvcMapBindResponseBody) GoString() string {
	return s.String()
}

func (s *FindSvcMapBindResponseBody) GetCode() *string {
	return s.Code
}

func (s *FindSvcMapBindResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *FindSvcMapBindResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *FindSvcMapBindResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *FindSvcMapBindResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *FindSvcMapBindResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FindSvcMapBindResponseBody) GetPages() *int32 {
	return s.Pages
}

func (s *FindSvcMapBindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindSvcMapBindResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *FindSvcMapBindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FindSvcMapBindResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *FindSvcMapBindResponseBody) SetCode(v string) *FindSvcMapBindResponseBody {
	s.Code = &v
	return s
}

func (s *FindSvcMapBindResponseBody) SetCurrent(v int32) *FindSvcMapBindResponseBody {
	s.Current = &v
	return s
}

func (s *FindSvcMapBindResponseBody) SetData(v map[string]interface{}) *FindSvcMapBindResponseBody {
	s.Data = v
	return s
}

func (s *FindSvcMapBindResponseBody) SetErrorName(v string) *FindSvcMapBindResponseBody {
	s.ErrorName = &v
	return s
}

func (s *FindSvcMapBindResponseBody) SetHttpCode(v int32) *FindSvcMapBindResponseBody {
	s.HttpCode = &v
	return s
}

func (s *FindSvcMapBindResponseBody) SetMessage(v string) *FindSvcMapBindResponseBody {
	s.Message = &v
	return s
}

func (s *FindSvcMapBindResponseBody) SetPages(v int32) *FindSvcMapBindResponseBody {
	s.Pages = &v
	return s
}

func (s *FindSvcMapBindResponseBody) SetRequestId(v string) *FindSvcMapBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindSvcMapBindResponseBody) SetSize(v int32) *FindSvcMapBindResponseBody {
	s.Size = &v
	return s
}

func (s *FindSvcMapBindResponseBody) SetSuccess(v bool) *FindSvcMapBindResponseBody {
	s.Success = &v
	return s
}

func (s *FindSvcMapBindResponseBody) SetTotal(v int32) *FindSvcMapBindResponseBody {
	s.Total = &v
	return s
}

func (s *FindSvcMapBindResponseBody) Validate() error {
	return dara.Validate(s)
}
