// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToWordJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitConvertImageToWordJobResponseBody
	GetCode() *string
	SetData(v *SubmitConvertImageToWordJobResponseBodyData) *SubmitConvertImageToWordJobResponseBody
	GetData() *SubmitConvertImageToWordJobResponseBodyData
	SetMessage(v string) *SubmitConvertImageToWordJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitConvertImageToWordJobResponseBody
	GetRequestId() *string
}

type SubmitConvertImageToWordJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertImageToWordJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertImageToWordJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToWordJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToWordJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitConvertImageToWordJobResponseBody) GetData() *SubmitConvertImageToWordJobResponseBodyData {
	return s.Data
}

func (s *SubmitConvertImageToWordJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitConvertImageToWordJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitConvertImageToWordJobResponseBody) SetCode(v string) *SubmitConvertImageToWordJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertImageToWordJobResponseBody) SetData(v *SubmitConvertImageToWordJobResponseBodyData) *SubmitConvertImageToWordJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertImageToWordJobResponseBody) SetMessage(v string) *SubmitConvertImageToWordJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertImageToWordJobResponseBody) SetRequestId(v string) *SubmitConvertImageToWordJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitConvertImageToWordJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitConvertImageToWordJobResponseBodyData struct {
	// example:
	//
	// docmind-20220810-7c5f9dd4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertImageToWordJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToWordJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToWordJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitConvertImageToWordJobResponseBodyData) SetId(v string) *SubmitConvertImageToWordJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitConvertImageToWordJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
