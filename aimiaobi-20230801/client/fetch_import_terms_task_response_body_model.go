// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchImportTermsTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FetchImportTermsTaskResponseBody
	GetCode() *string
	SetData(v *FetchImportTermsTaskResponseBodyData) *FetchImportTermsTaskResponseBody
	GetData() *FetchImportTermsTaskResponseBodyData
	SetHttpStatusCode(v int32) *FetchImportTermsTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *FetchImportTermsTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *FetchImportTermsTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FetchImportTermsTaskResponseBody
	GetSuccess() *bool
}

type FetchImportTermsTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *FetchImportTermsTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FetchImportTermsTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchImportTermsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *FetchImportTermsTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *FetchImportTermsTaskResponseBody) GetData() *FetchImportTermsTaskResponseBodyData {
	return s.Data
}

func (s *FetchImportTermsTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *FetchImportTermsTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FetchImportTermsTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FetchImportTermsTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FetchImportTermsTaskResponseBody) SetCode(v string) *FetchImportTermsTaskResponseBody {
	s.Code = &v
	return s
}

func (s *FetchImportTermsTaskResponseBody) SetData(v *FetchImportTermsTaskResponseBodyData) *FetchImportTermsTaskResponseBody {
	s.Data = v
	return s
}

func (s *FetchImportTermsTaskResponseBody) SetHttpStatusCode(v int32) *FetchImportTermsTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FetchImportTermsTaskResponseBody) SetMessage(v string) *FetchImportTermsTaskResponseBody {
	s.Message = &v
	return s
}

func (s *FetchImportTermsTaskResponseBody) SetRequestId(v string) *FetchImportTermsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchImportTermsTaskResponseBody) SetSuccess(v bool) *FetchImportTermsTaskResponseBody {
	s.Success = &v
	return s
}

func (s *FetchImportTermsTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type FetchImportTermsTaskResponseBodyData struct {
	// example:
	//
	// PENDING-待执行、RUNNING-执行中、SUCCESSED-成功、SUSPENDED-暂停、FAILED-失败、CANCELLED-取消
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s FetchImportTermsTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FetchImportTermsTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *FetchImportTermsTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *FetchImportTermsTaskResponseBodyData) SetStatus(v string) *FetchImportTermsTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *FetchImportTermsTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
