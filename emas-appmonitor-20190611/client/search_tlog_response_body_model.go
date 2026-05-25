// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTlogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int64) *SearchTlogResponseBody
	GetErrorCode() *int64
	SetMessage(v string) *SearchTlogResponseBody
	GetMessage() *string
	SetModel(v *SearchTlogResponseBodyModel) *SearchTlogResponseBody
	GetModel() *SearchTlogResponseBodyModel
	SetRequestId(v string) *SearchTlogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchTlogResponseBody
	GetSuccess() *bool
}

type SearchTlogResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int64 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Successful
	Message *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *SearchTlogResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// A8313212-EB4E-4E15-A7F9-D9C8F3FE8E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchTlogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTlogResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTlogResponseBody) GetErrorCode() *int64 {
	return s.ErrorCode
}

func (s *SearchTlogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchTlogResponseBody) GetModel() *SearchTlogResponseBodyModel {
	return s.Model
}

func (s *SearchTlogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTlogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchTlogResponseBody) SetErrorCode(v int64) *SearchTlogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchTlogResponseBody) SetMessage(v string) *SearchTlogResponseBody {
	s.Message = &v
	return s
}

func (s *SearchTlogResponseBody) SetModel(v *SearchTlogResponseBodyModel) *SearchTlogResponseBody {
	s.Model = v
	return s
}

func (s *SearchTlogResponseBody) SetRequestId(v string) *SearchTlogResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTlogResponseBody) SetSuccess(v bool) *SearchTlogResponseBody {
	s.Success = &v
	return s
}

func (s *SearchTlogResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchTlogResponseBodyModel struct {
	Data []interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	Took *int32 `json:"Took,omitempty" xml:"Took,omitempty"`
	// example:
	//
	// 10
	Total *int64        `json:"Total,omitempty" xml:"Total,omitempty"`
	Trend []interface{} `json:"Trend,omitempty" xml:"Trend,omitempty" type:"Repeated"`
}

func (s SearchTlogResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s SearchTlogResponseBodyModel) GoString() string {
	return s.String()
}

func (s *SearchTlogResponseBodyModel) GetData() []interface{} {
	return s.Data
}

func (s *SearchTlogResponseBodyModel) GetPageNum() *int32 {
	return s.PageNum
}

func (s *SearchTlogResponseBodyModel) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTlogResponseBodyModel) GetTook() *int32 {
	return s.Took
}

func (s *SearchTlogResponseBodyModel) GetTotal() *int64 {
	return s.Total
}

func (s *SearchTlogResponseBodyModel) GetTrend() []interface{} {
	return s.Trend
}

func (s *SearchTlogResponseBodyModel) SetData(v []interface{}) *SearchTlogResponseBodyModel {
	s.Data = v
	return s
}

func (s *SearchTlogResponseBodyModel) SetPageNum(v int32) *SearchTlogResponseBodyModel {
	s.PageNum = &v
	return s
}

func (s *SearchTlogResponseBodyModel) SetPageSize(v int32) *SearchTlogResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *SearchTlogResponseBodyModel) SetTook(v int32) *SearchTlogResponseBodyModel {
	s.Took = &v
	return s
}

func (s *SearchTlogResponseBodyModel) SetTotal(v int64) *SearchTlogResponseBodyModel {
	s.Total = &v
	return s
}

func (s *SearchTlogResponseBodyModel) SetTrend(v []interface{}) *SearchTlogResponseBodyModel {
	s.Trend = v
	return s
}

func (s *SearchTlogResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
