// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNodeGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodeGroupsResponseBody
	GetNextToken() *string
	SetNodeGroups(v []*NodeGroup) *ListNodeGroupsResponseBody
	GetNodeGroups() []*NodeGroup
	SetRequestId(v string) *ListNodeGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNodeGroupsResponseBody
	GetTotalCount() *int32
}

type ListNodeGroupsResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Returns the location of the data that was read. Empty indicates that the data has been read.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The array of node groups.
	NodeGroups []*NodeGroup `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodeGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodeGroupsResponseBody) GetNodeGroups() []*NodeGroup {
	return s.NodeGroups
}

func (s *ListNodeGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodeGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNodeGroupsResponseBody) SetMaxResults(v int32) *ListNodeGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNodeGroupsResponseBody) SetNextToken(v string) *ListNodeGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupsResponseBody) SetNodeGroups(v []*NodeGroup) *ListNodeGroupsResponseBody {
	s.NodeGroups = v
	return s
}

func (s *ListNodeGroupsResponseBody) SetRequestId(v string) *ListNodeGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeGroupsResponseBody) SetTotalCount(v int32) *ListNodeGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNodeGroupsResponseBody) Validate() error {
	if s.NodeGroups != nil {
		for _, item := range s.NodeGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
