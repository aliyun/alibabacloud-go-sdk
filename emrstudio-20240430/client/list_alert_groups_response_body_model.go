// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAlertGroupsResponseBodyData) *ListAlertGroupsResponseBody
	GetData() []*ListAlertGroupsResponseBodyData
	SetMaxResults(v int32) *ListAlertGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAlertGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAlertGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAlertGroupsResponseBody
	GetTotalCount() *int32
}

type ListAlertGroupsResponseBody struct {
	Data []*ListAlertGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAlertGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertGroupsResponseBody) GetData() []*ListAlertGroupsResponseBodyData {
	return s.Data
}

func (s *ListAlertGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAlertGroupsResponseBody) SetData(v []*ListAlertGroupsResponseBodyData) *ListAlertGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListAlertGroupsResponseBody) SetMaxResults(v int32) *ListAlertGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAlertGroupsResponseBody) SetNextToken(v string) *ListAlertGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAlertGroupsResponseBody) SetRequestId(v string) *ListAlertGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertGroupsResponseBody) SetTotalCount(v int32) *ListAlertGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAlertGroupsResponseBody) Validate() error {
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

type ListAlertGroupsResponseBodyData struct {
	// example:
	//
	// ag-8mklwpevk74****
	AlertGroupId *string `json:"alertGroupId,omitempty" xml:"alertGroupId,omitempty"`
	// example:
	//
	// 50
	AlertInstanceIds *string `json:"alertInstanceIds,omitempty" xml:"alertInstanceIds,omitempty"`
	// example:
	//
	// 1726834240000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// name
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// 1726834240000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListAlertGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAlertGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlertGroupsResponseBodyData) GetAlertGroupId() *string {
	return s.AlertGroupId
}

func (s *ListAlertGroupsResponseBodyData) GetAlertInstanceIds() *string {
	return s.AlertInstanceIds
}

func (s *ListAlertGroupsResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListAlertGroupsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListAlertGroupsResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAlertGroupsResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListAlertGroupsResponseBodyData) SetAlertGroupId(v string) *ListAlertGroupsResponseBodyData {
	s.AlertGroupId = &v
	return s
}

func (s *ListAlertGroupsResponseBodyData) SetAlertInstanceIds(v string) *ListAlertGroupsResponseBodyData {
	s.AlertInstanceIds = &v
	return s
}

func (s *ListAlertGroupsResponseBodyData) SetCreateTime(v int64) *ListAlertGroupsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAlertGroupsResponseBodyData) SetDescription(v string) *ListAlertGroupsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListAlertGroupsResponseBodyData) SetGroupName(v string) *ListAlertGroupsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListAlertGroupsResponseBodyData) SetUpdateTime(v int64) *ListAlertGroupsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListAlertGroupsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
