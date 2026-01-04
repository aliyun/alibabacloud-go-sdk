// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListSchemasResponseBodyList) *ListSchemasResponseBody
	GetList() []*ListSchemasResponseBodyList
	SetMaxResults(v int32) *ListSchemasResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSchemasResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSchemasResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSchemasResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListSchemasResponseBody
	GetTotalCount() *int32
}

type ListSchemasResponseBody struct {
	List []*ListSchemasResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 9892074a2a89600ae4b0d5a34fb99a3f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 20250401102332e68e3d0b04ab4904
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchemasResponseBody) GetList() []*ListSchemasResponseBodyList {
	return s.List
}

func (s *ListSchemasResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSchemasResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSchemasResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSchemasResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSchemasResponseBody) SetList(v []*ListSchemasResponseBodyList) *ListSchemasResponseBody {
	s.List = v
	return s
}

func (s *ListSchemasResponseBody) SetMaxResults(v int32) *ListSchemasResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSchemasResponseBody) SetNextToken(v string) *ListSchemasResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSchemasResponseBody) SetRequestId(v string) *ListSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSchemasResponseBody) SetSuccess(v bool) *ListSchemasResponseBody {
	s.Success = &v
	return s
}

func (s *ListSchemasResponseBody) SetTotalCount(v int32) *ListSchemasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSchemasResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSchemasResponseBodyList struct {
	// example:
	//
	// 1708171905000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1048133943212399
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// [{\\"Type\\":\\"STRING\\",\\"AllowNull\\":true,\\"Name\\":\\"context\\"}]
	RecordSchema *string `json:"RecordSchema,omitempty" xml:"RecordSchema,omitempty"`
	// example:
	//
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// example:
	//
	// 0
	VersionId *int32 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListSchemasResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSchemasResponseBodyList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListSchemasResponseBodyList) GetCreator() *string {
	return s.Creator
}

func (s *ListSchemasResponseBodyList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListSchemasResponseBodyList) GetRecordSchema() *string {
	return s.RecordSchema
}

func (s *ListSchemasResponseBodyList) GetTopicName() *string {
	return s.TopicName
}

func (s *ListSchemasResponseBodyList) GetVersionId() *int32 {
	return s.VersionId
}

func (s *ListSchemasResponseBodyList) SetCreateTime(v int64) *ListSchemasResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *ListSchemasResponseBodyList) SetCreator(v string) *ListSchemasResponseBodyList {
	s.Creator = &v
	return s
}

func (s *ListSchemasResponseBodyList) SetProjectName(v string) *ListSchemasResponseBodyList {
	s.ProjectName = &v
	return s
}

func (s *ListSchemasResponseBodyList) SetRecordSchema(v string) *ListSchemasResponseBodyList {
	s.RecordSchema = &v
	return s
}

func (s *ListSchemasResponseBodyList) SetTopicName(v string) *ListSchemasResponseBodyList {
	s.TopicName = &v
	return s
}

func (s *ListSchemasResponseBodyList) SetVersionId(v int32) *ListSchemasResponseBodyList {
	s.VersionId = &v
	return s
}

func (s *ListSchemasResponseBodyList) Validate() error {
	return dara.Validate(s)
}
