// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocumentExtractJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitDocumentExtractJobResponseBody
	GetCode() *string
	SetData(v *SubmitDocumentExtractJobResponseBodyData) *SubmitDocumentExtractJobResponseBody
	GetData() *SubmitDocumentExtractJobResponseBodyData
	SetMessage(v string) *SubmitDocumentExtractJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitDocumentExtractJobResponseBody
	GetRequestId() *string
}

type SubmitDocumentExtractJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitDocumentExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDocumentExtractJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocumentExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocumentExtractJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitDocumentExtractJobResponseBody) GetData() *SubmitDocumentExtractJobResponseBodyData {
	return s.Data
}

func (s *SubmitDocumentExtractJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitDocumentExtractJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDocumentExtractJobResponseBody) SetCode(v string) *SubmitDocumentExtractJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDocumentExtractJobResponseBody) SetData(v *SubmitDocumentExtractJobResponseBodyData) *SubmitDocumentExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocumentExtractJobResponseBody) SetMessage(v string) *SubmitDocumentExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDocumentExtractJobResponseBody) SetRequestId(v string) *SubmitDocumentExtractJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDocumentExtractJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitDocumentExtractJobResponseBodyData struct {
	// example:
	//
	// docmind-20220816-15bc7965
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitDocumentExtractJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocumentExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocumentExtractJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitDocumentExtractJobResponseBodyData) SetId(v string) *SubmitDocumentExtractJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitDocumentExtractJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
