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
	// The rule for adding items to the whitelist. For example, to add a whitelist rule based on file MD5 where the file contains the string "a", set this parameter to {"field":"md5","operate":"contains","fieldValue":"aa"}.
	//
	// example:
	//
	// {"field":"md5","operate":"contains","fieldValue":"aa"}
	MarkMissParam *string `json:"MarkMissParam,omitempty" xml:"MarkMissParam,omitempty"`
	// The type of operation for batch processing alert events of the same type.
	//
	// >Call the [DescribeSecurityEventOperations](~~DescribeSecurityEventOperations~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// offline_handled
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	// The configuration of the sub-operation for handling alerting events. The value is in JSON format.
	//
	// > This parameter is required when **OperationCode*	- is set to **kill_and_quara**, **block_ip**, or **virus_quara**. For other values of **OperationCode**, this parameter can be left empty.
	//
	// > When **OperationCode*	- is set to **block_ip**, the following field is included:
	//
	// > - **expireTime**: the lock expiration time. Unit: milliseconds.
	//
	// >
	//
	// > When **OperationCode*	- is set to **kill_and_quara**, the following field is included:
	//
	// > - **subOperation**: the method used to scan and remove threats. Valid values:
	//
	// >     - **killAndQuaraFileByMd5andPath**: terminates the process and moves the file to the quarantined file.
	//
	// >     - **killByMd5andPath**: terminates the running process.
	//
	// >
	//
	// > When **OperationCode*	- is set to **virus_quara**, the following field is included:
	//
	// > - **subOperation**: the method used to scan and remove threats. Valid values:
	//
	// >    - **quaraFileByMd5andPath**: moves the source file of the process to the quarantined file.
	//
	// example:
	//
	// {"expireTime":1646208726195}
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	// The remarks for the operation.
	//
	// example:
	//
	// remark test.
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the task that batch processes all alert events of the same type.
	//
	// >Call the [CreateSimilarSecurityEventsQueryTask](~~CreateSimilarSecurityEventsQueryTask~~) operation to obtain this parameter.
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
