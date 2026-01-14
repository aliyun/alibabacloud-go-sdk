// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *PageResponseBody
	GetCode() *int64
	SetMessage(v string) *PageResponseBody
	GetMessage() *string
	SetModel(v *PageResponseBodyModel) *PageResponseBody
	GetModel() *PageResponseBodyModel
	SetRequestId(v string) *PageResponseBody
	GetRequestId() *string
	SetSuccess(v string) *PageResponseBody
	GetSuccess() *string
	SetTimestamp(v int64) *PageResponseBody
	GetTimestamp() *int64
}

type PageResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *PageResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s PageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PageResponseBody) GoString() string {
	return s.String()
}

func (s *PageResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *PageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PageResponseBody) GetModel() *PageResponseBodyModel {
	return s.Model
}

func (s *PageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PageResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *PageResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *PageResponseBody) SetCode(v int64) *PageResponseBody {
	s.Code = &v
	return s
}

func (s *PageResponseBody) SetMessage(v string) *PageResponseBody {
	s.Message = &v
	return s
}

func (s *PageResponseBody) SetModel(v *PageResponseBodyModel) *PageResponseBody {
	s.Model = v
	return s
}

func (s *PageResponseBody) SetRequestId(v string) *PageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageResponseBody) SetSuccess(v string) *PageResponseBody {
	s.Success = &v
	return s
}

func (s *PageResponseBody) SetTimestamp(v int64) *PageResponseBody {
	s.Timestamp = &v
	return s
}

func (s *PageResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PageResponseBodyModel struct {
	List []*PageResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 97
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 5
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s PageResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s PageResponseBodyModel) GoString() string {
	return s.String()
}

func (s *PageResponseBodyModel) GetList() []*PageResponseBodyModelList {
	return s.List
}

func (s *PageResponseBodyModel) GetPageNo() *int64 {
	return s.PageNo
}

func (s *PageResponseBodyModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *PageResponseBodyModel) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *PageResponseBodyModel) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *PageResponseBodyModel) SetList(v []*PageResponseBodyModelList) *PageResponseBodyModel {
	s.List = v
	return s
}

func (s *PageResponseBodyModel) SetPageNo(v int64) *PageResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *PageResponseBodyModel) SetPageSize(v int64) *PageResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *PageResponseBodyModel) SetTotalCount(v int64) *PageResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *PageResponseBodyModel) SetTotalPage(v int64) *PageResponseBodyModel {
	s.TotalPage = &v
	return s
}

func (s *PageResponseBodyModel) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PageResponseBodyModelList struct {
	// 添加时间
	//
	// example:
	//
	// 2020-03-06 10:10:10
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 过期时间
	//
	// example:
	//
	// 1
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// 手机号码
	//
	// example:
	//
	// 13314206082
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 手机号MD5
	//
	// example:
	//
	// e10adc3949ba59abbe56e057f20f883e
	NumberMD5 *string `json:"NumberMD5,omitempty" xml:"NumberMD5,omitempty"`
	// 备注
	//
	// example:
	//
	// 示例值
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s PageResponseBodyModelList) String() string {
	return dara.Prettify(s)
}

func (s PageResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *PageResponseBodyModelList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *PageResponseBodyModelList) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *PageResponseBodyModelList) GetNumber() *string {
	return s.Number
}

func (s *PageResponseBodyModelList) GetNumberMD5() *string {
	return s.NumberMD5
}

func (s *PageResponseBodyModelList) GetRemark() *string {
	return s.Remark
}

func (s *PageResponseBodyModelList) SetCreateTime(v string) *PageResponseBodyModelList {
	s.CreateTime = &v
	return s
}

func (s *PageResponseBodyModelList) SetExpirationTime(v string) *PageResponseBodyModelList {
	s.ExpirationTime = &v
	return s
}

func (s *PageResponseBodyModelList) SetNumber(v string) *PageResponseBodyModelList {
	s.Number = &v
	return s
}

func (s *PageResponseBodyModelList) SetNumberMD5(v string) *PageResponseBodyModelList {
	s.NumberMD5 = &v
	return s
}

func (s *PageResponseBodyModelList) SetRemark(v string) *PageResponseBodyModelList {
	s.Remark = &v
	return s
}

func (s *PageResponseBodyModelList) Validate() error {
	return dara.Validate(s)
}
