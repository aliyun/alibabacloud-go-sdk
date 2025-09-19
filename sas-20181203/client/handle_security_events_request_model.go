// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleSecurityEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarkBatch(v string) *HandleSecurityEventsRequest
	GetMarkBatch() *string
	SetMarkMissParam(v string) *HandleSecurityEventsRequest
	GetMarkMissParam() *string
	SetOperationCode(v string) *HandleSecurityEventsRequest
	GetOperationCode() *string
	SetOperationParams(v string) *HandleSecurityEventsRequest
	GetOperationParams() *string
	SetRemark(v string) *HandleSecurityEventsRequest
	GetRemark() *string
	SetResourceDirectoryAccountId(v int64) *HandleSecurityEventsRequest
	GetResourceDirectoryAccountId() *int64
	SetSecurityEventIds(v []*string) *HandleSecurityEventsRequest
	GetSecurityEventIds() []*string
	SetSourceIp(v string) *HandleSecurityEventsRequest
	GetSourceIp() *string
}

type HandleSecurityEventsRequest struct {
	// Specifies whether to add multiple alert events to the whitelist at a time. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	MarkBatch *string `json:"MarkBatch,omitempty" xml:"MarkBatch,omitempty"`
	// The whitelist rule. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **field**: The field based on which alert events are added to the whitelist.
	//
	// 	- **operate**: The method that is used to added alert events to the whitelist. Valid values:
	//
	//     	- **notContains**: does not contain
	//
	//     	- **contains**: contains
	//
	//     	- **regex**: matches by regular expression
	//
	//     	- **strEqual**: equals
	//
	//     	- **strNotEqual**: does not equal
	//
	// 	- **fieldValue**: The value of the field based on which alert events are added to the whitelist.
	//
	// 	- **uuid**: The application scope of the whitelist rule. Valid values:
	//
	//     	- **part**: the current asset
	//
	//     	- **ALL**: all assets
	//
	// >  You can call the [DescribeSecurityEventOperations](~~DescribeSecurityEventOperations~~) operation to obtain the fields that you can specify for **field**.
	//
	// example:
	//
	// [{"uuid":"part","field":"gmtModified","operate":"contains","fieldValue":"asd"},{"uuid":"part","field":"loginUser","operate":"contains","fieldValue":"vff"}]
	MarkMissParam *string `json:"MarkMissParam,omitempty" xml:"MarkMissParam,omitempty"`
	// The operation that you want to perform to handle the alert events. Valid values:
	//
	// 	- **block_ip**: blocks the source IP address.
	//
	// 	- **advance_mark_mis_info**: adds the alert events to the whitelist.
	//
	// 	- **ignore**: ignores the alert events.
	//
	// 	- **manual_handled**: marks the alert events as manually handled.
	//
	// 	- **kill_process**: terminates the malicious process.
	//
	// 	- **cleanup**: performs in-depth virus detection and removal.
	//
	// 	- **kill_and_quara**: kills the malicious processes and quarantines the source file.
	//
	// 	- **disable_malicious_defense**: stops the container on which the alerting files or processes exist.
	//
	// 	- **client_problem_check**: performs troubleshooting.
	//
	// 	- **quara**: quarantines the source file of the malicious process.
	//
	// This parameter is required.
	//
	// example:
	//
	// block_ip
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The configuration of the operation that you want to perform to handle the alert events.
	//
	// >  If you set OperationCode to `kill_and_quara` or `block_ip`, you must specify OperationParams. If you set OperationCode to other values, you can leave OperationParams empty.
	//
	// example:
	//
	// {}
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	// The remarks of the handling operation.
	//
	// example:
	//
	// remark test.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 16670360956*****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The IDs of the alert events.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["909361"]
	SecurityEventIds []*string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s HandleSecurityEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s HandleSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsRequest) GetMarkBatch() *string {
	return s.MarkBatch
}

func (s *HandleSecurityEventsRequest) GetMarkMissParam() *string {
	return s.MarkMissParam
}

func (s *HandleSecurityEventsRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *HandleSecurityEventsRequest) GetOperationParams() *string {
	return s.OperationParams
}

func (s *HandleSecurityEventsRequest) GetRemark() *string {
	return s.Remark
}

func (s *HandleSecurityEventsRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *HandleSecurityEventsRequest) GetSecurityEventIds() []*string {
	return s.SecurityEventIds
}

func (s *HandleSecurityEventsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *HandleSecurityEventsRequest) SetMarkBatch(v string) *HandleSecurityEventsRequest {
	s.MarkBatch = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetMarkMissParam(v string) *HandleSecurityEventsRequest {
	s.MarkMissParam = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetOperationCode(v string) *HandleSecurityEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetOperationParams(v string) *HandleSecurityEventsRequest {
	s.OperationParams = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetRemark(v string) *HandleSecurityEventsRequest {
	s.Remark = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetResourceDirectoryAccountId(v int64) *HandleSecurityEventsRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetSecurityEventIds(v []*string) *HandleSecurityEventsRequest {
	s.SecurityEventIds = v
	return s
}

func (s *HandleSecurityEventsRequest) SetSourceIp(v string) *HandleSecurityEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *HandleSecurityEventsRequest) Validate() error {
	return dara.Validate(s)
}
