// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDatasetDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddDatasetDocumentResponseBody
	GetCode() *string
	SetData(v *AddDatasetDocumentResponseBodyData) *AddDatasetDocumentResponseBody
	GetData() *AddDatasetDocumentResponseBodyData
	SetHttpStatusCode(v int32) *AddDatasetDocumentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddDatasetDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDatasetDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDatasetDocumentResponseBody
	GetSuccess() *bool
}

type AddDatasetDocumentResponseBody struct {
	// example:
	//
	// NoData
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddDatasetDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDatasetDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *AddDatasetDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddDatasetDocumentResponseBody) GetData() *AddDatasetDocumentResponseBodyData {
	return s.Data
}

func (s *AddDatasetDocumentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddDatasetDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDatasetDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDatasetDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDatasetDocumentResponseBody) SetCode(v string) *AddDatasetDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *AddDatasetDocumentResponseBody) SetData(v *AddDatasetDocumentResponseBodyData) *AddDatasetDocumentResponseBody {
	s.Data = v
	return s
}

func (s *AddDatasetDocumentResponseBody) SetHttpStatusCode(v int32) *AddDatasetDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddDatasetDocumentResponseBody) SetMessage(v string) *AddDatasetDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *AddDatasetDocumentResponseBody) SetRequestId(v string) *AddDatasetDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDatasetDocumentResponseBody) SetSuccess(v bool) *AddDatasetDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *AddDatasetDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddDatasetDocumentResponseBodyData struct {
	// example:
	//
	// 文档业务唯一标识
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 8df2d69d63a247b6b52ff455b2d426b6
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// Success
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AddDatasetDocumentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddDatasetDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddDatasetDocumentResponseBodyData) GetDocId() *string {
	return s.DocId
}

func (s *AddDatasetDocumentResponseBodyData) GetDocUuid() *string {
	return s.DocUuid
}

func (s *AddDatasetDocumentResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddDatasetDocumentResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddDatasetDocumentResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *AddDatasetDocumentResponseBodyData) SetDocId(v string) *AddDatasetDocumentResponseBodyData {
	s.DocId = &v
	return s
}

func (s *AddDatasetDocumentResponseBodyData) SetDocUuid(v string) *AddDatasetDocumentResponseBodyData {
	s.DocUuid = &v
	return s
}

func (s *AddDatasetDocumentResponseBodyData) SetErrorCode(v string) *AddDatasetDocumentResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *AddDatasetDocumentResponseBodyData) SetErrorMessage(v string) *AddDatasetDocumentResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *AddDatasetDocumentResponseBodyData) SetStatus(v int32) *AddDatasetDocumentResponseBodyData {
	s.Status = &v
	return s
}

func (s *AddDatasetDocumentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
