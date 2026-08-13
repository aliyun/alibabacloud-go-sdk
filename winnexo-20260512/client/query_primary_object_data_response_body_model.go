// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPrimaryObjectDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPrimaryObjectDataResponseBody
	GetCode() *string
	SetItems(v []map[string]*string) *QueryPrimaryObjectDataResponseBody
	GetItems() []map[string]*string
	SetMessage(v string) *QueryPrimaryObjectDataResponseBody
	GetMessage() *string
	SetPage(v int64) *QueryPrimaryObjectDataResponseBody
	GetPage() *int64
	SetPageSize(v int64) *QueryPrimaryObjectDataResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *QueryPrimaryObjectDataResponseBody
	GetRequestId() *string
	SetTotal(v int64) *QueryPrimaryObjectDataResponseBody
	GetTotal() *int64
}

type QueryPrimaryObjectDataResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code  *string              `json:"code,omitempty" xml:"code,omitempty"`
	Items []map[string]*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 当前页码
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数量
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 总数
	//
	// example:
	//
	// 0
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s QueryPrimaryObjectDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPrimaryObjectDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPrimaryObjectDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPrimaryObjectDataResponseBody) GetItems() []map[string]*string {
	return s.Items
}

func (s *QueryPrimaryObjectDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPrimaryObjectDataResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *QueryPrimaryObjectDataResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryPrimaryObjectDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPrimaryObjectDataResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *QueryPrimaryObjectDataResponseBody) SetCode(v string) *QueryPrimaryObjectDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPrimaryObjectDataResponseBody) SetItems(v []map[string]*string) *QueryPrimaryObjectDataResponseBody {
	s.Items = v
	return s
}

func (s *QueryPrimaryObjectDataResponseBody) SetMessage(v string) *QueryPrimaryObjectDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPrimaryObjectDataResponseBody) SetPage(v int64) *QueryPrimaryObjectDataResponseBody {
	s.Page = &v
	return s
}

func (s *QueryPrimaryObjectDataResponseBody) SetPageSize(v int64) *QueryPrimaryObjectDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryPrimaryObjectDataResponseBody) SetRequestId(v string) *QueryPrimaryObjectDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPrimaryObjectDataResponseBody) SetTotal(v int64) *QueryPrimaryObjectDataResponseBody {
	s.Total = &v
	return s
}

func (s *QueryPrimaryObjectDataResponseBody) Validate() error {
	return dara.Validate(s)
}
