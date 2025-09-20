// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTriggersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTriggersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTriggersRequest
	GetNextToken() *string
	SetOrder(v string) *ListTriggersRequest
	GetOrder() *string
	SetProjectName(v string) *ListTriggersRequest
	GetProjectName() *string
	SetSort(v string) *ListTriggersRequest
	GetSort() *string
	SetState(v string) *ListTriggersRequest
	GetState() *string
	SetTagSelector(v string) *ListTriggersRequest
	GetTagSelector() *string
}

type ListTriggersRequest struct {
	// The maximum number of entries to return. Valid values: 0 to 100.
	//
	// Default value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of triggers is greater than the value of MaxResults, you must specify NextToken.
	//
	// You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order. Default value: DESC.
	//
	// 	- ASC (default): ascending order.
	//
	// 	- DESC: descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The sort field. Valid values:
	//
	// 	- CreateTime: the point in time when the trigger is created.
	//
	// 	- UpdateTime: the most recent point in time when the trigger is updated.
	//
	// example:
	//
	// 2020-11-11T06:51:17.5Z
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The status of the trigger. Valid values:
	//
	// 	- Ready: The trigger is ready.
	//
	// 	- Running: The trigger is running.
	//
	// 	- Failed: The trigger failed and cannot be automatically recovered.
	//
	// 	- Suspended: The trigger is suspended.
	//
	// 	- Succeeded: The trigger is complete.
	//
	// example:
	//
	// Succeeded
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The custom tag. You can specify this parameter only if you specified Tags when you called the CreateTrigger operation.
	//
	// example:
	//
	// test=val1
	TagSelector *string `json:"TagSelector,omitempty" xml:"TagSelector,omitempty"`
}

func (s ListTriggersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTriggersRequest) GoString() string {
	return s.String()
}

func (s *ListTriggersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTriggersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTriggersRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTriggersRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListTriggersRequest) GetSort() *string {
	return s.Sort
}

func (s *ListTriggersRequest) GetState() *string {
	return s.State
}

func (s *ListTriggersRequest) GetTagSelector() *string {
	return s.TagSelector
}

func (s *ListTriggersRequest) SetMaxResults(v int32) *ListTriggersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTriggersRequest) SetNextToken(v string) *ListTriggersRequest {
	s.NextToken = &v
	return s
}

func (s *ListTriggersRequest) SetOrder(v string) *ListTriggersRequest {
	s.Order = &v
	return s
}

func (s *ListTriggersRequest) SetProjectName(v string) *ListTriggersRequest {
	s.ProjectName = &v
	return s
}

func (s *ListTriggersRequest) SetSort(v string) *ListTriggersRequest {
	s.Sort = &v
	return s
}

func (s *ListTriggersRequest) SetState(v string) *ListTriggersRequest {
	s.State = &v
	return s
}

func (s *ListTriggersRequest) SetTagSelector(v string) *ListTriggersRequest {
	s.TagSelector = &v
	return s
}

func (s *ListTriggersRequest) Validate() error {
	return dara.Validate(s)
}
