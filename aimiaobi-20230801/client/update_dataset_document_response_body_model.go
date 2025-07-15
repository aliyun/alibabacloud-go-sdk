// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDatasetDocumentResponseBody
	GetCode() *string
	SetData(v *UpdateDatasetDocumentResponseBodyData) *UpdateDatasetDocumentResponseBody
	GetData() *UpdateDatasetDocumentResponseBodyData
	SetHttpStatusCode(v int32) *UpdateDatasetDocumentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDatasetDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDatasetDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDatasetDocumentResponseBody
	GetSuccess() *bool
}

type UpdateDatasetDocumentResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateDatasetDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UpdateDatasetDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDatasetDocumentResponseBody) GetData() *UpdateDatasetDocumentResponseBodyData {
	return s.Data
}

func (s *UpdateDatasetDocumentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDatasetDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDatasetDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDatasetDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDatasetDocumentResponseBody) SetCode(v string) *UpdateDatasetDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDatasetDocumentResponseBody) SetData(v *UpdateDatasetDocumentResponseBodyData) *UpdateDatasetDocumentResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDatasetDocumentResponseBody) SetHttpStatusCode(v int32) *UpdateDatasetDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDatasetDocumentResponseBody) SetMessage(v string) *UpdateDatasetDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDatasetDocumentResponseBody) SetRequestId(v string) *UpdateDatasetDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasetDocumentResponseBody) SetSuccess(v bool) *UpdateDatasetDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDatasetDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateDatasetDocumentResponseBodyData struct {
	// example:
	//
	// 用户指定的文档唯一ID
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 内部文档唯一ID
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateDatasetDocumentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDatasetDocumentResponseBodyData) GetDocId() *string {
	return s.DocId
}

func (s *UpdateDatasetDocumentResponseBodyData) GetDocUuid() *string {
	return s.DocUuid
}

func (s *UpdateDatasetDocumentResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *UpdateDatasetDocumentResponseBodyData) SetDocId(v string) *UpdateDatasetDocumentResponseBodyData {
	s.DocId = &v
	return s
}

func (s *UpdateDatasetDocumentResponseBodyData) SetDocUuid(v string) *UpdateDatasetDocumentResponseBodyData {
	s.DocUuid = &v
	return s
}

func (s *UpdateDatasetDocumentResponseBodyData) SetTitle(v string) *UpdateDatasetDocumentResponseBodyData {
	s.Title = &v
	return s
}

func (s *UpdateDatasetDocumentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
