// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportDocumentsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ImportDocumentsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportDocumentsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ImportDocumentsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ImportDocumentsResponseBody
	GetRequestId() *string
}

type ImportDocumentsResponseBody struct {
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
	// 01B12EE4-6AF2-4730-8B78-EC15F4E5C025
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ImportDocumentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportDocumentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportDocumentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportDocumentsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ImportDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportDocumentsResponseBody) SetCode(v string) *ImportDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *ImportDocumentsResponseBody) SetHttpStatusCode(v int32) *ImportDocumentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportDocumentsResponseBody) SetMessage(v string) *ImportDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *ImportDocumentsResponseBody) SetParams(v []*string) *ImportDocumentsResponseBody {
	s.Params = v
	return s
}

func (s *ImportDocumentsResponseBody) SetRequestId(v string) *ImportDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportDocumentsResponseBody) Validate() error {
	return dara.Validate(s)
}
