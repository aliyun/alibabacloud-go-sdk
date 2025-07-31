// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKernelReleaseNotesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKernelVersion(v string) *DescribeKernelReleaseNotesRequest
	GetKernelVersion() *string
	SetOwnerAccount(v string) *DescribeKernelReleaseNotesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeKernelReleaseNotesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeKernelReleaseNotesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeKernelReleaseNotesRequest
	GetResourceOwnerId() *int64
}

type DescribeKernelReleaseNotesRequest struct {
	// The minor version number of the instance. Example: **mongodb_20180522_0.4.8**.
	//
	// 	- This parameter is required. After you specify a version number for this parameter in a request, the release notes of the versions later than this version are returned.
	//
	// example:
	//
	// mongodb_20180522_0.4.8
	KernelVersion        *string `json:"KernelVersion,omitempty" xml:"KernelVersion,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeKernelReleaseNotesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKernelReleaseNotesRequest) GoString() string {
	return s.String()
}

func (s *DescribeKernelReleaseNotesRequest) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *DescribeKernelReleaseNotesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeKernelReleaseNotesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeKernelReleaseNotesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeKernelReleaseNotesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeKernelReleaseNotesRequest) SetKernelVersion(v string) *DescribeKernelReleaseNotesRequest {
	s.KernelVersion = &v
	return s
}

func (s *DescribeKernelReleaseNotesRequest) SetOwnerAccount(v string) *DescribeKernelReleaseNotesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeKernelReleaseNotesRequest) SetOwnerId(v int64) *DescribeKernelReleaseNotesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeKernelReleaseNotesRequest) SetResourceOwnerAccount(v string) *DescribeKernelReleaseNotesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeKernelReleaseNotesRequest) SetResourceOwnerId(v int64) *DescribeKernelReleaseNotesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeKernelReleaseNotesRequest) Validate() error {
	return dara.Validate(s)
}
