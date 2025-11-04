// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddFileResponseBody
	GetCode() *string
	SetData(v *AddFileResponseBodyData) *AddFileResponseBody
	GetData() *AddFileResponseBodyData
	SetMessage(v string) *AddFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddFileResponseBody
	GetRequestId() *string
	SetStatus(v string) *AddFileResponseBody
	GetStatus() *string
	SetSuccess(v string) *AddFileResponseBody
	GetSuccess() *string
}

type AddFileResponseBody struct {
	// The status code.
	//
	// example:
	//
	// DataCenter.FileTooLarge
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data fields.
	Data *AddFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 778C0B3B-xxxx-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indications whether the call is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFileResponseBody) GoString() string {
	return s.String()
}

func (s *AddFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddFileResponseBody) GetData() *AddFileResponseBodyData {
	return s.Data
}

func (s *AddFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *AddFileResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AddFileResponseBody) SetCode(v string) *AddFileResponseBody {
	s.Code = &v
	return s
}

func (s *AddFileResponseBody) SetData(v *AddFileResponseBodyData) *AddFileResponseBody {
	s.Data = v
	return s
}

func (s *AddFileResponseBody) SetMessage(v string) *AddFileResponseBody {
	s.Message = &v
	return s
}

func (s *AddFileResponseBody) SetRequestId(v string) *AddFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFileResponseBody) SetStatus(v string) *AddFileResponseBody {
	s.Status = &v
	return s
}

func (s *AddFileResponseBody) SetSuccess(v string) *AddFileResponseBody {
	s.Success = &v
	return s
}

func (s *AddFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFileResponseBodyData struct {
	// The primary key ID of the document. We recommend that you store the ID because it is required for all subsequent API operations related to this document.
	//
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The parser that is used to parse the document. Valid value:
	//
	// 	- DASHSCOPE_DOCMIND: Intelligent document parsing by Alibaba Cloud.
	//
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
}

func (s AddFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFileResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *AddFileResponseBodyData) GetParser() *string {
	return s.Parser
}

func (s *AddFileResponseBodyData) SetFileId(v string) *AddFileResponseBodyData {
	s.FileId = &v
	return s
}

func (s *AddFileResponseBodyData) SetParser(v string) *AddFileResponseBodyData {
	s.Parser = &v
	return s
}

func (s *AddFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
