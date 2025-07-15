// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDocumentsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteDocumentsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteDocumentsResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteDocumentsResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteDocumentsResponseBody
	GetRequestId() *string
}

type DeleteDocumentsResponseBody struct {
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
	// A450574A-337F-43E2-BC59-9C6594C994C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocumentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDocumentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDocumentsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocumentsResponseBody) SetCode(v string) *DeleteDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDocumentsResponseBody) SetHttpStatusCode(v int32) *DeleteDocumentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDocumentsResponseBody) SetMessage(v string) *DeleteDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDocumentsResponseBody) SetParams(v []*string) *DeleteDocumentsResponseBody {
	s.Params = v
	return s
}

func (s *DeleteDocumentsResponseBody) SetRequestId(v string) *DeleteDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentsResponseBody) Validate() error {
	return dara.Validate(s)
}
