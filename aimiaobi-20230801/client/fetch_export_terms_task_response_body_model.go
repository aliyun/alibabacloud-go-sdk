// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchExportTermsTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FetchExportTermsTaskResponseBody
	GetCode() *string
	SetData(v *FetchExportTermsTaskResponseBodyData) *FetchExportTermsTaskResponseBody
	GetData() *FetchExportTermsTaskResponseBodyData
	SetHttpStatusCode(v int32) *FetchExportTermsTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *FetchExportTermsTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *FetchExportTermsTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FetchExportTermsTaskResponseBody
	GetSuccess() *bool
}

type FetchExportTermsTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *FetchExportTermsTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s FetchExportTermsTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchExportTermsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *FetchExportTermsTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *FetchExportTermsTaskResponseBody) GetData() *FetchExportTermsTaskResponseBodyData {
	return s.Data
}

func (s *FetchExportTermsTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *FetchExportTermsTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FetchExportTermsTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FetchExportTermsTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FetchExportTermsTaskResponseBody) SetCode(v string) *FetchExportTermsTaskResponseBody {
	s.Code = &v
	return s
}

func (s *FetchExportTermsTaskResponseBody) SetData(v *FetchExportTermsTaskResponseBodyData) *FetchExportTermsTaskResponseBody {
	s.Data = v
	return s
}

func (s *FetchExportTermsTaskResponseBody) SetHttpStatusCode(v int32) *FetchExportTermsTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FetchExportTermsTaskResponseBody) SetMessage(v string) *FetchExportTermsTaskResponseBody {
	s.Message = &v
	return s
}

func (s *FetchExportTermsTaskResponseBody) SetRequestId(v string) *FetchExportTermsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchExportTermsTaskResponseBody) SetSuccess(v bool) *FetchExportTermsTaskResponseBody {
	s.Success = &v
	return s
}

func (s *FetchExportTermsTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FetchExportTermsTaskResponseBodyData struct {
	// example:
	//
	// https//xxxx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// PENDING-待执行、RUNNING-执行中、SUCCESSED-成功、SUSPENDED-暂停、FAILED-失败、CANCELLED-取消
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s FetchExportTermsTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FetchExportTermsTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *FetchExportTermsTaskResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *FetchExportTermsTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *FetchExportTermsTaskResponseBodyData) SetFileUrl(v string) *FetchExportTermsTaskResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *FetchExportTermsTaskResponseBodyData) SetStatus(v string) *FetchExportTermsTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *FetchExportTermsTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
