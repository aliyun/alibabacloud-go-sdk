// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListStackConfigsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListStackConfigsRequest
	GetNextToken() *string
	SetStatus(v string) *ListStackConfigsRequest
	GetStatus() *string
	SetVersion(v string) *ListStackConfigsRequest
	GetVersion() *string
}

type ListStackConfigsRequest struct {
	// The maximum number of records to return for the current query. Default value: 20. Maximum value: 200.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token that specifies the starting position for the query. If this parameter is left empty, the query starts from the beginning.
	//
	// example:
	//
	// s8UVlnE23gZvjCvCwkoZ7Z4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The status of the stack configuration. Valid values:
	//
	// | Name | Description |
	//
	// |------|-------------|
	//
	// | Creating | Creating |
	//
	// | Created | Created |
	//
	// | Waiting | Waiting for deployment |
	//
	// | Deploying | Deploying |
	//
	// | Deployed | Deployed |
	//
	// | Errored | Deployment failed |
	//
	// | Deleting | Deleting |
	//
	// | Deleted | Deleted |
	//
	// | DeleteFailed | Deletion failed |
	//
	// | DetectTriggered | Drift detection triggered successfully |
	//
	// example:
	//
	// Deployed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The version number of the stack configuration, such as v1. The initial value is v1. The version number increments each time the stack is updated or refreshed and the configuration is changed.
	//
	// example:
	//
	// v4
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListStackConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStackConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListStackConfigsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListStackConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListStackConfigsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListStackConfigsRequest) GetVersion() *string {
	return s.Version
}

func (s *ListStackConfigsRequest) SetMaxResults(v int32) *ListStackConfigsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListStackConfigsRequest) SetNextToken(v string) *ListStackConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListStackConfigsRequest) SetStatus(v string) *ListStackConfigsRequest {
	s.Status = &v
	return s
}

func (s *ListStackConfigsRequest) SetVersion(v string) *ListStackConfigsRequest {
	s.Version = &v
	return s
}

func (s *ListStackConfigsRequest) Validate() error {
	return dara.Validate(s)
}
