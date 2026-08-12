// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVersionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeVersionConfigRequest
	GetRegionId() *string
	SetSdkRequest(v *DescribeVersionConfigRequestSdkRequest) *DescribeVersionConfigRequest
	GetSdkRequest() *DescribeVersionConfigRequestSdkRequest
}

type DescribeVersionConfigRequest struct {
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-guangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Security Center SDK request.
	SdkRequest *DescribeVersionConfigRequestSdkRequest `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty" type:"Struct"`
}

func (s DescribeVersionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVersionConfigRequest) GetSdkRequest() *DescribeVersionConfigRequestSdkRequest {
	return s.SdkRequest
}

func (s *DescribeVersionConfigRequest) SetRegionId(v string) *DescribeVersionConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVersionConfigRequest) SetSdkRequest(v *DescribeVersionConfigRequestSdkRequest) *DescribeVersionConfigRequest {
	s.SdkRequest = v
	return s
}

func (s *DescribeVersionConfigRequest) Validate() error {
	if s.SdkRequest != nil {
		if err := s.SdkRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVersionConfigRequestSdkRequest struct {
	// The ID of the Alibaba Cloud account of the resource folder member accounts.
	//
	// >Invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 5815612291408486
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 2409:8a55:3827:cb50:5ad9:d5ff:fe87:f48c
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeVersionConfigRequestSdkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionConfigRequestSdkRequest) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigRequestSdkRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeVersionConfigRequestSdkRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeVersionConfigRequestSdkRequest) SetResourceDirectoryAccountId(v int64) *DescribeVersionConfigRequestSdkRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeVersionConfigRequestSdkRequest) SetSourceIp(v string) *DescribeVersionConfigRequestSdkRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVersionConfigRequestSdkRequest) Validate() error {
	return dara.Validate(s)
}
