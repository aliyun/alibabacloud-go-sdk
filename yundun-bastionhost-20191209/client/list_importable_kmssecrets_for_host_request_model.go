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
	// The ID of the host.
	//
	// > Call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to get this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *int32 `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the bastion host instance.
	//
	// > Call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to get this ID.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the secret to search for. Fuzzy matching is supported.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results.
	//
	// > You do not need to specify this parameter for the first request. For subsequent requests, use the `NextToken` value from the previous response.
	//
	// example:
	//
	// 42bc833a0a0002dae0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the bastion host.
	//
	// > For details about the mapping between region IDs and region names, see [Regions and availability zones](https://help.aliyun.com/document_detail/40654.html).
	//
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
