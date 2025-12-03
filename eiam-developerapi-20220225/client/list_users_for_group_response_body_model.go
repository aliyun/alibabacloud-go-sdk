// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListUsersForGroupResponseBodyData) *ListUsersForGroupResponseBody
	GetData() []*ListUsersForGroupResponseBodyData
	SetMaxResults(v int32) *ListUsersForGroupResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUsersForGroupResponseBody
	GetNextToken() *string
	SetTotalCount(v int64) *ListUsersForGroupResponseBody
	GetTotalCount() *int64
}

type ListUsersForGroupResponseBody struct {
	// The returned data.
	Data []*ListUsersForGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// NTxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListUsersForGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBody) GetData() []*ListUsersForGroupResponseBodyData {
	return s.Data
}

func (s *ListUsersForGroupResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersForGroupResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersForGroupResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUsersForGroupResponseBody) SetData(v []*ListUsersForGroupResponseBodyData) *ListUsersForGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersForGroupResponseBody) SetMaxResults(v int32) *ListUsersForGroupResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetNextToken(v string) *ListUsersForGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetTotalCount(v int64) *ListUsersForGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersForGroupResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersForGroupResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The account ID.
	//
	// example:
	//
	// user_001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListUsersForGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersForGroupResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersForGroupResponseBodyData) SetInstanceId(v string) *ListUsersForGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListUsersForGroupResponseBodyData) SetUserId(v string) *ListUsersForGroupResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListUsersForGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
