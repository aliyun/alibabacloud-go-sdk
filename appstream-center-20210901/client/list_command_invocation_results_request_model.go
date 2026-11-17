// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommandInvocationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *ListCommandInvocationResultsRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupIds(v []*string) *ListCommandInvocationResultsRequest
	GetAppInstanceGroupIds() []*string
	SetAppInstancePersistentIds(v []*string) *ListCommandInvocationResultsRequest
	GetAppInstancePersistentIds() []*string
	SetInvocationId(v string) *ListCommandInvocationResultsRequest
	GetInvocationId() *string
	SetMaxResults(v int32) *ListCommandInvocationResultsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCommandInvocationResultsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListCommandInvocationResultsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCommandInvocationResultsRequest
	GetPageSize() *int32
	SetProductType(v string) *ListCommandInvocationResultsRequest
	GetProductType() *string
}

type ListCommandInvocationResultsRequest struct {
	// The ID of the delivery group. Queries the command execution results under this delivery group. This parameter must be specified when AppInstanceGroupIds is not specified. If both are specified, AppInstanceGroupIds takes precedence and this parameter does not take effect.
	//
	// example:
	//
	// ws-0bxls9m9arax6****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The list of delivery group IDs. A maximum of 200 IDs can be specified, and all delivery groups must be in the same region. When specified, it takes precedence over AppInstanceGroupId and queries the command execution results of all persistent sessions under these delivery groups. The list cannot contain null values.
	AppInstanceGroupIds []*string `json:"AppInstanceGroupIds,omitempty" xml:"AppInstanceGroupIds,omitempty" type:"Repeated"`
	// The list of persistent session IDs. A maximum of 100 IDs can be specified. It is used to query only the command execution results of the specified persistent sessions. When AppInstanceGroupIds is specified, only persistent sessions belonging to these delivery groups are retained. If no matching persistent sessions are found, an empty list is returned.
	AppInstancePersistentIds []*string `json:"AppInstancePersistentIds,omitempty" xml:"AppInstancePersistentIds,omitempty" type:"Repeated"`
	// The command execution ID, which is the InvocationId returned when you call RunAppInstanceCommand to deliver a command.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-0bxls9m9ara****
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	// This parameter does not take effect in the current version and does not affect the query results if passed. Use PageNumber and PageSize for pagination.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// This parameter does not take effect in the current version and does not affect the query results if passed. Use PageNumber and PageSize for pagination.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kDw2Ml****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number of the query results to display. It is used together with PageSize for paging in a paged query. If not specified, the default value is 0, which indicates the first page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of query results per page. Default value: 200.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. Currently, only WuyingServer is supported. If you specify other values, a parameter error is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// WuyingServer
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListCommandInvocationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCommandInvocationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListCommandInvocationResultsRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListCommandInvocationResultsRequest) GetAppInstanceGroupIds() []*string {
	return s.AppInstanceGroupIds
}

func (s *ListCommandInvocationResultsRequest) GetAppInstancePersistentIds() []*string {
	return s.AppInstancePersistentIds
}

func (s *ListCommandInvocationResultsRequest) GetInvocationId() *string {
	return s.InvocationId
}

func (s *ListCommandInvocationResultsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCommandInvocationResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCommandInvocationResultsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCommandInvocationResultsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCommandInvocationResultsRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListCommandInvocationResultsRequest) SetAppInstanceGroupId(v string) *ListCommandInvocationResultsRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListCommandInvocationResultsRequest) SetAppInstanceGroupIds(v []*string) *ListCommandInvocationResultsRequest {
	s.AppInstanceGroupIds = v
	return s
}

func (s *ListCommandInvocationResultsRequest) SetAppInstancePersistentIds(v []*string) *ListCommandInvocationResultsRequest {
	s.AppInstancePersistentIds = v
	return s
}

func (s *ListCommandInvocationResultsRequest) SetInvocationId(v string) *ListCommandInvocationResultsRequest {
	s.InvocationId = &v
	return s
}

func (s *ListCommandInvocationResultsRequest) SetMaxResults(v int32) *ListCommandInvocationResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCommandInvocationResultsRequest) SetNextToken(v string) *ListCommandInvocationResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCommandInvocationResultsRequest) SetPageNumber(v int32) *ListCommandInvocationResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCommandInvocationResultsRequest) SetPageSize(v int32) *ListCommandInvocationResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCommandInvocationResultsRequest) SetProductType(v string) *ListCommandInvocationResultsRequest {
	s.ProductType = &v
	return s
}

func (s *ListCommandInvocationResultsRequest) Validate() error {
	return dara.Validate(s)
}
