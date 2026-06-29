// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationSubmitStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOperationSubmitStatusResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetOperationSubmitStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetOperationSubmitStatusResponseBody
	GetMessage() *string
	SetOperationSubmitJob(v *GetOperationSubmitStatusResponseBodyOperationSubmitJob) *GetOperationSubmitStatusResponseBody
	GetOperationSubmitJob() *GetOperationSubmitStatusResponseBodyOperationSubmitJob
	SetRequestId(v string) *GetOperationSubmitStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOperationSubmitStatusResponseBody
	GetSuccess() *bool
}

type GetOperationSubmitStatusResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The submit status of the operation.
	OperationSubmitJob *GetOperationSubmitStatusResponseBodyOperationSubmitJob `json:"OperationSubmitJob,omitempty" xml:"OperationSubmitJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOperationSubmitStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOperationSubmitStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationSubmitStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOperationSubmitStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetOperationSubmitStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOperationSubmitStatusResponseBody) GetOperationSubmitJob() *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	return s.OperationSubmitJob
}

func (s *GetOperationSubmitStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOperationSubmitStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOperationSubmitStatusResponseBody) SetCode(v string) *GetOperationSubmitStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBody) SetHttpStatusCode(v int32) *GetOperationSubmitStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBody) SetMessage(v string) *GetOperationSubmitStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBody) SetOperationSubmitJob(v *GetOperationSubmitStatusResponseBodyOperationSubmitJob) *GetOperationSubmitStatusResponseBody {
	s.OperationSubmitJob = v
	return s
}

func (s *GetOperationSubmitStatusResponseBody) SetRequestId(v string) *GetOperationSubmitStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBody) SetSuccess(v bool) *GetOperationSubmitStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBody) Validate() error {
	if s.OperationSubmitJob != nil {
		if err := s.OperationSubmitJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOperationSubmitStatusResponseBodyOperationSubmitJob struct {
	// The business ID. If the operation is a data backfill, this value is the data backfill workflow ID.
	//
	// example:
	//
	// f_122_232342
	ExternalBizId *string `json:"ExternalBizId,omitempty" xml:"ExternalBizId,omitempty"`
	// The submit ID of the data backfill request.
	//
	// example:
	//
	// 123456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The operation type. Valid values:
	//
	// - SUPPLY_DATA: data backfill.
	//
	// example:
	//
	// SUPPLY_DATA
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The status.
	//
	// example:
	//
	// SUCCESS
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// The user ID of the submitter.
	//
	// example:
	//
	// 132344
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The progress.
	//
	// example:
	//
	// 80
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s GetOperationSubmitStatusResponseBodyOperationSubmitJob) String() string {
	return dara.Prettify(s)
}

func (s GetOperationSubmitStatusResponseBodyOperationSubmitJob) GoString() string {
	return s.String()
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) GetExternalBizId() *string {
	return s.ExternalBizId
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) GetJobId() *string {
	return s.JobId
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) GetOperation() *string {
	return s.Operation
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) GetOperator() *string {
	return s.Operator
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) GetProgress() *string {
	return s.Progress
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) SetExternalBizId(v string) *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	s.ExternalBizId = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) SetJobId(v string) *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	s.JobId = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) SetOperation(v string) *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	s.Operation = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) SetOperationStatus(v string) *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	s.OperationStatus = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) SetOperator(v string) *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	s.Operator = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) SetProgress(v string) *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	s.Progress = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) Validate() error {
	return dara.Validate(s)
}
