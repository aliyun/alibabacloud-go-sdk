// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListWorkspaceUsersResponseBodyData) *ListWorkspaceUsersResponseBody
	GetData() []*ListWorkspaceUsersResponseBodyData
	SetMaxResults(v int32) *ListWorkspaceUsersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkspaceUsersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkspaceUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkspaceUsersResponseBody
	GetSuccess() *bool
	SetTotalCount(v string) *ListWorkspaceUsersResponseBody
	GetTotalCount() *string
}

type ListWorkspaceUsersResponseBody struct {
	Data []*ListWorkspaceUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// NesLoKLEdIZrKhDT7I2gS****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWorkspaceUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUsersResponseBody) GetData() []*ListWorkspaceUsersResponseBodyData {
	return s.Data
}

func (s *ListWorkspaceUsersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspaceUsersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspaceUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspaceUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkspaceUsersResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListWorkspaceUsersResponseBody) SetData(v []*ListWorkspaceUsersResponseBodyData) *ListWorkspaceUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkspaceUsersResponseBody) SetMaxResults(v int32) *ListWorkspaceUsersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspaceUsersResponseBody) SetNextToken(v string) *ListWorkspaceUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspaceUsersResponseBody) SetRequestId(v string) *ListWorkspaceUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspaceUsersResponseBody) SetSuccess(v bool) *ListWorkspaceUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkspaceUsersResponseBody) SetTotalCount(v string) *ListWorkspaceUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspaceUsersResponseBody) Validate() error {
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

type ListWorkspaceUsersResponseBodyData struct {
	// dms user id
	//
	// example:
	//
	// 12345
	DmsUserId *int64 `json:"DmsUserId,omitempty" xml:"DmsUserId,omitempty"`
	// example:
	//
	// li
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// example:
	//
	// dynamic_320514_20250429102514
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
}

func (s ListWorkspaceUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUsersResponseBodyData) GetDmsUserId() *int64 {
	return s.DmsUserId
}

func (s *ListWorkspaceUsersResponseBodyData) GetLoginName() *string {
	return s.LoginName
}

func (s *ListWorkspaceUsersResponseBodyData) GetNickName() *string {
	return s.NickName
}

func (s *ListWorkspaceUsersResponseBodyData) SetDmsUserId(v int64) *ListWorkspaceUsersResponseBodyData {
	s.DmsUserId = &v
	return s
}

func (s *ListWorkspaceUsersResponseBodyData) SetLoginName(v string) *ListWorkspaceUsersResponseBodyData {
	s.LoginName = &v
	return s
}

func (s *ListWorkspaceUsersResponseBodyData) SetNickName(v string) *ListWorkspaceUsersResponseBodyData {
	s.NickName = &v
	return s
}

func (s *ListWorkspaceUsersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
