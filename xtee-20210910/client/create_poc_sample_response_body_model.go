// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePocSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePocSampleResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *CreatePocSampleResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CreatePocSampleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePocSampleResponseBody
	GetRequestId() *string
	SetResultObject(v *CreatePocSampleResponseBodyResultObject) *CreatePocSampleResponseBody
	GetResultObject() *CreatePocSampleResponseBodyResultObject
}

type CreatePocSampleResponseBody struct {
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
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *CreatePocSampleResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s CreatePocSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePocSampleResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePocSampleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePocSampleResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreatePocSampleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePocSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePocSampleResponseBody) GetResultObject() *CreatePocSampleResponseBodyResultObject {
	return s.ResultObject
}

func (s *CreatePocSampleResponseBody) SetCode(v string) *CreatePocSampleResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePocSampleResponseBody) SetHttpStatusCode(v string) *CreatePocSampleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePocSampleResponseBody) SetMessage(v string) *CreatePocSampleResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePocSampleResponseBody) SetRequestId(v string) *CreatePocSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePocSampleResponseBody) SetResultObject(v *CreatePocSampleResponseBodyResultObject) *CreatePocSampleResponseBody {
	s.ResultObject = v
	return s
}

func (s *CreatePocSampleResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePocSampleResponseBodyResultObject struct {
	// example:
	//
	// true
	HasWarnings *bool `json:"HasWarnings,omitempty" xml:"HasWarnings,omitempty"`
	// example:
	//
	// 1
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

func (s CreatePocSampleResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CreatePocSampleResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CreatePocSampleResponseBodyResultObject) GetHasWarnings() *bool {
	return s.HasWarnings
}

func (s *CreatePocSampleResponseBodyResultObject) GetSampleId() *int32 {
	return s.SampleId
}

func (s *CreatePocSampleResponseBodyResultObject) GetSampleName() *string {
	return s.SampleName
}

func (s *CreatePocSampleResponseBodyResultObject) GetTab() *string {
	return s.Tab
}

func (s *CreatePocSampleResponseBodyResultObject) GetWarningMessage() *string {
	return s.WarningMessage
}

func (s *CreatePocSampleResponseBodyResultObject) SetHasWarnings(v bool) *CreatePocSampleResponseBodyResultObject {
	s.HasWarnings = &v
	return s
}

func (s *CreatePocSampleResponseBodyResultObject) SetSampleId(v int32) *CreatePocSampleResponseBodyResultObject {
	s.SampleId = &v
	return s
}

func (s *CreatePocSampleResponseBodyResultObject) SetSampleName(v string) *CreatePocSampleResponseBodyResultObject {
	s.SampleName = &v
	return s
}

func (s *CreatePocSampleResponseBodyResultObject) SetTab(v string) *CreatePocSampleResponseBodyResultObject {
	s.Tab = &v
	return s
}

func (s *CreatePocSampleResponseBodyResultObject) SetWarningMessage(v string) *CreatePocSampleResponseBodyResultObject {
	s.WarningMessage = &v
	return s
}

func (s *CreatePocSampleResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
