// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceInspectionResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListInstanceInspectionResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstanceInspectionResourcesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListInstanceInspectionResourcesRequest
	GetRegionId() *string
	SetRunId(v string) *ListInstanceInspectionResourcesRequest
	GetRunId() *string
	SetTemplateCode(v string) *ListInstanceInspectionResourcesRequest
	GetTemplateCode() *string
}

type ListInstanceInspectionResourcesRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// inspection-run-001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
}

func (s ListInstanceInspectionResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceInspectionResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceInspectionResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceInspectionResourcesRequest) GetRunId() *string {
	return s.RunId
}

func (s *ListInstanceInspectionResourcesRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListInstanceInspectionResourcesRequest) SetMaxResults(v int32) *ListInstanceInspectionResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceInspectionResourcesRequest) SetNextToken(v string) *ListInstanceInspectionResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstanceInspectionResourcesRequest) SetRegionId(v string) *ListInstanceInspectionResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceInspectionResourcesRequest) SetRunId(v string) *ListInstanceInspectionResourcesRequest {
	s.RunId = &v
	return s
}

func (s *ListInstanceInspectionResourcesRequest) SetTemplateCode(v string) *ListInstanceInspectionResourcesRequest {
	s.TemplateCode = &v
	return s
}

func (s *ListInstanceInspectionResourcesRequest) Validate() error {
	return dara.Validate(s)
}
