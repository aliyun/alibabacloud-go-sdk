// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImportableKMSSecretsForHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostId(v int32) *ListImportableKMSSecretsForHostRequest
	GetHostId() *int32
	SetInstanceId(v string) *ListImportableKMSSecretsForHostRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListImportableKMSSecretsForHostRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListImportableKMSSecretsForHostRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListImportableKMSSecretsForHostRequest
	GetNextToken() *string
	SetRegionId(v string) *ListImportableKMSSecretsForHostRequest
	GetRegionId() *string
}

type ListImportableKMSSecretsForHostRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *int32 `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 42bc833a0a0002dae0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListImportableKMSSecretsForHostRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImportableKMSSecretsForHostRequest) GoString() string {
	return s.String()
}

func (s *ListImportableKMSSecretsForHostRequest) GetHostId() *int32 {
	return s.HostId
}

func (s *ListImportableKMSSecretsForHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListImportableKMSSecretsForHostRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListImportableKMSSecretsForHostRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListImportableKMSSecretsForHostRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListImportableKMSSecretsForHostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImportableKMSSecretsForHostRequest) SetHostId(v int32) *ListImportableKMSSecretsForHostRequest {
	s.HostId = &v
	return s
}

func (s *ListImportableKMSSecretsForHostRequest) SetInstanceId(v string) *ListImportableKMSSecretsForHostRequest {
	s.InstanceId = &v
	return s
}

func (s *ListImportableKMSSecretsForHostRequest) SetKeyword(v string) *ListImportableKMSSecretsForHostRequest {
	s.Keyword = &v
	return s
}

func (s *ListImportableKMSSecretsForHostRequest) SetMaxResults(v int32) *ListImportableKMSSecretsForHostRequest {
	s.MaxResults = &v
	return s
}

func (s *ListImportableKMSSecretsForHostRequest) SetNextToken(v string) *ListImportableKMSSecretsForHostRequest {
	s.NextToken = &v
	return s
}

func (s *ListImportableKMSSecretsForHostRequest) SetRegionId(v string) *ListImportableKMSSecretsForHostRequest {
	s.RegionId = &v
	return s
}

func (s *ListImportableKMSSecretsForHostRequest) Validate() error {
	return dara.Validate(s)
}
