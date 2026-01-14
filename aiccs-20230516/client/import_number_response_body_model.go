// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ImportNumberResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int64) *ImportNumberResponseBody
	GetCode() *int64
	SetMessage(v string) *ImportNumberResponseBody
	GetMessage() *string
	SetModel(v *ImportNumberResponseBodyModel) *ImportNumberResponseBody
	GetModel() *ImportNumberResponseBodyModel
	SetRequestId(v string) *ImportNumberResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ImportNumberResponseBody
	GetSuccess() *string
	SetTimestamp(v int64) *ImportNumberResponseBody
	GetTimestamp() *int64
}

type ImportNumberResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *ImportNumberResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
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

func (s ImportNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ImportNumberResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ImportNumberResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ImportNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportNumberResponseBody) GetModel() *ImportNumberResponseBodyModel {
	return s.Model
}

func (s *ImportNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportNumberResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ImportNumberResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ImportNumberResponseBody) SetAccessDeniedDetail(v string) *ImportNumberResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ImportNumberResponseBody) SetCode(v int64) *ImportNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ImportNumberResponseBody) SetMessage(v string) *ImportNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ImportNumberResponseBody) SetModel(v *ImportNumberResponseBodyModel) *ImportNumberResponseBody {
	s.Model = v
	return s
}

func (s *ImportNumberResponseBody) SetRequestId(v string) *ImportNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportNumberResponseBody) SetSuccess(v string) *ImportNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ImportNumberResponseBody) SetTimestamp(v int64) *ImportNumberResponseBody {
	s.Timestamp = &v
	return s
}

func (s *ImportNumberResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportNumberResponseBodyModel struct {
	// example:
	//
	// 54
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// example:
	//
	// 94
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 26
	ImportNum *int64 `json:"ImportNum,omitempty" xml:"ImportNum,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ImportNumberResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s ImportNumberResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ImportNumberResponseBodyModel) GetBatchId() *int64 {
	return s.BatchId
}

func (s *ImportNumberResponseBodyModel) GetCode() *int64 {
	return s.Code
}

func (s *ImportNumberResponseBodyModel) GetData() *string {
	return s.Data
}

func (s *ImportNumberResponseBodyModel) GetImportNum() *int64 {
	return s.ImportNum
}

func (s *ImportNumberResponseBodyModel) GetMessage() *string {
	return s.Message
}

func (s *ImportNumberResponseBodyModel) SetBatchId(v int64) *ImportNumberResponseBodyModel {
	s.BatchId = &v
	return s
}

func (s *ImportNumberResponseBodyModel) SetCode(v int64) *ImportNumberResponseBodyModel {
	s.Code = &v
	return s
}

func (s *ImportNumberResponseBodyModel) SetData(v string) *ImportNumberResponseBodyModel {
	s.Data = &v
	return s
}

func (s *ImportNumberResponseBodyModel) SetImportNum(v int64) *ImportNumberResponseBodyModel {
	s.ImportNum = &v
	return s
}

func (s *ImportNumberResponseBodyModel) SetMessage(v string) *ImportNumberResponseBodyModel {
	s.Message = &v
	return s
}

func (s *ImportNumberResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
