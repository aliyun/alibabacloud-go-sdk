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
	// The error code.
	//
	// example:
	//
	// Index.InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned by the operation.
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
	// The status code returned by the operation.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitIndexJobResponseBodyData struct {
	// The task ID, which is the `JobId` required when calling the **GetIndexJobStatus*	- operation.
	//
	// example:
	//
	// eFDr2fGRzP9gdDZWAdo3xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The knowledge base ID.
	//
	// example:
	//
	// 79c0alxxxx
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
