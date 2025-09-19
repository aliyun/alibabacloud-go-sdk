// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAntiBruteForceRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewRuleId(v int64) *ModifyInstanceAntiBruteForceRuleRequest
	GetNewRuleId() *int64
	SetResourceOwnerId(v int64) *ModifyInstanceAntiBruteForceRuleRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *ModifyInstanceAntiBruteForceRuleRequest
	GetSourceIp() *string
	SetUuid(v string) *ModifyInstanceAntiBruteForceRuleRequest
	GetUuid() *string
}

type ModifyInstanceAntiBruteForceRuleRequest struct {
	// The ID of the defense rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65767
	NewRuleId       *int64 `json:"NewRuleId,omitempty" xml:"NewRuleId,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the server for which you want to modify the defense rule. You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to query the UUIDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7cc91747-2845-40d4-bb69-c077597f****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyInstanceAntiBruteForceRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAntiBruteForceRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) GetNewRuleId() *int64 {
	return s.NewRuleId
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) SetNewRuleId(v int64) *ModifyInstanceAntiBruteForceRuleRequest {
	s.NewRuleId = &v
	return s
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) SetResourceOwnerId(v int64) *ModifyInstanceAntiBruteForceRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) SetSourceIp(v string) *ModifyInstanceAntiBruteForceRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) SetUuid(v string) *ModifyInstanceAntiBruteForceRuleRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) Validate() error {
	return dara.Validate(s)
}
