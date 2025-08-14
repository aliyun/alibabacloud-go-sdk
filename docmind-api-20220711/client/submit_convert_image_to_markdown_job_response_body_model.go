// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToMarkdownJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitConvertImageToMarkdownJobResponseBody
	GetCode() *string
	SetData(v *SubmitConvertImageToMarkdownJobResponseBodyData) *SubmitConvertImageToMarkdownJobResponseBody
	GetData() *SubmitConvertImageToMarkdownJobResponseBodyData
	SetMessage(v string) *SubmitConvertImageToMarkdownJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitConvertImageToMarkdownJobResponseBody
	GetRequestId() *string
}

type SubmitConvertImageToMarkdownJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertImageToMarkdownJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertImageToMarkdownJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToMarkdownJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) GetData() *SubmitConvertImageToMarkdownJobResponseBodyData {
	return s.Data
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) SetCode(v string) *SubmitConvertImageToMarkdownJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) SetData(v *SubmitConvertImageToMarkdownJobResponseBodyData) *SubmitConvertImageToMarkdownJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) SetMessage(v string) *SubmitConvertImageToMarkdownJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) SetRequestId(v string) *SubmitConvertImageToMarkdownJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitConvertImageToMarkdownJobResponseBodyData struct {
	// example:
	//
	// docmind-20220810-7c5f9dd4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertImageToMarkdownJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToMarkdownJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToMarkdownJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitConvertImageToMarkdownJobResponseBodyData) SetId(v string) *SubmitConvertImageToMarkdownJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
