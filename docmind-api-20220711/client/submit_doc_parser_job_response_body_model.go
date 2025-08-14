// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocParserJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitDocParserJobResponseBody
	GetCode() *string
	SetData(v *SubmitDocParserJobResponseBodyData) *SubmitDocParserJobResponseBody
	GetData() *SubmitDocParserJobResponseBodyData
	SetMessage(v string) *SubmitDocParserJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitDocParserJobResponseBody
	GetRequestId() *string
}

type SubmitDocParserJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitDocParserJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDocParserJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitDocParserJobResponseBody) GetData() *SubmitDocParserJobResponseBodyData {
	return s.Data
}

func (s *SubmitDocParserJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitDocParserJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDocParserJobResponseBody) SetCode(v string) *SubmitDocParserJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDocParserJobResponseBody) SetData(v *SubmitDocParserJobResponseBodyData) *SubmitDocParserJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocParserJobResponseBody) SetMessage(v string) *SubmitDocParserJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDocParserJobResponseBody) SetRequestId(v string) *SubmitDocParserJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDocParserJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitDocParserJobResponseBodyData struct {
	// example:
	//
	// docmind-20220816-15bc7965
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitDocParserJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitDocParserJobResponseBodyData) SetId(v string) *SubmitDocParserJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitDocParserJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
