// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaterialDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMaterialDocumentResponseBody
	GetCode() *string
	SetData(v int64) *UpdateMaterialDocumentResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateMaterialDocumentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMaterialDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMaterialDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMaterialDocumentResponseBody
	GetSuccess() *bool
}

type UpdateMaterialDocumentResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 82
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

func (s UpdateMaterialDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaterialDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMaterialDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMaterialDocumentResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateMaterialDocumentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMaterialDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMaterialDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMaterialDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMaterialDocumentResponseBody) SetCode(v string) *UpdateMaterialDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetData(v int64) *UpdateMaterialDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetHttpStatusCode(v int32) *UpdateMaterialDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetMessage(v string) *UpdateMaterialDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetRequestId(v string) *UpdateMaterialDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetSuccess(v bool) *UpdateMaterialDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
