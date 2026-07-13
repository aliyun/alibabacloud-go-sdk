// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListModelsRequest
	GetId() *string
	SetInstanceId(v string) *ListModelsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListModelsRequest
	GetMaxResults() *int32
	SetName(v string) *ListModelsRequest
	GetName() *string
	SetNextToken(v string) *ListModelsRequest
	GetNextToken() *string
	SetProviderName(v string) *ListModelsRequest
	GetProviderName() *string
}

type ListModelsRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AgentTeams
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s ListModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModelsRequest) GetId() *string {
	return s.Id
}

func (s *ListModelsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListModelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelsRequest) GetName() *string {
	return s.Name
}

func (s *ListModelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelsRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListModelsRequest) SetId(v string) *ListModelsRequest {
	s.Id = &v
	return s
}

func (s *ListModelsRequest) SetInstanceId(v string) *ListModelsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListModelsRequest) SetMaxResults(v int32) *ListModelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelsRequest) SetName(v string) *ListModelsRequest {
	s.Name = &v
	return s
}

func (s *ListModelsRequest) SetNextToken(v string) *ListModelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelsRequest) SetProviderName(v string) *ListModelsRequest {
	s.ProviderName = &v
	return s
}

func (s *ListModelsRequest) Validate() error {
	return dara.Validate(s)
}
