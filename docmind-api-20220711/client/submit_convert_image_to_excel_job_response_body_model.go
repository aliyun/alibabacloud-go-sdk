// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToExcelJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitConvertImageToExcelJobResponseBody
	GetCode() *string
	SetData(v *SubmitConvertImageToExcelJobResponseBodyData) *SubmitConvertImageToExcelJobResponseBody
	GetData() *SubmitConvertImageToExcelJobResponseBodyData
	SetMessage(v string) *SubmitConvertImageToExcelJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitConvertImageToExcelJobResponseBody
	GetRequestId() *string
}

type SubmitConvertImageToExcelJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertImageToExcelJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertImageToExcelJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToExcelJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToExcelJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitConvertImageToExcelJobResponseBody) GetData() *SubmitConvertImageToExcelJobResponseBodyData {
	return s.Data
}

func (s *SubmitConvertImageToExcelJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitConvertImageToExcelJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitConvertImageToExcelJobResponseBody) SetCode(v string) *SubmitConvertImageToExcelJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertImageToExcelJobResponseBody) SetData(v *SubmitConvertImageToExcelJobResponseBodyData) *SubmitConvertImageToExcelJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertImageToExcelJobResponseBody) SetMessage(v string) *SubmitConvertImageToExcelJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertImageToExcelJobResponseBody) SetRequestId(v string) *SubmitConvertImageToExcelJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitConvertImageToExcelJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitConvertImageToExcelJobResponseBodyData struct {
	// example:
	//
	// docmind-20220810-7c5f9dd4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertImageToExcelJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToExcelJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToExcelJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitConvertImageToExcelJobResponseBodyData) SetId(v string) *SubmitConvertImageToExcelJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitConvertImageToExcelJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
