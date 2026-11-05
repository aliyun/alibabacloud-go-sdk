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
	// - **repoName**: The name of the repository.
	//
	// - **repoNamespace**: The namespace of the repository.
	//
	// example:
	//
	// repoName
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The name of the repository or the namespace of the repository.
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
	// The parameters for configuring proactive defense for servers, including the following parameters:
	//
	// - **targetType**: The dimension of the defense configuration. Currently, only the UUID dimension is supported. Set the value to **uuid**.
	//
	// - **target**: The UUID of the server for which you want to configure the defense settings.
	//
	// - **flag**: Specifies whether to enable or disable proactive defense for the server. Valid values: **add*	- (enable) and **del*	- (disable).
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
	// The target Asset Type. Valid values:
	//
	// - **uuid**: The UUID of the server.
	//
	// - **Cluster**: The ID of the cluster.
	//
	// - **image_repo**: The name of the image repository.
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The switch type. Valid values:
	//
	// - **alidetect-scan-enable**: Local file detection is performed only locally.
	//
	// - **ACTION-TRIAL-PERMISSION**: ActionTrail data delivery.
	//
	// - **alidetect**: Local file detection engine.
	//
	// - **container_prevent_escape**: Container escape prevention.
	//
	// - **image_repo**: Repository image scanning.
	//
	// - **proc_filter_switch**: Log filtering.
	//
	// - **agentless**: Agentless detection.
	//
	// - **rasp**: Application defense.
	//
	// - **sensitiveFile**: Sensitive information scanning.
	//
	// - **aliscriptengine**: Deep detection engine.
	//
	// - **containerNetwork**: Container visualization.
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
