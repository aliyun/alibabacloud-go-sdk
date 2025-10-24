// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllMediaBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaximumPageSize(v int32) *ListAllMediaBucketRequest
	GetMaximumPageSize() *int32
	SetNextPageToken(v string) *ListAllMediaBucketRequest
	GetNextPageToken() *string
	SetOwnerAccount(v string) *ListAllMediaBucketRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListAllMediaBucketRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListAllMediaBucketRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAllMediaBucketRequest
	GetResourceOwnerId() *int64
}

type ListAllMediaBucketRequest struct {
	// The maximum number of media buckets to return. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 10
	MaximumPageSize *int32 `json:"MaximumPageSize,omitempty" xml:"MaximumPageSize,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. The response to the first request contains this parameter, which is added to the next request.
	//
	// example:
	//
	// P2Zqo1PLGhZdygo-ajSsjUX5zrBHCgXy6j4hEvv****
	NextPageToken        *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAllMediaBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllMediaBucketRequest) GoString() string {
	return s.String()
}

func (s *ListAllMediaBucketRequest) GetMaximumPageSize() *int32 {
	return s.MaximumPageSize
}

func (s *ListAllMediaBucketRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListAllMediaBucketRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAllMediaBucketRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAllMediaBucketRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAllMediaBucketRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAllMediaBucketRequest) SetMaximumPageSize(v int32) *ListAllMediaBucketRequest {
	s.MaximumPageSize = &v
	return s
}

func (s *ListAllMediaBucketRequest) SetNextPageToken(v string) *ListAllMediaBucketRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListAllMediaBucketRequest) SetOwnerAccount(v string) *ListAllMediaBucketRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAllMediaBucketRequest) SetOwnerId(v int64) *ListAllMediaBucketRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAllMediaBucketRequest) SetResourceOwnerAccount(v string) *ListAllMediaBucketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAllMediaBucketRequest) SetResourceOwnerId(v int64) *ListAllMediaBucketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAllMediaBucketRequest) Validate() error {
	return dara.Validate(s)
}
