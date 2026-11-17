// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody
	GetCode() *string
	SetData(v *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody
	GetData() *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData
	SetMessage(v string) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody
	GetRequestId() *string
}

type CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	//
	// example:
	//
	// {
	//
	//   "taskId": "finops-export-abcdef1234567890",
	//
	//   "status": "Pending"
	//
	// }
	Data *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID of this OpenAPI call, which is used to troubleshoot call issues. This is different from the request_id of each consumption record in the exported file.
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) GetData() *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	return s.Data
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) SetCode(v string) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) SetData(v *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) SetMessage(v string) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) SetRequestId(v string) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData struct {
	// The current status of create task.
	//
	// - Pending: Create task is submitted and waiting to be processed.
	//
	// - Running: The system is querying details, generating the file, or uploading the file.
	//
	// - Success: The export is successful. Call the query operation to obtain the file information.
	//
	// - Failed: The export fails. Call the query operation to view the errorMessage.
	//
	// A newly created task returns Pending. When an existing task is reused through clientToken, the status of create task at that time is returned, which is not guaranteed to be Pending.
	//
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the export task. When you call GetGatewayQuotaRuleSubjectUsageExportTask, pass this value along with the gatewayId, ruleId, and subjectId from this request.
	//
	// example:
	//
	// finops-export-abcdef1234567890
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetStatus(v string) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetTaskId(v string) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
