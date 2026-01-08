// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListProductCatalogResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListProductCatalogResponseBody
	GetCode() *string
	SetMessage(v string) *ListProductCatalogResponseBody
	GetMessage() *string
	SetModel(v *ListProductCatalogResponseBodyModel) *ListProductCatalogResponseBody
	GetModel() *ListProductCatalogResponseBodyModel
	SetRequestId(v string) *ListProductCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProductCatalogResponseBody
	GetSuccess() *bool
}

type ListProductCatalogResponseBody struct {
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
	Model *ListProductCatalogResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
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

func (s ListProductCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductCatalogResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListProductCatalogResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProductCatalogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProductCatalogResponseBody) GetModel() *ListProductCatalogResponseBodyModel {
	return s.Model
}

func (s *ListProductCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProductCatalogResponseBody) SetAccessDeniedDetail(v string) *ListProductCatalogResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListProductCatalogResponseBody) SetCode(v string) *ListProductCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductCatalogResponseBody) SetMessage(v string) *ListProductCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductCatalogResponseBody) SetModel(v *ListProductCatalogResponseBodyModel) *ListProductCatalogResponseBody {
	s.Model = v
	return s
}

func (s *ListProductCatalogResponseBody) SetRequestId(v string) *ListProductCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductCatalogResponseBody) SetSuccess(v bool) *ListProductCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *ListProductCatalogResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductCatalogResponseBodyModel struct {
	// The returned data.
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination details.
	Paging *ListProductCatalogResponseBodyModelPaging `json:"Paging,omitempty" xml:"Paging,omitempty" type:"Struct"`
}

func (s ListProductCatalogResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s ListProductCatalogResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ListProductCatalogResponseBodyModel) GetData() []map[string]interface{} {
	return s.Data
}

func (s *ListProductCatalogResponseBodyModel) GetPaging() *ListProductCatalogResponseBodyModelPaging {
	return s.Paging
}

func (s *ListProductCatalogResponseBodyModel) SetData(v []map[string]interface{}) *ListProductCatalogResponseBodyModel {
	s.Data = v
	return s
}

func (s *ListProductCatalogResponseBodyModel) SetPaging(v *ListProductCatalogResponseBodyModelPaging) *ListProductCatalogResponseBodyModel {
	s.Paging = v
	return s
}

func (s *ListProductCatalogResponseBodyModel) Validate() error {
	if s.Paging != nil {
		if err := s.Paging.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductCatalogResponseBodyModelPaging struct {
	// The cursors for pagination.
	Cursors *ListProductCatalogResponseBodyModelPagingCursors `json:"Cursors,omitempty" xml:"Cursors,omitempty" type:"Struct"`
}

func (s ListProductCatalogResponseBodyModelPaging) String() string {
	return dara.Prettify(s)
}

func (s ListProductCatalogResponseBodyModelPaging) GoString() string {
	return s.String()
}

func (s *ListProductCatalogResponseBodyModelPaging) GetCursors() *ListProductCatalogResponseBodyModelPagingCursors {
	return s.Cursors
}

func (s *ListProductCatalogResponseBodyModelPaging) SetCursors(v *ListProductCatalogResponseBodyModelPagingCursors) *ListProductCatalogResponseBodyModelPaging {
	s.Cursors = v
	return s
}

func (s *ListProductCatalogResponseBodyModelPaging) Validate() error {
	if s.Cursors != nil {
		if err := s.Cursors.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductCatalogResponseBodyModelPagingCursors struct {
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

func (s ListProductCatalogResponseBodyModelPagingCursors) String() string {
	return dara.Prettify(s)
}

func (s ListProductCatalogResponseBodyModelPagingCursors) GoString() string {
	return s.String()
}

func (s *ListProductCatalogResponseBodyModelPagingCursors) GetAfter() *string {
	return s.After
}

func (s *ListProductCatalogResponseBodyModelPagingCursors) GetBefore() *string {
	return s.Before
}

func (s *ListProductCatalogResponseBodyModelPagingCursors) SetAfter(v string) *ListProductCatalogResponseBodyModelPagingCursors {
	s.After = &v
	return s
}

func (s *ListProductCatalogResponseBodyModelPagingCursors) SetBefore(v string) *ListProductCatalogResponseBodyModelPagingCursors {
	s.Before = &v
	return s
}

func (s *ListProductCatalogResponseBodyModelPagingCursors) Validate() error {
	return dara.Validate(s)
}
