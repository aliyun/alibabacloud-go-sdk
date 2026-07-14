// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNotifyPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNotifyPoliciesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNotifyPoliciesResponseBody
	GetNextToken() *string
	SetNotifyPolicyList(v []*NotifyPolicySummary) *ListNotifyPoliciesResponseBody
	GetNotifyPolicyList() []*NotifyPolicySummary
	SetRequestId(v string) *ListNotifyPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListNotifyPoliciesResponseBody
	GetTotalCount() *int64
}

type ListNotifyPoliciesResponseBody struct {
	// The maximum number of entries returned in this request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page. This parameter is empty if no more data is available.
	//
	// example:
	//
	// eyJjdXJzb3IiOjEwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of notify policies. Each entry is of the NotifyPolicySummary type (lightweight view).
	//
	// example:
	//
	// [{"uuid":"np-12345678-1234-1234-1234-123456789012","name":"prod-alert","description":"生产环境告警策略","enabled":true,"version":1,"workspace":"default-cms-xxxx-cn-hangzhou","userId":"107640","createTime":"1711792800000","updateTime":"1711792800000"}]
	NotifyPolicyList []*NotifyPolicySummary `json:"notifyPolicyList,omitempty" xml:"notifyPolicyList,omitempty" type:"Repeated"`
	// The request ID. You can use this ID for troubleshooting and ticket submission.
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries. The actual total is returned on the first page. A fixed value of -1 is returned on subsequent pages.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListNotifyPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNotifyPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNotifyPoliciesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNotifyPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNotifyPoliciesResponseBody) GetNotifyPolicyList() []*NotifyPolicySummary {
	return s.NotifyPolicyList
}

func (s *ListNotifyPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNotifyPoliciesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListNotifyPoliciesResponseBody) SetMaxResults(v int32) *ListNotifyPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNotifyPoliciesResponseBody) SetNextToken(v string) *ListNotifyPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNotifyPoliciesResponseBody) SetNotifyPolicyList(v []*NotifyPolicySummary) *ListNotifyPoliciesResponseBody {
	s.NotifyPolicyList = v
	return s
}

func (s *ListNotifyPoliciesResponseBody) SetRequestId(v string) *ListNotifyPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNotifyPoliciesResponseBody) SetTotalCount(v int64) *ListNotifyPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNotifyPoliciesResponseBody) Validate() error {
	if s.NotifyPolicyList != nil {
		for _, item := range s.NotifyPolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
