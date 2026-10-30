// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerKnowledgeFileParseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TriggerKnowledgeFileParseResponseBodyData) *TriggerKnowledgeFileParseResponseBody
	GetData() *TriggerKnowledgeFileParseResponseBodyData
	SetRequestId(v string) *TriggerKnowledgeFileParseResponseBody
	GetRequestId() *string
}

type TriggerKnowledgeFileParseResponseBody struct {
	// The returned data.
	Data *TriggerKnowledgeFileParseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerKnowledgeFileParseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerKnowledgeFileParseResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerKnowledgeFileParseResponseBody) GetData() *TriggerKnowledgeFileParseResponseBodyData {
	return s.Data
}

func (s *TriggerKnowledgeFileParseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerKnowledgeFileParseResponseBody) SetData(v *TriggerKnowledgeFileParseResponseBodyData) *TriggerKnowledgeFileParseResponseBody {
	s.Data = v
	return s
}

func (s *TriggerKnowledgeFileParseResponseBody) SetRequestId(v string) *TriggerKnowledgeFileParseResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerKnowledgeFileParseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TriggerKnowledgeFileParseResponseBodyData struct {
	// The ID of the file submitted for reparsing, which is a decimal string. This parameter is returned only when the request is successful.
	//
	// example:
	//
	// 123
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The description of the operation result. If the operation fails, this parameter indicates the cause or the security error code.
	//
	// example:
	//
	// Reparse task submitted
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The status before submission. Valid values: PENDING, COMPLETED, FAILED, SKIPPED, and DUPLICATED. This parameter is returned only when the request is successful.
	//
	// example:
	//
	// COMPLETED
	PreviousStatus *string `json:"PreviousStatus,omitempty" xml:"PreviousStatus,omitempty"`
	// The status after submission. The value is fixed to PENDING. This parameter is returned only when the request is successful.
	//
	// example:
	//
	// PENDING
	ProcessStatus *string `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// Indicates whether the operation is successful. If the operation fails, the Data parameter contains only the Success and Message parameters.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TriggerKnowledgeFileParseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TriggerKnowledgeFileParseResponseBodyData) GoString() string {
	return s.String()
}

func (s *TriggerKnowledgeFileParseResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *TriggerKnowledgeFileParseResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *TriggerKnowledgeFileParseResponseBodyData) GetPreviousStatus() *string {
	return s.PreviousStatus
}

func (s *TriggerKnowledgeFileParseResponseBodyData) GetProcessStatus() *string {
	return s.ProcessStatus
}

func (s *TriggerKnowledgeFileParseResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *TriggerKnowledgeFileParseResponseBodyData) SetFileId(v string) *TriggerKnowledgeFileParseResponseBodyData {
	s.FileId = &v
	return s
}

func (s *TriggerKnowledgeFileParseResponseBodyData) SetMessage(v string) *TriggerKnowledgeFileParseResponseBodyData {
	s.Message = &v
	return s
}

func (s *TriggerKnowledgeFileParseResponseBodyData) SetPreviousStatus(v string) *TriggerKnowledgeFileParseResponseBodyData {
	s.PreviousStatus = &v
	return s
}

func (s *TriggerKnowledgeFileParseResponseBodyData) SetProcessStatus(v string) *TriggerKnowledgeFileParseResponseBodyData {
	s.ProcessStatus = &v
	return s
}

func (s *TriggerKnowledgeFileParseResponseBodyData) SetSuccess(v bool) *TriggerKnowledgeFileParseResponseBodyData {
	s.Success = &v
	return s
}

func (s *TriggerKnowledgeFileParseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
