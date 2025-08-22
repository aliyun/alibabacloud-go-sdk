// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobScriptHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListJobScriptHistoryRequest
	GetAppName() *string
	SetClusterId(v string) *ListJobScriptHistoryRequest
	GetClusterId() *string
	SetJobId(v int64) *ListJobScriptHistoryRequest
	GetJobId() *int64
	SetMaxResults(v int32) *ListJobScriptHistoryRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListJobScriptHistoryRequest
	GetNextToken() *string
}

type ListJobScriptHistoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListJobScriptHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobScriptHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListJobScriptHistoryRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListJobScriptHistoryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListJobScriptHistoryRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListJobScriptHistoryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListJobScriptHistoryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListJobScriptHistoryRequest) SetAppName(v string) *ListJobScriptHistoryRequest {
	s.AppName = &v
	return s
}

func (s *ListJobScriptHistoryRequest) SetClusterId(v string) *ListJobScriptHistoryRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobScriptHistoryRequest) SetJobId(v int64) *ListJobScriptHistoryRequest {
	s.JobId = &v
	return s
}

func (s *ListJobScriptHistoryRequest) SetMaxResults(v int32) *ListJobScriptHistoryRequest {
	s.MaxResults = &v
	return s
}

func (s *ListJobScriptHistoryRequest) SetNextToken(v string) *ListJobScriptHistoryRequest {
	s.NextToken = &v
	return s
}

func (s *ListJobScriptHistoryRequest) Validate() error {
	return dara.Validate(s)
}
