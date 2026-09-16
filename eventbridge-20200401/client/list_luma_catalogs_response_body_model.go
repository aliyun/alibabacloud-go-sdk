// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaCatalogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLumaCatalogsResponseBody
	GetCode() *string
	SetData(v *ListLumaCatalogsResponseBodyData) *ListLumaCatalogsResponseBody
	GetData() *ListLumaCatalogsResponseBodyData
	SetMessage(v string) *ListLumaCatalogsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLumaCatalogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLumaCatalogsResponseBody
	GetSuccess() *bool
}

type ListLumaCatalogsResponseBody struct {
	// The response code. A value of Success indicates a successful call. A specific error code is returned when the call fails.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of data catalogs bound to the Agent. All results are returned at once without pagination.
	Data *ListLumaCatalogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation. The value is Operation success when the call succeeds, or a specific error description when the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, used for troubleshooting and ticket submission.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLumaCatalogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLumaCatalogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLumaCatalogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLumaCatalogsResponseBody) GetData() *ListLumaCatalogsResponseBodyData {
	return s.Data
}

func (s *ListLumaCatalogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLumaCatalogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLumaCatalogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLumaCatalogsResponseBody) SetCode(v string) *ListLumaCatalogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListLumaCatalogsResponseBody) SetData(v *ListLumaCatalogsResponseBodyData) *ListLumaCatalogsResponseBody {
	s.Data = v
	return s
}

func (s *ListLumaCatalogsResponseBody) SetMessage(v string) *ListLumaCatalogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListLumaCatalogsResponseBody) SetRequestId(v string) *ListLumaCatalogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLumaCatalogsResponseBody) SetSuccess(v bool) *ListLumaCatalogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListLumaCatalogsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLumaCatalogsResponseBodyData struct {
	// The list of data catalogs bound to the Agent.
	//
	// example:
	//
	// [{"Name":"my_catalog"}]
	Catalogs []*Catalog `json:"Catalogs,omitempty" xml:"Catalogs,omitempty" type:"Repeated"`
	// 本次请求实际生效的每页数量。未传 Limit 时为服务端默认值，超出上限时为收敛后的值
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 下一页起始Token，传入下次请求的 NextToken 可获取下一页；为空表示已无更多数据
	//
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Agent 绑定的数据目录总数，与本页返回条数无关
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLumaCatalogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLumaCatalogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLumaCatalogsResponseBodyData) GetCatalogs() []*Catalog {
	return s.Catalogs
}

func (s *ListLumaCatalogsResponseBodyData) GetLimit() *int32 {
	return s.Limit
}

func (s *ListLumaCatalogsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaCatalogsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLumaCatalogsResponseBodyData) SetCatalogs(v []*Catalog) *ListLumaCatalogsResponseBodyData {
	s.Catalogs = v
	return s
}

func (s *ListLumaCatalogsResponseBodyData) SetLimit(v int32) *ListLumaCatalogsResponseBodyData {
	s.Limit = &v
	return s
}

func (s *ListLumaCatalogsResponseBodyData) SetNextToken(v string) *ListLumaCatalogsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListLumaCatalogsResponseBodyData) SetTotalCount(v int32) *ListLumaCatalogsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLumaCatalogsResponseBodyData) Validate() error {
	if s.Catalogs != nil {
		for _, item := range s.Catalogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
