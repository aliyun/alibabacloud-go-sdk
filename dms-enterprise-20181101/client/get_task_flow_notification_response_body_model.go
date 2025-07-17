// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskFlowNotificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetTaskFlowNotificationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTaskFlowNotificationResponseBody
	GetErrorMessage() *string
	SetNotification(v *GetTaskFlowNotificationResponseBodyNotification) *GetTaskFlowNotificationResponseBody
	GetNotification() *GetTaskFlowNotificationResponseBodyNotification
	SetRequestId(v string) *GetTaskFlowNotificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskFlowNotificationResponseBody
	GetSuccess() *bool
}

type GetTaskFlowNotificationResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The notification settings specified by the user.
	Notification *GetTaskFlowNotificationResponseBodyNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// F19D575F-EBEA-5683-AFA3-A8F6D9A7DE03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTaskFlowNotificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskFlowNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskFlowNotificationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskFlowNotificationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTaskFlowNotificationResponseBody) GetNotification() *GetTaskFlowNotificationResponseBodyNotification {
	return s.Notification
}

func (s *GetTaskFlowNotificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskFlowNotificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskFlowNotificationResponseBody) SetErrorCode(v string) *GetTaskFlowNotificationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskFlowNotificationResponseBody) SetErrorMessage(v string) *GetTaskFlowNotificationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTaskFlowNotificationResponseBody) SetNotification(v *GetTaskFlowNotificationResponseBodyNotification) *GetTaskFlowNotificationResponseBody {
	s.Notification = v
	return s
}

func (s *GetTaskFlowNotificationResponseBody) SetRequestId(v string) *GetTaskFlowNotificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskFlowNotificationResponseBody) SetSuccess(v bool) *GetTaskFlowNotificationResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskFlowNotificationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTaskFlowNotificationResponseBodyNotification struct {
	// Indicates whether notifications for failed task flows are enabled. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disabled
	//
	// example:
	//
	// true
	DagNotificationFail *bool `json:"DagNotificationFail,omitempty" xml:"DagNotificationFail,omitempty"`
	// Indicates whether service level agreement (SLA) global notifications for task flows are enabled. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disabled
	//
	// example:
	//
	// true
	DagNotificationSla *bool `json:"DagNotificationSla,omitempty" xml:"DagNotificationSla,omitempty"`
	// Indicates whether notifications for successful task flows are enabled. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disabled
	//
	// example:
	//
	// true
	DagNotificationSuccess *bool `json:"DagNotificationSuccess,omitempty" xml:"DagNotificationSuccess,omitempty"`
}

func (s GetTaskFlowNotificationResponseBodyNotification) String() string {
	return dara.Prettify(s)
}

func (s GetTaskFlowNotificationResponseBodyNotification) GoString() string {
	return s.String()
}

func (s *GetTaskFlowNotificationResponseBodyNotification) GetDagNotificationFail() *bool {
	return s.DagNotificationFail
}

func (s *GetTaskFlowNotificationResponseBodyNotification) GetDagNotificationSla() *bool {
	return s.DagNotificationSla
}

func (s *GetTaskFlowNotificationResponseBodyNotification) GetDagNotificationSuccess() *bool {
	return s.DagNotificationSuccess
}

func (s *GetTaskFlowNotificationResponseBodyNotification) SetDagNotificationFail(v bool) *GetTaskFlowNotificationResponseBodyNotification {
	s.DagNotificationFail = &v
	return s
}

func (s *GetTaskFlowNotificationResponseBodyNotification) SetDagNotificationSla(v bool) *GetTaskFlowNotificationResponseBodyNotification {
	s.DagNotificationSla = &v
	return s
}

func (s *GetTaskFlowNotificationResponseBodyNotification) SetDagNotificationSuccess(v bool) *GetTaskFlowNotificationResponseBodyNotification {
	s.DagNotificationSuccess = &v
	return s
}

func (s *GetTaskFlowNotificationResponseBodyNotification) Validate() error {
	return dara.Validate(s)
}
