// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDatasetDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImportDatasetDataResponseBody
	GetCode() *int32
	SetData(v map[string]interface{}) *ImportDatasetDataResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *ImportDatasetDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportDatasetDataResponseBody
	GetRequestId() *string
}

type ImportDatasetDataResponseBody struct {
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

func (s ImportDatasetDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportDatasetDataResponseBody) GoString() string {
	return s.String()
}

func (s *ImportDatasetDataResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImportDatasetDataResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *ImportDatasetDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportDatasetDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportDatasetDataResponseBody) SetCode(v int32) *ImportDatasetDataResponseBody {
	s.Code = &v
	return s
}

func (s *ImportDatasetDataResponseBody) SetData(v map[string]interface{}) *ImportDatasetDataResponseBody {
	s.Data = v
	return s
}

func (s *ImportDatasetDataResponseBody) SetMessage(v string) *ImportDatasetDataResponseBody {
	s.Message = &v
	return s
}

func (s *ImportDatasetDataResponseBody) SetRequestId(v string) *ImportDatasetDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportDatasetDataResponseBody) Validate() error {
	return dara.Validate(s)
}
