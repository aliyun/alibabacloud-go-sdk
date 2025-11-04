// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIndexAddDocumentsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitIndexAddDocumentsJobResponseBody
	GetCode() *string
	SetData(v *SubmitIndexAddDocumentsJobResponseBodyData) *SubmitIndexAddDocumentsJobResponseBody
	GetData() *SubmitIndexAddDocumentsJobResponseBodyData
	SetMessage(v string) *SubmitIndexAddDocumentsJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitIndexAddDocumentsJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *SubmitIndexAddDocumentsJobResponseBody
	GetStatus() *string
	SetSuccess(v bool) *SubmitIndexAddDocumentsJobResponseBody
	GetSuccess() *bool
}

type SubmitIndexAddDocumentsJobResponseBody struct {
	// HTTP status code
	//
	// example:
	//
	// Index.InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *SubmitIndexAddDocumentsJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 778C0B3B-03C1-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indications whether the API call is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitIndexAddDocumentsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitIndexAddDocumentsJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIndexAddDocumentsJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitIndexAddDocumentsJobResponseBody) GetData() *SubmitIndexAddDocumentsJobResponseBodyData {
	return s.Data
}

func (s *SubmitIndexAddDocumentsJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitIndexAddDocumentsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitIndexAddDocumentsJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SubmitIndexAddDocumentsJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitIndexAddDocumentsJobResponseBody) SetCode(v string) *SubmitIndexAddDocumentsJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponseBody) SetData(v *SubmitIndexAddDocumentsJobResponseBodyData) *SubmitIndexAddDocumentsJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponseBody) SetMessage(v string) *SubmitIndexAddDocumentsJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponseBody) SetRequestId(v string) *SubmitIndexAddDocumentsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponseBody) SetStatus(v string) *SubmitIndexAddDocumentsJobResponseBody {
	s.Status = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponseBody) SetSuccess(v bool) *SubmitIndexAddDocumentsJobResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitIndexAddDocumentsJobResponseBodyData struct {
	// The primary key ID of the task, `JobId`.
	//
	// example:
	//
	// 42687eb254a34802bed398357f5498ae
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitIndexAddDocumentsJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitIndexAddDocumentsJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitIndexAddDocumentsJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitIndexAddDocumentsJobResponseBodyData) SetId(v string) *SubmitIndexAddDocumentsJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
