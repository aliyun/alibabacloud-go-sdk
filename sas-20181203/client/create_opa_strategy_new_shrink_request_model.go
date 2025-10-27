// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpaStrategyNewShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmDetailShrink(v string) *CreateOpaStrategyNewShrinkRequest
	GetAlarmDetailShrink() *string
	SetClusterId(v string) *CreateOpaStrategyNewShrinkRequest
	GetClusterId() *string
	SetClusterName(v string) *CreateOpaStrategyNewShrinkRequest
	GetClusterName() *string
	SetDescription(v string) *CreateOpaStrategyNewShrinkRequest
	GetDescription() *string
	SetImageName(v []*string) *CreateOpaStrategyNewShrinkRequest
	GetImageName() []*string
	SetLabel(v []*string) *CreateOpaStrategyNewShrinkRequest
	GetLabel() []*string
	SetMaliciousImage(v bool) *CreateOpaStrategyNewShrinkRequest
	GetMaliciousImage() *bool
	SetRuleAction(v int32) *CreateOpaStrategyNewShrinkRequest
	GetRuleAction() *int32
	SetScopes(v []*CreateOpaStrategyNewShrinkRequestScopes) *CreateOpaStrategyNewShrinkRequest
	GetScopes() []*CreateOpaStrategyNewShrinkRequestScopes
	SetStrategyId(v int64) *CreateOpaStrategyNewShrinkRequest
	GetStrategyId() *int64
	SetStrategyName(v string) *CreateOpaStrategyNewShrinkRequest
	GetStrategyName() *string
	SetStrategyTemplateId(v int64) *CreateOpaStrategyNewShrinkRequest
	GetStrategyTemplateId() *int64
	SetUnScanedImage(v bool) *CreateOpaStrategyNewShrinkRequest
	GetUnScanedImage() *bool
	SetWhiteList(v []*string) *CreateOpaStrategyNewShrinkRequest
	GetWhiteList() []*string
}

type CreateOpaStrategyNewShrinkRequest struct {
	// The risks that you want to detect by using the rule.
	AlarmDetailShrink *string `json:"AlarmDetail,omitempty" xml:"AlarmDetail,omitempty"`
	// The cluster ID.
	//
	// > This parameter is deprecated.
	//
	// example:
	//
	// cfa7e2fb8c221483ba59e098c34c6****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// > This parameter is deprecated.
	//
	// example:
	//
	// *
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The rule description.
	//
	// example:
	//
	// default policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image names.
	ImageName []*string `json:"ImageName,omitempty" xml:"ImageName,omitempty" type:"Repeated"`
	// The container tags.
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
	// 	- **1**: trigger alerts
	//
	// 	- **2**: block
	//
	// 	- **3**: allow
	//
	// example:
	//
	// 1
	RuleAction *int32 `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The application scope of the rule.
	Scopes []*CreateOpaStrategyNewShrinkRequestScopes `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// >  You can call the [ListOpaClusterStrategyNew](https://help.aliyun.com/document_detail/2623574.html) operation to query the rule ID.
	//
	// > This parameter is invalid when you create a rule.
	//
	// example:
	//
	// 16
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// default
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
	// The whitelist.
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s CreateOpaStrategyNewShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewShrinkRequest) GetAlarmDetailShrink() *string {
	return s.AlarmDetailShrink
}

func (s *CreateOpaStrategyNewShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateOpaStrategyNewShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateOpaStrategyNewShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateOpaStrategyNewShrinkRequest) GetImageName() []*string {
	return s.ImageName
}

func (s *CreateOpaStrategyNewShrinkRequest) GetLabel() []*string {
	return s.Label
}

func (s *CreateOpaStrategyNewShrinkRequest) GetMaliciousImage() *bool {
	return s.MaliciousImage
}

func (s *CreateOpaStrategyNewShrinkRequest) GetRuleAction() *int32 {
	return s.RuleAction
}

func (s *CreateOpaStrategyNewShrinkRequest) GetScopes() []*CreateOpaStrategyNewShrinkRequestScopes {
	return s.Scopes
}

func (s *CreateOpaStrategyNewShrinkRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *CreateOpaStrategyNewShrinkRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *CreateOpaStrategyNewShrinkRequest) GetStrategyTemplateId() *int64 {
	return s.StrategyTemplateId
}

func (s *CreateOpaStrategyNewShrinkRequest) GetUnScanedImage() *bool {
	return s.UnScanedImage
}

func (s *CreateOpaStrategyNewShrinkRequest) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *CreateOpaStrategyNewShrinkRequest) SetAlarmDetailShrink(v string) *CreateOpaStrategyNewShrinkRequest {
	s.AlarmDetailShrink = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetClusterId(v string) *CreateOpaStrategyNewShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetClusterName(v string) *CreateOpaStrategyNewShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetDescription(v string) *CreateOpaStrategyNewShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetImageName(v []*string) *CreateOpaStrategyNewShrinkRequest {
	s.ImageName = v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetLabel(v []*string) *CreateOpaStrategyNewShrinkRequest {
	s.Label = v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetMaliciousImage(v bool) *CreateOpaStrategyNewShrinkRequest {
	s.MaliciousImage = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetRuleAction(v int32) *CreateOpaStrategyNewShrinkRequest {
	s.RuleAction = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetScopes(v []*CreateOpaStrategyNewShrinkRequestScopes) *CreateOpaStrategyNewShrinkRequest {
	s.Scopes = v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetStrategyId(v int64) *CreateOpaStrategyNewShrinkRequest {
	s.StrategyId = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetStrategyName(v string) *CreateOpaStrategyNewShrinkRequest {
	s.StrategyName = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetStrategyTemplateId(v int64) *CreateOpaStrategyNewShrinkRequest {
	s.StrategyTemplateId = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetUnScanedImage(v bool) *CreateOpaStrategyNewShrinkRequest {
	s.UnScanedImage = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) SetWhiteList(v []*string) *CreateOpaStrategyNewShrinkRequest {
	s.WhiteList = v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequest) Validate() error {
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

type CreateOpaStrategyNewShrinkRequestScopes struct {
	// The ID of the cluster node to which the rule is applied.
	//
	// > This parameter is not required when you create the instance.
	//
	// example:
	//
	// ack-p-1
	AckPolicyInstanceId *string `json:"AckPolicyInstanceId,omitempty" xml:"AckPolicyInstanceId,omitempty"`
	// Specifies whether to include all namespaces. Valid values:
	//
	// 	- **1**: includes all namespaces.
	//
	// 	- **0**: does not include all namespaces.
	//
	// example:
	//
	// 1
	AllNamespace *int32 `json:"AllNamespace,omitempty" xml:"AllNamespace,omitempty"`
	// The ID of the cluster that is specified in the rule.
	//
	// >  You can call the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/421736.html) operation to query the cluster ID.
	//
	// example:
	//
	// cc50d***015d2
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespaces.
	//
	// > This parameter is valid only when the AllNamespace parameter is set to 0.
	NamespaceList []*string `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
}

func (s CreateOpaStrategyNewShrinkRequestScopes) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewShrinkRequestScopes) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewShrinkRequestScopes) GetAckPolicyInstanceId() *string {
	return s.AckPolicyInstanceId
}

func (s *CreateOpaStrategyNewShrinkRequestScopes) GetAllNamespace() *int32 {
	return s.AllNamespace
}

func (s *CreateOpaStrategyNewShrinkRequestScopes) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateOpaStrategyNewShrinkRequestScopes) GetNamespaceList() []*string {
	return s.NamespaceList
}

func (s *CreateOpaStrategyNewShrinkRequestScopes) SetAckPolicyInstanceId(v string) *CreateOpaStrategyNewShrinkRequestScopes {
	s.AckPolicyInstanceId = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequestScopes) SetAllNamespace(v int32) *CreateOpaStrategyNewShrinkRequestScopes {
	s.AllNamespace = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequestScopes) SetClusterId(v string) *CreateOpaStrategyNewShrinkRequestScopes {
	s.ClusterId = &v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequestScopes) SetNamespaceList(v []*string) *CreateOpaStrategyNewShrinkRequestScopes {
	s.NamespaceList = v
	return s
}

func (s *CreateOpaStrategyNewShrinkRequestScopes) Validate() error {
	return dara.Validate(s)
}
