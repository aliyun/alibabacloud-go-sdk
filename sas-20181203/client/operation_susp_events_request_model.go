// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationSuspEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *OperationSuspEventsRequest
	GetFrom() *string
	SetOperation(v string) *OperationSuspEventsRequest
	GetOperation() *string
	SetSourceIp(v string) *OperationSuspEventsRequest
	GetSourceIp() *string
	SetSubOperation(v string) *OperationSuspEventsRequest
	GetSubOperation() *string
	SetSuspiciousEventIds(v string) *OperationSuspEventsRequest
	GetSuspiciousEventIds() *string
	SetWarnType(v string) *OperationSuspEventsRequest
	GetWarnType() *string
}

type OperationSuspEventsRequest struct {
	// The ID of the request source.
	//
	// Set the value to **sas**, which indicates that the request is sent from Security Center.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The operation that you want to perform on alerts. Valid values:
	//
	// 	- **deal**: quarantines the source file of the malicious process.
	//
	// 	- **ignore**: ignores the alerts.
	//
	// 	- **mark_mis_info**: marks the alerts as false positives by adding the alerts to the whitelist.
	//
	// 	- **rm_mark_mis_info**: cancels false positives by removing the alerts from the whitelist.
	//
	// 	- **offline_handled**: marks the alerts as handled.
	//
	// This parameter is required.
	//
	// example:
	//
	// deal
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The suboperation that you want to perform when you quarantine the source file of the malicious process. Valid values:
	//
	// 	- **killAndQuaraFileByPidAndMd5andPath**: terminates the process based on its process ID (PID) and quarantines the source file of the process.
	//
	// 	- **quaraFileByMd5andPath**: quarantines the source file of the process.
	//
	// 	- **killAndQuaraFileByMd5andPath**: terminates the process and quarantines the source file of the process.
	//
	// example:
	//
	// killAndQuaraFileByPidAndMd5andPath
	SubOperation *string `json:"SubOperation,omitempty" xml:"SubOperation,omitempty"`
	// The IDs of alert events.
	//
	// > You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to obtain the IDs of alert events from the SecurityEventIds response parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 290852
	SuspiciousEventIds *string `json:"SuspiciousEventIds,omitempty" xml:"SuspiciousEventIds,omitempty"`
	// The type of the exceptions. Valid values:
	//
	// 	- **alarm**: alerts
	//
	// 	- **null**: exceptions
	//
	// example:
	//
	// alarm
	WarnType *string `json:"WarnType,omitempty" xml:"WarnType,omitempty"`
}

func (s OperationSuspEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s OperationSuspEventsRequest) GoString() string {
	return s.String()
}

func (s *OperationSuspEventsRequest) GetFrom() *string {
	return s.From
}

func (s *OperationSuspEventsRequest) GetOperation() *string {
	return s.Operation
}

func (s *OperationSuspEventsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *OperationSuspEventsRequest) GetSubOperation() *string {
	return s.SubOperation
}

func (s *OperationSuspEventsRequest) GetSuspiciousEventIds() *string {
	return s.SuspiciousEventIds
}

func (s *OperationSuspEventsRequest) GetWarnType() *string {
	return s.WarnType
}

func (s *OperationSuspEventsRequest) SetFrom(v string) *OperationSuspEventsRequest {
	s.From = &v
	return s
}

func (s *OperationSuspEventsRequest) SetOperation(v string) *OperationSuspEventsRequest {
	s.Operation = &v
	return s
}

func (s *OperationSuspEventsRequest) SetSourceIp(v string) *OperationSuspEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *OperationSuspEventsRequest) SetSubOperation(v string) *OperationSuspEventsRequest {
	s.SubOperation = &v
	return s
}

func (s *OperationSuspEventsRequest) SetSuspiciousEventIds(v string) *OperationSuspEventsRequest {
	s.SuspiciousEventIds = &v
	return s
}

func (s *OperationSuspEventsRequest) SetWarnType(v string) *OperationSuspEventsRequest {
	s.WarnType = &v
	return s
}

func (s *OperationSuspEventsRequest) Validate() error {
	return dara.Validate(s)
}
