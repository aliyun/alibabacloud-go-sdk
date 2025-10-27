// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContainerDefenseRuleClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterList(v []*ListContainerDefenseRuleClustersResponseBodyClusterList) *ListContainerDefenseRuleClustersResponseBody
	GetClusterList() []*ListContainerDefenseRuleClustersResponseBodyClusterList
	SetCode(v string) *ListContainerDefenseRuleClustersResponseBody
	GetCode() *string
	SetCount(v int32) *ListContainerDefenseRuleClustersResponseBody
	GetCount() *int32
	SetHttpStatusCode(v int32) *ListContainerDefenseRuleClustersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListContainerDefenseRuleClustersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContainerDefenseRuleClustersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContainerDefenseRuleClustersResponseBody
	GetSuccess() *bool
}

type ListContainerDefenseRuleClustersResponseBody struct {
	// The clusters.
	ClusterList []*ListContainerDefenseRuleClustersResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Repeated"`
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1F995515-CAF3-5F84-8D82-C9F706AD5070
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListContainerDefenseRuleClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContainerDefenseRuleClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListContainerDefenseRuleClustersResponseBody) GetClusterList() []*ListContainerDefenseRuleClustersResponseBodyClusterList {
	return s.ClusterList
}

func (s *ListContainerDefenseRuleClustersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContainerDefenseRuleClustersResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListContainerDefenseRuleClustersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListContainerDefenseRuleClustersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContainerDefenseRuleClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContainerDefenseRuleClustersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContainerDefenseRuleClustersResponseBody) SetClusterList(v []*ListContainerDefenseRuleClustersResponseBodyClusterList) *ListContainerDefenseRuleClustersResponseBody {
	s.ClusterList = v
	return s
}

func (s *ListContainerDefenseRuleClustersResponseBody) SetCode(v string) *ListContainerDefenseRuleClustersResponseBody {
	s.Code = &v
	return s
}

func (s *ListContainerDefenseRuleClustersResponseBody) SetCount(v int32) *ListContainerDefenseRuleClustersResponseBody {
	s.Count = &v
	return s
}

func (s *ListContainerDefenseRuleClustersResponseBody) SetHttpStatusCode(v int32) *ListContainerDefenseRuleClustersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContainerDefenseRuleClustersResponseBody) SetMessage(v string) *ListContainerDefenseRuleClustersResponseBody {
	s.Message = &v
	return s
}

func (s *ListContainerDefenseRuleClustersResponseBody) SetRequestId(v string) *ListContainerDefenseRuleClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContainerDefenseRuleClustersResponseBody) SetSuccess(v bool) *ListContainerDefenseRuleClustersResponseBody {
	s.Success = &v
	return s
}

func (s *ListContainerDefenseRuleClustersResponseBody) Validate() error {
	if s.ClusterList != nil {
		for _, item := range s.ClusterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContainerDefenseRuleClustersResponseBodyClusterList struct {
	// Indicates whether all namespaces are included. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AllNamespace *int32 `json:"AllNamespace,omitempty" xml:"AllNamespace,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// cfeb7a9f99ce740e98c5595d0fe37****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespaces.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The ID of the rule.
	//
	// >  You can call the [ListInterceptionRulePage](https://help.aliyun.com/document_detail/2590599.html) operation to query the IDs of rules.
	//
	// example:
	//
	// 403178
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListContainerDefenseRuleClustersResponseBodyClusterList) String() string {
	return dara.Prettify(s)
}

func (s ListContainerDefenseRuleClustersResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *ListContainerDefenseRuleClustersResponseBodyClusterList) GetAllNamespace() *int32 {
	return s.AllNamespace
}

func (s *ListContainerDefenseRuleClustersResponseBodyClusterList) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListContainerDefenseRuleClustersResponseBodyClusterList) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *ListContainerDefenseRuleClustersResponseBodyClusterList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListContainerDefenseRuleClustersResponseBodyClusterList) SetAllNamespace(v int32) *ListContainerDefenseRuleClustersResponseBodyClusterList {
	s.AllNamespace = &v
	return s
}

func (s *ListContainerDefenseRuleClustersResponseBodyClusterList) SetClusterId(v string) *ListContainerDefenseRuleClustersResponseBodyClusterList {
	s.ClusterId = &v
	return s
}

func (s *ListContainerDefenseRuleClustersResponseBodyClusterList) SetNamespaces(v []*string) *ListContainerDefenseRuleClustersResponseBodyClusterList {
	s.Namespaces = v
	return s
}

func (s *ListContainerDefenseRuleClustersResponseBodyClusterList) SetRuleId(v int64) *ListContainerDefenseRuleClustersResponseBodyClusterList {
	s.RuleId = &v
	return s
}

func (s *ListContainerDefenseRuleClustersResponseBodyClusterList) Validate() error {
	return dara.Validate(s)
}
