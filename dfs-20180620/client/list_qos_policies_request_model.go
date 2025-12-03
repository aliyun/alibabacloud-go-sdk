// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQosPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFederationId(v string) *ListQosPoliciesRequest
	GetFederationId() *string
	SetFileSystemId(v string) *ListQosPoliciesRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *ListQosPoliciesRequest
	GetInputRegionId() *string
	SetMaxResults(v int32) *ListQosPoliciesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListQosPoliciesRequest
	GetNextToken() *string
}

type ListQosPoliciesRequest struct {
	FederationId *string `json:"FederationId,omitempty" xml:"FederationId,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListQosPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQosPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListQosPoliciesRequest) GetFederationId() *string {
	return s.FederationId
}

func (s *ListQosPoliciesRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListQosPoliciesRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *ListQosPoliciesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQosPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQosPoliciesRequest) SetFederationId(v string) *ListQosPoliciesRequest {
	s.FederationId = &v
	return s
}

func (s *ListQosPoliciesRequest) SetFileSystemId(v string) *ListQosPoliciesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListQosPoliciesRequest) SetInputRegionId(v string) *ListQosPoliciesRequest {
	s.InputRegionId = &v
	return s
}

func (s *ListQosPoliciesRequest) SetMaxResults(v int32) *ListQosPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQosPoliciesRequest) SetNextToken(v string) *ListQosPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListQosPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
