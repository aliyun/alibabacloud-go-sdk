// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *ListSnapshotsRequest
	GetFunctionName() *string
	SetLimit(v int32) *ListSnapshotsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListSnapshotsRequest
	GetNextToken() *string
	SetQualifier(v string) *ListSnapshotsRequest
	GetQualifier() *string
	SetSessionId(v string) *ListSnapshotsRequest
	GetSessionId() *string
}

type ListSnapshotsRequest struct {
	// The function name.
	//
	// example:
	//
	// my-func
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The maximum number of snapshots to return. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token used to retrieve more results.
	//
	// example:
	//
	// caeba0be03****f84eb48b699f0a4883
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The function alias.
	//
	// example:
	//
	// alias
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// The source session ID from which the snapshot was created. When specified, functionName must also be specified.
	//
	// example:
	//
	// test-session-id-1
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s ListSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListSnapshotsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListSnapshotsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSnapshotsRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *ListSnapshotsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListSnapshotsRequest) SetFunctionName(v string) *ListSnapshotsRequest {
	s.FunctionName = &v
	return s
}

func (s *ListSnapshotsRequest) SetLimit(v int32) *ListSnapshotsRequest {
	s.Limit = &v
	return s
}

func (s *ListSnapshotsRequest) SetNextToken(v string) *ListSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSnapshotsRequest) SetQualifier(v string) *ListSnapshotsRequest {
	s.Qualifier = &v
	return s
}

func (s *ListSnapshotsRequest) SetSessionId(v string) *ListSnapshotsRequest {
	s.SessionId = &v
	return s
}

func (s *ListSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
