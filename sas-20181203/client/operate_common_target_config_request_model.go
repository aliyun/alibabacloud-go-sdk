// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateCommonTargetConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFieldName(v string) *OperateCommonTargetConfigRequest
	GetFieldName() *string
	SetFieldValue(v string) *OperateCommonTargetConfigRequest
	GetFieldValue() *string
	SetSourceIp(v string) *OperateCommonTargetConfigRequest
	GetSourceIp() *string
	SetTargetOperations(v string) *OperateCommonTargetConfigRequest
	GetTargetOperations() *string
	SetTargetType(v string) *OperateCommonTargetConfigRequest
	GetTargetType() *string
	SetType(v string) *OperateCommonTargetConfigRequest
	GetType() *string
}

type OperateCommonTargetConfigRequest struct {
	// The target type of the image switch. Valid values:
	//
	// - **repoName**: repository name.
	//
	// - **repoNamespace**: repository namespace name.
	//
	// example:
	//
	// repoName
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The repository name or repository namespace name.
	//
	// example:
	//
	// cafcmc-dev
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 182.92.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The parameters for configuring proactive defense on servers. The following parameters are included:
	//
	// - **targetType**: the dimension of the defense configuration. Currently, only the UUID dimension is supported. Fixed value: **uuid**.
	//
	// - **target**: the UUID of the server for which you want to configure proactive defense.
	//
	// - **flag**: specifies whether to enable or disable proactive defense for the server. Valid values: **add*	- (enable) and **del*	- (disable).
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "targetType": "uuid",
	//
	//             "target": "f329a044-6a2f-49a0-9d33-352f6c1d****",
	//
	//             "flag": "del"
	//
	//       }
	//
	// ]
	TargetOperations *string `json:"TargetOperations,omitempty" xml:"TargetOperations,omitempty"`
	// The Asset Type of the target. Valid values:
	//
	// - **uuid**: server UUID.
	//
	// - **Cluster**: cluster ID.
	//
	// - **image_repo**: image repository name.
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The switch type. Valid values:
	//
	// - **alidetect-scan-enable**: local file detection that performs detection only locally.
	//
	// - **ACTION-TRIAL-PERMISSION**: ActionTrail data delivery.
	//
	// - **alidetect**: local file detection engine.
	//
	// - **container_prevent_escape**: container escape prevention.
	//
	// - **image_repo**: repository image scanning.
	//
	// - **proc_filter_switch**: log filtering.
	//
	// - **agentless**: agentless detection.
	//
	// - **rasp**: application protection.
	//
	// - **sensitiveFile**: sensitive information scanning.
	//
	// - **aliscriptengine**: deep detection engine.
	//
	// - **containerNetwork**: container visualization.
	//
	// This parameter is required.
	//
	// example:
	//
	// alidetect
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OperateCommonTargetConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateCommonTargetConfigRequest) GoString() string {
	return s.String()
}

func (s *OperateCommonTargetConfigRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *OperateCommonTargetConfigRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *OperateCommonTargetConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *OperateCommonTargetConfigRequest) GetTargetOperations() *string {
	return s.TargetOperations
}

func (s *OperateCommonTargetConfigRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *OperateCommonTargetConfigRequest) GetType() *string {
	return s.Type
}

func (s *OperateCommonTargetConfigRequest) SetFieldName(v string) *OperateCommonTargetConfigRequest {
	s.FieldName = &v
	return s
}

func (s *OperateCommonTargetConfigRequest) SetFieldValue(v string) *OperateCommonTargetConfigRequest {
	s.FieldValue = &v
	return s
}

func (s *OperateCommonTargetConfigRequest) SetSourceIp(v string) *OperateCommonTargetConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *OperateCommonTargetConfigRequest) SetTargetOperations(v string) *OperateCommonTargetConfigRequest {
	s.TargetOperations = &v
	return s
}

func (s *OperateCommonTargetConfigRequest) SetTargetType(v string) *OperateCommonTargetConfigRequest {
	s.TargetType = &v
	return s
}

func (s *OperateCommonTargetConfigRequest) SetType(v string) *OperateCommonTargetConfigRequest {
	s.Type = &v
	return s
}

func (s *OperateCommonTargetConfigRequest) Validate() error {
	return dara.Validate(s)
}
