// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDatasetImportDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OpenDatasetImportDataResponseBody
	GetCode() *int32
	SetData(v map[string]interface{}) *OpenDatasetImportDataResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *OpenDatasetImportDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *OpenDatasetImportDataResponseBody
	GetRequestId() *string
}

type OpenDatasetImportDataResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// []
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 1a0f40dd17774641794394269ec0e9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OpenDatasetImportDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenDatasetImportDataResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDatasetImportDataResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OpenDatasetImportDataResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *OpenDatasetImportDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenDatasetImportDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenDatasetImportDataResponseBody) SetCode(v int32) *OpenDatasetImportDataResponseBody {
	s.Code = &v
	return s
}

func (s *OpenDatasetImportDataResponseBody) SetData(v map[string]interface{}) *OpenDatasetImportDataResponseBody {
	s.Data = v
	return s
}

func (s *OpenDatasetImportDataResponseBody) SetMessage(v string) *OpenDatasetImportDataResponseBody {
	s.Message = &v
	return s
}

func (s *OpenDatasetImportDataResponseBody) SetRequestId(v string) *OpenDatasetImportDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenDatasetImportDataResponseBody) Validate() error {
	return dara.Validate(s)
}
