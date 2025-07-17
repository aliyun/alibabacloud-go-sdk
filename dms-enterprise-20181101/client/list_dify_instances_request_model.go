// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDifyInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListDifyInstancesRequest
	GetClientToken() *string
	SetDataRegion(v string) *ListDifyInstancesRequest
	GetDataRegion() *string
	SetMaxResults(v int32) *ListDifyInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDifyInstancesRequest
	GetNextToken() *string
}

type ListDifyInstancesRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataRegion  *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListDifyInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDifyInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListDifyInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListDifyInstancesRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *ListDifyInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDifyInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDifyInstancesRequest) SetClientToken(v string) *ListDifyInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListDifyInstancesRequest) SetDataRegion(v string) *ListDifyInstancesRequest {
	s.DataRegion = &v
	return s
}

func (s *ListDifyInstancesRequest) SetMaxResults(v int32) *ListDifyInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDifyInstancesRequest) SetNextToken(v string) *ListDifyInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDifyInstancesRequest) Validate() error {
	return dara.Validate(s)
}
