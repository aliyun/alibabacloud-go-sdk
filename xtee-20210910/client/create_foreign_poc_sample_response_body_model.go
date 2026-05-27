// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForeignPocSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateForeignPocSampleResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *CreateForeignPocSampleResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CreateForeignPocSampleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateForeignPocSampleResponseBody
	GetRequestId() *string
	SetResultObject(v *CreateForeignPocSampleResponseBodyResultObject) *CreateForeignPocSampleResponseBody
	GetResultObject() *CreateForeignPocSampleResponseBodyResultObject
}

type CreateForeignPocSampleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *CreateForeignPocSampleResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CreateForeignPocSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateForeignPocSampleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateForeignPocSampleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateForeignPocSampleResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateForeignPocSampleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateForeignPocSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateForeignPocSampleResponseBody) GetResultObject() *CreateForeignPocSampleResponseBodyResultObject {
	return s.ResultObject
}

func (s *CreateForeignPocSampleResponseBody) SetCode(v string) *CreateForeignPocSampleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateForeignPocSampleResponseBody) SetHttpStatusCode(v string) *CreateForeignPocSampleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateForeignPocSampleResponseBody) SetMessage(v string) *CreateForeignPocSampleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateForeignPocSampleResponseBody) SetRequestId(v string) *CreateForeignPocSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateForeignPocSampleResponseBody) SetResultObject(v *CreateForeignPocSampleResponseBodyResultObject) *CreateForeignPocSampleResponseBody {
	s.ResultObject = v
	return s
}

func (s *CreateForeignPocSampleResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateForeignPocSampleResponseBodyResultObject struct {
	// example:
	//
	// true
	HasWarnings *bool `json:"HasWarnings,omitempty" xml:"HasWarnings,omitempty"`
	// example:
	//
	// 174
	SampleId *int32 `json:"SampleId,omitempty" xml:"SampleId,omitempty"`
	// example:
	//
	// SampleNameTest
	SampleName *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// example:
	//
	// test
	WarningMessage *string `json:"WarningMessage,omitempty" xml:"WarningMessage,omitempty"`
}

func (s CreateForeignPocSampleResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CreateForeignPocSampleResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CreateForeignPocSampleResponseBodyResultObject) GetHasWarnings() *bool {
	return s.HasWarnings
}

func (s *CreateForeignPocSampleResponseBodyResultObject) GetSampleId() *int32 {
	return s.SampleId
}

func (s *CreateForeignPocSampleResponseBodyResultObject) GetSampleName() *string {
	return s.SampleName
}

func (s *CreateForeignPocSampleResponseBodyResultObject) GetTab() *string {
	return s.Tab
}

func (s *CreateForeignPocSampleResponseBodyResultObject) GetWarningMessage() *string {
	return s.WarningMessage
}

func (s *CreateForeignPocSampleResponseBodyResultObject) SetHasWarnings(v bool) *CreateForeignPocSampleResponseBodyResultObject {
	s.HasWarnings = &v
	return s
}

func (s *CreateForeignPocSampleResponseBodyResultObject) SetSampleId(v int32) *CreateForeignPocSampleResponseBodyResultObject {
	s.SampleId = &v
	return s
}

func (s *CreateForeignPocSampleResponseBodyResultObject) SetSampleName(v string) *CreateForeignPocSampleResponseBodyResultObject {
	s.SampleName = &v
	return s
}

func (s *CreateForeignPocSampleResponseBodyResultObject) SetTab(v string) *CreateForeignPocSampleResponseBodyResultObject {
	s.Tab = &v
	return s
}

func (s *CreateForeignPocSampleResponseBodyResultObject) SetWarningMessage(v string) *CreateForeignPocSampleResponseBodyResultObject {
	s.WarningMessage = &v
	return s
}

func (s *CreateForeignPocSampleResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
