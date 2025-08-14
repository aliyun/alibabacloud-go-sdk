// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertPdfToExcelJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitConvertPdfToExcelJobResponseBody
	GetCode() *string
	SetData(v *SubmitConvertPdfToExcelJobResponseBodyData) *SubmitConvertPdfToExcelJobResponseBody
	GetData() *SubmitConvertPdfToExcelJobResponseBodyData
	SetMessage(v string) *SubmitConvertPdfToExcelJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitConvertPdfToExcelJobResponseBody
	GetRequestId() *string
}

type SubmitConvertPdfToExcelJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertPdfToExcelJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertPdfToExcelJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToExcelJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToExcelJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitConvertPdfToExcelJobResponseBody) GetData() *SubmitConvertPdfToExcelJobResponseBodyData {
	return s.Data
}

func (s *SubmitConvertPdfToExcelJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitConvertPdfToExcelJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitConvertPdfToExcelJobResponseBody) SetCode(v string) *SubmitConvertPdfToExcelJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponseBody) SetData(v *SubmitConvertPdfToExcelJobResponseBodyData) *SubmitConvertPdfToExcelJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponseBody) SetMessage(v string) *SubmitConvertPdfToExcelJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponseBody) SetRequestId(v string) *SubmitConvertPdfToExcelJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitConvertPdfToExcelJobResponseBodyData struct {
	// example:
	//
	// docmind-20220816-15bc7965
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertPdfToExcelJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToExcelJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToExcelJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitConvertPdfToExcelJobResponseBodyData) SetId(v string) *SubmitConvertPdfToExcelJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
