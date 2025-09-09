// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPublicKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListUserPublicKeysRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListUserPublicKeysRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListUserPublicKeysRequest
	GetPageSize() *string
	SetRegionId(v string) *ListUserPublicKeysRequest
	GetRegionId() *string
	SetUserId(v string) *ListUserPublicKeysRequest
	GetUserId() *string
}

type ListUserPublicKeysRequest struct {
	// The ID of the bastion host on which you want to query all public keys of the user.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.\\
	//
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 50
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host on which you want to query all public keys of the user.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose public keys you want to query.
	//
	// example:
	//
	// 2
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserPublicKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserPublicKeysRequest) GoString() string {
	return s.String()
}

func (s *ListUserPublicKeysRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserPublicKeysRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListUserPublicKeysRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListUserPublicKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUserPublicKeysRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListUserPublicKeysRequest) SetInstanceId(v string) *ListUserPublicKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserPublicKeysRequest) SetPageNumber(v string) *ListUserPublicKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserPublicKeysRequest) SetPageSize(v string) *ListUserPublicKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserPublicKeysRequest) SetRegionId(v string) *ListUserPublicKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserPublicKeysRequest) SetUserId(v string) *ListUserPublicKeysRequest {
	s.UserId = &v
	return s
}

func (s *ListUserPublicKeysRequest) Validate() error {
	return dara.Validate(s)
}
