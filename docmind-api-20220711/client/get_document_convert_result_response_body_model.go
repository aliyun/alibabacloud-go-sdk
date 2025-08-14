// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentConvertResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocumentConvertResultResponseBody
	GetCode() *string
	SetCompleted(v bool) *GetDocumentConvertResultResponseBody
	GetCompleted() *bool
	SetData(v []*GetDocumentConvertResultResponseBodyData) *GetDocumentConvertResultResponseBody
	GetData() []*GetDocumentConvertResultResponseBodyData
	SetMessage(v string) *GetDocumentConvertResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocumentConvertResultResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetDocumentConvertResultResponseBody
	GetStatus() *string
}

type GetDocumentConvertResultResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	Completed *bool                                       `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      []*GetDocumentConvertResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDocumentConvertResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentConvertResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentConvertResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocumentConvertResultResponseBody) GetCompleted() *bool {
	return s.Completed
}

func (s *GetDocumentConvertResultResponseBody) GetData() []*GetDocumentConvertResultResponseBodyData {
	return s.Data
}

func (s *GetDocumentConvertResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentConvertResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentConvertResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetDocumentConvertResultResponseBody) SetCode(v string) *GetDocumentConvertResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentConvertResultResponseBody) SetCompleted(v bool) *GetDocumentConvertResultResponseBody {
	s.Completed = &v
	return s
}

func (s *GetDocumentConvertResultResponseBody) SetData(v []*GetDocumentConvertResultResponseBodyData) *GetDocumentConvertResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentConvertResultResponseBody) SetMessage(v string) *GetDocumentConvertResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentConvertResultResponseBody) SetRequestId(v string) *GetDocumentConvertResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentConvertResultResponseBody) SetStatus(v string) *GetDocumentConvertResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetDocumentConvertResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDocumentConvertResultResponseBodyData struct {
	// example:
	//
	// e6d83e55df218650b9a296bfbc300076
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// 2355965
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// pdf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// http://example.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetDocumentConvertResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentConvertResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentConvertResultResponseBodyData) GetMd5() *string {
	return s.Md5
}

func (s *GetDocumentConvertResultResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *GetDocumentConvertResultResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetDocumentConvertResultResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetDocumentConvertResultResponseBodyData) SetMd5(v string) *GetDocumentConvertResultResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *GetDocumentConvertResultResponseBodyData) SetSize(v int64) *GetDocumentConvertResultResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetDocumentConvertResultResponseBodyData) SetType(v string) *GetDocumentConvertResultResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetDocumentConvertResultResponseBodyData) SetUrl(v string) *GetDocumentConvertResultResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetDocumentConvertResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
