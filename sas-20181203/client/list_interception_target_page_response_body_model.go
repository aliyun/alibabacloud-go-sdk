// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterceptionTargetPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListInterceptionTargetPageResponseBodyPageInfo) *ListInterceptionTargetPageResponseBody
	GetPageInfo() *ListInterceptionTargetPageResponseBodyPageInfo
	SetRequestId(v string) *ListInterceptionTargetPageResponseBody
	GetRequestId() *string
	SetRuleTargetList(v []*ListInterceptionTargetPageResponseBodyRuleTargetList) *ListInterceptionTargetPageResponseBody
	GetRuleTargetList() []*ListInterceptionTargetPageResponseBodyRuleTargetList
}

type ListInterceptionTargetPageResponseBody struct {
	// The pagination information.
	PageInfo *ListInterceptionTargetPageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 028CF634-5268-5660-9575-48C9ED6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the network objects.
	RuleTargetList []*ListInterceptionTargetPageResponseBodyRuleTargetList `json:"RuleTargetList,omitempty" xml:"RuleTargetList,omitempty" type:"Repeated"`
}

func (s ListInterceptionTargetPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionTargetPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterceptionTargetPageResponseBody) GetPageInfo() *ListInterceptionTargetPageResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListInterceptionTargetPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInterceptionTargetPageResponseBody) GetRuleTargetList() []*ListInterceptionTargetPageResponseBodyRuleTargetList {
	return s.RuleTargetList
}

func (s *ListInterceptionTargetPageResponseBody) SetPageInfo(v *ListInterceptionTargetPageResponseBodyPageInfo) *ListInterceptionTargetPageResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListInterceptionTargetPageResponseBody) SetRequestId(v string) *ListInterceptionTargetPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterceptionTargetPageResponseBody) SetRuleTargetList(v []*ListInterceptionTargetPageResponseBodyRuleTargetList) *ListInterceptionTargetPageResponseBody {
	s.RuleTargetList = v
	return s
}

func (s *ListInterceptionTargetPageResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RuleTargetList != nil {
		for _, item := range s.RuleTargetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInterceptionTargetPageResponseBodyPageInfo struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
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
	// 45
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInterceptionTargetPageResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionTargetPageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListInterceptionTargetPageResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListInterceptionTargetPageResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterceptionTargetPageResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInterceptionTargetPageResponseBodyPageInfo) SetCurrentPage(v int32) *ListInterceptionTargetPageResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyPageInfo) SetPageSize(v int32) *ListInterceptionTargetPageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyPageInfo) SetTotalCount(v int32) *ListInterceptionTargetPageResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListInterceptionTargetPageResponseBodyRuleTargetList struct {
	// The name of the application of the network object.
	//
	// example:
	//
	// frontend
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the container cluster.
	//
	// example:
	//
	// c3e2eae03eb064d2ebf940cd5e1b17****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the container cluster.
	//
	// example:
	//
	// sas-test-cnnf
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The images of the network object.
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// The namespace to which the network object belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The type of the rule. Valid value:
	//
	// 	- customize: custom rule
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The tags specified for the network object.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The ID of the network object.
	//
	// > You can call the [ListInterceptionTargetPage](~~ListInterceptionTargetPage~~) operation to query the IDs of network objects.
	//
	// example:
	//
	// 400914
	TargetId *int64 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the network object.
	//
	// example:
	//
	// destination-test-obj-Na3cF
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the network object. Valid value:
	//
	// 	- IMAGE
	//
	// example:
	//
	// IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListInterceptionTargetPageResponseBodyRuleTargetList) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionTargetPageResponseBodyRuleTargetList) GoString() string {
	return s.String()
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) GetAppName() *string {
	return s.AppName
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) GetImageList() []*string {
	return s.ImageList
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) GetNamespace() *string {
	return s.Namespace
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) GetRuleType() *string {
	return s.RuleType
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) GetTagList() []*string {
	return s.TagList
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) GetTargetId() *int64 {
	return s.TargetId
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) GetTargetName() *string {
	return s.TargetName
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) GetTargetType() *string {
	return s.TargetType
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) SetAppName(v string) *ListInterceptionTargetPageResponseBodyRuleTargetList {
	s.AppName = &v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) SetClusterId(v string) *ListInterceptionTargetPageResponseBodyRuleTargetList {
	s.ClusterId = &v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) SetClusterName(v string) *ListInterceptionTargetPageResponseBodyRuleTargetList {
	s.ClusterName = &v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) SetImageList(v []*string) *ListInterceptionTargetPageResponseBodyRuleTargetList {
	s.ImageList = v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) SetNamespace(v string) *ListInterceptionTargetPageResponseBodyRuleTargetList {
	s.Namespace = &v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) SetRuleType(v string) *ListInterceptionTargetPageResponseBodyRuleTargetList {
	s.RuleType = &v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) SetTagList(v []*string) *ListInterceptionTargetPageResponseBodyRuleTargetList {
	s.TagList = v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) SetTargetId(v int64) *ListInterceptionTargetPageResponseBodyRuleTargetList {
	s.TargetId = &v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) SetTargetName(v string) *ListInterceptionTargetPageResponseBodyRuleTargetList {
	s.TargetName = &v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) SetTargetType(v string) *ListInterceptionTargetPageResponseBodyRuleTargetList {
	s.TargetType = &v
	return s
}

func (s *ListInterceptionTargetPageResponseBodyRuleTargetList) Validate() error {
	return dara.Validate(s)
}
