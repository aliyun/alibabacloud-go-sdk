// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVmsVirtualNumberRelationByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryVmsVirtualNumberRelationByPageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryVmsVirtualNumberRelationByPageResponseBody
	GetCode() *string
	SetMessage(v string) *QueryVmsVirtualNumberRelationByPageResponseBody
	GetMessage() *string
	SetModel(v *QueryVmsVirtualNumberRelationByPageResponseBodyModel) *QueryVmsVirtualNumberRelationByPageResponseBody
	GetModel() *QueryVmsVirtualNumberRelationByPageResponseBodyModel
	SetRequestId(v string) *QueryVmsVirtualNumberRelationByPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryVmsVirtualNumberRelationByPageResponseBody
	GetSuccess() *bool
}

type QueryVmsVirtualNumberRelationByPageResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *QueryVmsVirtualNumberRelationByPageResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 01A58FA5-422C-58E0-AA71-B307A665416F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryVmsVirtualNumberRelationByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVmsVirtualNumberRelationByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) GetModel() *QueryVmsVirtualNumberRelationByPageResponseBodyModel {
	return s.Model
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) SetAccessDeniedDetail(v string) *QueryVmsVirtualNumberRelationByPageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) SetCode(v string) *QueryVmsVirtualNumberRelationByPageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) SetMessage(v string) *QueryVmsVirtualNumberRelationByPageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) SetModel(v *QueryVmsVirtualNumberRelationByPageResponseBodyModel) *QueryVmsVirtualNumberRelationByPageResponseBody {
	s.Model = v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) SetRequestId(v string) *QueryVmsVirtualNumberRelationByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) SetSuccess(v bool) *QueryVmsVirtualNumberRelationByPageResponseBody {
	s.Success = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryVmsVirtualNumberRelationByPageResponseBodyModel struct {
	Data []*QueryVmsVirtualNumberRelationByPageResponseBodyModelData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 65
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 231
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryVmsVirtualNumberRelationByPageResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s QueryVmsVirtualNumberRelationByPageResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModel) GetData() []*QueryVmsVirtualNumberRelationByPageResponseBodyModelData {
	return s.Data
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModel) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModel) GetTotal() *int64 {
	return s.Total
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModel) SetData(v []*QueryVmsVirtualNumberRelationByPageResponseBodyModelData) *QueryVmsVirtualNumberRelationByPageResponseBodyModel {
	s.Data = v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModel) SetPageNo(v int64) *QueryVmsVirtualNumberRelationByPageResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModel) SetPageSize(v int64) *QueryVmsVirtualNumberRelationByPageResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModel) SetTotal(v int64) *QueryVmsVirtualNumberRelationByPageResponseBodyModel {
	s.Total = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModel) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryVmsVirtualNumberRelationByPageResponseBodyModelData struct {
	// 号码归属地--城市
	//
	// example:
	//
	// 示例值示例值
	NumberCity *string `json:"NumberCity,omitempty" xml:"NumberCity,omitempty"`
	// 号码归属地--省
	//
	// example:
	//
	// 示例值示例值
	NumberProvince *string `json:"NumberProvince,omitempty" xml:"NumberProvince,omitempty"`
	// 真实号码
	//
	// example:
	//
	// 131****1234
	RealNumber *string `json:"RealNumber,omitempty" xml:"RealNumber,omitempty"`
	// 状态 1:有效；0:无效；-1:注销
	//
	// example:
	//
	// 1
	State *int64 `json:"State,omitempty" xml:"State,omitempty"`
	// 虚拟号码
	//
	// example:
	//
	// 0571***1234
	VirtualNumber *string `json:"VirtualNumber,omitempty" xml:"VirtualNumber,omitempty"`
}

func (s QueryVmsVirtualNumberRelationByPageResponseBodyModelData) String() string {
	return dara.Prettify(s)
}

func (s QueryVmsVirtualNumberRelationByPageResponseBodyModelData) GoString() string {
	return s.String()
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModelData) GetNumberCity() *string {
	return s.NumberCity
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModelData) GetNumberProvince() *string {
	return s.NumberProvince
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModelData) GetRealNumber() *string {
	return s.RealNumber
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModelData) GetState() *int64 {
	return s.State
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModelData) GetVirtualNumber() *string {
	return s.VirtualNumber
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModelData) SetNumberCity(v string) *QueryVmsVirtualNumberRelationByPageResponseBodyModelData {
	s.NumberCity = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModelData) SetNumberProvince(v string) *QueryVmsVirtualNumberRelationByPageResponseBodyModelData {
	s.NumberProvince = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModelData) SetRealNumber(v string) *QueryVmsVirtualNumberRelationByPageResponseBodyModelData {
	s.RealNumber = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModelData) SetState(v int64) *QueryVmsVirtualNumberRelationByPageResponseBodyModelData {
	s.State = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModelData) SetVirtualNumber(v string) *QueryVmsVirtualNumberRelationByPageResponseBodyModelData {
	s.VirtualNumber = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponseBodyModelData) Validate() error {
	return dara.Validate(s)
}
