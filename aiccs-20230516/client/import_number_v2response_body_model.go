// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNumberV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ImportNumberV2ResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int64) *ImportNumberV2ResponseBody
	GetCode() *int64
	SetMessage(v string) *ImportNumberV2ResponseBody
	GetMessage() *string
	SetModel(v *ImportNumberV2ResponseBodyModel) *ImportNumberV2ResponseBody
	GetModel() *ImportNumberV2ResponseBodyModel
	SetRequestId(v string) *ImportNumberV2ResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ImportNumberV2ResponseBody
	GetSuccess() *string
	SetTimestamp(v int64) *ImportNumberV2ResponseBody
	GetTimestamp() *int64
}

type ImportNumberV2ResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *ImportNumberV2ResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
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
	// 98
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ImportNumberV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportNumberV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ImportNumberV2ResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ImportNumberV2ResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ImportNumberV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportNumberV2ResponseBody) GetModel() *ImportNumberV2ResponseBodyModel {
	return s.Model
}

func (s *ImportNumberV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportNumberV2ResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ImportNumberV2ResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ImportNumberV2ResponseBody) SetAccessDeniedDetail(v string) *ImportNumberV2ResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ImportNumberV2ResponseBody) SetCode(v int64) *ImportNumberV2ResponseBody {
	s.Code = &v
	return s
}

func (s *ImportNumberV2ResponseBody) SetMessage(v string) *ImportNumberV2ResponseBody {
	s.Message = &v
	return s
}

func (s *ImportNumberV2ResponseBody) SetModel(v *ImportNumberV2ResponseBodyModel) *ImportNumberV2ResponseBody {
	s.Model = v
	return s
}

func (s *ImportNumberV2ResponseBody) SetRequestId(v string) *ImportNumberV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportNumberV2ResponseBody) SetSuccess(v string) *ImportNumberV2ResponseBody {
	s.Success = &v
	return s
}

func (s *ImportNumberV2ResponseBody) SetTimestamp(v int64) *ImportNumberV2ResponseBody {
	s.Timestamp = &v
	return s
}

func (s *ImportNumberV2ResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportNumberV2ResponseBodyModel struct {
	// example:
	//
	// 74
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 58
	ImportNum *int64 `json:"ImportNum,omitempty" xml:"ImportNum,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ImportNumberV2ResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s ImportNumberV2ResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ImportNumberV2ResponseBodyModel) GetBatchId() *int64 {
	return s.BatchId
}

func (s *ImportNumberV2ResponseBodyModel) GetCode() *int64 {
	return s.Code
}

func (s *ImportNumberV2ResponseBodyModel) GetData() *string {
	return s.Data
}

func (s *ImportNumberV2ResponseBodyModel) GetImportNum() *int64 {
	return s.ImportNum
}

func (s *ImportNumberV2ResponseBodyModel) GetMessage() *string {
	return s.Message
}

func (s *ImportNumberV2ResponseBodyModel) SetBatchId(v int64) *ImportNumberV2ResponseBodyModel {
	s.BatchId = &v
	return s
}

func (s *ImportNumberV2ResponseBodyModel) SetCode(v int64) *ImportNumberV2ResponseBodyModel {
	s.Code = &v
	return s
}

func (s *ImportNumberV2ResponseBodyModel) SetData(v string) *ImportNumberV2ResponseBodyModel {
	s.Data = &v
	return s
}

func (s *ImportNumberV2ResponseBodyModel) SetImportNum(v int64) *ImportNumberV2ResponseBodyModel {
	s.ImportNum = &v
	return s
}

func (s *ImportNumberV2ResponseBodyModel) SetMessage(v string) *ImportNumberV2ResponseBodyModel {
	s.Message = &v
	return s
}

func (s *ImportNumberV2ResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
