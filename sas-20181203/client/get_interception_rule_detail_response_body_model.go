// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterceptionRuleDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInterceptionRuleDetail(v *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) *GetInterceptionRuleDetailResponseBody
	GetInterceptionRuleDetail() *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail
	SetRequestId(v string) *GetInterceptionRuleDetailResponseBody
	GetRequestId() *string
}

type GetInterceptionRuleDetailResponseBody struct {
	// The details of the rule.
	InterceptionRuleDetail *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail `json:"InterceptionRuleDetail,omitempty" xml:"InterceptionRuleDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9FBC6E47-7508-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInterceptionRuleDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionRuleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterceptionRuleDetailResponseBody) GetInterceptionRuleDetail() *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail {
	return s.InterceptionRuleDetail
}

func (s *GetInterceptionRuleDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInterceptionRuleDetailResponseBody) SetInterceptionRuleDetail(v *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) *GetInterceptionRuleDetailResponseBody {
	s.InterceptionRuleDetail = v
	return s
}

func (s *GetInterceptionRuleDetailResponseBody) SetRequestId(v string) *GetInterceptionRuleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBody) Validate() error {
	if s.InterceptionRuleDetail != nil {
		if err := s.InterceptionRuleDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail struct {
	// The destination network object.
	DstTarget *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget `json:"DstTarget,omitempty" xml:"DstTarget,omitempty" type:"Struct"`
	// The interception mode. Valid values:
	//
	// 	- **0**: monitor
	//
	// 	- **1**: block
	//
	// 	- **2**: alert
	//
	// 	- **3**: allow
	//
	// example:
	//
	// 1
	InterceptType *int64 `json:"InterceptType,omitempty" xml:"InterceptType,omitempty"`
	// The priority of the rule. Valid values: 1 to 1000. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	OrderIndex *int64 `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 467
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// dmz-frontend-accept
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the defense rule. Valid values:
	//
	// 	- **1**: The rule is enabled.
	//
	// 	- **0**: The rule is disabled.
	//
	// example:
	//
	// 1
	RuleSwitch *int32 `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **suggest**: a suggestion rule
	//
	// 	- **customize**: a custom rule
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The source network object.
	SrcTarget *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget `json:"SrcTarget,omitempty" xml:"SrcTarget,omitempty" type:"Struct"`
}

func (s GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) GoString() string {
	return s.String()
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) GetDstTarget() *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget {
	return s.DstTarget
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) GetInterceptType() *int64 {
	return s.InterceptType
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) GetOrderIndex() *int64 {
	return s.OrderIndex
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) GetRuleName() *string {
	return s.RuleName
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) GetRuleType() *string {
	return s.RuleType
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) GetSrcTarget() *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget {
	return s.SrcTarget
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) SetDstTarget(v *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail {
	s.DstTarget = v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) SetInterceptType(v int64) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail {
	s.InterceptType = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) SetOrderIndex(v int64) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail {
	s.OrderIndex = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) SetRuleId(v int64) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail {
	s.RuleId = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) SetRuleName(v string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail {
	s.RuleName = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) SetRuleSwitch(v int32) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail {
	s.RuleSwitch = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) SetRuleType(v string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail {
	s.RuleType = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) SetSrcTarget(v *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail {
	s.SrcTarget = v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetail) Validate() error {
	if s.DstTarget != nil {
		if err := s.DstTarget.Validate(); err != nil {
			return err
		}
	}
	if s.SrcTarget != nil {
		if err := s.SrcTarget.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget struct {
	// The name of the application.
	//
	// example:
	//
	// console
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// An array that consists of the name of the image specified for the network object.
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// The namespace to which the network object belongs.
	//
	// example:
	//
	// kube-system
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// An array that consists of the port range of the destination network object.
	Ports []*string `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The type of the rule.
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// An array that consists of the labels specified for the network object.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The ID of the network object.
	//
	// example:
	//
	// 200014
	TargetId *int32 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the object.
	//
	// example:
	//
	// demo4-be1
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the network object.
	//
	// example:
	//
	// IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) GoString() string {
	return s.String()
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) GetAppName() *string {
	return s.AppName
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) GetImageList() []*string {
	return s.ImageList
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) GetNamespace() *string {
	return s.Namespace
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) GetPorts() []*string {
	return s.Ports
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) GetRuleType() *string {
	return s.RuleType
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) GetTagList() []*string {
	return s.TagList
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) GetTargetId() *int32 {
	return s.TargetId
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) GetTargetName() *string {
	return s.TargetName
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) SetAppName(v string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget {
	s.AppName = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) SetImageList(v []*string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget {
	s.ImageList = v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) SetNamespace(v string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget {
	s.Namespace = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) SetPorts(v []*string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget {
	s.Ports = v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) SetRuleType(v string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget {
	s.RuleType = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) SetTagList(v []*string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget {
	s.TagList = v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) SetTargetId(v int32) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget {
	s.TargetId = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) SetTargetName(v string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget {
	s.TargetName = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) SetTargetType(v string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget {
	s.TargetType = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailDstTarget) Validate() error {
	return dara.Validate(s)
}

type GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget struct {
	// The name of the application.
	//
	// example:
	//
	// console
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The images of the network object.
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// The namespace to which the network object belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The type of the rule.
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The labels specified for the network object.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The ID of the network object.
	//
	// example:
	//
	// 300635
	TargetId *int32 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the object.
	//
	// example:
	//
	// dmz
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the network object.
	//
	// example:
	//
	// IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) GoString() string {
	return s.String()
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) GetAppName() *string {
	return s.AppName
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) GetImageList() []*string {
	return s.ImageList
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) GetNamespace() *string {
	return s.Namespace
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) GetRuleType() *string {
	return s.RuleType
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) GetTagList() []*string {
	return s.TagList
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) GetTargetId() *int32 {
	return s.TargetId
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) GetTargetName() *string {
	return s.TargetName
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) SetAppName(v string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget {
	s.AppName = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) SetImageList(v []*string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget {
	s.ImageList = v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) SetNamespace(v string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget {
	s.Namespace = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) SetRuleType(v string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget {
	s.RuleType = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) SetTagList(v []*string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget {
	s.TagList = v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) SetTargetId(v int32) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget {
	s.TargetId = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) SetTargetName(v string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget {
	s.TargetName = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) SetTargetType(v string) *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget {
	s.TargetType = &v
	return s
}

func (s *GetInterceptionRuleDetailResponseBodyInterceptionRuleDetailSrcTarget) Validate() error {
	return dara.Validate(s)
}
