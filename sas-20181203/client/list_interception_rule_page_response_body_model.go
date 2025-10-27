// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterceptionRulePageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInterceptionRuleList(v []*ListInterceptionRulePageResponseBodyInterceptionRuleList) *ListInterceptionRulePageResponseBody
	GetInterceptionRuleList() []*ListInterceptionRulePageResponseBodyInterceptionRuleList
	SetPageInfo(v *ListInterceptionRulePageResponseBodyPageInfo) *ListInterceptionRulePageResponseBody
	GetPageInfo() *ListInterceptionRulePageResponseBodyPageInfo
	SetRequestId(v string) *ListInterceptionRulePageResponseBody
	GetRequestId() *string
}

type ListInterceptionRulePageResponseBody struct {
	// An array that consists of information about the defense rules.
	InterceptionRuleList []*ListInterceptionRulePageResponseBodyInterceptionRuleList `json:"InterceptionRuleList,omitempty" xml:"InterceptionRuleList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListInterceptionRulePageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// ACF97412-FD09-4D1F-994F-34DF12BR****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInterceptionRulePageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionRulePageResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterceptionRulePageResponseBody) GetInterceptionRuleList() []*ListInterceptionRulePageResponseBodyInterceptionRuleList {
	return s.InterceptionRuleList
}

func (s *ListInterceptionRulePageResponseBody) GetPageInfo() *ListInterceptionRulePageResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListInterceptionRulePageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInterceptionRulePageResponseBody) SetInterceptionRuleList(v []*ListInterceptionRulePageResponseBodyInterceptionRuleList) *ListInterceptionRulePageResponseBody {
	s.InterceptionRuleList = v
	return s
}

func (s *ListInterceptionRulePageResponseBody) SetPageInfo(v *ListInterceptionRulePageResponseBodyPageInfo) *ListInterceptionRulePageResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListInterceptionRulePageResponseBody) SetRequestId(v string) *ListInterceptionRulePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterceptionRulePageResponseBody) Validate() error {
	if s.InterceptionRuleList != nil {
		for _, item := range s.InterceptionRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInterceptionRulePageResponseBodyInterceptionRuleList struct {
	// The destination network object.
	DstTarget *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget `json:"DstTarget,omitempty" xml:"DstTarget,omitempty" type:"Struct"`
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
	// The order in which the entries are sorted.
	//
	// example:
	//
	// 1
	OrderIndex *int64 `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	// The ID of the defense rule.
	//
	// example:
	//
	// 30****
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the defense rule.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the defense rule. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	RuleSwitch *int32 `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The type of the defense rule.
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The source network object.
	SrcTarget *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget `json:"SrcTarget,omitempty" xml:"SrcTarget,omitempty" type:"Struct"`
}

func (s ListInterceptionRulePageResponseBodyInterceptionRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionRulePageResponseBodyInterceptionRuleList) GoString() string {
	return s.String()
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) GetDstTarget() *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget {
	return s.DstTarget
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) GetInterceptType() *int64 {
	return s.InterceptType
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) GetOrderIndex() *int64 {
	return s.OrderIndex
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) GetRuleType() *string {
	return s.RuleType
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) GetSrcTarget() *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget {
	return s.SrcTarget
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) SetDstTarget(v *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) *ListInterceptionRulePageResponseBodyInterceptionRuleList {
	s.DstTarget = v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) SetInterceptType(v int64) *ListInterceptionRulePageResponseBodyInterceptionRuleList {
	s.InterceptType = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) SetOrderIndex(v int64) *ListInterceptionRulePageResponseBodyInterceptionRuleList {
	s.OrderIndex = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) SetRuleId(v int64) *ListInterceptionRulePageResponseBodyInterceptionRuleList {
	s.RuleId = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) SetRuleName(v string) *ListInterceptionRulePageResponseBodyInterceptionRuleList {
	s.RuleName = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) SetRuleSwitch(v int32) *ListInterceptionRulePageResponseBodyInterceptionRuleList {
	s.RuleSwitch = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) SetRuleType(v string) *ListInterceptionRulePageResponseBodyInterceptionRuleList {
	s.RuleType = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) SetSrcTarget(v *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) *ListInterceptionRulePageResponseBodyInterceptionRuleList {
	s.SrcTarget = v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleList) Validate() error {
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

type ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget struct {
	// The name of the application.
	//
	// example:
	//
	// console
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// An array that consists of the affected images.
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// The namespace.
	//
	// example:
	//
	// test
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// An array that consists of information about the ports used by the destination server.
	Ports []*string `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The type of the defense rule. Valid values:
	//
	// 	- **suggest**: intelligently recommended rule
	//
	// 	- **customize**: custom rule
	//
	// 	- **system**: system rule
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// An array that consists of tags added to the destination network object.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The ID of the network object.
	//
	// example:
	//
	// 302001
	TargetId *int32 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the network object.
	//
	// example:
	//
	// demo4****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the network object.
	//
	// example:
	//
	// IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) GoString() string {
	return s.String()
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) GetAppName() *string {
	return s.AppName
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) GetImageList() []*string {
	return s.ImageList
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) GetNamespace() *string {
	return s.Namespace
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) GetPorts() []*string {
	return s.Ports
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) GetRuleType() *string {
	return s.RuleType
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) GetTagList() []*string {
	return s.TagList
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) GetTargetId() *int32 {
	return s.TargetId
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) GetTargetName() *string {
	return s.TargetName
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) SetAppName(v string) *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget {
	s.AppName = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) SetImageList(v []*string) *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget {
	s.ImageList = v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) SetNamespace(v string) *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget {
	s.Namespace = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) SetPorts(v []*string) *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget {
	s.Ports = v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) SetRuleType(v string) *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget {
	s.RuleType = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) SetTagList(v []*string) *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget {
	s.TagList = v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) SetTargetId(v int32) *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget {
	s.TargetId = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) SetTargetName(v string) *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget {
	s.TargetName = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) SetTargetType(v string) *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget {
	s.TargetType = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListDstTarget) Validate() error {
	return dara.Validate(s)
}

type ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget struct {
	// The name of the application.
	//
	// example:
	//
	// ack-jenkins-lawr****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// An array that consists of the images of the network object.
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// The namespace.
	//
	// example:
	//
	// jenkins
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The type of the defense rule. Valid values:
	//
	// 	- **suggest**: intelligently recommended rule
	//
	// 	- **customize**: custom rule
	//
	// 	- **system**: system rule
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// An array that consists of tags added to the source network object.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The ID of the network object.
	//
	// example:
	//
	// 40****
	TargetId *int32 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the network object.
	//
	// example:
	//
	// mhh-te****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the affected assets.
	//
	// example:
	//
	// containerId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) GoString() string {
	return s.String()
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) GetAppName() *string {
	return s.AppName
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) GetImageList() []*string {
	return s.ImageList
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) GetNamespace() *string {
	return s.Namespace
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) GetRuleType() *string {
	return s.RuleType
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) GetTagList() []*string {
	return s.TagList
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) GetTargetId() *int32 {
	return s.TargetId
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) GetTargetName() *string {
	return s.TargetName
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) SetAppName(v string) *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget {
	s.AppName = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) SetImageList(v []*string) *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget {
	s.ImageList = v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) SetNamespace(v string) *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget {
	s.Namespace = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) SetRuleType(v string) *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget {
	s.RuleType = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) SetTagList(v []*string) *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget {
	s.TagList = v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) SetTargetId(v int32) *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget {
	s.TargetId = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) SetTargetName(v string) *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget {
	s.TargetName = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) SetTargetType(v string) *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget {
	s.TargetType = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyInterceptionRuleListSrcTarget) Validate() error {
	return dara.Validate(s)
}

type ListInterceptionRulePageResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 19
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInterceptionRulePageResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionRulePageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListInterceptionRulePageResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListInterceptionRulePageResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListInterceptionRulePageResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterceptionRulePageResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInterceptionRulePageResponseBodyPageInfo) SetCount(v int32) *ListInterceptionRulePageResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyPageInfo) SetCurrentPage(v int32) *ListInterceptionRulePageResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyPageInfo) SetPageSize(v int32) *ListInterceptionRulePageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyPageInfo) SetTotalCount(v int32) *ListInterceptionRulePageResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListInterceptionRulePageResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
