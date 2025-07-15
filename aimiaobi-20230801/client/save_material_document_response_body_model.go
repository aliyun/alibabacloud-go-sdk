// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMaterialDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveMaterialDocumentResponseBody
	GetCode() *string
	SetData(v int64) *SaveMaterialDocumentResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *SaveMaterialDocumentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveMaterialDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveMaterialDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveMaterialDocumentResponseBody
	GetSuccess() *bool
}

type SaveMaterialDocumentResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 12
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveMaterialDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveMaterialDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *SaveMaterialDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveMaterialDocumentResponseBody) GetData() *int64 {
	return s.Data
}

func (s *SaveMaterialDocumentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveMaterialDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveMaterialDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveMaterialDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveMaterialDocumentResponseBody) SetCode(v string) *SaveMaterialDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetData(v int64) *SaveMaterialDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetHttpStatusCode(v int32) *SaveMaterialDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetMessage(v string) *SaveMaterialDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetRequestId(v string) *SaveMaterialDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetSuccess(v bool) *SaveMaterialDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
