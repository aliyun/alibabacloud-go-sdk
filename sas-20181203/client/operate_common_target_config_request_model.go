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
	// The type of the image. Valid values:
	//
	// 	- **repoName**: the name of the image repository
	//
	// 	- **repoNamespace**: the namespace of the image repository
	//
	// example:
	//
	// repoName
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The name of the image repository or the namespace of the image repository.
	//
	// example:
	//
	// cafcmc-dev
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 182.92.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The configuration of proactive defense for your server. The value includes the following fields:
	//
	// 	- **targetType**: specifies the dimension from which you manage proactive defense. UUIDs are supported. Set the value to **uuid**.
	//
	// 	- **target**: specifies the UUID of the server for which you want to configure proactive defense.
	//
	// 	- **flag**: specifies whether to enable or disable proactive defense for your server. Valid values are **add*	- and **del**. The value add indicates that proactive defense will be enabled for your server. The value del indicates that proactive defense will be disabled for your server.
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
	// The dimension based on which the asset is selected. Valid values:
	//
	// 	- **uuid**: the UUID of the server
	//
	// 	- **Cluster**: the ID of the cluster
	//
	// 	- **image_repo**: the name of the image repository
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the feature. Valid values:
	//
	// 	- **alidetect-scan-enable**: local file detection
	//
	// 	- **ACTION-TRIAL-PERMISSION**: data delivery to ActionTrail
	//
	// 	- **alidetect**: local file detection engine
	//
	// 	- **container_prevent_escape**: container escape prevention
	//
	// 	- **image_repo**: repository image scan
	//
	// 	- **proc_filter_switch**: log filtering
	//
	// 	- **agentless**: agentless detection
	//
	// 	- **rasp**: application protection
	//
	// 	- **sensitiveFile**: sensitive file detection
	//
	// 	- **aliscriptengine**: in-depth detection engine
	//
	// 	- **containerNetwork**: container network visualization
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
