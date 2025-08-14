// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToPdfJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitConvertImageToPdfJobResponseBody
	GetCode() *string
	SetData(v *SubmitConvertImageToPdfJobResponseBodyData) *SubmitConvertImageToPdfJobResponseBody
	GetData() *SubmitConvertImageToPdfJobResponseBodyData
	SetMessage(v string) *SubmitConvertImageToPdfJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitConvertImageToPdfJobResponseBody
	GetRequestId() *string
}

type SubmitConvertImageToPdfJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitConvertImageToPdfJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitConvertImageToPdfJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToPdfJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToPdfJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitConvertImageToPdfJobResponseBody) GetData() *SubmitConvertImageToPdfJobResponseBodyData {
	return s.Data
}

func (s *SubmitConvertImageToPdfJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitConvertImageToPdfJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitConvertImageToPdfJobResponseBody) SetCode(v string) *SubmitConvertImageToPdfJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitConvertImageToPdfJobResponseBody) SetData(v *SubmitConvertImageToPdfJobResponseBodyData) *SubmitConvertImageToPdfJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitConvertImageToPdfJobResponseBody) SetMessage(v string) *SubmitConvertImageToPdfJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitConvertImageToPdfJobResponseBody) SetRequestId(v string) *SubmitConvertImageToPdfJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitConvertImageToPdfJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitConvertImageToPdfJobResponseBodyData struct {
	// example:
	//
	// docmind-20220810-7c5f9dd4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitConvertImageToPdfJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToPdfJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToPdfJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitConvertImageToPdfJobResponseBodyData) SetId(v string) *SubmitConvertImageToPdfJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitConvertImageToPdfJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
