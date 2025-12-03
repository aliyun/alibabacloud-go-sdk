// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFederationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFederationId(v string) *ListFederationsRequest
	GetFederationId() *string
	SetFileSystemId(v string) *ListFederationsRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *ListFederationsRequest
	GetInputRegionId() *string
	SetMaxResults(v int32) *ListFederationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListFederationsRequest
	GetNextToken() *string
}

type ListFederationsRequest struct {
	FederationId *string `json:"FederationId,omitempty" xml:"FederationId,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListFederationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFederationsRequest) GoString() string {
	return s.String()
}

func (s *ListFederationsRequest) GetFederationId() *string {
	return s.FederationId
}

func (s *ListFederationsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListFederationsRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *ListFederationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFederationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFederationsRequest) SetFederationId(v string) *ListFederationsRequest {
	s.FederationId = &v
	return s
}

func (s *ListFederationsRequest) SetFileSystemId(v string) *ListFederationsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListFederationsRequest) SetInputRegionId(v string) *ListFederationsRequest {
	s.InputRegionId = &v
	return s
}

func (s *ListFederationsRequest) SetMaxResults(v int32) *ListFederationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFederationsRequest) SetNextToken(v string) *ListFederationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFederationsRequest) Validate() error {
	return dara.Validate(s)
}
