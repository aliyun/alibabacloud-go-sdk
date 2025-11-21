// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInstancesV2ResponseBody
	GetRequestId() *string
	SetInstances(v []*InstanceDetail) *ListInstancesV2ResponseBody
	GetInstances() []*InstanceDetail
	SetMaxResults(v int32) *ListInstancesV2ResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstancesV2ResponseBody
	GetNextToken() *string
	SetSuccess(v bool) *ListInstancesV2ResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListInstancesV2ResponseBody
	GetTotal() *int32
}

type ListInstancesV2ResponseBody struct {
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Instances []*InstanceDetail `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// None
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 15
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListInstancesV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesV2ResponseBody) GetInstances() []*InstanceDetail {
	return s.Instances
}

func (s *ListInstancesV2ResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancesV2ResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancesV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstancesV2ResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListInstancesV2ResponseBody) SetRequestId(v string) *ListInstancesV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesV2ResponseBody) SetInstances(v []*InstanceDetail) *ListInstancesV2ResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesV2ResponseBody) SetMaxResults(v int32) *ListInstancesV2ResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesV2ResponseBody) SetNextToken(v string) *ListInstancesV2ResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancesV2ResponseBody) SetSuccess(v bool) *ListInstancesV2ResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesV2ResponseBody) SetTotal(v int32) *ListInstancesV2ResponseBody {
	s.Total = &v
	return s
}

func (s *ListInstancesV2ResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
