// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertPdfToImageJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitConvertPdfToImageJobResponseBody
	GetCode() *string
	SetData(v *SubmitConvertPdfToImageJobResponseBodyData) *SubmitConvertPdfToImageJobResponseBody
	GetData() *SubmitConvertPdfToImageJobResponseBodyData
	SetMessage(v string) *SubmitConvertPdfToImageJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitConvertPdfToImageJobResponseBody
	GetRequestId() *string
}

type SubmitConvertPdfToImageJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertPdfToImageJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertPdfToImageJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToImageJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToImageJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitConvertPdfToImageJobResponseBody) GetData() *SubmitConvertPdfToImageJobResponseBodyData {
	return s.Data
}

func (s *SubmitConvertPdfToImageJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitConvertPdfToImageJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitConvertPdfToImageJobResponseBody) SetCode(v string) *SubmitConvertPdfToImageJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertPdfToImageJobResponseBody) SetData(v *SubmitConvertPdfToImageJobResponseBodyData) *SubmitConvertPdfToImageJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertPdfToImageJobResponseBody) SetMessage(v string) *SubmitConvertPdfToImageJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertPdfToImageJobResponseBody) SetRequestId(v string) *SubmitConvertPdfToImageJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitConvertPdfToImageJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitConvertPdfToImageJobResponseBodyData struct {
	// example:
	//
	// docmind-20220810-7c5f9dd4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertPdfToImageJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToImageJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToImageJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitConvertPdfToImageJobResponseBodyData) SetId(v string) *SubmitConvertPdfToImageJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitConvertPdfToImageJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
