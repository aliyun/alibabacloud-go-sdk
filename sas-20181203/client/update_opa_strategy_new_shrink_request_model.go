// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpaStrategyNewShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmDetailShrink(v string) *UpdateOpaStrategyNewShrinkRequest
	GetAlarmDetailShrink() *string
	SetClusterId(v string) *UpdateOpaStrategyNewShrinkRequest
	GetClusterId() *string
	SetClusterName(v string) *UpdateOpaStrategyNewShrinkRequest
	GetClusterName() *string
	SetDescription(v string) *UpdateOpaStrategyNewShrinkRequest
	GetDescription() *string
	SetImageName(v []*string) *UpdateOpaStrategyNewShrinkRequest
	GetImageName() []*string
	SetLabel(v []*string) *UpdateOpaStrategyNewShrinkRequest
	GetLabel() []*string
	SetMaliciousImage(v bool) *UpdateOpaStrategyNewShrinkRequest
	GetMaliciousImage() *bool
	SetRuleAction(v int32) *UpdateOpaStrategyNewShrinkRequest
	GetRuleAction() *int32
	SetScopes(v []*UpdateOpaStrategyNewShrinkRequestScopes) *UpdateOpaStrategyNewShrinkRequest
	GetScopes() []*UpdateOpaStrategyNewShrinkRequestScopes
	SetStrategyId(v int64) *UpdateOpaStrategyNewShrinkRequest
	GetStrategyId() *int64
	SetStrategyName(v string) *UpdateOpaStrategyNewShrinkRequest
	GetStrategyName() *string
	SetStrategyTemplateId(v int64) *UpdateOpaStrategyNewShrinkRequest
	GetStrategyTemplateId() *int64
	SetUnScanedImage(v bool) *UpdateOpaStrategyNewShrinkRequest
	GetUnScanedImage() *bool
	SetWhiteList(v []*string) *UpdateOpaStrategyNewShrinkRequest
	GetWhiteList() []*string
}

type UpdateOpaStrategyNewShrinkRequest struct {
	// The risks that you want to detect by using the rule.
	AlarmDetailShrink *string `json:"AlarmDetail,omitempty" xml:"AlarmDetail,omitempty"`
	// The cluster ID.
	//
	// > This parameter is deprecated. You can use the Scopes parameter to specify a scope in which cluster parameters take effect.
	//
	// example:
	//
	// c870ec78ecbcb41d2a35c679823ef****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// > This parameter is deprecated.
	//
	// example:
	//
	// docker-law
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The rule description.
	//
	// example:
	//
	// 4566
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image names.
	ImageName []*string `json:"ImageName,omitempty" xml:"ImageName,omitempty" type:"Repeated"`
	// The image tags.
	Label []*string `json:"Label,omitempty" xml:"Label,omitempty" type:"Repeated"`
	// Specifies whether the rule supports malicious Internet images. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	MaliciousImage *bool `json:"MaliciousImage,omitempty" xml:"MaliciousImage,omitempty"`
	// The action that is performed when the rule is hit. Valid values:
	//
	// 	- **1**: alert
	//
	// 	- **2**: block
	//
	// 	- **3**: allow
	//
	// example:
	//
	// 1
	RuleAction *int32 `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The application scope.
	Scopes []*UpdateOpaStrategyNewShrinkRequestScopes `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
	// The ID of the rule.
	//
	// >  You can call the [ListOpaClusterStrategyNew](https://help.aliyun.com/document_detail/2623574.html) operation to query the ID.
	//
	// example:
	//
	// 1003
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// test
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The ID of the rule template.
	//
	// >  You can call the [GetOpaStrategyTemplateSummary](https://help.aliyun.com/document_detail/2539952.html) operation to query the ID of the rule template.
	//
	// example:
	//
	// 109
	StrategyTemplateId *int64 `json:"StrategyTemplateId,omitempty" xml:"StrategyTemplateId,omitempty"`
	// Specifies whether the rule supports unscanned images. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	UnScanedImage *bool `json:"UnScanedImage,omitempty" xml:"UnScanedImage,omitempty"`
	// The whitelists.
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s UpdateOpaStrategyNewShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetAlarmDetailShrink() *string {
	return s.AlarmDetailShrink
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetImageName() []*string {
	return s.ImageName
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetLabel() []*string {
	return s.Label
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetMaliciousImage() *bool {
	return s.MaliciousImage
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetRuleAction() *int32 {
	return s.RuleAction
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetScopes() []*UpdateOpaStrategyNewShrinkRequestScopes {
	return s.Scopes
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetStrategyTemplateId() *int64 {
	return s.StrategyTemplateId
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetUnScanedImage() *bool {
	return s.UnScanedImage
}

func (s *UpdateOpaStrategyNewShrinkRequest) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetAlarmDetailShrink(v string) *UpdateOpaStrategyNewShrinkRequest {
	s.AlarmDetailShrink = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetClusterId(v string) *UpdateOpaStrategyNewShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetClusterName(v string) *UpdateOpaStrategyNewShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetDescription(v string) *UpdateOpaStrategyNewShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetImageName(v []*string) *UpdateOpaStrategyNewShrinkRequest {
	s.ImageName = v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetLabel(v []*string) *UpdateOpaStrategyNewShrinkRequest {
	s.Label = v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetMaliciousImage(v bool) *UpdateOpaStrategyNewShrinkRequest {
	s.MaliciousImage = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetRuleAction(v int32) *UpdateOpaStrategyNewShrinkRequest {
	s.RuleAction = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetScopes(v []*UpdateOpaStrategyNewShrinkRequestScopes) *UpdateOpaStrategyNewShrinkRequest {
	s.Scopes = v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetStrategyId(v int64) *UpdateOpaStrategyNewShrinkRequest {
	s.StrategyId = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetStrategyName(v string) *UpdateOpaStrategyNewShrinkRequest {
	s.StrategyName = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetStrategyTemplateId(v int64) *UpdateOpaStrategyNewShrinkRequest {
	s.StrategyTemplateId = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetUnScanedImage(v bool) *UpdateOpaStrategyNewShrinkRequest {
	s.UnScanedImage = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) SetWhiteList(v []*string) *UpdateOpaStrategyNewShrinkRequest {
	s.WhiteList = v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequest) Validate() error {
	if s.Scopes != nil {
		for _, item := range s.Scopes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateOpaStrategyNewShrinkRequestScopes struct {
	// The ID of the cluster node to which the rule is applied.
	//
	// >  You can call the [GetOpaStrategyDetailNew](~~GetOpaStrategyDetailNew~~) operation to query the ID of the cluster node to which the rule is applied.
	//
	// example:
	//
	// ack-1
	AckPolicyInstanceId *string `json:"AckPolicyInstanceId,omitempty" xml:"AckPolicyInstanceId,omitempty"`
	// Specifies whether all namespaces are included. Valid values:
	//
	// 	- **0**: Not all namespaces are included.
	//
	// 	- **1**: All namespaces are included.
	//
	// example:
	//
	// 1
	AllNamespace *int32 `json:"AllNamespace,omitempty" xml:"AllNamespace,omitempty"`
	// The cluster ID.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the cluster ID.
	//
	// example:
	//
	// cdcb56a931c**
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespaces.
	//
	// > This parameter is valid only when the AllNamespace parameter is set to 0.
	NamespaceList []*string `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
}

func (s UpdateOpaStrategyNewShrinkRequestScopes) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewShrinkRequestScopes) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewShrinkRequestScopes) GetAckPolicyInstanceId() *string {
	return s.AckPolicyInstanceId
}

func (s *UpdateOpaStrategyNewShrinkRequestScopes) GetAllNamespace() *int32 {
	return s.AllNamespace
}

func (s *UpdateOpaStrategyNewShrinkRequestScopes) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateOpaStrategyNewShrinkRequestScopes) GetNamespaceList() []*string {
	return s.NamespaceList
}

func (s *UpdateOpaStrategyNewShrinkRequestScopes) SetAckPolicyInstanceId(v string) *UpdateOpaStrategyNewShrinkRequestScopes {
	s.AckPolicyInstanceId = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequestScopes) SetAllNamespace(v int32) *UpdateOpaStrategyNewShrinkRequestScopes {
	s.AllNamespace = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequestScopes) SetClusterId(v string) *UpdateOpaStrategyNewShrinkRequestScopes {
	s.ClusterId = &v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequestScopes) SetNamespaceList(v []*string) *UpdateOpaStrategyNewShrinkRequestScopes {
	s.NamespaceList = v
	return s
}

func (s *UpdateOpaStrategyNewShrinkRequestScopes) Validate() error {
	return dara.Validate(s)
}
