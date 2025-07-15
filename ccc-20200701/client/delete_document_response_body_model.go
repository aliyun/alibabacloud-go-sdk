// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDocumentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteDocumentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDocumentResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteDocumentResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteDocumentResponseBody
	GetRequestId() *string
}

type DeleteDocumentResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 8707EB29-BAED-4302-B999-40BA61877437
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDocumentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDocumentResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocumentResponseBody) SetCode(v string) *DeleteDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetHttpStatusCode(v int32) *DeleteDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetMessage(v string) *DeleteDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetParams(v []*string) *DeleteDocumentResponseBody {
	s.Params = v
	return s
}

func (s *DeleteDocumentResponseBody) SetRequestId(v string) *DeleteDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
