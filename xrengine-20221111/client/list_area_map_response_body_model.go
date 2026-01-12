// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAreaMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAreaMapResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListAreaMapResponseBody
	GetCurrent() *int32
	SetData(v map[string]interface{}) *ListAreaMapResponseBody
	GetData() map[string]interface{}
	SetErrorName(v string) *ListAreaMapResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *ListAreaMapResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *ListAreaMapResponseBody
	GetMessage() *string
	SetPages(v int32) *ListAreaMapResponseBody
	GetPages() *int32
	SetRequestId(v string) *ListAreaMapResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListAreaMapResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListAreaMapResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListAreaMapResponseBody
	GetTotal() *int32
}

type ListAreaMapResponseBody struct {
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

func (s ListAreaMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAreaMapResponseBody) GoString() string {
	return s.String()
}

func (s *ListAreaMapResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAreaMapResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListAreaMapResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *ListAreaMapResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *ListAreaMapResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListAreaMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAreaMapResponseBody) GetPages() *int32 {
	return s.Pages
}

func (s *ListAreaMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAreaMapResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListAreaMapResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAreaMapResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAreaMapResponseBody) SetCode(v string) *ListAreaMapResponseBody {
	s.Code = &v
	return s
}

func (s *ListAreaMapResponseBody) SetCurrent(v int32) *ListAreaMapResponseBody {
	s.Current = &v
	return s
}

func (s *ListAreaMapResponseBody) SetData(v map[string]interface{}) *ListAreaMapResponseBody {
	s.Data = v
	return s
}

func (s *ListAreaMapResponseBody) SetErrorName(v string) *ListAreaMapResponseBody {
	s.ErrorName = &v
	return s
}

func (s *ListAreaMapResponseBody) SetHttpCode(v int32) *ListAreaMapResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAreaMapResponseBody) SetMessage(v string) *ListAreaMapResponseBody {
	s.Message = &v
	return s
}

func (s *ListAreaMapResponseBody) SetPages(v int32) *ListAreaMapResponseBody {
	s.Pages = &v
	return s
}

func (s *ListAreaMapResponseBody) SetRequestId(v string) *ListAreaMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAreaMapResponseBody) SetSize(v int32) *ListAreaMapResponseBody {
	s.Size = &v
	return s
}

func (s *ListAreaMapResponseBody) SetSuccess(v bool) *ListAreaMapResponseBody {
	s.Success = &v
	return s
}

func (s *ListAreaMapResponseBody) SetTotal(v int32) *ListAreaMapResponseBody {
	s.Total = &v
	return s
}

func (s *ListAreaMapResponseBody) Validate() error {
	return dara.Validate(s)
}
