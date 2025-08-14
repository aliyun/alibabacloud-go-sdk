// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertPdfToWordJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitConvertPdfToWordJobResponseBody
	GetCode() *string
	SetData(v *SubmitConvertPdfToWordJobResponseBodyData) *SubmitConvertPdfToWordJobResponseBody
	GetData() *SubmitConvertPdfToWordJobResponseBodyData
	SetMessage(v string) *SubmitConvertPdfToWordJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitConvertPdfToWordJobResponseBody
	GetRequestId() *string
}

type SubmitConvertPdfToWordJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertPdfToWordJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertPdfToWordJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToWordJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToWordJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitConvertPdfToWordJobResponseBody) GetData() *SubmitConvertPdfToWordJobResponseBodyData {
	return s.Data
}

func (s *SubmitConvertPdfToWordJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitConvertPdfToWordJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitConvertPdfToWordJobResponseBody) SetCode(v string) *SubmitConvertPdfToWordJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertPdfToWordJobResponseBody) SetData(v *SubmitConvertPdfToWordJobResponseBodyData) *SubmitConvertPdfToWordJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertPdfToWordJobResponseBody) SetMessage(v string) *SubmitConvertPdfToWordJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertPdfToWordJobResponseBody) SetRequestId(v string) *SubmitConvertPdfToWordJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitConvertPdfToWordJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitConvertPdfToWordJobResponseBodyData struct {
	// example:
	//
	// docmind-20220816-15bc7965
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertPdfToWordJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToWordJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToWordJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitConvertPdfToWordJobResponseBodyData) SetId(v string) *SubmitConvertPdfToWordJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitConvertPdfToWordJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
