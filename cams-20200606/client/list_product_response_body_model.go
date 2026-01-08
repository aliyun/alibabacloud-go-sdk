// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListProductResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListProductResponseBody
	GetCode() *string
	SetMessage(v string) *ListProductResponseBody
	GetMessage() *string
	SetModel(v *ListProductResponseBodyModel) *ListProductResponseBody
	GetModel() *ListProductResponseBodyModel
	SetRequestId(v string) *ListProductResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProductResponseBody
	GetSuccess() *bool
}

type ListProductResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	Model *ListProductResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProductResponseBody) GetModel() *ListProductResponseBodyModel {
	return s.Model
}

func (s *ListProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProductResponseBody) SetAccessDeniedDetail(v string) *ListProductResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListProductResponseBody) SetCode(v string) *ListProductResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductResponseBody) SetMessage(v string) *ListProductResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductResponseBody) SetModel(v *ListProductResponseBodyModel) *ListProductResponseBody {
	s.Model = v
	return s
}

func (s *ListProductResponseBody) SetRequestId(v string) *ListProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductResponseBody) SetSuccess(v bool) *ListProductResponseBody {
	s.Success = &v
	return s
}

func (s *ListProductResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductResponseBodyModel struct {
	// The returned data.
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination details.
	Paging *ListProductResponseBodyModelPaging `json:"Paging,omitempty" xml:"Paging,omitempty" type:"Struct"`
}

func (s ListProductResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s ListProductResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ListProductResponseBodyModel) GetData() []map[string]interface{} {
	return s.Data
}

func (s *ListProductResponseBodyModel) GetPaging() *ListProductResponseBodyModelPaging {
	return s.Paging
}

func (s *ListProductResponseBodyModel) SetData(v []map[string]interface{}) *ListProductResponseBodyModel {
	s.Data = v
	return s
}

func (s *ListProductResponseBodyModel) SetPaging(v *ListProductResponseBodyModelPaging) *ListProductResponseBodyModel {
	s.Paging = v
	return s
}

func (s *ListProductResponseBodyModel) Validate() error {
	if s.Paging != nil {
		if err := s.Paging.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductResponseBodyModelPaging struct {
	// The cursors for pagination.
	Cursors *ListProductResponseBodyModelPagingCursors `json:"Cursors,omitempty" xml:"Cursors,omitempty" type:"Struct"`
}

func (s ListProductResponseBodyModelPaging) String() string {
	return dara.Prettify(s)
}

func (s ListProductResponseBodyModelPaging) GoString() string {
	return s.String()
}

func (s *ListProductResponseBodyModelPaging) GetCursors() *ListProductResponseBodyModelPagingCursors {
	return s.Cursors
}

func (s *ListProductResponseBodyModelPaging) SetCursors(v *ListProductResponseBodyModelPagingCursors) *ListProductResponseBodyModelPaging {
	s.Cursors = v
	return s
}

func (s *ListProductResponseBodyModelPaging) Validate() error {
	if s.Cursors != nil {
		if err := s.Cursors.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductResponseBodyModelPagingCursors struct {
	// The cursor that points to the end of the page of the returned data.
	//
	// example:
	//
	// sjsuueu83838
	After *string `json:"After,omitempty" xml:"After,omitempty"`
	// The cursor that points to the beginning of the page of the returned data.
	//
	// example:
	//
	// sjjsjdjjdjd83883
	Before *string `json:"Before,omitempty" xml:"Before,omitempty"`
}

func (s ListProductResponseBodyModelPagingCursors) String() string {
	return dara.Prettify(s)
}

func (s ListProductResponseBodyModelPagingCursors) GoString() string {
	return s.String()
}

func (s *ListProductResponseBodyModelPagingCursors) GetAfter() *string {
	return s.After
}

func (s *ListProductResponseBodyModelPagingCursors) GetBefore() *string {
	return s.Before
}

func (s *ListProductResponseBodyModelPagingCursors) SetAfter(v string) *ListProductResponseBodyModelPagingCursors {
	s.After = &v
	return s
}

func (s *ListProductResponseBodyModelPagingCursors) SetBefore(v string) *ListProductResponseBodyModelPagingCursors {
	s.Before = &v
	return s
}

func (s *ListProductResponseBodyModelPagingCursors) Validate() error {
	return dara.Validate(s)
}
