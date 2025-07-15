// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDatasetDocumentResponseBody
	GetCode() *string
	SetData(v []*string) *DeleteDatasetDocumentResponseBody
	GetData() []*string
	SetHttpStatusCode(v int32) *DeleteDatasetDocumentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDatasetDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDatasetDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDatasetDocumentResponseBody
	GetSuccess() *bool
}

type DeleteDatasetDocumentResponseBody struct {
	// example:
	//
	// NoData
	Code *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DeleteDatasetDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDatasetDocumentResponseBody) GetData() []*string {
	return s.Data
}

func (s *DeleteDatasetDocumentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDatasetDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDatasetDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasetDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDatasetDocumentResponseBody) SetCode(v string) *DeleteDatasetDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDatasetDocumentResponseBody) SetData(v []*string) *DeleteDatasetDocumentResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDatasetDocumentResponseBody) SetHttpStatusCode(v int32) *DeleteDatasetDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDatasetDocumentResponseBody) SetMessage(v string) *DeleteDatasetDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDatasetDocumentResponseBody) SetRequestId(v string) *DeleteDatasetDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetDocumentResponseBody) SetSuccess(v bool) *DeleteDatasetDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDatasetDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
