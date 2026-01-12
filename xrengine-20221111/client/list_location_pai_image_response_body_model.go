// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLocationPaiImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLocationPaiImageResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListLocationPaiImageResponseBody
	GetCurrent() *int32
	SetData(v map[string]interface{}) *ListLocationPaiImageResponseBody
	GetData() map[string]interface{}
	SetErrorName(v string) *ListLocationPaiImageResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *ListLocationPaiImageResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *ListLocationPaiImageResponseBody
	GetMessage() *string
	SetPages(v int32) *ListLocationPaiImageResponseBody
	GetPages() *int32
	SetRequestId(v string) *ListLocationPaiImageResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListLocationPaiImageResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListLocationPaiImageResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListLocationPaiImageResponseBody
	GetTotal() *int32
}

type ListLocationPaiImageResponseBody struct {
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

func (s ListLocationPaiImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLocationPaiImageResponseBody) GoString() string {
	return s.String()
}

func (s *ListLocationPaiImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLocationPaiImageResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListLocationPaiImageResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *ListLocationPaiImageResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *ListLocationPaiImageResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListLocationPaiImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLocationPaiImageResponseBody) GetPages() *int32 {
	return s.Pages
}

func (s *ListLocationPaiImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLocationPaiImageResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListLocationPaiImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLocationPaiImageResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListLocationPaiImageResponseBody) SetCode(v string) *ListLocationPaiImageResponseBody {
	s.Code = &v
	return s
}

func (s *ListLocationPaiImageResponseBody) SetCurrent(v int32) *ListLocationPaiImageResponseBody {
	s.Current = &v
	return s
}

func (s *ListLocationPaiImageResponseBody) SetData(v map[string]interface{}) *ListLocationPaiImageResponseBody {
	s.Data = v
	return s
}

func (s *ListLocationPaiImageResponseBody) SetErrorName(v string) *ListLocationPaiImageResponseBody {
	s.ErrorName = &v
	return s
}

func (s *ListLocationPaiImageResponseBody) SetHttpCode(v int32) *ListLocationPaiImageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListLocationPaiImageResponseBody) SetMessage(v string) *ListLocationPaiImageResponseBody {
	s.Message = &v
	return s
}

func (s *ListLocationPaiImageResponseBody) SetPages(v int32) *ListLocationPaiImageResponseBody {
	s.Pages = &v
	return s
}

func (s *ListLocationPaiImageResponseBody) SetRequestId(v string) *ListLocationPaiImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLocationPaiImageResponseBody) SetSize(v int32) *ListLocationPaiImageResponseBody {
	s.Size = &v
	return s
}

func (s *ListLocationPaiImageResponseBody) SetSuccess(v bool) *ListLocationPaiImageResponseBody {
	s.Success = &v
	return s
}

func (s *ListLocationPaiImageResponseBody) SetTotal(v int32) *ListLocationPaiImageResponseBody {
	s.Total = &v
	return s
}

func (s *ListLocationPaiImageResponseBody) Validate() error {
	return dara.Validate(s)
}
