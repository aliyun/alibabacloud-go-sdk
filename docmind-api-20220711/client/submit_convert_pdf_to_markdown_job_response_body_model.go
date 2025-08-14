// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertPdfToMarkdownJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitConvertPdfToMarkdownJobResponseBody
	GetCode() *string
	SetData(v *SubmitConvertPdfToMarkdownJobResponseBodyData) *SubmitConvertPdfToMarkdownJobResponseBody
	GetData() *SubmitConvertPdfToMarkdownJobResponseBodyData
	SetMessage(v string) *SubmitConvertPdfToMarkdownJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitConvertPdfToMarkdownJobResponseBody
	GetRequestId() *string
}

type SubmitConvertPdfToMarkdownJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertPdfToMarkdownJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertPdfToMarkdownJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToMarkdownJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) GetData() *SubmitConvertPdfToMarkdownJobResponseBodyData {
	return s.Data
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) SetCode(v string) *SubmitConvertPdfToMarkdownJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) SetData(v *SubmitConvertPdfToMarkdownJobResponseBodyData) *SubmitConvertPdfToMarkdownJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) SetMessage(v string) *SubmitConvertPdfToMarkdownJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) SetRequestId(v string) *SubmitConvertPdfToMarkdownJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitConvertPdfToMarkdownJobResponseBodyData struct {
	// example:
	//
	// docmind-20220816-15bc7965
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertPdfToMarkdownJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToMarkdownJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToMarkdownJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitConvertPdfToMarkdownJobResponseBodyData) SetId(v string) *SubmitConvertPdfToMarkdownJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
