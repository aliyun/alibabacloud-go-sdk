// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterceptionTargetDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetInterceptionTargetDetailResponseBody
	GetRequestId() *string
	SetRuleTarget(v *GetInterceptionTargetDetailResponseBodyRuleTarget) *GetInterceptionTargetDetailResponseBody
	GetRuleTarget() *GetInterceptionTargetDetailResponseBodyRuleTarget
}

type GetInterceptionTargetDetailResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53C0DC1F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the network object.
	RuleTarget *GetInterceptionTargetDetailResponseBodyRuleTarget `json:"RuleTarget,omitempty" xml:"RuleTarget,omitempty" type:"Struct"`
}

func (s GetInterceptionTargetDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionTargetDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterceptionTargetDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInterceptionTargetDetailResponseBody) GetRuleTarget() *GetInterceptionTargetDetailResponseBodyRuleTarget {
	return s.RuleTarget
}

func (s *GetInterceptionTargetDetailResponseBody) SetRequestId(v string) *GetInterceptionTargetDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterceptionTargetDetailResponseBody) SetRuleTarget(v *GetInterceptionTargetDetailResponseBodyRuleTarget) *GetInterceptionTargetDetailResponseBody {
	s.RuleTarget = v
	return s
}

func (s *GetInterceptionTargetDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInterceptionTargetDetailResponseBodyRuleTarget struct {
	// The name of the application to which the network object belongs.
	//
	// example:
	//
	// netperf-client
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the container cluster.
	//
	// example:
	//
	// ca6e6594def8d4be8b2795fd12c32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the container cluster.
	//
	// example:
	//
	// sas-test-cnnf
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// An array that consists of the images of the network object.
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// The namespace.
	//
	// example:
	//
	// secondary
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// An array that consists of the labels specified for the network object.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The ID of the network object.
	//
	// example:
	//
	// 400723
	TargetId *int64 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the network object.
	//
	// example:
	//
	// destination-test-obj-Na3cF
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the network object. Valid values:
	//
	// 	- **IMAGE**: image
	//
	// example:
	//
	// IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetInterceptionTargetDetailResponseBodyRuleTarget) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionTargetDetailResponseBodyRuleTarget) GoString() string {
	return s.String()
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) GetAppName() *string {
	return s.AppName
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) GetImageList() []*string {
	return s.ImageList
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) GetNamespace() *string {
	return s.Namespace
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) GetTagList() []*string {
	return s.TagList
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) GetTargetId() *int64 {
	return s.TargetId
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) GetTargetName() *string {
	return s.TargetName
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) SetAppName(v string) *GetInterceptionTargetDetailResponseBodyRuleTarget {
	s.AppName = &v
	return s
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) SetClusterId(v string) *GetInterceptionTargetDetailResponseBodyRuleTarget {
	s.ClusterId = &v
	return s
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) SetClusterName(v string) *GetInterceptionTargetDetailResponseBodyRuleTarget {
	s.ClusterName = &v
	return s
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) SetImageList(v []*string) *GetInterceptionTargetDetailResponseBodyRuleTarget {
	s.ImageList = v
	return s
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) SetNamespace(v string) *GetInterceptionTargetDetailResponseBodyRuleTarget {
	s.Namespace = &v
	return s
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) SetTagList(v []*string) *GetInterceptionTargetDetailResponseBodyRuleTarget {
	s.TagList = v
	return s
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) SetTargetId(v int64) *GetInterceptionTargetDetailResponseBodyRuleTarget {
	s.TargetId = &v
	return s
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) SetTargetName(v string) *GetInterceptionTargetDetailResponseBodyRuleTarget {
	s.TargetName = &v
	return s
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) SetTargetType(v string) *GetInterceptionTargetDetailResponseBodyRuleTarget {
	s.TargetType = &v
	return s
}

func (s *GetInterceptionTargetDetailResponseBodyRuleTarget) Validate() error {
	return dara.Validate(s)
}
