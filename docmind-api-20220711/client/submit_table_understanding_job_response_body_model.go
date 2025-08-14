// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTableUnderstandingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitTableUnderstandingJobResponseBody
	GetCode() *string
	SetData(v *SubmitTableUnderstandingJobResponseBodyData) *SubmitTableUnderstandingJobResponseBody
	GetData() *SubmitTableUnderstandingJobResponseBodyData
	SetMessage(v string) *SubmitTableUnderstandingJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitTableUnderstandingJobResponseBody
	GetRequestId() *string
}

type SubmitTableUnderstandingJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitTableUnderstandingJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitTableUnderstandingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTableUnderstandingJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTableUnderstandingJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitTableUnderstandingJobResponseBody) GetData() *SubmitTableUnderstandingJobResponseBodyData {
	return s.Data
}

func (s *SubmitTableUnderstandingJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitTableUnderstandingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTableUnderstandingJobResponseBody) SetCode(v string) *SubmitTableUnderstandingJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitTableUnderstandingJobResponseBody) SetData(v *SubmitTableUnderstandingJobResponseBodyData) *SubmitTableUnderstandingJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTableUnderstandingJobResponseBody) SetMessage(v string) *SubmitTableUnderstandingJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTableUnderstandingJobResponseBody) SetRequestId(v string) *SubmitTableUnderstandingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTableUnderstandingJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitTableUnderstandingJobResponseBodyData struct {
	// example:
	//
	// docmind-20220816-15bc7965
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitTableUnderstandingJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitTableUnderstandingJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTableUnderstandingJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitTableUnderstandingJobResponseBodyData) SetId(v string) *SubmitTableUnderstandingJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitTableUnderstandingJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
