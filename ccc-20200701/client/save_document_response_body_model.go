// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveDocumentResponseBody
	GetCode() *string
	SetData(v string) *SaveDocumentResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *SaveDocumentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveDocumentResponseBody
	GetMessage() *string
	SetParams(v []*string) *SaveDocumentResponseBody
	GetParams() []*string
	SetRequestId(v string) *SaveDocumentResponseBody
	GetRequestId() *string
}

type SaveDocumentResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 0101234****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 01B12EE4-6AF2-4730-8B78-EC15F4E5C025
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *SaveDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveDocumentResponseBody) GetData() *string {
	return s.Data
}

func (s *SaveDocumentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveDocumentResponseBody) GetParams() []*string {
	return s.Params
}

func (s *SaveDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveDocumentResponseBody) SetCode(v string) *SaveDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *SaveDocumentResponseBody) SetData(v string) *SaveDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *SaveDocumentResponseBody) SetHttpStatusCode(v int32) *SaveDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveDocumentResponseBody) SetMessage(v string) *SaveDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *SaveDocumentResponseBody) SetParams(v []*string) *SaveDocumentResponseBody {
	s.Params = v
	return s
}

func (s *SaveDocumentResponseBody) SetRequestId(v string) *SaveDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
