// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMaintainWindowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaintainWindowList(v []*MaintainWindowForView) *ListMaintainWindowsResponseBody
	GetMaintainWindowList() []*MaintainWindowForView
	SetMaxResults(v int32) *ListMaintainWindowsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListMaintainWindowsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMaintainWindowsResponseBody
	GetRequestId() *string
}

type ListMaintainWindowsResponseBody struct {
	// The list of silence policies. An empty list is returned when no items match. The list items do not include workspaceFilterSetting.
	MaintainWindowList []*MaintainWindowForView `json:"maintainWindowList,omitempty" xml:"maintainWindowList,omitempty" type:"Repeated"`
	// The maximum number of records returned in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page. A non-empty value indicates that more results may be available, but the next page may still be empty. An empty value indicates the end of pagination.
	//
	// example:
	//
	// Y21zRXZlbnRCYXNlUGFnZT0y
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The unique ID of this request, used for troubleshooting and ticket tracking.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A0D1C36
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMaintainWindowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMaintainWindowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMaintainWindowsResponseBody) GetMaintainWindowList() []*MaintainWindowForView {
	return s.MaintainWindowList
}

func (s *ListMaintainWindowsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMaintainWindowsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMaintainWindowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMaintainWindowsResponseBody) SetMaintainWindowList(v []*MaintainWindowForView) *ListMaintainWindowsResponseBody {
	s.MaintainWindowList = v
	return s
}

func (s *ListMaintainWindowsResponseBody) SetMaxResults(v int32) *ListMaintainWindowsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMaintainWindowsResponseBody) SetNextToken(v string) *ListMaintainWindowsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMaintainWindowsResponseBody) SetRequestId(v string) *ListMaintainWindowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMaintainWindowsResponseBody) Validate() error {
	if s.MaintainWindowList != nil {
		for _, item := range s.MaintainWindowList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
