// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleSimilarSecurityEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarkMissParam(v string) *HandleSimilarSecurityEventsRequest
	GetMarkMissParam() *string
	SetOperationCode(v string) *HandleSimilarSecurityEventsRequest
	GetOperationCode() *string
	SetOperationParams(v string) *HandleSimilarSecurityEventsRequest
	GetOperationParams() *string
	SetRemark(v string) *HandleSimilarSecurityEventsRequest
	GetRemark() *string
	SetResourceOwnerId(v int64) *HandleSimilarSecurityEventsRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *HandleSimilarSecurityEventsRequest
	GetSourceIp() *string
	SetTaskId(v int64) *HandleSimilarSecurityEventsRequest
	GetTaskId() *int64
}

type HandleSimilarSecurityEventsRequest struct {
	// The whitelist rule. For example, if you want to add a file that contains the string a to the whitelist based on the MD5 hash value, set this parameter to {"field":"md5","operate":"contains","fieldValue":"aa"}.
	//
	// example:
	//
	// {"field":"md5","operate":"contains","fieldValue":"aa"}
	MarkMissParam *string `json:"MarkMissParam,omitempty" xml:"MarkMissParam,omitempty"`
	// The operation that you want to perform to handle the alert events.
	//
	// >  You can call the [DescribeSecurityEventOperations](~~DescribeSecurityEventOperations~~) operation to query the operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// offline_handled
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The configuration of the operation that you want to perform to handle the alert events. The value of this parameter is in the JSON format.
	//
	// >  If you set **OperationCode*	- to **kill_and_quara**, **block_ip**, or **virus_quara**, you must specify OperationParams. If you set **OperationCode*	- to other values, you can leave OperationParams empty. If you set **OperationCode*	- to **block_ip**, the value of OperationParams must consist of the following fields:
	//
	// > 	- **expireTime**: the end time of locking. Unit: milliseconds.
	//
	// >  If you set **OperationCode*	- to **kill_and_quara**, the value of OperationParams must consist of the following fields:
	//
	// > 	- **subOperation**: the method of detection and removal. Valid values:
	//
	// >     	- **killAndQuaraFileByMd5andPath**: terminates the process and quarantines the source file of the process.
	//
	// >     	- **killByMd5andPath**: terminates the running process.
	//
	// >  If you set **OperationCode*	- to **virus_quara**, the value of OperationParams consists of the following fields:
	//
	// > 	- **subOperation**: the method of detection and removal. Valid values:
	//
	// >     	- **quaraFileByMd5andPath**: quarantines the source file of the process.
	//
	// example:
	//
	// {"expireTime":1646208726195}
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	// The remark of the operation.
	//
	// example:
	//
	// remark test.
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the task that handles the alert events at a time.
	//
	// >  You can call the [CreateSimilarSecurityEventsQueryTask](~~CreateSimilarSecurityEventsQueryTask~~) operation to query the IDs of tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 666038
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s HandleSimilarSecurityEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s HandleSimilarSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *HandleSimilarSecurityEventsRequest) GetMarkMissParam() *string {
	return s.MarkMissParam
}

func (s *HandleSimilarSecurityEventsRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *HandleSimilarSecurityEventsRequest) GetOperationParams() *string {
	return s.OperationParams
}

func (s *HandleSimilarSecurityEventsRequest) GetRemark() *string {
	return s.Remark
}

func (s *HandleSimilarSecurityEventsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *HandleSimilarSecurityEventsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *HandleSimilarSecurityEventsRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *HandleSimilarSecurityEventsRequest) SetMarkMissParam(v string) *HandleSimilarSecurityEventsRequest {
	s.MarkMissParam = &v
	return s
}

func (s *HandleSimilarSecurityEventsRequest) SetOperationCode(v string) *HandleSimilarSecurityEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *HandleSimilarSecurityEventsRequest) SetOperationParams(v string) *HandleSimilarSecurityEventsRequest {
	s.OperationParams = &v
	return s
}

func (s *HandleSimilarSecurityEventsRequest) SetRemark(v string) *HandleSimilarSecurityEventsRequest {
	s.Remark = &v
	return s
}

func (s *HandleSimilarSecurityEventsRequest) SetResourceOwnerId(v int64) *HandleSimilarSecurityEventsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *HandleSimilarSecurityEventsRequest) SetSourceIp(v string) *HandleSimilarSecurityEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *HandleSimilarSecurityEventsRequest) SetTaskId(v int64) *HandleSimilarSecurityEventsRequest {
	s.TaskId = &v
	return s
}

func (s *HandleSimilarSecurityEventsRequest) Validate() error {
	return dara.Validate(s)
}
