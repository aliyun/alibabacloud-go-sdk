// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmartAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitSmartAuditResponseBody
	GetCode() *string
	SetData(v *SubmitSmartAuditResponseBodyData) *SubmitSmartAuditResponseBody
	GetData() *SubmitSmartAuditResponseBodyData
	SetHttpStatusCode(v int32) *SubmitSmartAuditResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitSmartAuditResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitSmartAuditResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitSmartAuditResponseBody
	GetSuccess() *bool
}

type SubmitSmartAuditResponseBody struct {
	// Error code
	//
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Task response object
	Data *SubmitSmartAuditResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitSmartAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartAuditResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSmartAuditResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitSmartAuditResponseBody) GetData() *SubmitSmartAuditResponseBodyData {
	return s.Data
}

func (s *SubmitSmartAuditResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitSmartAuditResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitSmartAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSmartAuditResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitSmartAuditResponseBody) SetCode(v string) *SubmitSmartAuditResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitSmartAuditResponseBody) SetData(v *SubmitSmartAuditResponseBodyData) *SubmitSmartAuditResponseBody {
	s.Data = v
	return s
}

func (s *SubmitSmartAuditResponseBody) SetHttpStatusCode(v int32) *SubmitSmartAuditResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitSmartAuditResponseBody) SetMessage(v string) *SubmitSmartAuditResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitSmartAuditResponseBody) SetRequestId(v string) *SubmitSmartAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSmartAuditResponseBody) SetSuccess(v bool) *SubmitSmartAuditResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSmartAuditResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitSmartAuditResponseBodyData struct {
	// Task ID
	//
	// example:
	//
	// xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitSmartAuditResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartAuditResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitSmartAuditResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitSmartAuditResponseBodyData) SetTaskId(v string) *SubmitSmartAuditResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitSmartAuditResponseBodyData) Validate() error {
	return dara.Validate(s)
}
