// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableUnderstandingResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTableUnderstandingResultResponseBody
	GetCode() *string
	SetCompleted(v bool) *GetTableUnderstandingResultResponseBody
	GetCompleted() *bool
	SetData(v map[string]interface{}) *GetTableUnderstandingResultResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *GetTableUnderstandingResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTableUnderstandingResultResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetTableUnderstandingResultResponseBody
	GetStatus() *string
}

type GetTableUnderstandingResultResponseBody struct {
	// example:
	//
	// noPermission
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool                  `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTableUnderstandingResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableUnderstandingResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableUnderstandingResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTableUnderstandingResultResponseBody) GetCompleted() *bool {
	return s.Completed
}

func (s *GetTableUnderstandingResultResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetTableUnderstandingResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTableUnderstandingResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableUnderstandingResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetTableUnderstandingResultResponseBody) SetCode(v string) *GetTableUnderstandingResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableUnderstandingResultResponseBody) SetCompleted(v bool) *GetTableUnderstandingResultResponseBody {
	s.Completed = &v
	return s
}

func (s *GetTableUnderstandingResultResponseBody) SetData(v map[string]interface{}) *GetTableUnderstandingResultResponseBody {
	s.Data = v
	return s
}

func (s *GetTableUnderstandingResultResponseBody) SetMessage(v string) *GetTableUnderstandingResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableUnderstandingResultResponseBody) SetRequestId(v string) *GetTableUnderstandingResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableUnderstandingResultResponseBody) SetStatus(v string) *GetTableUnderstandingResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetTableUnderstandingResultResponseBody) Validate() error {
	return dara.Validate(s)
}
