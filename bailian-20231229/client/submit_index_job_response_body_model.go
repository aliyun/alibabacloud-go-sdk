// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIndexJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitIndexJobResponseBody
	GetCode() *string
	SetData(v *SubmitIndexJobResponseBodyData) *SubmitIndexJobResponseBody
	GetData() *SubmitIndexJobResponseBodyData
	SetMessage(v string) *SubmitIndexJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitIndexJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *SubmitIndexJobResponseBody
	GetStatus() *string
	SetSuccess(v bool) *SubmitIndexJobResponseBody
	GetSuccess() *bool
}

type SubmitIndexJobResponseBody struct {
	// HTTP status code
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *SubmitIndexJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indications whether the API call is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitIndexJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitIndexJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIndexJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitIndexJobResponseBody) GetData() *SubmitIndexJobResponseBodyData {
	return s.Data
}

func (s *SubmitIndexJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitIndexJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitIndexJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SubmitIndexJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitIndexJobResponseBody) SetCode(v string) *SubmitIndexJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitIndexJobResponseBody) SetData(v *SubmitIndexJobResponseBodyData) *SubmitIndexJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitIndexJobResponseBody) SetMessage(v string) *SubmitIndexJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitIndexJobResponseBody) SetRequestId(v string) *SubmitIndexJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIndexJobResponseBody) SetStatus(v string) *SubmitIndexJobResponseBody {
	s.Status = &v
	return s
}

func (s *SubmitIndexJobResponseBody) SetSuccess(v bool) *SubmitIndexJobResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitIndexJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitIndexJobResponseBodyData struct {
	// The primary key ID of the job, which is the `JobId` parameter of the [GetIndexJobStatus](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-getindexjobstatus) operation.
	//
	// example:
	//
	// eFDr2fGRzP9gdDZWAdo3YQ==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The primary key ID of the knowledge base.
	//
	// example:
	//
	// khdyak1uuj
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s SubmitIndexJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitIndexJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitIndexJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitIndexJobResponseBodyData) GetIndexId() *string {
	return s.IndexId
}

func (s *SubmitIndexJobResponseBodyData) SetId(v string) *SubmitIndexJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitIndexJobResponseBodyData) SetIndexId(v string) *SubmitIndexJobResponseBodyData {
	s.IndexId = &v
	return s
}

func (s *SubmitIndexJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
